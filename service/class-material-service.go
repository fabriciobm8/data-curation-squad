package service

import (
	"context"
	"data-curation-squad/model"
	"data-curation-squad/repository"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
    "strings"
)

type ClassMaterialService struct {
	repo           repository.ClassMaterialRepository
	keywordService *KeywordService
}

func NewClassMaterialService(repo repository.ClassMaterialRepository, keywordService *KeywordService) *ClassMaterialService {
	return &ClassMaterialService{repo: repo, keywordService: keywordService}
}

func (s *ClassMaterialService) validateClassMaterial(cm *model.ClassMaterial) error {
	if cm.MaterialId == "" {
		return errors.New("MaterialId é obrigatório")
	}
	if cm.MaterialType == "" {
		return errors.New("materialType é obrigatório")
	}
	if cm.MaterialType != "video" && cm.MaterialType != "pdf" {
		return errors.New("materialType deve ser 'video' ou 'pdf'")
	}
	return nil
}

func (s *ClassMaterialService) Create(ctx context.Context, classMaterial *model.ClassMaterial) error {
	if err := s.validateClassMaterial(classMaterial); err != nil {
		return err
	}

	existingCM, _ := s.repo.FindByID(ctx, classMaterial.Id)
	if existingCM != nil {
		return errors.New("classMaterial já existe")
	}

	return s.repo.Create(ctx, classMaterial)
}

func (s *ClassMaterialService) FindAll(ctx context.Context) ([]model.ClassMaterial, error) {
	return s.repo.FindAll(ctx)
}

func (s *ClassMaterialService) FindByID(ctx context.Context, id string) (*model.ClassMaterial, error) {
    if s.keywordService == nil {
        return nil, errors.New("keywordService is nil")
    }

	classMaterial, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("classMaterial não encontrado")
		}
		return nil, err
	}

	return classMaterial, nil
}

func (s *ClassMaterialService) UpdateTranscriptTime(ctx context.Context, classMaterialID string, transcriptTimes []model.TranscriptTime) error {
	_, err := s.FindByID(ctx, classMaterialID)
	if err != nil {
		return err
	}

	for i, tt := range transcriptTimes {
		for j, keyword := range tt.Keywords {
			// Converta a keyword para minúsculas
			keywordLower := strings.ToLower(keyword)
			keywordID, err := s.keywordService.FindKeywordIDByName(ctx, keywordLower)
			if err == nil {
				transcriptTimes[i].Keywords[j] = keywordID
			} else {
				return err
			}
		}
	}

	return s.repo.UpdateTranscriptTime(ctx, classMaterialID, transcriptTimes)
}


func (s *ClassMaterialService) UpdateKeywords(ctx context.Context, classMaterialID string, newKeywords []string) error {
	classMaterial, err := s.FindByID(ctx, classMaterialID)
	if err != nil {
		return err
	}

	existingKeywordMap := make(map[string]struct{})
	for _, id := range classMaterial.Keyword {
		existingKeywordMap[id] = struct{}{}
	}

	for _, newID := range newKeywords {
		if _, exists := existingKeywordMap[newID]; !exists {
			classMaterial.Keyword = append(classMaterial.Keyword, newID)
			existingKeywordMap[newID] = struct{}{}
		}
	}

	return s.repo.UpdateKeywords(ctx, classMaterial)
}
