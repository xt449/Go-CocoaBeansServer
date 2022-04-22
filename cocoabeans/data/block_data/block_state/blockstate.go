package block_state

import "CocoaBeansServer/cocoabeans/data"

type BlockState interface {
	GetNamespaceId() data.NamespaceId
	GetProperties() BlockStateProperties
}

type BlockStateProperties map[string]string
