package main

import (
	"hermannvincent/forgottenkingdom/components"
	"hermannvincent/forgottenkingdom/ecs"
	"hermannvincent/forgottenkingdom/systems"
	"log"
)

func Start() {
	var newWorld ecs.IWorld = &ecs.World{}

	var db Database = Database{
		Entities: []ecs.ModelEntity{},
	}

	modelEntities, err := db.LoadReturnEntities()
	if err != nil {
		log.Panicln(err.Error())
	}

	for _, entityModel := range modelEntities {
		ecs.CEntityFromData(&newWorld, entityModel)
	}

	var playerEntity ecs.IEntity = &ecs.Entity{
		Id:    "Yeah",
		World: &newWorld,
		Components: []*ecs.IComponent{
			components.PositionComponent(1, 0, 0),
			components.DimensionComponent(32, 32),
			components.SolidComponent(true),
		},
	}

	var otherPlayerEntity ecs.IEntity = &ecs.Entity{
		Id: "Yeah2",
		Components: []*ecs.IComponent{
			components.PositionComponent(10, 0, 0),
			components.DimensionComponent(32, 32),
			components.SolidComponent(true),
		},
	}

	var collisionSystem ecs.ISystem = &systems.CollisionSystem{
		World: &newWorld,
	}

	for _, component := range playerEntity.GetComponents() {
		localisedComponent := *component
		log.Println(localisedComponent.GetId())
	}

	newWorld.AddSystem(&collisionSystem)

	newWorld.AddEntity(&playerEntity)
	newWorld.AddEntity(&otherPlayerEntity)

	newWorld.Update()

	var fks ForgottenKingdomServer = ForgottenKingdomServer{}

	fks.Start()
}
