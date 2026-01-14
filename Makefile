
run:
	air
	
build:
	go build -o bin/app ./cmd/main.go

clean:
	rm -rf bin/