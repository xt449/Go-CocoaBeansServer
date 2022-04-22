package trait

type Trait interface {
	CollectableEntity | FallingBlockEntity | LivingEntity | ProjectileEntity | ServerEntity | SpecialEntity
}
