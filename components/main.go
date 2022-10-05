package components

import "hermannvincent/deliveryservice/ecs"

func PositionComponent(x, y int) *ecs.IComponent {
	component := ecs.CreateComponent("position", map[string]interface{}{"x": x, "y": y})
	return component
}

func DimensionComponent(width, height int) *ecs.IComponent {
	component := ecs.CreateComponent("dimension", map[string]interface{}{"width": width, "height": height})
	return component
}

func SolidComponent(isSolid bool) *ecs.IComponent {
	component := ecs.CreateComponent("solid", map[string]interface{}{"isSolid": isSolid})
	return component
}
