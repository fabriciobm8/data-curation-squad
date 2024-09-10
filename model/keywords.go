package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Keyword struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Keyword    string             `bson:"keyword" json:"keyword"`
    UsageCount int                `bson:"usageCount" json:"usageCount"`
}
