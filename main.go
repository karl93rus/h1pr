package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bufio"
  "strings"
  "strconv"
  "sync"
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
    fmt.Println(param)
    param++
  }
}

var id int = 104543

func main() {
  // loopURLs()
  var wg sync.WaitGroup
  for j := 0; j < 10; j++ {
    wg.Add(1)
    go func() {
      for i := 0; i < 1100; i++ {
        resp, err := http.Get(host + strconv.Itoa(id))
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
            fmt.Println(host + strconv.Itoa(id))
            break
          }
        }
        id++
        // fmt.Println(id)
      }
      wg.Done()
    }()
  }
  wg.Wait()
}















