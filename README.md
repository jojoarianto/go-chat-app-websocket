# Chat APP

Simple chat app using gorilla websocket websocket

> Demo [screenshot app](#demo-screenshot) 

## Goal

- [X] [API for sending a message](#api-endpoint-for-sending-a-message)
- [X] [API for collect message that has been sent out](#api-for-collect-message-that-has-been-sent-out)
- [X] [API for display message in real time](#display-message-on-websocket-realtime)

## Installation & Run

First, Make sure you have set up \$GOPATH.

```bash
# Download this project
go get github.com/jojoarianto/go-chat-app-websocket

# It's take several minute to download project
```

Set project environment and run

```bash
# move to project directory
cd $GOPATH/src/github.com/jojoarianto/go-chat-app-websocket

# run golang project
go run main.go

# Visit : http://localhost:8000/
```

## Structure Design

- Domain
  - Define struct represent mapping to data model
      - message.go
- Interfaces
  - HTTP handler
  - Websocket handler
  - Respond handler
- Static
  - Static assets
      - .png .jpeg
      - .css 
  - index.html

## Model Data
Message (Pesan)
```go 
type Message struct {
	ID             xid.ID       `json:"id"`
	Sender         string       `json:"sender,omitempty"`
	ContentMessage string       `json:"content_message"`
	CreatedAt      time.Time    `json:"created_at"`
}
```

### API endpoint for sending a message

> `POST` localhost:8000/sent

#### Request with data json 
```json
{
    "content_message": "Hi Chat App"
}
```

or using curl

```bash
curl --request POST \
  --url http://localhost:8000/sent \
  --header 'content-type: application/json' \
  --data '{
  	    "content_message": "Hi Chat App"
    }'
```

#### Respond
```json
{
  "id": "bhgi9mml0s15jsg5c260",
  "sender": "Anonim",
  "content_message": "Hi Chat App",
  "created_at": "2019-02-11T14:31:38.711328+07:00"
}
```

### API for collect message that has been sent out

> `GET` localhost:8000/sent

#### Request with hit http://localhost:8000/sent

or using curl

```bash
curl --request GET \
  --url http://localhost:8000/sent
```
#### Respond

```json
[
  ...
  {
    "id": "bhgibael0s15jsg5c27g",
    "sender": "Anonim",
    "content_message": "Ada apa",
    "created_at": "2019-02-11T14:35:05.886498+07:00"
  },
  {
    "id": "bhgibcml0s15jsg5c28g",
    "sender": "Anonim",
    "content_message": "tidak ada apa apa",
    "created_at": "2019-02-11T14:35:14.373446+07:00"
  }
  ...
]
```

### Display Message on Websocket (Realtime)

> `Websocket` ws://localhost:8000/ws

#### Open on browser http://localhost:8000

or using curl

```bash
curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Host: localhost:8000" \
     --header "Origin: http://localhost:8000" \
     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
     --header "Sec-WebSocket-Version: 13" \
     http://localhost:8000/ws
```

## Unit Test

```bash
# move to interfaces folder
cd $GOPATH/src/github.com/jojoarianto/go-chat-app-websocket/interfaces

# run golang test
go test -V
```

## Libraries
- Gorilla Websocket from https://github.com/gorilla/websocket
- HttpRouter from https://github.com/julienschmidt/httprouter
- Globally Unique ID Generator from https://github.com/rs/xid
- End-to-end HTTP and REST API testing from https://github.com/gavv/httpexpect

## Demo Screenshot

![demo-png](https://user-images.githubusercontent.com/5858756/52549176-79d10880-2e04-11e9-9c09-902bf5db00bb.png)

