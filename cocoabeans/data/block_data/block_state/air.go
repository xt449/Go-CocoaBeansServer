package block_state

import "Learning2/cocoabeans/data"

type Air struct {
}

func (Air) GetNamespaceId() data.NamespaceId {
	return [2]string{"minecraft", "air"}
}

func (Air) GetProperties() BlockStateProperties {
	return nil
}
