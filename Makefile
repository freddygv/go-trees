SOURCE = .
APP = redblack

all: build

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/$(APP)-linux-amd64 -i $(SOURCE)
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(APP)-darwin-amd64 -i $(SOURCE)
	env GOOS=windows GOARCH=amd64 go build -o bin/$(APP)-windows-amd64.exe -i $(SOURCE)

clean:
	rm -r bin/*