package date

import (
	"strings"
	"time"
)

var (
	HolidayMap    map[string]map[time.Time]bool
	plantLocation map[string]*time.Location
	LocalHourDiff map[string]int
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

	waHolidays = `2020-01-01
2020-01-27
2020-03-02
2020-04-10
2020-04-13
2020-04-25
2020-04-27
2020-06-01
2020-09-28
2020-12-25
2020-12-26
2020-12-28
2021-01-01
2021-01-26
2021-03-01
2021-04-02
2021-04-05
2021-04-25
2021-04-26
2021-06-07
2021-09-27
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

	plantLocation = map[string]*time.Location{}
	plantLocation["2000"], _ = time.LoadLocation("Australia/Sydney")
	plantLocation["3000"], _ = time.LoadLocation("Australia/Melbourne")
	plantLocation["4000"], _ = time.LoadLocation("Australia/Brisbane")
	plantLocation["6000"], _ = time.LoadLocation("Australia/Perth")

	t := time.Now()
	t2 := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, plantLocation["2000"])
	LocalHourDiff = map[string]int{}
	LocalHourDiff["2000"] = 0
	LocalHourDiff["3000"] = 0
	LocalHourDiff["4000"] = t2.Hour() - t2.In(plantLocation["4000"]).Hour()
	LocalHourDiff["6000"] = t2.Hour() - t2.In(plantLocation["6000"]).Hour()

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
	case "6000":
		days = waHolidays
	}

	for _, v := range strings.Split(days, "\n") {
		d, _ := time.Parse("2006-01-02", v)
		m[d] = true
	}
	//log.Println(m)
	return m
}

func PreviousWorkDay(plant string) (time.Time, int) {
	return NWorkDaysAgo(plant, 1)
}

func NWorkDaysAgo(plant string, n int) (time.Time, int) {
	var offset int
	var workdays int
	loc := plantLocation[plant]
	offset = -1
	d, _ := time.Parse("2006-01-02", time.Now().Add(-24*time.Hour).In(loc).Format("2006-01-02"))
	for workdays < n {
		for d.Weekday() == time.Saturday || d.Weekday() == time.Sunday || HolidayMap[plant][d] {
			d = d.Add(-24 * time.Hour)
			offset--
		}
		workdays++
		if workdays == n {
			break
		}
		d = d.Add(-24 * time.Hour)
		offset--
	}
	return d, offset
}

// Need to ADD result to get to the 2000 time.  Need to be careful with not crossing date boundaries.
func HourLocalTo2000(plant string, localHour int) int {
	return (localHour + LocalHourDiff[plant]) % 24
}
