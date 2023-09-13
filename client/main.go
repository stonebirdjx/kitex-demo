package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/stonebirdjx/kitex-demo/kitex_gen/api"
	"github.com/stonebirdjx/kitex-demo/kitex_gen/api/echo"
)

func main() {
	c, err := echo.NewClient("example-server", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
