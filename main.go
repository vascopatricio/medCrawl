package main

import (
	"fmt"
)

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