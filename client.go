package main

import (
	"cowboy-server/net"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:9999", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	request := &net.C2SLogin{}
	request.Uid = "1"
	request.Sessionkey = "xxxx"

	brequest, _ := proto.Marshal(request)

	packet := &net.Packet{}
	packet.Pbbody = brequest
	packet.Cmd = net.NetCMD_C2S_MSG_LOGIN

	bpacket, _ := proto.Marshal(packet)

	c.WriteMessage(websocket.BinaryMessage, bpacket)
}
