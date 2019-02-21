package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"github.com/yields/phony/pkg/phony"
)

func main() {
	key := flag.String("api-key", "", "the api key to send to the endpoint")
	endpoint := flag.String("endpoint", "", "the endpoint to send to (https://api.yourtool.com)")
	dir := flag.String("dir", "", "the templates directory to use. Default is: ../../templates")

	flag.Parse()

	if *key == "" {
		log.Fatal("Error, must specify a valid --api-key <...>")
	}

	if *endpoint == "" {
		log.Fatal("Error, must specify a valid --endpoint <...>")
	}

	box := packr.NewBox("../../templates")
	if *dir == "" {
		box = packr.NewBox(*dir)
	}

	lines := readdir(box)

	errors := 0
	success := 0
	linesCounter := 0

	ch := make(chan string, len(lines))

	for line := range lines {
		linesCounter += 1
		go createAndSendRequest(line, endpoint, key, ch)
	}

	for i := 0; i < linesCounter; i++ {
		result := <-ch

		if result == "success" {
			success += 1
		} else {
			errors += 1
		}
	}

	fmt.Printf("errors: %d, successes: %d \n", errors, success)
}

func createAndSendRequest(data string, endpoint *string, key *string, ch chan string) {
	gen, err := compile(data)
	if err != nil {
		log.Fatal(err)
	}
	err = request(*key, *endpoint, gen())
	if err != nil {
		log.Printf("Warn: error hitting endpoint %v\n", err)
		ch <- "error"
	} else {
		ch <- "success"
	}
}

func readdir(box packr.Box) chan string {
	ch := make(chan string, 100)

	go func() {
		box.Walk(func(path string, f packd.File) error {
			sc := bufio.NewScanner(f)
			for sc.Scan() {
				ch <- sc.Text()
			}
			return nil
		})
		close(ch)
	}()

	return ch
}

func request(key, endpoint, data string) error {
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.SetBasicAuth(key, "")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		errStr := "Bad response, got: " + resp.Status + " " + string(body) + ", sent: " + data
		return errors.New(errStr)
	}

	return nil
}

func compile(tmpl string) (func() string, error) {
	expr, err := regexp.Compile(`({{ *(([a-zA-Z0-9]+(\.[a-zA-Z0-9]+)?)+) *}})`)
	if err != nil {
		return nil, err
	}

	return func() string {
		return expr.ReplaceAllStringFunc(tmpl, func(s string) string {
			path := strings.Trim(s[2:len(s)-2], " ")
			return phony.Get(path)
		})
	}, nil
}
