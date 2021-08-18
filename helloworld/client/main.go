package main

import (
	"github.com/micro/micro/v3/service"
	"time"
	proto "github.com/chyidl/services/helloworld/proto"
)

func main() {
	// create an initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHello
}
