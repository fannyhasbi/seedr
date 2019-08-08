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
	db := config.InitCockroachDB()
	defer db.Close()

	file, err := ioutil.ReadFile("comments.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &comments)
	if err != nil {
		log.Fatal(err)
	}

	worker.SeedComment(db, comments)

	log.Println("DONE")
}
