package components

import (
	"hermannvincent/forgottenkingdom/ecs"
	"hermannvincent/forgottenkingdom/utils"
)

func PositionComponent(x, y, z int) *ecs.IComponent {
	return ecs.CreateComponent("position", map[string]interface{}{"x": x, "y": y, "z": z})
}

func PositionComponentByVec3(vector utils.Vec3) *ecs.IComponent {
	return ecs.CreateComponent("position", map[string]interface{}{"x": vector[0], "y": vector[1], "z": vector[2]})
}

func DimensionComponent(width, height int) *ecs.IComponent {
	return ecs.CreateComponent("dimension", map[string]interface{}{"width": width, "height": height})
}

func SolidComponent(isSolid bool) *ecs.IComponent {
	return ecs.CreateComponent("solid", map[string]interface{}{"isSolid": isSolid})
}

func ItemStorageComponent([]*ecs.IItem) *ecs.IComponent {
	return ecs.CreateComponent("item_storage", map[string]interface{}{})
}
