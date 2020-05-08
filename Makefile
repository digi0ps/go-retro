APPNAME = "go-retro"
TARGET = "./out/$(APPNAME)"

all: pre build run

pre:
	gofmt -w .
	golint ./...

test:
	gotest -v -cover ./...

build:
	go build -o $(TARGET) ./main.go

run:
	./$(TARGET)
