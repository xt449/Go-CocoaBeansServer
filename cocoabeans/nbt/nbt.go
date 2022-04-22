package nbt

type Tag interface {
	GetId() uint8
}

type Container struct {
	RootName string
	Root     CompoundTag
}

func NewContainer(name string) Container {
	return Container{name, NewCompoundTag()}
}

// Holder NBT implementors should use this
type Holder interface {
	BuildNBTContainer() Container
}
