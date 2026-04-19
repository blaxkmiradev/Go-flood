build:
	go build -o Go-flood main.go

run:
	go run main.go -config config.json

clean:
	rm -f Go-flood
