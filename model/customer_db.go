package model

import (
	"context"
	"log"
	"time"

	"github.com/oxycoder/struct2bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE_NAME   = "prod"
	COLLECTION_NAME = "customers"
)

//-------------------------------------MONGO CORE FUNCTIONS---------------------------------------------

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns a mongo.Client, context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database
// operation. context.Context will be used set
// deadlines for process. context.CancelFunc will
// be used to cancel context and resource
// associated with it.
func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context,
	dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return result, err
}

// ------------------------------------Customer related DB Functions------------------------------------

//Create
func InsertCustomerToDB(data CustomerDB) error {

	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	insertOneResult, err := insertOne(client, ctx, DATABASE_NAME, COLLECTION_NAME, insertdata)
	if err != nil {
		return err
	}
	log.Print("Result of InsertOne : ")
	log.Println(insertOneResult.InsertedID)
	return nil
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

//READ
func GetCustomersFromDB() []CustomerDB {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	var filter, option interface{}
	filter = bson.D{}
	//option remove id field from all documents
	option = bson.D{{"_id", 0}}

	cursor, err := query(client, ctx, DATABASE_NAME, COLLECTION_NAME, filter, option)
	if err != nil {
		panic(err)
	}

	var results []CustomerDB
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	// printing the result of query.
	// fmt.Println("Query Result");
	// for _, doc := range results {
	// 	fmt.Println(doc)
	// }
	return results
}

func GetCustomerbyIDFromDB(Id int) CustomerDB {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	filter := bson.D{{"cid", Id}}
	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)

	cursor := collection.FindOne(ctx, filter)

	var result CustomerDB
	if err := cursor.Decode(&result); err != nil {
		panic(err)
	}
	return result
}

//UPDATE
func UpdateCustomerInDB(data CustomerDB) error {
	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	filter := bson.D{
		{"cid", data.Cid},
	}

	update := bson.D{
		{"$set", insertdata},
	}

	result, err := UpdateOne(client, ctx, DATABASE_NAME, COLLECTION_NAME, filter, update)
	if err != nil {
		panic(err)
	}

	log.Println("update single document")
	log.Println(result.ModifiedCount)
	return nil
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase,
	col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateOne(ctx, filter, update)
	return result, err
}

//DELETE
func DeleteCustomerFromDB(Id int) {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	query := bson.D{{"cid", Id}}
	result, err := deleteOne(client, ctx, DATABASE_NAME, COLLECTION_NAME, query)
	if err != nil {
		panic(err)
	}
	log.Print("No.of rows affected by DeleteOne() : ")
	log.Println(result.DeletedCount)
}

func deleteOne(client *mongo.Client, ctx context.Context,
	dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}
