APPNAME = "go-retro"
TARGET = "./out/$(APPNAME)"

all: lint build run

docker.build: .
	docker-compose build

docker.up:
	docker-compose up

docker.down:
	docker-compose down

lint:
	gofmt -w .
	golint ./...

test:
	gotest -v -cover ./...

build:
	go build -o $(TARGET) ./main.go

run:
	./$(TARGET)
