package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/jnst/go-world/protobuf/user"
)

func main() {
	getUser := &user.GetUser{
		UserId: proto.Int32(1),
		Name:   proto.String("Taro"),
		Lv:     proto.Int32(1),
		Exp:    proto.Int32(0),
	}
	data, err := proto.Marshal(getUser)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newGetUser := &user.GetUser{}
	err = proto.Unmarshal(data, newGetUser)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if getUser.GetName() != newGetUser.GetName() {
		log.Fatalf("data mismatch %q != %q", getUser.GetName(), newGetUser.GetName())
	}
}
