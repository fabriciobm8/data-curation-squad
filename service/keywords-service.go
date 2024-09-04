package service

import (
    "context"
    "data-curation-squad/model"
    "data-curation-squad/repository"
    "errors"
	"go.mongodb.org/mongo-driver/mongo"

)

type KeywordService struct {
    repo repository.KeywordRepository
}

func NewKeywordService(repo repository.KeywordRepository) *KeywordService {
    return &KeywordService{repo: repo}
}

func (s *KeywordService) validateKeyword(kw *model.Keyword) error {
    if kw.CourseId == "" {
        return errors.New("courseID é obrigatório")
    }
    if kw.Keyword == "" {
        return errors.New("keyword é obrigatório")
    }
    if kw.ClassMaterialId == "" {
        return errors.New("classMaterialId é obrigatório")
    }
    return nil
}

func (s *KeywordService) Create(ctx context.Context, keyword *model.Keyword) error {
    if err := s.validateKeyword(keyword); err != nil {
        return err
    }   
    
    return s.repo.Create(ctx, keyword)
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