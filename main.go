package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var fullUrl string

	fmt.Print("Enter full url: ")
	_, err := fmt.Scanf("%s", &fullUrl)
	if err != nil {
		log.Fatalln(err)

	}

	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	//create folder
	err = os.MkdirAll("google-fonts", 0755)
	if err != nil {
		fmt.Println("Can not create folder:", err)
		return
	}
	//replace and write to file
	var cssFile = filepath.Join("google-fonts", "fonts.css")
	f, err := os.OpenFile(cssFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Can not open file:", err)
		return
	}
	defer f.Close()

	content := strings.Replace(string(body), "https://fonts.gstatic.com/", "/google-fonts/", -1)
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Fatalln("Can not write to file:", err)
		return
	}

	/**Download files**/
	re := regexp.MustCompile(`src:\s*url\((.*?)\)`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		var fileUrl = match[1]
		var file = filepath.Join("google-fonts", strings.Replace(fileUrl, "https://fonts.gstatic.com/", "", 1))
		downloadFile(fileUrl, file)
	}
}
