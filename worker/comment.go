package worker

import (
	"database/sql"
	"fmt"
	"log"
)

// Comment store single comment
type Comment struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	CreatedTime string `json:"created_time"`
	Body        string `json:"body"`
}

// SeedComment seed comment data
func SeedComment(db *sql.DB, comments []Comment) {
	log.Println("Seeding comment table")

	query := `INSERT INTO comment (id, body, created_at, article_id, author_id) VALUES ($1, $2, $3, $4, $5)`

	for _, v := range comments {
		t := fmt.Sprintf("%s %s", v.CreatedAt, v.CreatedTime)

		_, err := db.Exec(query, v.ID, v.Body, t, "2e846a2d-dfef-42e9-a52c-b638786eecfb", "18371a10-29f9-4fb6-bab8-71e313890ea0")
		if err != nil {
			log.Println("[Comment][ERR]", err)
		}
	}

	log.Println("Comment done")
}
