package main

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"strings"
)

/*
	// child pipeline initializer
	var childInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			// the maximum allowable packet length is 128 bytes，use \n to split, strip delimiter
			AddLast(frame.DelimiterCodec(128, "\n", true)).
			// convert to string
			AddLast(format.TextCodec()).
			// LoggerHandler, print connected/disconnected event and received messages
			AddLast(LoggerHandler{}).
			// UpperHandler (string to upper case)
			AddLast(UpperHandler{})
	}

	// create bootstrap & listening & accepting
	netty.NewBootstrap(netty.WithChildInitializer(childInitializer)).Listen("0.0.0.0:25565").Sync()
*/

type LoggerHandler struct{}

func (LoggerHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Println("go-netty:", "->", "active:", ctx.Channel().RemoteAddr())
	// write welcome message
	ctx.Write("Hello I'm " + "go-netty")
}

func (LoggerHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println("go-netty:", "->", "handle read:", message)
	// leave it to the next handler(UpperHandler)
	ctx.HandleRead(message)
}

func (LoggerHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Println("go-netty:", "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	// disconnected，the default processing is to close the connection
	ctx.HandleInactive(ex)
}

type UpperHandler struct{}

func (UpperHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	// text to upper case
	text := message.(string)
	upText := strings.ToUpper(text)
	// write the result to the client
	ctx.Write(text + " -> " + upText)
}
