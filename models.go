package main

import (
	"time"
)

type Article struct {
	name string
	link string
	date time.Time
	dateStr string
}

// SetName receives a pointer to Foo so it can modify it.
func (f *Article) SetDate(date time.Time) {
    f.date = date
}

func (f Article) GetDate() time.Time {
    return f.date
}
func (f Article) GetDateStr() string {
    return f.dateStr
}