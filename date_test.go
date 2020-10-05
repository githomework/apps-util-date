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
	fmt.Println("6000 hour to 2000", date.HourLocalTo2000("6000", 7))
}
