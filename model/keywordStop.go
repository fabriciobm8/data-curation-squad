package model

type KeywordStop struct {	
	Id 				string	`bson:"_id" validate:"required"`
	CourseId		string	`bson:"CourseId"`
	Keyword     	string	`bson:"Keyword"`
	UsageCount  	int		`bson:"UsageCount"`
	Inserted		bool	`bson:"Inserted"`
}
