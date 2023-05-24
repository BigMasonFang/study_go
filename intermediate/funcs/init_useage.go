// 1. reset package level var

// 2. initialize package var, make it useable

// 3. register mode (like factory pattern), see blow
// PrintInitDB and

// 3. check if everthing needed is success loaded befor
// the main runs
package funcs

import (
	"database/sql"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func PrintInitDB() {

	db, err := sql.Open("postgres", "user=test dbname=test sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	data, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func ImageSize(imageFile string) (int, int, error) {
	// support gif jpeg and png
	// cause import _ "image/{format}" calls the corresponded
	// init eg: $GOROOT/src/image/png/reader.go
	f, _ := os.Open(imageFile)
	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		return 0, 0, err
	}
	b := img.Bounds()
	return b.Max.X, b.Max.Y, nil
}
