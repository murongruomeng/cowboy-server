package model

import (
	"cowboy-server/db"
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	Uid        string    `bson:"uid"`
	SessionKey string    `bson:"sessionkey"`
	Deviceid   string    `bson:"deviceid"`
	Name       string    `bson:"name"`
	Sex        int64     `bson:"sex"`
	Avatar     string    `bson:"avatar"`
	Score      int       `bson:"score"`
	Headshot   int       `bson:"headshot"`
	Coin       int64     `bson:"coin"`
	Diamond    int       `bson:"diamond"`
	Level      uint32    `bson:"level"`
	CreateTime time.Time `bson:"create_time"`
}

func (user User) Save() error {
	session, collection := db.Connect("user")
	defer session.Close()
	return collection.Insert(user)
}

func (user User) FindOne(m bson.M) (User, error) {
	session, collection := db.Connect("user")
	defer session.Close()
	err := collection.Find(m).One(&user)
	return user, err
}
