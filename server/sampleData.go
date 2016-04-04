package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
    "time"
    "strconv"
)

func ObtainDataFromExcel () ([]Article,  []Topic, []Article_Topic_Rel,  error){

    articlesArray := make([]Article, 256)
    topicsArray := make([]Topic, 256)
    relsArray := make([]Article_Topic_Rel, 256)

    articlesSlice := articlesArray[0:0]
    topicsSlice := topicsArray[0:0]
    relsSlice := relsArray[0:0]

    lenArticles := 0
    lenTopics := 0
    lenRels := 0

    excelFileName := "./sampleData/SampleArticles.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)

    fmt.Println(xlFile)

    if err != nil {
        fmt.Println("Error opening sample articles Excel",err)
        return articlesArray[0:lenArticles], topicsArray[0:lenTopics], relsArray[0:lenRels], err
    }

    fmt.Println("Excel file successfully opened")

    for sheetIndex, sheet := range xlFile.Sheets {
        if sheetIndex == 0 {
            lenArticles, articlesSlice = addArticlesArrayFromExcelSheet(*sheet, 
                lenArticles, articlesSlice, articlesArray)

        } else if(sheetIndex == 1){
            lenTopics, topicsSlice = addTopicsArrayFromExcelSheet(*sheet, 
                lenTopics, topicsSlice, topicsArray)

        }else if(sheetIndex == 2){
            lenRels, relsSlice = addRelsArrayFromExcelSheet(*sheet, 
                lenRels, relsSlice, relsArray)
        }
    }

    return articlesArray[0:lenArticles], topicsArray[0:lenTopics], relsArray[0:lenRels], nil
}

func addTopicsArrayFromExcelSheet(sheet xlsx.Sheet, lenTopics int, 
    topicsSlice []Topic, topicsArray []Topic) (int, []Topic) {

    for i, row := range sheet.Rows {
        if i == 0 {
            continue
        }

        curTopic := Topic{}
        lenTopics += 1

        for cellIndex, cell := range row.Cells {
            cellValue, err := cell.String()

            if cellIndex == 0 {
                id, _ := strconv.Atoi(cellValue)
                curTopic.SetId(id)
            }else if cellIndex == 1{
                curTopic.SetName(cellValue)
            }

            if err != nil{
                continue
            }
        }
        //fmt.Println("Current article:")
        topicsSlice = append(topicsSlice, curTopic)
    }

    return lenTopics, topicsSlice
}

func addRelsArrayFromExcelSheet(sheet xlsx.Sheet, lenRels int, 
    relsSlice []Article_Topic_Rel, relsArray []Article_Topic_Rel) (int, []Article_Topic_Rel) {

    for i, row := range sheet.Rows {
        if i == 0 {
            continue
        }

        curRel := Article_Topic_Rel{}
        lenRels += 1

        for cellIndex, cell := range row.Cells {
            cellValue, err := cell.String()

            if cellIndex == 0 {
                id, _ := strconv.Atoi(cellValue)
                curRel.SetId(id)
            }else if cellIndex == 1{
                articleId, _ := strconv.Atoi(cellValue)
                curRel.SetArticleId(articleId)
            }else if cellIndex == 2{
                topicId, _ := strconv.Atoi(cellValue)
                curRel.SetTopicId(topicId)
            }else if cellIndex == 3{
                strength, _ := strconv.Atoi(cellValue)
                curRel.SetStrength(strength)
            }else if cellIndex == 4{
                curRel.SetDescription(cellValue)
            }

            if err != nil{
                continue
            }
        }
        //fmt.Println("Current article:")
        relsSlice = append(relsSlice, curRel)

    }

    return lenRels, relsSlice
}

func addArticlesArrayFromExcelSheet(sheet xlsx.Sheet, lenArticles int, 
    articlesSlice []Article, articlesArray []Article) (int, []Article) {

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
                id, _ := strconv.Atoi(cellValue)
                curArticle.SetId(id)
            }else if cellIndex == 1{
                curArticle.SetName(cellValue)
            }else if cellIndex == 2{
                curArticle.SetLink(cellValue)
            }else if cellIndex == 3{
                curArticle.SetDateStr(cellValue)
            }

            if err != nil{
                continue
            }
        }
        //fmt.Println("Current article:")
        articlesSlice = append(articlesSlice, curArticle)
    }

    return lenArticles, articlesSlice
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