package service

import (
    "context"
    "data-curation-squad/model"
    "data-curation-squad/repository"
    "errors"
	"go.mongodb.org/mongo-driver/mongo"

)

type TranscriptTimeService struct {
    repo repository.TranscriptTimeRepository
}

func NewTranscriptTimeService(repo repository.TranscriptTimeRepository) *TranscriptTimeService {
    return &TranscriptTimeService{repo: repo}
}

// Valida se o TranscriptTime possui todos os campos obrigatórios
func (s *TranscriptTimeService) validateTranscriptTime(tt *model.TranscriptTime) error {
    if tt.Transcript == "" {
        return errors.New("transcript é obrigatório")
    }
    if tt.ClassMaterialId == "" {
        return errors.New("classMaterialId é obrigatório")
    }    
    return nil
}

// Cria um novo TranscriptTime
func (s *TranscriptTimeService) Create(ctx context.Context, transcriptTime *model.TranscriptTime) error {
    if err := s.validateTranscriptTime(transcriptTime); err != nil {
        return err
    }  
    
    return s.repo.Create(ctx, transcriptTime)
}

func (s *TranscriptTimeService) FindAll(ctx context.Context) ([]model.TranscriptTime, error) {
    return s.repo.FindAll(ctx)
}

func (s *TranscriptTimeService) FindByID(ctx context.Context, id string) (*model.TranscriptTime, error) {
    if id == "" {
        return nil, errors.New("id é obrigatório")
    }
    
    transcriptTime, err := s.repo.FindByID(ctx, id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("transcriptTime não encontrado")
        }
        return nil, err
    }
    
    return transcriptTime, nil
}
