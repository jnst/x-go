package protobuf

import (
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/jnst/x-go/protobuf/internal"
)

// Check will be serialize and deserialize.
func Check() {
	user := &pb.User{
		UserId: 1,
		Name:   "Taro",
		Lv:     1,
		Exp:    0,
	}

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newUser := &pb.User{}
	if err = proto.Unmarshal(data, newUser); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if user.Name != newUser.Name {
		log.Fatalf("data mismatch %q != %q", user.Name, newUser.Name)
	}
}
