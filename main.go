package main

import (
	"hermannvincent/deliveryservice/components"
	"hermannvincent/deliveryservice/ecs"
	"hermannvincent/deliveryservice/systems"
)

func main() {
	newWorld := ecs.World{}

	var playerEntity ecs.IEntity = &ecs.Entity{
		Id: "Yeah",
		Components: []ecs.IComponent{
			&ecs.Component{Id: "position", Data: map[string]interface{}{"x": 0, "y": 0}},
			ecs.CreateComponent("dimension", map[string]interface{}{"width": 32, "height": 32}),
		},
	}

	var otherPlayerEntity ecs.IEntity = &ecs.Entity{
		Id: "Yeah2",
		Components: []ecs.IComponent{
			components.PositionComponent(0, 0),
			components.DimensionComponent(32, 32),
		},
	}

	collisionSystem := systems.CollisionSystem{
		World: &newWorld,
	}

	newWorld.AddSystem(&collisionSystem)

	newWorld.AddEntity(playerEntity)
	newWorld.AddEntity(otherPlayerEntity)

	newWorld.Update()
}
