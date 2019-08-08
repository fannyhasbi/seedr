package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fannyhasbi/seeder/config"
	"github.com/fannyhasbi/seeder/worker"
	_ "github.com/lib/pq"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var comments []worker.Comment
	var categories []worker.Category
	var articles []worker.Article

	commentDB := config.InitCommentDB()
	articleDB := config.InitArticleDB()
	defer commentDB.Close()
	defer articleDB.Close()

	file, err := ioutil.ReadFile("comments.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &comments)
	if err != nil {
		log.Fatal(err)
	}

	fileCategory, err := ioutil.ReadFile("categories.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileCategory, &categories)
	if err != nil {
		log.Fatal(err)
	}

	fileArticle, err := ioutil.ReadFile("articles.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileArticle, &articles)

	// worker.SeedComment(commentDB, comments)
	// worker.SeedCategory(articleDB, categories)
	worker.SeedArticle(articleDB, articles)

	log.Println("DONE")
}
