package repository

import (
    "context"
    "data-curation-squad/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type KeywordRepository interface {
    FindAll(ctx context.Context) ([]model.Keyword, error)
    FindByID(ctx context.Context, id string) (*model.Keyword, error)
    FindByKeyword(ctx context.Context, keyword string) (*model.Keyword, error)
    SaveKeywords(ctx context.Context, keywords []model.Keyword) error
}

type keywordRepository struct {
    collection *mongo.Collection
}

func NewKeywordRepository(client *mongo.Client) KeywordRepository {
    collection := client.Database("class").Collection("keyword")
    return &keywordRepository{collection: collection}
}

func (r *keywordRepository) FindAll(ctx context.Context) ([]model.Keyword, error) {
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var keywords []model.Keyword
    err = cursor.All(ctx, &keywords)
    return keywords, err
}

func (r *keywordRepository) FindByID(ctx context.Context, id string) (*model.Keyword, error) {
    // Converter o ID de string para primitive.ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err // Retorna erro se a conversão falhar
    }

    var keyword model.Keyword
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&keyword)
    if err != nil {
        return nil, err
    }
    return &keyword, nil
}

// Busca uma keyword pelo nome
func (r *keywordRepository) FindByKeyword(ctx context.Context, keyword string) (*model.Keyword, error) {
    var result model.Keyword
    filter := bson.M{"keyword": keyword}
    err := r.collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil // Se não houver documento, retorna nil
        }
        return nil, err
    }
    return &result, nil
}

// Salva a lista de keywords, criando apenas as que não existem e atualizando o UsageCount se já existirem
func (r *keywordRepository) SaveKeywords(ctx context.Context, keywords []model.Keyword) error {
    var err error
    for _, keyword := range keywords {
        existingKeyword, err := r.FindByKeyword(ctx, keyword.Keyword)
        if err != nil {
            continue
        }

        if existingKeyword == nil {
            keyword.ID = primitive.NewObjectID()
            keyword.UsageCount = 1
            _, err = r.collection.InsertOne(ctx, keyword)
        } else {
            _, err = r.collection.UpdateOne(ctx, bson.M{"_id": existingKeyword.ID}, bson.M{"$inc": bson.M{"usageCount": 1}})
        }

        if err != nil {
            break
        }
    }
    return err
}