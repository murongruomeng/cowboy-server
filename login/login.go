package login

import (
	"cowboy-server/model"
	"cowboy-server/net"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func HandleLogin(conn *websocket.Conn, packet *net.Packet) {
	c2sLogin := &net.C2SLogin{}
	proto.Unmarshal(packet.Pbbody, c2sLogin)
	user, err := model.User{}.FindOne(bson.M{"uid": c2sLogin.Uid, "sessionkey": c2sLogin.Sessionkey})

	response := &net.Packet{}
	packet.Cmd = net.NetCMD_S2C_MSG_LOGIN
	s2cLogin := &net.S2CLogin{}
	if err == mgo.ErrNotFound {
		s2cLogin.Errcode = 2000
	} else {
		s2cLogin.Uid = user.Uid
		s2cLogin.Coin = user.Coin
		s2cLogin.Diamond = user.Coin
		s2cLogin.Sex = user.Sex
		s2cLogin.Level = user.Level
		s2cLogin.Name = user.Name
	}

	pbody, _ := proto.Marshal(c2sLogin)
	response.Pbbody = pbody

	data, _ := proto.Marshal(response)
	conn.WriteMessage(websocket.BinaryMessage, data)
}
