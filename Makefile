
install-windows:
	choco install air
	choco install golang

install-linux:
	sudo apt-get update
	sudo apt-get install -y golang-go
	go install github.com/cosmtrek/air@latest

install-mac:
	brew install golang
	brew install air

run:
	air
	
build:
	go build -o bin/app ./cmd/main.go

clean:
	rm -rf bin/