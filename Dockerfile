FROM golang:1.14.3-alpine AS build
WORKDIR /code
COPY . .
RUN go build -o /out/server ./main.go
EXPOSE 8000
CMD ["/out/server"]
