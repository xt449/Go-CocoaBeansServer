package block_state

import "CocoaBeansServer/cocoabeans/data"

type Air struct {
}

func (Air) GetNamespaceId() data.NamespaceId {
	return [2]string{"minecraft", "air"}
}

func (Air) GetProperties() BlockStateProperties {
	return nil
}
