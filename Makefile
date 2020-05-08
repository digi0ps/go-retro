APPNAME = "go-retro"
TARGET = "./out/$(APPNAME)"

all: pre build run

pre:
	gofmt -w .
	golint ./...

build:
	go build -o $(TARGET) ./main.go

run:
	./$(TARGET)
