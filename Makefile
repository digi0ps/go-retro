APPNAME = "go-retro"
TARGET = "./out/$(APPNAME)"

all: lint build run

lint:
	gofmt -w .
	golint ./...

test:
	gotest -v -cover ./...

build:
	go build -o $(TARGET) ./main.go

run:
	./$(TARGET)
