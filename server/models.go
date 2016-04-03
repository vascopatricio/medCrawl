package main

import (
	"time"
)

type Article struct {
	Name string
	Link string
	date time.Time
	dateStr string
}

type Topic struct {
	Name string
	specificity int
}

type Article_Topic_Rel struct {
	article Article
	topic Topic
	description string
}

// SetName receives a pointer to Foo so it can modify it.
func (f *Article) SetName(name string) {
    f.Name = name
}
func (f *Topic) SetName(name string) {
    f.Name = name
}
func (f *Article_Topic_Rel) SetDescription(description string) {
    f.description = description
}
func (f *Article) SetLink(link string) {
    f.Link = link
}
func (f *Article) SetDateStr(dateStr string) {
    f.dateStr = dateStr
}
func (f *Article) SetDate(date time.Time) {
    f.date = date
}

func (f Article) GetDate() time.Time {
    return f.date
}
func (f Article) GetDateStr() string {
    return f.dateStr
}