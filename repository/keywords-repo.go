package repository

import (
    "context"
    "data-curation-squad/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type KeywordRepository interface {
    Create(ctx context.Context, keyword *model.Keyword) error
    FindAll(ctx context.Context) ([]model.Keyword, error)
    FindByID(ctx context.Context, id string) (*model.Keyword, error)    
}

type keywordRepository struct {
    collection *mongo.Collection
}

func NewKeywordRepository(client *mongo.Client) KeywordRepository {
    collection := client.Database("class").Collection("keyword")
    return &keywordRepository{collection: collection}
}

func (r *keywordRepository) Create(ctx context.Context, keyword *model.Keyword) error {
    _, err := r.collection.InsertOne(ctx, keyword)
    return err
}

func (r *keywordRepository) FindAll(ctx context.Context) ([]model.Keyword, error) {
    var keywords []model.Keyword
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var keyword model.Keyword
        if err := cursor.Decode(&keyword); err != nil {
            return nil, err
        }
        keywords = append(keywords, keyword)
    }
    return keywords, nil
}

func (r *keywordRepository) FindByID(ctx context.Context, id string) (*model.Keyword, error) {
    var keyword model.Keyword
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&keyword)
    if err != nil {
        return nil, err
    }
    return &keyword, nil
}