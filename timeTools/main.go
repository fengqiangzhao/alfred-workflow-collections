package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AlfredItem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     struct {
		Path string `json:"path"`
	} `json:"icon"`
}

var Keywords = map[string]int{
	"NOW":       1,
	"YESTERDAY": 2,
	"YSD":       2,
	"TODAY":     3,
	"TOMORROW":  4,
	"TMW":       4,
}

func str2TS(arg string) int64 {
	if value, ok := Keywords[strings.ToUpper(arg)]; ok {
		t := time.Now()
		ts := t.Unix()
		switch value {
		case 1:
			return ts
		case 2:
			_, offset := t.Local().Zone()
			return ts - ts%86400 - int64(offset) - 86400
		case 3:
			_, offset := t.Local().Zone()
			return ts - ts%86400 - int64(offset)
		case 4:
			_, offset := t.Local().Zone()
			return ts - ts%86400 - int64(offset) + 86400
		default:
			return 0
		}
	} else if ok, _ := regexp.Match(`\d{4}-\d{2}-\d{2}$`, []byte(arg)); ok {
		loc, _ := time.LoadLocation("Local")
		ts, _ := time.ParseInLocation("2006-01-02", arg, loc)
		return ts.Unix()
	} else if ok, _ := regexp.Match(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`, []byte(arg)); ok {
		loc, _ := time.LoadLocation("Local")
		ts, _ := time.ParseInLocation("2006-01-02 15:04:05", arg, loc)
		return ts.Unix()
	} else {
		return 0
	}
}

func generateResponse(ts int64) {
	r := make([]AlfredItem, 0)
	r = append(r, AlfredItem{
		Type:     "file",
		Title:    strconv.FormatInt(ts, 10),
		Subtitle: "时间戳",
		Arg:      fmt.Sprintf("%d", ts),
		Icon:     struct{ Path string `json:"path"` }{Path: "icon.png"},
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
	ts := str2TS(arg)
	generateResponse(ts)
}
