package main

import (
	"concurrency/common"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var zipFileName string

func main() {
	//date := time.Now()
	date := time.Date(2019, 8, 30, 0, 0, 0, 0, time.UTC)
	start := date.AddDate(0, 0, 0)
	end := date.AddDate(0, 0, 0)
	fmt.Printf("Start Date : %s and End Date : %s\n",
		start.Format("2006-Jan-02"), end.Format("2006-Jan-02"))
	dates := getDates(start, end)

	for _, dt := range dates {
		url := getNSEURL(dt)
		fmt.Printf("Got URL: %s \n", url)
		zipFileName = date.Format("20060102150405")
		startDownload(url)
		fmt.Println("zip file downloaded")
		z := readZipFile()
		files := z.File
		for _, file := range files {
			fmt.Printf("found file %s", file.Name)
			r, e := readCSVFile(file)
			if e == nil {
				fmt.Println("Saving to db")
				save(r)
			}
		}
		defer z.Close()
	}
	defer os.Remove(zipFileName)
}

func startDownload(specURL string) {
	resp, err := http.Get(specURL)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		fmt.Printf("err: %d", resp.StatusCode)
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
}

//getNSEURL construct the url, from where the zip file will be downloaded
func getNSEURL(date time.Time) string {
	month := date.Month().String()
	shortMonth := strings.ToUpper(month[0:3])
	year := strconv.Itoa((date.Year()))
	day := strconv.Itoa(date.Day())
	if date.Day() < 10 {
		day = "0" + day
	}

	fileName := "cm" + day + shortMonth + year + "bhav.csv.zip"
	specURL := common.BaseURL + year + "/" +
		shortMonth + "/" + fileName
	return specURL
}
