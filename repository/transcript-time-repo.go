package repository

import (
    "context"
    "data-curation-squad/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type TranscriptTimeRepository interface {
    Create(ctx context.Context, transcriptTime *model.TranscriptTime) error
    FindAll(ctx context.Context) ([]model.TranscriptTime, error)
    FindByID(ctx context.Context, id string) (*model.TranscriptTime, error)
    
}

type transcriptTimeRepository struct {
    collection *mongo.Collection
}

func NewTranscriptTimeRepository(client *mongo.Client) TranscriptTimeRepository {
    collection := client.Database("class").Collection("transcriptTime")
    return &transcriptTimeRepository{collection: collection}
}

func (r *transcriptTimeRepository) Create(ctx context.Context, transcriptTime *model.TranscriptTime) error {
    _, err := r.collection.InsertOne(ctx, transcriptTime)
    return err
}

func (r *transcriptTimeRepository) FindAll(ctx context.Context) ([]model.TranscriptTime, error) {
    var transcriptTimes []model.TranscriptTime
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var transcriptTime model.TranscriptTime
        if err := cursor.Decode(&transcriptTime); err != nil {
            return nil, err
        }
        transcriptTimes = append(transcriptTimes, transcriptTime)
    }
    return transcriptTimes, nil
}

func (r *transcriptTimeRepository) FindByID(ctx context.Context, id string) (*model.TranscriptTime, error) {
    var transcriptTime model.TranscriptTime
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transcriptTime)
    if err != nil {
        return nil, err
    }
    return &transcriptTime, nil
}
