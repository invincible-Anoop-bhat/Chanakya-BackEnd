package api

import (
	"context"
	"log"
	"time"

	"Chanakya-BackEnd/model"

	"github.com/oxycoder/struct2bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE_NAME   = "prod"
	COLLECTION_NAME = "customers"
)

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

// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.

//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limits
// the field being returned.

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context,
	dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return result, err
}

//UPDATE
func updateCustomerInDB(data model.CustomerDB) error {
	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)
	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)
	// filter object is used to select a single
	// document matching that matches.
	filter := bson.D{
		{"cid", bson.D{{"$eq", data.Cid}}},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{"$set", insertdata},
	}

	// Returns result of updated document and a error.
	result, err := UpdateOne(client, ctx, DATABASE_NAME,
		COLLECTION_NAME, filter, update)

	// handle error
	if err != nil {
		panic(err)
	}

	// print count of documents that affected
	log.Println("update single document")
	log.Println(result.ModifiedCount)
	return nil
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase,
	col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the database and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

//INSERT
func insertCustomerToDB(data model.CustomerDB) error {
	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)
	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
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

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func GetCustomersFromDB() []model.CustomerDB {

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)
	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	filter = bson.D{}

	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns momngo.cursor and error if any.
	cursor, err := query(client, ctx, DATABASE_NAME,
		COLLECTION_NAME, filter, option)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	var results []model.CustomerDB

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &results); err != nil {
		// handle the error
		panic(err)
	}

	// printing the result of query.
	// fmt.Println("Query Result");
	// for _, doc := range results {
	// 	fmt.Println(doc)
	// }
	return results
}
func GetCustomerbyIDFromDB(Id int) model.CustomerDB {
	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)

	filter := bson.D{{"cid", Id}}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	cursor := collection.FindOne(ctx, filter)

	var result model.CustomerDB
	if err := cursor.Decode(&result); err != nil {
		panic(err)
	}
	return result
}

//DELETE
func DeleteCustomerFromDB(Id int) {

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)
	// This query delete document when the maths
	// field is greater than  60
	query := bson.D{{"cid", Id}}

	// Returns result of deletion and error
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
