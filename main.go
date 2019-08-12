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

//var id int = 104500

type id struct {
	id int
	m  sync.Mutex
}
func (i *id) Inc() {
	i.m.Lock()
	i.id++
	i.m.Unlock()
}
func (i *id) Write(ch chan string) {
	i.m.Lock()
	ch <- strconv.Itoa(i.id)
	i.m.Unlock()
}

//func checkUrl(host string, id int, ch chan string) {
//	for i := 0; i < 500; i++ {
//		resp, err := http.Get(host + strconv.Itoa(id))
//		if err != nil {
//			fmt.Println(err)
//		}
//		defer resp.Body.Close()
//
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			fmt.Println(err)
//		}
//		res := string(body)
//		scan := bufio.NewScanner(strings.NewReader(res))
//		for scan.Scan() {
//			line := scan.Text()
//			if strings.Index(line, "<meta name=\"twitter") >= 0 {
//				// fmt.Println(line)
//				// fmt.Println(host + strconv.Itoa(id))
//				ch <- host + strconv.Itoa(id)
//				break
//			}
//		}
//		id++
//		fmt.Println(id)
//	}
//}


func checkUrl(host string, id id, ch chan string) {
	for i := 0; i < 500; i++ {
		resp, err := http.Get(host + strconv.Itoa(id.id))
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
				// fmt.Println(host + strconv.Itoa(id))
        ch <- host + strconv.Itoa(id.id)
        id.Write(ch)
				break
			}
		}
		id.id++
		fmt.Println(id.id)
	}
}

func main() {
	// loopURLs()
	ch := make(chan string)
  count := id{id: 104540}

	for i := 0; i < 2; i++ {
		go checkUrl(host, count, ch)
	}

	// for j := 0; j < 5; j++ {
	//   go func() {
	//     for i := 0; i < 500; i++ {
	//       resp, err := http.Get(host + strconv.Itoa(id))
	//       if err != nil {
	//         fmt.Println(err)
	//       }
	//       defer resp.Body.Close()

	//       body, err := ioutil.ReadAll(resp.Body)
	//       if err != nil {
	//         fmt.Println(err)
	//       }
	//       res := string(body)
	//       scan := bufio.NewScanner(strings.NewReader(res))
	//       for scan.Scan() {
	//         line := scan.Text()
	//         if strings.Index(line, "<meta name=\"twitter") >= 0 {
	//           // fmt.Println(line)
	//           // fmt.Println(host + strconv.Itoa(id))
	//           ch <- host + strconv.Itoa(id)
	//           break
	//         }
	//       }
	//       id++
	//       fmt.Println(id)
	//     }
	//   }()
	// }

	for i := range ch {
		fmt.Println(i)
	}
}
