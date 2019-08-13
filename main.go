package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var host = "https://hackerone.com/reports/"

var id int = 104500

func checkUrl(host string, id *int, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
    postId := *id
		resp, err := http.Get(host + strconv.Itoa(postId))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		res := string(body)
		scan := bufio.NewScanner(strings.NewReader(res))
		for scan.Scan() {
			line := scan.Text()
			if strings.Index(line, "<meta name=\"twitter") >= 0 {
				fmt.Println(host + strconv.Itoa(postId))
				break
			}
		}
		*id = *id + 1
		fmt.Println(*id)
	}
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  for i := 0; i < 25; i++ {
    wg.Add(1)
    go checkUrl(host, &id, &wg)
  }
  wg.Wait()
}
