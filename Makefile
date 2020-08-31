APPNAME = "go-retro"
TARGET = "./out/$(APPNAME)"

all: fmt lint build test run

docker.build: .
	docker-compose build

docker.up:
	docker-compose up

docker.down:
	docker-compose down

fmt:
	gofmt -w ./

lint:
	golint ./...
	go vet ./...

test:
	gotest -cover ./...

make test.coverage:
	mkdir -p coverage
	gotest -coverprofile=coverage/coverage.out ./...
	gcov2lcov -infile=coverage/coverage.out -outfile=coverage/lcov.info

build:
	go build -v -o $(TARGET) ./main.go

run:
	./$(TARGET)
