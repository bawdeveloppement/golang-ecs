package ecs

type IItem interface {
	GetId() string
	GetName() string
	GetDescription() string
	Use(*IWorld) string
}
