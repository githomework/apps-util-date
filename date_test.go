package date_test

import (
	"fmt"
	"testing"

	date "github.com/githomework/apps-util-date"
)

func TestDate(t *testing.T) {
	fmt.Println(date.PreviousWorkDay("3000"))
}
