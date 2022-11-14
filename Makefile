install: 
	go clean
	go build -o ~/bin

test:
	go test -v