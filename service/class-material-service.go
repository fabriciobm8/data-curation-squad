package service

import (
	"context"
	"data-curation-squad/model"
	"data-curation-squad/repository"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type ClassMaterialService struct {
    repo repository.ClassMaterialRepository
}

// Cria uma nova instância de ClassMaterialService
func NewClassMaterialService(repo repository.ClassMaterialRepository) *ClassMaterialService {
    return &ClassMaterialService{repo: repo}
}

// Valida se o ClassMaterial possui todos os campos obrigatórios
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

// Cria um novo ClassMaterial
func (s *ClassMaterialService) Create(ctx context.Context, classMaterial *model.ClassMaterial) error {
    if err := s.validateClassMaterial(classMaterial); err != nil {
        return err
    }
    
    // Verifica se já existe um ClassMaterial com o mesmo Id
    existingCM, _ := s.repo.FindByID(ctx, classMaterial.Id)
    if existingCM != nil {
        return errors.New("classMaterial já existe")
    }
    
    return s.repo.Create(ctx, classMaterial)
}

// Retorna todos os ClassMaterials
func (s *ClassMaterialService) FindAll(ctx context.Context) ([]model.ClassMaterial, error) {
    return s.repo.FindAll(ctx)
}

// Encontra um ClassMaterial por ID
func (s *ClassMaterialService) FindByID(ctx context.Context, id string) (*model.ClassMaterial, error) {
    if id == "" {
        return nil, errors.New("id é obrigatório")
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

