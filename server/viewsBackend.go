package main

import (
"fmt"
"net/http"
)

func topicsResource(w http.ResponseWriter, r *http.Request) {

	topics := make([]Topic, 10)
	topic1 := Topic{Name:"Neurology"}
	topicsSlice := topics[0:0]
	topicsSlice = append(topicsSlice, topic1)

	err := serveTopicsJsonResponse(w, topicsSlice)
	if err != nil {
		fmt.Printf("Error obtaining topics: ",err)
		return
	}
}

func articlesResource(w http.ResponseWriter, r *http.Request) {

	excelArticles, articlesLen, err := ObtainArticlesFromExcel()
	fmt.Println("Obtained Excel articles: ",len(excelArticles), "length ",articlesLen)

	//articles, err := ObtainSampleArticles()
	
	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}

	err = serveArticlesJsonResponse(w, excelArticles[0:articlesLen])
	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}
}