package dataaccess

import (
	"context"
	"quiztopedia-backend/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Question struct {
	ID           string    `bson:"question_id"`
	CreatedData  time.Time `bson:"created_date"`
	QuestionText string    `bson:"question_text"`
	Type         string    `bson:"type"`
	Context      string    `bson:"context"`
}

func GetQuestionsByType(qype string) (Question, error) {
	result := Question{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "type", Value: qype}}
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func GetAllQuestions() ([]Question, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	var questions []Question
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return questions, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return questions, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		var t Question
		err := cur.Decode(&t)
		if err != nil {
			return questions, err
		}
		questions = append(questions, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(questions) == 0 {
		return questions, mongo.ErrNoDocuments
	}
	return questions, nil
}

func CreateIssue(qn Question) error {
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), qn)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func CreateMany(list []Question) error {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform InsertMany operation & validate against the error.
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func MarkCompleted(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)

	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func DeleteOne(code string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

//DeleteAll - Get All issues for collection
func DeleteAll() error {
	//Define filter query for fetching specific document from collection
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	//Get MongoDB connection using connectionhelper.
	client, err := helpers.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helpers.DB).Collection(helpers.QUESTIONS)
	//Perform DeleteMany operation & validate against the error.
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
