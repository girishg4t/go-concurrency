package main

import (
	"concurrency/common"
	"concurrency/utils"
	"time"
)

//getDateRange between start and end date
func getDateRange(start time.Time, end time.Time) []time.Time {
	dates := []time.Time{}
	for rd := utils.RangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		if !utils.Contains(common.HolidayList, date.Format("2006-Jan-02")) {
			weekday := date.Weekday()
			//exclude saturday(6) and sunday(0)
			if int(weekday) == 6 || int(weekday) == 0 {
				continue
			}
			dates = append(dates, date)
		}
	}
	return dates
}
