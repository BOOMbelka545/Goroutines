package main

import (
	di "goroutine/downloadImage"
	"os"

	"github.com/labstack/gommon/log"
)

var (
	pathFile = "links"
	pathLink = "https://www.freecodecamp.org/news/content/images/2021/10/golang.png"
	amount = 1000
)

func main() {
	writeLinks(amount, pathLink)
	di.DownloadWithGoroutine(pathFile, "")
}

func writeLinks(amount int, links string) {
	f, err := os.Create(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	for ; amount > 0; amount-- {
		_, err = f.WriteString("\n"+links)
		if err != nil {
			log.Fatal(err)
		}
	}
}
