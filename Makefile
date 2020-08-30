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
	gotest -cover ./...

build:
	go build -v -o $(TARGET) ./main.go

run:
	./$(TARGET)
