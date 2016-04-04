package main

import (
	"time"
)

type Article struct {
	Id int
	Name string
	Link string
	Date time.Time
	dateStr string

	TopicRels []Article_Topic_Rel
	Topics []Topic
}
func (f Article) GetDate() time.Time {
    return f.Date
}
func (f Article) GetDateStr() string {
    return f.dateStr
}
func (f *Article) SetId(id int) {
    f.Id = id
}
func (f *Article) SetName(name string) {
    f.Name = name
}
func (f *Article) SetLink(link string) {
    f.Link = link
}
func (f *Article) SetDateStr(dateStr string) {
    f.dateStr = dateStr
}
func (f *Article) SetDate(date time.Time) {
    f.Date = date
}

type Topic struct {
	Id int
	Name string
	specificity int

	ArticleRels []Article_Topic_Rel
	Articles []Article
}
func (f *Topic) SetId(id int) {
    f.Id = id
}
func (f *Topic) SetName(name string) {
    f.Name = name
}

type Article_Topic_Rel struct {
	Id int
	Article_ID int
	Topic_ID int
	Strength int //	Strength on a 1-5 scale. The more the article is related to topic, the higher the strength
	Description string
}
func (f *Article_Topic_Rel) SetDescription(desc string) {
    f.Description = desc
}
func (f *Article_Topic_Rel) SetId(id int) {
    f.Id = id
}
func (f *Article_Topic_Rel) SetArticleId(id int) {
    f.Article_ID = id
}
func (f *Article_Topic_Rel) SetTopicId(id int) {
    f.Topic_ID = id
}
func (f *Article_Topic_Rel) SetStrength(id int) {
    f.Strength = id
}

// SetName receives a pointer to Foo so it can modify it.




