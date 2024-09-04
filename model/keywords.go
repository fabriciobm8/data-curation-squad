package model

type Keyword struct {
	ClassMaterialId	string	`bson:"_id" validate:"required"`
	CourseId		string	`bson:"CourseId"`
	Keyword     	string	`bson:"Keyword"`
	UsageCount  	int		`bson:"UsageCount"`
}

