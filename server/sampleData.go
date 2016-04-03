package main

import (
	"fmt"
    "github.com/tealeg/xlsx"
	"time"
)

func ObtainArticlesFromExcel () ([]Article, int, error){

    articlesArray := make([]Article, 256)

	excelFileName := "./sampleData/SampleArticles.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)

    fmt.Println(xlFile)

    if err != nil {
        fmt.Println("Error opening sample articles Excel",err)
        return articlesArray, 0, err
    }

    fmt.Println("Excel file successfully opened")

    articlesSlice := articlesArray[0:0]
    lenArticles := 0

    for _, sheet := range xlFile.Sheets {
        for i, row := range sheet.Rows {
        	if i == 0 {
        		continue
        	}

        	fmt.Println("Found row ")

        	curArticle := Article{}
        	lenArticles += 1
            	
            for cellIndex, cell := range row.Cells {
            	cellValue, err := cell.String()


            	if cellIndex == 0 {
            		curArticle.SetName(cellValue)
        		}else if cellIndex == 1{
            		curArticle.SetLink(cellValue)
        		}else if cellIndex == 2{
            		curArticle.SetDateStr(cellValue)
        		}

            	if err != nil{
            		fmt.Printf("Error obtaining cell contents")
            	}
                fmt.Printf("%s\n", cellValue)
            }
            fmt.Println("Current article:")
            articlesSlice = append(articlesSlice, curArticle)
        }
    }

	return articlesArray, lenArticles, nil
}

func ObtainSampleArticles () ([]Article, error) {

	art1 := Article{
		Name:"Women in Science: Neurology professor inspires sophomore to pursue her dreams",
		Link: "http://www.wildcat.arizona.edu/article/2016/03/neurology-professor-inspires-sophomore-to-pursue-her-dreams",
		dateStr: "03/30/2016"}

	art2 := Article{
		Name: "MU Neurology's art auction part of comprehensive approach",
		Link: "http://www.herald-dispatch.com/news/mu-neurology-s-art-auction-part-of-comprehensive-approach/article_7ced5b49-ba53-56ac-a8a6-32a26f0fbe20.html",
		dateStr: "03/31/2016"}

	articles := make([]Article, 2)
	articles[0] = art1
	articles[1] = art2

	for i := range articles{
		dateStr := articles[i].GetDateStr()
		//fmt.Println("Current date formatted:",dateStr)

		t, e := time.Parse("01/02/2006", dateStr)
		if e != nil {
			fmt.Println("Error obtaining date from ",dateStr)
		}else {
			articles[i].SetDate(t)
		}
		//fmt.Println("Date parsed:",articles[i].GetDate())
	}

	return articles, nil
}