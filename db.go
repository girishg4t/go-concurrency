package main

import (
	//"bufio"
	"context" // manage multiple requests
	"encoding/csv"
	"fmt" // Println() function
	"io"
	"log" // os.Exit(1) on Error
	"os"
	"sync"

	// get an object type

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func setDbConnection() {
	if db == nil {
		fmt.Println("Called one's")
		clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

		mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
		db = mongoClient.Database("StockDB")
		if err != nil {
			fmt.Println("mongo.Connect() ERROR:", err)
			os.Exit(1)
		}
	}
}

func save(r *csv.Reader) {
	var wg sync.WaitGroup

	//ignore header
	stockQuotes, err := r.Read()
	setDbConnection()
	// Iterate through the records
	for {
		// Read each record from csv
		stockQuotes, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//take only equity
		if stockQuotes[1] != "EQ" {
			continue
		}
		sp := getObjToSave(stockQuotes)
		wg.Add(1)
		go insertIntoDb(stockQuotes[0]+".BO", sp, &wg)
	}
	wg.Wait()
}

func insertIntoDb(c string, s primitive.D, wg *sync.WaitGroup) {
	col := db.Collection(c)
	//ctx, _ := context..WithTimeout(context.Background(), 1000*time.Second)
	ctx := context.TODO()
	_, insertErr := col.InsertOne(ctx, s)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
	} else {
		fmt.Println("Done saving: ", c)
	}
	wg.Done()
}
