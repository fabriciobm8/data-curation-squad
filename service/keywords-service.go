package service

import (
	"context"
	"data-curation-squad/model"
	"data-curation-squad/repository"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type KeywordService struct {
	repo                   repository.KeywordRepository
	ClassMaterialService   *ClassMaterialService
}

func NewKeywordService(repo repository.KeywordRepository, classMaterialService *ClassMaterialService) *KeywordService {
	return &KeywordService{repo: repo, ClassMaterialService: classMaterialService}
}

func (s *KeywordService) FindAll(ctx context.Context) ([]model.Keyword, error) {
	return s.repo.FindAll(ctx)
}

func (s *KeywordService) FindByID(ctx context.Context, id string) (*model.Keyword, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	
	keyword, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("keyword não encontrado")
		}
		return nil, err
	}
	
	return keyword, nil
}

func (s *KeywordService) SaveKeywords(ctx context.Context, classMaterialID string, keywordList []string) error {

	var keywords []model.Keyword
	for _, kw := range keywordList {
		lowerKw := strings.ToLower(kw)
		keyword := model.Keyword{ Keyword: lowerKw }
		keywords = append(keywords, keyword)
	}

	for _, keyword := range keywords {
		_, err := s.repo.FindByKeyword(ctx, keyword.Keyword)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
	}

	err := s.repo.SaveKeywords(ctx, keywords)
	if err != nil {
		return err
	}

	classMaterial, err := s.ClassMaterialService.FindByID(ctx, classMaterialID)
    if err != nil {
        return err
    }

	for _, keyword := range keywords {
		savedKeyword, err := s.repo.FindByKeyword(ctx, keyword.Keyword)
		if err != nil {
			return err
		}
		classMaterial.Keyword = append(classMaterial.Keyword, savedKeyword.ID.Hex())
	}

	return s.ClassMaterialService.UpdateKeywords(ctx, classMaterialID, classMaterial.Keyword)
}

func (s *KeywordService) FindKeywordIDByName(ctx context.Context, keywordName string) (string, error) {
	keyword, err := s.repo.FindByKeyword(ctx, keywordName)
	if err != nil {
		return "", err
	}
	if keyword == nil {
		return "", errors.New("keyword não encontrada")
	}
	return keyword.ID.Hex(), nil
}
