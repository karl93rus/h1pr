package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bufio"
  "strings"
  "strconv"
)

var host = "https://hackerone.com/reports/"

func loopURLs() {
  var param int
  param = 104543
  for i := 0; i < 1100; i++ {
    resp, err := http.Get(host + strconv.Itoa(param))
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
        // fmt.Println(line)
        fmt.Println(host + strconv.Itoa(param))
        break
      }
    }
    //fmt.Println(i)
    param++
  }
}

func main() {
    loopURLs()
}
