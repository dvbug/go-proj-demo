package scripts

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//var buf bytes.Buffer
	//io.Copy(&buf, resp.Body)
	//nbytes := buf.Len()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	ms := time.Since(start).Milliseconds()
	ch <- fmt.Sprintf("%dms %7d %s", ms, nbytes, url)
}

func FetchMain() {
	urls := []string{"https://golang.org", "https://baidu.com", "https://ibug.info",
		"https://books.studygolang.com", "https://www.dsdsadfsdfadddd.com"}
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
