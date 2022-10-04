package components

import "hermannvincent/deliveryservice/ecs"

func PositionComponent(x, y int) ecs.IComponent {
	return ecs.CreateComponent("position", map[string]interface{}{"x": 0, "y": 0})
}

func DimensionComponent(width, height int) ecs.IComponent {
	return ecs.CreateComponent("dimension", map[string]interface{}{"width": 32, "height": 32})
}
