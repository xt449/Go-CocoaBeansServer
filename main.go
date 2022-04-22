package main

import (
	"CocoaBeansServer/cocoabeans/colors"
	"CocoaBeansServer/cocoabeans/nbt"
	"CocoaBeansServer/cocoabeans/protocol"
	"CocoaBeansServer/cocoabeans/protocol/packet"
	"fmt"
)

func main() {
	colorTest()

	fmt.Println(protocol.ACCEPTED)
	fmt.Println(protocol.DELAY_ACCEPT)
	fmt.Println(protocol.WAITING_FOR_START)

	var man packet.Manager
	man.GetIncomingPacketBuilder(1)

	var container = nbt.NewContainer("")
	container.Root["testb"] = nbt.NewByteTag(1)
	container.Root["testl"] = nbt.NewLongTag(2)

	fmt.Println(container)
}

func colorTest() {
	var color colors.Color
	var cc colors.Color

	color = colors.FromRGB(51, 153, 204)
	fmt.Println("#1 rgb", color)
	cc = color.ToHSL()
	fmt.Println("#1 hsl", cc)
	cc = color.ToHSV()
	fmt.Println("#1 hsv", cc)
}
