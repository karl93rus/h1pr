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

// func loopURLs() {
// 	var param int
// 	param = 104543
// 	for i := 0; i < 1100; i++ {
// 		resp, err := http.Get(host + strconv.Itoa(param))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 
// 		defer resp.Body.Close()
// 
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		res := string(body)
// 
// 		scan := bufio.NewScanner(strings.NewReader(res))
// 		for scan.Scan() {
// 			line := scan.Text()
// 			if strings.Index(line, "<meta name=\"twitter") >= 0 {
// 				// fmt.Println(line)
// 				fmt.Println(host + strconv.Itoa(param))
// 				break
// 			}
// 		}
// 		fmt.Println(param)
// 		param++
// 	}
// }

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


// type id struct {
// 	id int
// 	m  *sync.Mutex
// }
// func (i *id) Inc() {
// 	i.m.Lock()
// 	i.id++
// 	i.m.Unlock()
// }
// func (i *id) Write(ch chan string) {
// 	i.m.Lock()
// 	ch <- strconv.Itoa(i.id)
// 	i.m.Unlock()
// }
// func (i *id) WriteStr(r string) string {
// 	i.m.Lock()
//   fmt.Println(r)
// 	i.m.Unlock()
//   return r
// }

// func checkUrl(host string, id *id, ch chan string) {
// 	//for i := 0; i < 500; i++ {
// 		resp, err := http.Get(host + strconv.Itoa(id.id))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		defer resp.Body.Close()
// 
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		res := string(body)
// 		scan := bufio.NewScanner(strings.NewReader(res))
// 		for scan.Scan() {
// 			line := scan.Text()
// 			if strings.Index(line, "<meta name=\"twitter") >= 0 {
// 				// fmt.Println(line)
// 				// fmt.Println(host + strconv.Itoa(id))
//         //ch <- host + strconv.Itoa(id.id)
//         id.Write(ch)
//         //id.WriteStr(host + strconv.Itoa(id.id))
// 				break
// 			}
// 		}
// 		id.Inc()
// 		//fmt.Println(id.id)
// 	//}
// }

func main() {
	// loopURLs()
	//ch := make(chan string)
  //count := id{id: 104540, m: new(sync.Mutex)}
  var wg sync.WaitGroup
  for i := 0; i < 25; i++ {
    wg.Add(1)
    go checkUrl(host, &id, &wg)
  }
  wg.Wait()

	// for i := 0; i < 5; i++ {
	// 	go checkUrl(host, &count, ch)
	// }

	// for i := range ch {
	// 	fmt.Println(<-i)
	// }

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

}
