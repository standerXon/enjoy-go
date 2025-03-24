package main
import(
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchall() {
	start := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]
	for _,url := range urls {
		// 并发执行 fetch 函数
		go fetchOne(url,ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchOne(url string, ch chan<- string) {
	start := time.Now()
	// 内容
	resp,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 大小
	nbytes,err := io.Copy( io.Discard,resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v",url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs, nbytes, url)

}