// Code generated by Kitex v0.7.1. DO NOT EDIT.
package echo

import (
	server "github.com/cloudwego/kitex/server"
	api "github.com/stonebirdjx/kitex-demo/kitex_gen/api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.Echo, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
