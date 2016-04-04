package main

import (
	"fmt"
)

func embedArticleRelationInfo(mode string, rels []Article_Topic_Rel, 
	articles []Article, topics []Topic) ([]Article, []Topic){

	var relsMap map[int]Article_Topic_Rel

	if(mode == "articlesInTopics"){
		fmt.Println("Matched mode articlesInTopics")

		relsMap = obtainArticleTopicRelsByArticleId("byTopicId", rels)
		for _, element := range topics {
			fmt.Println("Current element: ", element)
			curTopicId := element.Id
			element.ArticleRels = append(element.ArticleRels, relsMap[curTopicId])
		}

	}else if(mode == "topicsInArticles"){
		fmt.Println("Matched mode topicsInArticles")

		relsMap = obtainArticleTopicRelsByArticleId("byArticleId", rels)
		for _, element := range articles {
			fmt.Println("Current element: ", element)
			curArticleId := element.Id
			element.TopicRels = append(element.TopicRels, relsMap[curArticleId])
		}
	}

	fmt.Println(relsMap)

	return articles, topics
}

func obtainArticleTopicRelsByArticleId(mode string, rels []Article_Topic_Rel) map[int]Article_Topic_Rel{

	relationsById := make(map[int]Article_Topic_Rel)

	for _, element := range rels {
		selectedKey := -1

		fmt.Println("Current rel: ",element)

		if(mode == "byArticleId"){
			selectedKey = element.Article_ID

		}else if(mode == "byTopicId"){
			selectedKey = element.Topic_ID
		}

		fmt.Println("Selected key: ",selectedKey)

		if selectedKey > -1 {
			relationsById[selectedKey] = element
		}
	}

	return relationsById
}