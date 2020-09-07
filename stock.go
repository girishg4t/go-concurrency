package main

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getObjToSave(stockQuotes []string) primitive.D {
	layOut := "2-Jan-2006"

	date, _ := time.Parse(layOut, stockQuotes[10])
	open, _ := strconv.ParseFloat(stockQuotes[2], 64)
	high, _ := strconv.ParseFloat(stockQuotes[3], 64)
	low, _ := strconv.ParseFloat(stockQuotes[4], 64)
	close, _ := strconv.ParseFloat(stockQuotes[5], 64)
	volume, _ := strconv.ParseInt(stockQuotes[9], 20, 64)
	stockName := stockQuotes[0] + ".BO"
	sp := bson.D{
		{"date", date},
		{"open", open},
		{"high", high},
		{"low", low},
		{"close", close},
		{"volume", volume},
		{"adjClose", nil},
		{"symbol", stockName},
	}
	return sp
}

//stock to store in db
type todayStockBhav struct {
	Date     time.Time
	open     float64
	high     float64
	low      float64
	close    float64
	volume   int64
	adjClose *float64
	symbol   string
}
