run/touch:
	go run cmd/touch/main.go

build/touch:
	go build -o .out/touch.exe cmd/touch/main.go

clean:
	go clean