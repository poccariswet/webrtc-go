package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

const (
	WS_SERVER = "8088"
)

type MSG struct {
	Type string
	Msg  string
	From string
}

type Connection struct {
	Conn    *websocket.Conn
	sendch  chan MSG
	Room    string
	Channel string
}

var connections map[string]*Connection

func wshandle(ws *websocket.Conn) {
	con := New(ws)
	connections[conn.Channel] = con
	wsMSGhandler(con)
}

func main() {
	conn = make(map[string]*Connection)
	http.HandleFunc("/websocket", func(w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(wshandle)}
		s.ServeHTTP(w, req)
	})
	if err := http.ListenAndServe(WS_SERVER, nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
