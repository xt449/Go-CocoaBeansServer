package block_state

import "CocoaBeansServer/cocoabeans/data"

type Stone struct {
}

func (Stone) GetNamespaceId() data.NamespaceId {
	return [2]string{"minecraft", "stone"}
}

func (Stone) GetProperties() BlockStateProperties {
	return nil
}
