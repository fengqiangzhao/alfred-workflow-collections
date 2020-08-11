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

func str2TS(arg string) int64 {
	if strings.ToUpper(arg) == "NOW" {
		return time.Now().Unix()
	} else if strings.ToUpper(arg) == "TODAY" {
		t := time.Now()
		ts := t.Unix()
		_, offset := t.Local().Zone()
		fmt.Println(t, ts, offset)
		return ts - ts%86400 - int64(offset)
	} else if strings.ToUpper(arg) == "TOMORROW" {
		t := time.Now()
		ts := t.Unix()
		_, offset := t.Local().Zone()
		fmt.Println(t, ts, offset)
		return ts - ts%86400 - int64(offset) + 86400
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
