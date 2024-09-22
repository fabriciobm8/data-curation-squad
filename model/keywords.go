package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Keyword struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    Keyword    string             `bson:"keyword"`
    UsageCount int                `bson:"usageCount"`
    Inserted   bool               `bson:"inserted" default:"false"` 
}
