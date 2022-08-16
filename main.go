package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mistralll/gakujo-chromedp/scrape"
)

func main() {

	ans, err := scrape.GetHomeHTML(`eh022046`, `g27ETh45Et`)
	if err != nil {
		log.Fatal(err)
	}
	SaveToFileAsText(`res.txt`, ans)
}

func SaveToFileAsText(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)

	fmt.Println("Complited!")

	return nil
}
