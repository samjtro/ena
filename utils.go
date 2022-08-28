package main

import (
	"fmt"
	"time"
)

func Now(t time.Time) string {
	str := fmt.Sprintf("%d-%d-%dT%d:%d:%d%s",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		utcDiffFlag)

	return str
}
