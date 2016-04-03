package main

import (
	"fmt"
)

func main() {

	excelArticles, lenArticles, err := ObtainArticlesFromExcel()

	//articles, err := ObtainSampleArticles()
	
	if err != nil {
		fmt.Printf("Error obtaining sample articles: ",err)
		return
	}
	fmt.Println("Obtained Excel articles: ",len(excelArticles), "length ",lenArticles)

	// for i := range excelArticles{
	// 	fmt.Println("Article",excelArticles[i])
	// }

}