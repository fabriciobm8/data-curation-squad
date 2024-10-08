package repository

import (
	"context"
	"data-curation-squad/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassMaterialRepository interface {
	Create(ctx context.Context, classMaterial *model.ClassMaterial) error
	FindAll(ctx context.Context) ([]model.ClassMaterial, error)
	FindByID(ctx context.Context, id string) (*model.ClassMaterial, error)
	UpdateTranscriptTime(ctx context.Context, id string, transcriptTime []model.TranscriptTime) error
	UpdateKeywords(ctx context.Context, classMaterial *model.ClassMaterial) error
	GetByCourseId(ctx context.Context, courseId string) ([]model.ClassMaterial, error)
	GetByObjectiveId(ctx context.Context, objectiveId string) ([]model.ClassMaterial, error)
	GetByMaterialId(ctx context.Context, materialId string) ([]model.ClassMaterial, error)
	FindByKeywordIds(ctx context.Context, keywordIds []string) ([]model.ClassMaterial, error)
}

type classMaterialRepository struct {
	collection *mongo.Collection
}

func NewClassMaterialRepository(client *mongo.Client) ClassMaterialRepository {
	collection := client.Database("class").Collection("classMaterial")
	return &classMaterialRepository{collection: collection}
}

func (r *classMaterialRepository) Create(ctx context.Context, classMaterial *model.ClassMaterial) error {
	_, err := r.collection.InsertOne(ctx, classMaterial)
	return err
}

func (r *classMaterialRepository) FindAll(ctx context.Context) ([]model.ClassMaterial, error) {
	var classMaterials []model.ClassMaterial
	cursor, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var classMaterial model.ClassMaterial
		if err := cursor.Decode(&classMaterial); err != nil {
			return nil, err
		}
		classMaterials = append(classMaterials, classMaterial)
	}
	return classMaterials, nil
}

func (r *classMaterialRepository) FindByID(ctx context.Context, id string) (*model.ClassMaterial, error) {
	var classMaterial model.ClassMaterial
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&classMaterial)
	if err != nil {
		return nil, err
	}
	return &classMaterial, nil
}

func (r *classMaterialRepository) UpdateTranscriptTime(ctx context.Context, id string, transcriptTime []model.TranscriptTime) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"TranscriptTime": transcriptTime}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *classMaterialRepository) UpdateKeywords(ctx context.Context, classMaterial *model.ClassMaterial) error {
    // Filtra pelo ID do ClassMaterial
    filter := bson.M{"_id": classMaterial.Id}

    // Recupera o ClassMaterial existente para verificar as keywords
    existingClassMaterial := &model.ClassMaterial{}
    err := r.collection.FindOne(ctx, filter).Decode(existingClassMaterial)
    if err != nil {
        return err
    }

    // Cria um mapa para verificar se a keyword já existe
    existingKeywordMap := make(map[string]struct{})
    for _, keyword := range existingClassMaterial.Keyword {
        existingKeywordMap[keyword] = struct{}{}
    }

    // Adiciona apenas novas keywords
    for _, newKeyword := range classMaterial.Keyword {
        if _, exists := existingKeywordMap[newKeyword]; !exists {
            existingClassMaterial.Keyword = append(existingClassMaterial.Keyword, newKeyword)
            existingKeywordMap[newKeyword] = struct{}{} // Adiciona ao mapa para futuras verificações
        }
    }

    // Atualiza o ClassMaterial com a lista de keywords ajustada
    update := bson.M{"$set": bson.M{
        "Keyword":       existingClassMaterial.Keyword,
    }}
    _, err = r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *classMaterialRepository) GetByCourseId(ctx context.Context, courseId string) ([]model.ClassMaterial, error) {
    var classMaterials []model.ClassMaterial
    cursor, err := r.collection.Find(ctx, bson.M{"CourseId": courseId})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var classMaterial model.ClassMaterial
        if err := cursor.Decode(&classMaterial); err != nil {
            return nil, err
        }
        classMaterials = append(classMaterials, classMaterial)
    }
    return classMaterials, nil
}

func (r *classMaterialRepository) GetByObjectiveId(ctx context.Context, objectiveId string) ([]model.ClassMaterial, error) {
    var classMaterials []model.ClassMaterial
    cursor, err := r.collection.Find(ctx, bson.M{"ObjectiveId": objectiveId})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var classMaterial model.ClassMaterial
        if err := cursor.Decode(&classMaterial); err != nil {
            return nil, err
        }
        classMaterials = append(classMaterials, classMaterial)
    }
    return classMaterials, nil
}

func (r *classMaterialRepository) GetByMaterialId(ctx context.Context, materialId string) ([]model.ClassMaterial, error) {
    var classMaterials []model.ClassMaterial
    cursor, err := r.collection.Find(ctx, bson.M{"MaterialId": materialId})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var classMaterial model.ClassMaterial
        if err := cursor.Decode(&classMaterial); err != nil {
            return nil, err
        }
        classMaterials = append(classMaterials, classMaterial)
    }
    return classMaterials, nil
}

func (r *classMaterialRepository) FindByKeywordIds(ctx context.Context, keywordIds []string) ([]model.ClassMaterial, error) {
    filter := bson.M{"Keyword": bson.M{"$in": keywordIds}}
    var classMaterials []model.ClassMaterial
    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var classMaterial model.ClassMaterial
        if err := cursor.Decode(&classMaterial); err != nil {
            return nil, err
        }
        classMaterials = append(classMaterials, classMaterial)
    }
    return classMaterials, nil
}
