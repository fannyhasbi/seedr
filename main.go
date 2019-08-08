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

	commentDB := config.InitCommentDB()
	categoryDB := config.InitCategoryDB()
	defer commentDB.Close()
	defer categoryDB.Close()

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

	// worker.SeedComment(commentDB, comments)
	worker.SeedCategory(categoryDB, categories)

	log.Println("DONE")
}
