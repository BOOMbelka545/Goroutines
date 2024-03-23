package downloadimage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
)

func DownloadWithoutGoroutine(pathFile, pathImage string) {
	os.Mkdir("Image", 0777)
	links, err := os.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	for i, link := range strings.Split(string(links), "\n")[1:]{
		response, e := http.Get(link)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		file, err := os.Create(fmt.Sprintf("Image/%v.png", i))

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Success!")
	}

}

// func DownloadWithGoroutine(pathFile, pathImage string) {

// }