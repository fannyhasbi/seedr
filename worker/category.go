package worker

import (
	"database/sql"
	"log"
	"strings"
)

// Category store single category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// SeedCategory seed category data
func SeedCategory(db *sql.DB, categories []Category) {
	log.Println("Seeding category table")

	query := `INSERT INTO category (id, name, slug, author_id) VALUES ($1, $2, $3, $4)`

	for _, v := range categories {
		s := strings.ToLower(v.Name)
		s = strings.ReplaceAll(s, " ", "-")

		_, err := db.Exec(query, v.ID, v.Name, s, "18371a10-29f9-4fb6-bab8-71e313890ea0")
		if err != nil {
			log.Println("[Category][ERR]", err)
		}
	}

	log.Println("Category done")
}
