package date_test

import (
	"fmt"
	"testing"

	date "github.com/githomework/apps-util-date"
)

func TestDate(t *testing.T) {
	fmt.Println("2000", date.PreviousWorkDay("2000"))
	fmt.Println("3000", date.PreviousWorkDay("3000"))
	fmt.Println("4000", date.PreviousWorkDay("4000"))
}
