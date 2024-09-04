package model
type ClassMaterial struct{
	Id 				string				`bson:"_id" validate:"required"`
	CourseId 		string 				`bson:"CourseId" validate:"required"`
	ObjectiveId 	string				`bson:"ObjectiveId" validate:"required"`
	MaterialId 		string 				`bson:"MaterialId" validate:"required"`
	Transcript 		string 				`bson:"Transcript" validate:"required"`
	MaterialType 	string				`bson:"MaterialType" validate:"required"`
	IsSuccessful	bool				`bson:"IsSuccessful" default:"false"`
	TranscriptTime 	[]TranscriptTime	`bson:"TranscriptTime"`
	Keyword 		[]Keyword 			`bson:"Keyword"`
}