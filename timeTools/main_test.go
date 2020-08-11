package main

import "testing"

func TestStr2TS(t *testing.T) {

	ts := str2TS("today")
	if ts == 0 {
		t.Error("today parse error")
		t.FailNow()
	}

	ts = str2TS("now")
	if ts == 0 {
		t.Error("now parse error")
		t.FailNow()
	}

	ts = str2TS("tomorrow")
	if ts == 0 {
		t.Error("tomorrow parse error")
		t.FailNow()
	}

	ts = str2TS("2020-01-01")
	if ts == 0 {
		t.Error("2020-01-01 parse error")
		t.FailNow()
	}

	ts = str2TS("2020-01-01 00:00:01")
	if ts == 0 {
		t.Error("2020-01-01 00:00:01 parse error")
		t.FailNow()
	}
}

func TestGenerateResponse(t *testing.T) {

}
