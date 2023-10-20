package model

import (
	"log"

	"github.com/oxycoder/struct2bson"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	ORDER_COLLECTION_NAME = "orders"
)

//-------------------------------------MONGO CORE FUNCTIONS---------------------------------------------

// ------------------------------------Order related DB Functions------------------------------------

//Create
func InsertOrderToDB(data OrderDB) error {

	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	insertOneResult, err := insertOne(client, ctx, DATABASE_NAME, ORDER_COLLECTION_NAME, insertdata)
	if err != nil {
		return err
	}
	log.Print("Result of InsertOne : ")
	log.Println(insertOneResult.InsertedID)
	return nil
}

//READ
func GetOrdersFromDB() []OrderDB {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	var filter, option interface{}
	filter = bson.D{}
	//option remove id field from all documents
	option = bson.D{{"_id", 0}}

	cursor, err := query(client, ctx, DATABASE_NAME, ORDER_COLLECTION_NAME, filter, option)
	if err != nil {
		panic(err)
	}

	var results []OrderDB
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

func GetOrderbyIDFromDB(Id int) OrderDB {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	filter := bson.D{{"oid", Id}}
	collection := client.Database(DATABASE_NAME).Collection(ORDER_COLLECTION_NAME)

	cursor := collection.FindOne(ctx, filter)

	var result OrderDB
	if err := cursor.Decode(&result); err != nil {
		panic(err)
	}
	return result
}

//UPDATE
func UpdateOrderInDB(data OrderDB) error {
	insertdata := struct2bson.ConvertStructToBSONMap(data, nil)

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	filter := bson.D{
		{"oid", data.Cid},
	}

	update := bson.D{
		{"$set", insertdata},
	}

	result, err := UpdateOne(client, ctx, DATABASE_NAME, ORDER_COLLECTION_NAME, filter, update)
	if err != nil {
		panic(err)
	}

	log.Println("update single document")
	log.Println(result.ModifiedCount)
	return nil
}

//DELETE
func DeleteOrderFromDB(Id int) {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	query := bson.D{{"oid", Id}}
	result, err := deleteOne(client, ctx, DATABASE_NAME, ORDER_COLLECTION_NAME, query)
	if err != nil {
		panic(err)
	}
	log.Print("No.of rows affected by DeleteOne() : ")
	log.Println(result.DeletedCount)
}
