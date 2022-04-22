package packet

import (
	"bytes"
	"go/types"
)

// instances

type Packet interface {
}

type IncomingServerboundPacket interface {
	Handle(listener Listener)
	Underlying() types.Type
	String() string
}

type OutgoingClientboundPacket interface {
	WriteTo(buffer bytes.Buffer)
}

type Builder[P Packet] interface {
	Build(buffer bytes.Buffer) P
}

// management

/*type Manager interface {
	AddIncomingPacket(packetType types.Type, builder Builder[IncomingServerboundPacket])
	GetIncomingPacketBuilder(id int) Builder[OutgoingClientboundPacket]

	//AddOutgoingPacket(packet chan OutgoingClientboundPacket)
	//GetOutgoingPacketId(packet chan OutgoingClientboundPacket) int
}*/

type Manager struct {
	incomingPacketToId        map[types.Type]int
	incomingPacketIdToBuilder []Builder[IncomingServerboundPacket]

	//outgoingPacketToId      map[chan OutgoingClientboundPacket]int
	//outgoingPacketIdCounter int
}

func (manager Manager) AddIncomingPacket(packetType types.Type, builder Builder[IncomingServerboundPacket]) {
	manager.incomingPacketToId[packetType] = len(manager.incomingPacketIdToBuilder)
	manager.incomingPacketIdToBuilder = append(manager.incomingPacketIdToBuilder, builder)
}

func (manager Manager) GetIncomingPacketBuilder(id int) Builder[IncomingServerboundPacket] {
	if len(manager.incomingPacketIdToBuilder) > id {
		return manager.incomingPacketIdToBuilder[id]
	}
	return nil
}
