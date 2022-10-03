package main

import (
	"hermannvincent/deliveryservice/ecs"
	"hermannvincent/deliveryservice/systems"
)

func main() {
	newWorld := ecs.World{}

	var playerEntity ecs.IEntity = ecs.Entity{Id: "Yeah"}
	var otherPlayerEntity ecs.IEntity = ecs.Entity{Id: "Yeah", Components: []ecs.IComponent{}}

	collisionSystem := systems.CollisionSystem{
		World: newWorld,
	}

	newWorld.AddSystem(collisionSystem)
	newWorld.AddEntity(playerEntity)

	newWorld.Update()
}
