package main

import (
	"fmt"
	"time"
)

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

		t, e := time.Parse("01/02/2006", dateStr)
		if e != nil {
			fmt.Println("Error obtaining date from ",dateStr)
		}else {
			articles[i].SetDate(t)
		}
		fmt.Println("Date parsed:",articles[i].GetDate())
	}

	return articles, nil
}