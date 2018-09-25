package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

func dispatchWebSocketClose(ctx *context.DispatchContext, webSocket *websocket.WebSocket) error {
	log.Printf("websocket [ %s ] to remote [ %s ] terminated\n", webSocket.Id, webSocket.RemoteAddr())
	// unregister the closed websocket
	connectionManager.Unregister(webSocket)
	return nil
}
