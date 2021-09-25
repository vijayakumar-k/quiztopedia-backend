package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	QUESTIONS = "questions"
)

//Question Model
var questionCollection *mongo.Collection

func QuestionCollection(c *mongo.Database) {
	questionCollection = c.Collection(QUESTIONS)
}

type Question struct {
	ID               primitive.ObjectID `bson:"_id"`
	Text             string             `bson:"text"`
	Type             string             `bson:"type"`
	Area             string             `bson:"area"`
	Weightage        int                `bson:"weightage"`
	Tags             []string           `bson:"tags"`
	SuggestedAnswers []string           `bson:"suggested_answers"`
	AcceptedAnswers  []string           `bson:"accepted_answers"`
	AuthorId         string             `bson:"author_id"`
	CreatedDate      time.Time          `bson:"created_date"`
	ModifiedDate     time.Time          `bson:"modified_date"`
}

//Answer Model

//Author Model

func Initialize(c *mongo.Database) {
	QuestionCollection(c)
}
