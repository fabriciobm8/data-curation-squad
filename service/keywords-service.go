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
	repo           repository.KeywordRepository
	classMaterialService *ClassMaterialService
}

func NewKeywordService(repo repository.KeywordRepository, classMaterialService *ClassMaterialService) *KeywordService {
	return &KeywordService{repo: repo, classMaterialService: classMaterialService}
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
        // Converte a palavra-chave para minúsculas
        lowerKw := strings.ToLower(kw)
        keyword := model.Keyword{
            Keyword: lowerKw,
        }
        keywords = append(keywords, keyword)
    }

    // Verifica se a keyword já existe na coleção
    for _, keyword := range keywords {
        _, err := s.repo.FindByKeyword(ctx, keyword.Keyword)
        if err != nil {
            if err != mongo.ErrNoDocuments {
                return err
            }
        }
    }

    // Salva as keywords no banco de dados
    err := s.repo.SaveKeywords(ctx, keywords)
    if err != nil {
        return err
    }

    // Adiciona o ID da keyword ao ClassMaterial correspondente
    classMaterial, err := s.classMaterialService.FindByID(ctx, classMaterialID)
    if err != nil {
        return err
    }

    for _, keyword := range keywords {
        // Encontre o ID da keyword que foi salva no banco de dados
        savedKeyword, err := s.repo.FindByKeyword(ctx, keyword.Keyword)
        if err != nil {
            return err
        }
        classMaterial.Keyword = append(classMaterial.Keyword, savedKeyword.ID.Hex())
    }

    // Atualiza o ClassMaterial com as novas keywords
    return s.classMaterialService.UpdateKeywords(ctx, classMaterialID, classMaterial.Keyword)
}
