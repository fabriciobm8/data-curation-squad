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
