package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuestionById(id string) (Question, error) {
	result := Question{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	err = questionCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllQuestions() ([]Question, error) {
	filter := bson.D{{}}
	var questions []Question
	cur, findError := questionCollection.Find(context.Background(), filter)
	if findError != nil {
		return questions, findError
	}
	for cur.Next(context.Background()) {
		var t Question
		err := cur.Decode(&t)
		if err != nil {
			return questions, err
		}
		questions = append(questions, t)
	}
	cur.Close(context.Background())
	if len(questions) == 0 {
		return questions, mongo.ErrNoDocuments
	}
	return questions, nil
}

func CreateQuestion(qn *Question) error {
	qn.ID = primitive.NewObjectID()
	qn.CreatedDate = time.Now()
	qn.ModifiedDate = time.Now()
	_, err := questionCollection.InsertOne(context.Background(), qn)
	if err != nil {
		return err
	}
	return nil
}
