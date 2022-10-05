package main

import (
	"hermannvincent/deliveryservice/components"
	"hermannvincent/deliveryservice/ecs"
	"hermannvincent/deliveryservice/systems"
	"log"
)

func main() {
	newWorld := ecs.World{}

	var playerEntity ecs.IEntity = &ecs.Entity{
		Id: "Yeah",
		Components: []*ecs.IComponent{
			components.PositionComponent(1, 0),
			components.DimensionComponent(32, 32),
			components.SolidComponent(true),
		},
	}

	var otherPlayerEntity ecs.IEntity = &ecs.Entity{
		Id: "Yeah2",
		Components: []*ecs.IComponent{
			components.PositionComponent(10, 0),
			components.DimensionComponent(32, 32),
			components.SolidComponent(true),
		},
	}

	collisionSystem := systems.CollisionSystem{
		World: &newWorld,
	}
	components, err := playerEntity.GetComponents()
	if err != nil {
		log.Panicln(err)
	}
	for _, component := range components {
		localisedComponent := *component
		log.Println(localisedComponent.GetId())
	}
	log.Println(components)
	// log.Println(playerEntity.GetComponent("position"))
	newWorld.AddSystem(&collisionSystem)

	newWorld.AddEntity(&playerEntity)
	newWorld.AddEntity(&otherPlayerEntity)

	newWorld.Update()
}
