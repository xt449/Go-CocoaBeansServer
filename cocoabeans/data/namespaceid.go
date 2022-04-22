package data

type NamespaceId [2]string

func (id NamespaceId) String() string {
	return id[0] + ":" + id[1]
}

func MinecraftNamespaceId(path string) NamespaceId {
	return [2]string{"minecraft", path}
}
