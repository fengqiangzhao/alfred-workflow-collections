package main

import (
	"fmt"
	"testing"
)

func TestStr2TS(t *testing.T) {
	test_arr := []string{
		"now",
		"today",
		"ysd",
		"yesterday",
		"tmw",
		"tomorrow",
		"2020-01-01",
		"2020-01-01 00:00:01",
	}
	for _, item := range test_arr {
		if ts := str2TS(item); ts == 0 {
			t.Error(fmt.Sprintf("%s parse error", item))
			t.FailNow()
		}
	}
}
