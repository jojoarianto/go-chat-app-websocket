## Chat APP

Libraries
- HttpRouter from github.com/julienschmidt/httprouter
- GorillaWebsocket from github.com/gorilla/websocket

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