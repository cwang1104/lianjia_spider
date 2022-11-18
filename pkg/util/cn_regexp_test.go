package util

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetNumberStr(t *testing.T) {
	a := "16,334元/平"

	s := GetNumberStr(a)

	sn, err := strconv.Atoi(s)
	if err != nil {
		t.Fatalf("err")
	}
	fmt.Println(sn)

}
