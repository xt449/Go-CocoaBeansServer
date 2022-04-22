package block_state

import "Learning2/cocoabeans/data"

type BlockState interface {
	GetNamespaceId() data.NamespaceId
	GetProperties() BlockStateProperties
}

type BlockStateProperties map[string]string
