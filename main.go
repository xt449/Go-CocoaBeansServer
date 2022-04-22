package main

import (
	"Learning2/cocoabeans/data/colors"
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

	var tag nbt.Container = nbt.NewContainer()
	tag.GetRoot()["test"] = nbt.NewByteTag(1)

	var test = []int{1, 2, 3}
	fmt.Printf("%v %T\n", len(test), len(test))

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
