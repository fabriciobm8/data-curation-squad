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

// Transforma uma lista de strings em uma lista de models.Keyword e chama o repositório para salvá-las
func (s *KeywordService) SaveKeywords(ctx context.Context, keywordList []string) error {
    var keywords []model.Keyword
    for _, kw := range keywordList {
        keyword := model.Keyword{
            Keyword:    kw,
            UsageCount: 0, // Inicializa usageCount como 0
        }
        keywords = append(keywords, keyword)
    }
    return s.repo.SaveKeywords(ctx, keywords)
}