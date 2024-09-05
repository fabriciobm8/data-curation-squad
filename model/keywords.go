package model

type Keyword struct {
	keywordId		string	`bson:"_id" validate:"required"`
	Keyword     	string	`bson:"Keyword"`
	UsageCount  	int		`bson:"UsageCount"`
}

