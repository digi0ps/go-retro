# Go Retro

## About

Backend for Retro built using Go which serves both REST and
Websockets.

## Installation

1. Make sure Docker is installed
2. Run `make docker.build` whenever you want to build
3. Run `make docker.up` to start server
4. Run `make docker.down` to stop server

## Contract

### Ping

URL: /ping

Method: **GET**

Expected Response: `PONG`

### Create Board

URL: /api/board

Method: **PUT**

Body:

- **title** : _string_

Expected Response:

```json
{
  "success": true,
  "data": {
    "board": "5f28eed5ce811e0f52aa2cc5"
  }
}
```

### Get Board

URL: /api/board?id=string

Method: **GET**

Params:

- **id** : _string_

Expected Response:

```json
{
  "success": true,
  "data": {
    "id": "5f28eed5ce811e0f52aa2cc5",
    "title": "",
    "columns": [],
    "created_at": 1596518101422902733
  }
}
```

## Major Tasks

- [x] Setup Go
- [x] Setup routing using Gorilla Mux
- [x] Setup Mongo datastore
- [x] Setup REST endpoints
- [x] Setup WS endpoint
- [x] Setup WS handlers
- [ ] Swap Columns

## Minor Tasks

- [x] Dockerise
- [ ] Log, Recover middleware
- [ ] Reporting
- [ ] Better logger
