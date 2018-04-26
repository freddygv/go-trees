SOURCE = .
APP = redblack

all: build

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/$(APP)-linux-amd64 $(SOURCE) 
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(APP)-darwin-amd64 $(SOURCE) 

clean:
	rm -r bin/*