package main

import (
	colors2 "Learning2/cocoabeans/colors"
	"Learning2/cocoabeans/nbt"
	"Learning2/cocoabeans/protocol"
	"Learning2/cocoabeans/protocol/packet"
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
	var color colors2.Color
	var cc colors2.Color

	color = colors2.FromRGB(51, 153, 204)
	fmt.Println("#1 rgb", color)
	cc = color.ToHSL()
	fmt.Println("#1 hsl", cc)
	cc = color.ToHSV()
	fmt.Println("#1 hsv", cc)
}
