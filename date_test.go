package date_test

import (
	"fmt"
	"testing"

	date "github.com/githomework/apps-util-date"
)

func TestDate(t *testing.T) {
	d, _ := date.PreviousWorkDay("3000")
	fmt.Println("3000", d)
	d, _ = date.NWorkDaysAgo("3000", 1)
	fmt.Println("3000", d)
	d, offset := date.NWorkDaysAgo("3000", 60)
	fmt.Println("3000", d, offset)
	fmt.Println("6000 hour to 2000 7am", date.HourLocalTo2000("6000", 7))
	fmt.Println("6000 hour to 2000 11pm", date.HourLocalTo2000("6000", 23))
}

func TestLocalTime(t *testing.T) {
	fmt.Println("Local time 3000", date.LocalTime("3000").Format("15:04"))
	fmt.Println("Local time 4000", date.LocalTime("4000").Format("15:04"))
	fmt.Println("Local time 6000", date.LocalTime("6000").Format("15:04"))
}

func TestAdhoc(t *testing.T) {
	d, offset := date.PreviousWorkDay("2000")
	fmt.Println(d, offset)
}
