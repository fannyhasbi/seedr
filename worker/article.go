package worker

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// Article store single article
type Article struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	CreatedTime string `json:"created_time"`
	Name        string `json:"name"`
}

// SeedArticle seed article data
func SeedArticle(db *sql.DB, articles []Article) {
	log.Println("Seeding article table")

	body := "{\"blocks\":[{\"key\":\"5vdda\",\"text\":\"Catatan Redaksi: Judul berita ini sudah mengalami ubahan dari 'AHM Sodorkan Honda PCX Listrik untuk Armada Gojek' ke 'Honda PCX Listrik Jadi Armada Gojek'. Penggunaan Honda PCX Listrik sebagai unit driver Gojek merupakan tindak lanjut dari kolaborasi Astra International Tbk ke layanan transportasi berbasis aplikasi itu.\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[{\"offset\":0,\"length\":321,\"style\":\"ITALIC\"},{\"offset\":0,\"length\":17,\"style\":\"BOLD\"}],\"entityRanges\":[],\"data\":{}},{\"key\":\"4cgeb\",\"text\":\"Layanan transportasi berbasis aplikasi Gojek resmi berkolaborasi dengan Astra, untuk mengoperasikan beberapa unit Honda PCX Listrik yang disediakan PT Astra Honda Motor (AHM)  buat mengangkut penumpang.\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":114,\"length\":6,\"key\":0}],\"data\":{}},{\"key\":\"fdfb6\",\"text\":\"Namun, berdasarkan unggahan Gojek pada akun resmi Instagram-nya, ini masih tahap uji coba awal. Nantinya hasil dari tes ini, akan menjadi masukan untuk pengembangan ke tahap berikutnya.⁣\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}}],\"entityMap\":{\"0\":{\"type\":\"LINK\",\"mutability\":\"MUTABLE\",\"data\":{\"href\":\"https://kumparan.com/@kumparanoto/gesits-laku-1-200-unit-honda-pcx-listrik-justru-sepi-peminat-1r2444Qrgp7\",\"rel\":\"noopener noreferrer\",\"target\":\"_blank\",\"url\":\"https://kumparan.com/@kumparanoto/gesits-laku-1-200-unit-honda-pcx-listrik-justru-sepi-peminat-1r2444Qrgp7\"}}}}"
	imageURL := "https://www.thamesvalley.police.uk/SysSiteAssets/media/images/thames-valley/featured-content/news_website_box_640x394px.png"

	query := `INSERT INTO article (id, slug, name, body, status, created_at, updated_at, author_id, comments_count, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	for _, v := range articles {
		s := strings.ToLower(v.Name)
		s = strings.ReplaceAll(s, " ", "-")
		t := fmt.Sprintf("%s %s", v.CreatedAt, v.CreatedTime)

		_, err := db.Exec(query, v.ID, s, v.Name, body, "PUBLISHED", t, t, "589575ef-1989-4415-9e1c-e17b7ed6275d", 0, imageURL)
		if err != nil {
			log.Println("[Article][ERR]", err)
		}
	}

	log.Println("Article done")
}
