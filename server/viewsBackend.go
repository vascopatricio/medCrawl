package main

import (
"fmt"
"net/http"
)

func topicsResource(w http.ResponseWriter, r *http.Request) {

	excelArticles, excelTopics, excelRels, err := ObtainDataFromExcel()
	fmt.Println("Obtained Excel articles: ",len(excelArticles))
	fmt.Println("Obtained Excel topics: ",len(excelTopics))
	fmt.Println("Obtained Excel rels: ",len(excelRels))

	_, excelTopics = embedArticleRelationInfo("articlesInTopics", excelRels, excelArticles, excelTopics)
	
	err = serveTopicsJsonResponse(w, excelTopics)
	if err != nil {
		fmt.Printf("Error obtaining topics: ",err)
		return
	}
}

func articlesResource(w http.ResponseWriter, r *http.Request) {

	excelArticles, excelTopics, excelRels, err := ObtainDataFromExcel()
	fmt.Println("Obtained Excel articles: ",len(excelArticles))
	fmt.Println("Obtained Excel topics: ",len(excelTopics))
	fmt.Println("Obtained Excel rels: ",len(excelRels))

	excelArticles, _ = embedArticleRelationInfo("topicsInArticles", excelRels, excelArticles, excelTopics)

	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}

	err = serveArticlesJsonResponse(w, excelArticles)
	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}
}