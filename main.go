package main

import (
	"fmt"
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

func ObtainSampleArticles () ([]Article, error) {

	art1 := Article{
		name:"Women in Science: Neurology professor inspires sophomore to pursue her dreams",
		link: "http://www.wildcat.arizona.edu/article/2016/03/neurology-professor-inspires-sophomore-to-pursue-her-dreams",
		dateStr: "03/30/2016"}

	art2 := Article{
		name: "MU Neurology's art auction part of comprehensive approach",
		link: "http://www.herald-dispatch.com/news/mu-neurology-s-art-auction-part-of-comprehensive-approach/article_7ced5b49-ba53-56ac-a8a6-32a26f0fbe20.html",
		dateStr: "03/31/2016"}

	articles := make([]Article, 2)
	articles[0] = art1
	articles[1] = art2

	for i := range articles{
		dateStr := articles[i].GetDateStr()
		fmt.Println("Current date formatted:",dateStr)

		t, _ := time.Parse(dateStr)
		if t != nil {
			articles[i].SetDate(t)
		}
	}

	return articles, nil
}

func main() {

	articles, err := ObtainSampleArticles()
	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}
	fmt.Println("Obtained sample articles: ")
	for i := range articles{
		fmt.Println("Article",articles[i])
	}

}