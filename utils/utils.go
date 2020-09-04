package utils

import (
	"strconv"
	"strings"
	"time"
)

func getURL(date time.Time) string {
	month := date.Month().String()
	shortMonth := strings.ToUpper(month[0:3])
	year := strconv.Itoa((date.Year()))
	day := strconv.Itoa(date.Day())
	if date.Day() < 10 {
		day = "0" + day
	}

	fileName := "cm" + day + shortMonth + year + "bhav.csv.zip"
	specURL := common.baseURL + year + "/" +
		shortMonth + "/" + fileName
	return specURL
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func date(start time.Time, end time.Time) []time.Time {
	dates := []time.Time{}
	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		if !contains(holidayList, date.Format("2006-Jan-02")) {
			weekday := date.Weekday()
			if int(weekday) == 6 || int(weekday) == 0 {
				continue
			}
			dates = append(dates, date)
			// fmt.Println(date.Format("2006-Jan-02"))
			// url := getURL(date)
			// startDownload(url)
		}
	}
	return dates
}

// rangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
