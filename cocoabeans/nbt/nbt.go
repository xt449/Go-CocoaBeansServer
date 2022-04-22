package nbt

type Tag interface {
	GetId() uint8
}

// Container verbose nbt tag container
type Container interface {
	GetRoot() CompoundTag
}

type nbtContainer struct {
	root CompoundTag
}

func (container nbtContainer) GetRoot() CompoundTag {
	return container.root
}

func NewContainer() Container {
	return nbtContainer{NewCompoundTag()}
}

// MiniContainer short, less structured tag container
type MiniContainer CompoundTag

func (container MiniContainer) GetRoot() CompoundTag {
	return CompoundTag(container)
}

func NewMiniContainer() MiniContainer {
	return MiniContainer(NewCompoundTag())
}
