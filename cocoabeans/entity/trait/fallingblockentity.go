package trait

import (
	"CocoaBeansServer/cocoabeans/data/block_data/block_state"
)

// Sand, Gravel, ConcretePowder, DragonEgg, Anvil, Scafolding, etc.

type FallingBlockEntity interface {
	GetBlockState() block_state.BlockState
}
