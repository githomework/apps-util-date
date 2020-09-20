package date

import (
	"strings"
	"time"
)

var (
	HolidayMap map[string]map[time.Time]bool
)

const (
	vicHolidays = `2019-01-01
2019-01-28
2019-03-11
2019-04-19
2019-04-20
2019-04-21
2019-04-22
2019-04-25
2019-06-10
2019-09-27
2019-11-05
2019-12-25
2019-12-26
2020-01-01
2020-01-27
2020-03-09
2020-04-10
2020-04-13
2020-04-25
2020-06-08
2020-11-03
2020-12-25
2020-12-28`

	nswHolidays = `2019-01-01
2019-01-28
2019-04-19
2019-04-20
2019-04-21
2019-04-22
2019-04-25
2019-06-10
2019-10-07
2019-12-25
2019-12-26
2020-01-01
2020-01-27
2020-04-10
2020-04-13
2020-04-25
2020-06-08
2020-10-05
2020-12-25
2020-12-28`

	qldHolidays = `2019-01-01
2019-01-28
2019-04-19
2019-04-20
2019-04-21
2019-04-22
2019-04-25
2019-05-06
2019-08-14
2019-10-07
2019-12-24
2019-12-25
2019-12-26
2020-01-01
2020-01-27
2020-04-10
2020-04-11
2020-04-12
2020-04-13
2020-04-25
2020-05-04
2020-08-12
2020-10-05
2020-12-25
2020-12-26
2020-12-28
2021-01-01
2021-01-26
2021-04-02
2021-04-03
2021-04-04
2021-04-05
2021-04-26
2021-05-03
2021-08-11
2021-10-04
2021-12-24
2021-12-25
2021-12-26
2021-12-27
2021-12-28`
)

func init() {
	HolidayMap = map[string]map[time.Time]bool{}

	HolidayMap["2000"] = holidays("2000")
	HolidayMap["3000"] = holidays("3000")
	HolidayMap["4000"] = holidays("4000")

}
func holidays(plant string) map[time.Time]bool {
	days := ""
	m := map[time.Time]bool{}
	switch plant {
	case "2000":
		days = nswHolidays
	case "3000":
		days = vicHolidays
	case "4000":
		days = qldHolidays
	}

	for _, v := range strings.Split(days, "\n") {
		d, _ := time.Parse("2006-01-02", v)
		m[d] = true
	}
	//log.Println(m)
	return m
}


func PreviousWorkDay(plant string) time.Time {
	var loc *time.Location
	switch plant {
	case "2000","3000":
		loc ,_ = time.LoadLocation("Australia/Sydney")
	case "4000":
		loc, _ = time.LoadLocation("Australia/Brisbane")
	}
	d,_ := time.Parse("2006-01-02", time.Now().Add(-24*time.Hour).In(loc).Format("2006-01-02"))

	for d.Weekday() == time.Saturday || d.Weekday() == time.Sunday || HolidayMap[plant][d] {
		d = d.Add(-24*time.Hour)
	}
	return d
}
