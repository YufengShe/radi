// Code generated by Kitex v0.2.0. DO NOT EDIT.

package ccoperation

import (
	"github.com/cloudwego/kitex/server"
	ccapi "radi/rpc/kitex_gen/ccAPI"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler ccapi.CCOperation, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
