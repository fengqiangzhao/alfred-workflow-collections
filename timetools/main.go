package main

import (
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const DatetimeFormatString = "Mon Jan 2006-01-02 15:04:05"

var keywords = map[string]int{
	"YESTERDAY": -1,
	"YSD":       -1,
	"NOW":       0,
	"TODAY":     1,
	"TOMORROW":  2,
	"TMW":       2,
}

type AlfredItem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     struct {
		Path string `json:"path"`
	} `json:"icon"`
}

func keywords2TS(value int) int64 {
	t := time.Now()
	ts := t.Unix()
	_, offset := t.Local().Zone()
	switch value {
	case -1:
		return ts - ts%86400 - int64(offset) - 86400
	case 0:
		return ts
	case 1:
		return ts - ts%86400 - int64(offset)
	case 2:
		return ts - ts%86400 - int64(offset) + 86400
	default:
		return 0
	}
}

func parse(timeStr string) []string {
	var ts int64
	var dt, tsStr string

	if value, ok := keywords[strings.ToUpper(timeStr)]; ok {
		ts = keywords2TS(value)
	} else if t, err := dateparse.ParseLocal(timeStr); err == nil {
		ts = t.Unix()
	} else {
		ts = 0
	}

	if ts == 0 {
		tsStr = "数据格式有误"
		dt = "0"
		return []string{tsStr, dt}
	} else {
		dt = time.Unix(ts, 0).Format(DatetimeFormatString)
		tsStr = strconv.FormatInt(ts, 10)
	}

	if ok, _ := regexp.Match(`^\d{10,19}$`, []byte(timeStr)); ok {
		return []string{dt, tsStr}
	}
	return []string{tsStr, dt}
}

func generateAlfredItems(dtSlice []string) {
	r := make([]AlfredItem, 0)

	r = append(r, AlfredItem{
		Type:     "file",
		Title:    dtSlice[0],
		Subtitle: dtSlice[1],
		Arg:      dtSlice[0],
		Icon: struct {
			Path string `json:"path"`
		}{Path: "icon.png"},
	})

	finalRes, _ := json.Marshal(struct {
		Items []AlfredItem `json:"items"`
	}{
		Items: r,
	})
	fmt.Println(string(finalRes))
}

func main() {
	arg := os.Args[1]
	generateAlfredItems(parse(arg))
}
