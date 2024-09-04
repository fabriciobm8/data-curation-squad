package model

type TranscriptTime struct {
	ClassMaterialId		string		`bson:"Id"`
	StartTime			float64		`bson:"StartTime"`
	EndTime				float64		`bson:"EndTime"`
	Transcript			string		`bson:"Transcript"`
	Keyword				[]Keyword	`bson:"Keyword"`
}