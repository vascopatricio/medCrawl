package main

import (
"fmt"
"net/http"
	"encoding/json"
)

type TopicsResponse struct {
	Topics []Topic
    Successful string
}

type ArticlesResponse struct {
	Articles []Article
    Successful string
}

func serveTopicsJsonResponse(w http.ResponseWriter, topics []Topic) (error) {

	response := TopicsResponse{topics, "true"}

	res, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error mashaling JSON topics",err)
		return err
	}

	return serveGenericJsonResponse(w, res)
}

func serveArticlesJsonResponse(w http.ResponseWriter, articles []Article) (error) {

	response := ArticlesResponse{articles, "true"}

	res, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error mashaling JSON topics",err)
		return err
	}

	return serveGenericJsonResponse(w, res)
}

func serveGenericJsonResponse(w http.ResponseWriter, res []byte) error{

	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, string(res))

	return nil
}