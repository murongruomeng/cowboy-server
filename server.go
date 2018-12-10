package main

import (
	"cowboy-server/login"
	"cowboy-server/net"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	startServer()
}

func startServer() {
	log.Println("Server starting...")
	http.HandleFunc("/", handleWs)
	http.ListenAndServe(":9999", nil)
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade: ", err)
	}

	defer conn.Close()

	msgType, data, err := conn.ReadMessage()
	if msgType == websocket.BinaryMessage {
		packet := &net.Packet{}
		proto.Unmarshal(data, packet)
		switch packet.Cmd {
		case net.NetCMD_C2S_MSG_LOGIN:
			{
				go login.HandleLogin(conn, packet)
			}
		}
	}
}
