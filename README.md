# Chat APP

Simple chat app using gorilla websocket websocket

> Demo app

![demo-png](https://user-images.githubusercontent.com/5858756/52549176-79d10880-2e04-11e9-9c09-902bf5db00bb.png)

## Goal

- [X] API for sending a message
- [X] API for collect message that has been sent out
- [X] API for display message in real time
- [X] Create a static page to run Websocket
- [ ] Create a unit testing

## Model Data
Message (Pesan)
```go 
type Message struct {
	ID			   xid.ID       `json:"id"`
	Sender         string       `json:"sender,omitempty"`
	ContentMessage string       `json:"content_message"`
	CreatedAt      time.Time    `json:"created_at"`
}
```

## Libraries
- Gorilla Websocket from github.com/gorilla/websocket
- HttpRouter from github.com/julienschmidt/httprouter
- Globally Unique ID Generator from https://github.com/rs/xid

### API for sending a message

```bash
curl --request POST \
  --url http://localhost:8000/sent \
  --header 'content-type: application/json' \
  --data '{
  	    "content_message": "Hi Chat App"
    }'
```

### API for collect message that has been sent out

```bash
curl --request GET \
  --url http://localhost:8000/sent
```

### Display Message on Websocket (Realtime)

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

curl --include --no-buffer --header "Connection: Upgrade" --header "Upgrade: websocket" --header "Host: localhost:8000" --header "Origin: http://localhost:8000" --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" --header "Sec-WebSocket-Version: 13" http://localhost:8000/ws