package downloadimage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func DownloadWithoutGoroutine(pathFile, pathImage string) {
	os.Mkdir("Image", 0777)
	links, err := os.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range strings.Split(string(links), "\n")[1:] {
		downloadImage(link)

	}

}

func DownloadWithGoroutine(pathFile, pathImage string) {
	var wg sync.WaitGroup

	os.Mkdir("Image", 0777)
	links, err := os.ReadFile(pathFile)
	linksString := strings.Split(string(links), "\n")[1:]
	if err != nil {
		log.Fatal(err)
	}
	if len(linksString) > 10000 {
		wg.Add(10000)
	} else {
		wg.Add(len(linksString)-1)
	}
	
	for i, link := range linksString {

		go func(link string, i int) {
			defer wg.Done()
			name := uuid.NewString()
			response, e := http.Get(link)
			if e != nil {
				log.Fatal(e)
			}
			defer response.Body.Close()

			//open a file for writing
			file, err := os.Create(fmt.Sprintf("Image/%v.png", name))

			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			// Use io.Copy to just dump the response body to the file. This supports huge files
			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Success from %d goroutines! \n", i)
		} (link, i)

	}

	wg.Wait()
}

func downloadImage(link string) {
	name := uuid.NewString()
	response, e := http.Get(link)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(fmt.Sprintf("Image/%v.png", name))

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
