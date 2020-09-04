package main

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"bufio"
	// manage multiple requests
	"os" // os.Exit(1) on Error
	// get an object type
)

var zipFileName = time.Now().Format("20060102150405")
var db *mongo.Database

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	db = mongoClient.Database("StockDB")
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	date()
}

func readFile() {
	read, err := zip.OpenReader(zipFileName)
	defer os.Remove(zipFileName)
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for _, file := range read.File {
		listFiles(file)
	}
}

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	lines, err := ioutil.ReadAll(fileread)
	r := csv.NewReader(bytes.NewBuffer(lines))
	save(r)
	return nil
}

func getURL(date time.Time) string {
	month := date.Month().String()
	shortMonth := strings.ToUpper(month[0:3])
	year := strconv.Itoa((date.Year()))
	day := strconv.Itoa(date.Day())
	if date.Day() < 10 {
		day = "0" + day
	}

	fileName := "cm" + day + shortMonth + year + "bhav.csv.zip"
	specURL := "https://www1.nseindia.com/content/historical/EQUITIES/" + year + "/" +
		shortMonth + "/" + fileName
	return specURL
}

func startDownload(specURL string) {
	resp, err := http.Get(specURL)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create(zipFileName)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)
	readFile()
}
