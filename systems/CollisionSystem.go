package systems

import (
	"hermannvincent/forgottenkingdom/ecs"
	"log"
)

type CollisionSystem struct {
	Id    string
	World *ecs.IWorld
}

func (c *CollisionSystem) GetId() string {
	return c.Id
}

func (c *CollisionSystem) Update() {
	var world ecs.IWorld = *c.World
	for _, firstEntity := range world.GetEntitiesWithComponents("position", "dimension", "solid") {
		firstEntityLocalised := *firstEntity
		for _, secondEntity := range world.GetEntitiesWithComponents("position", "dimension", "solid") {
			secondEntityLocalised := *secondEntity
			if firstEntityLocalised.GetId() != secondEntityLocalised.GetId() {
				firstPositionComponentPointer, err := firstEntityLocalised.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				firstPositionComponent := *firstPositionComponentPointer
				firstX := firstPositionComponent.GetData("x").(int)
				firstY := firstPositionComponent.GetData("y").(int)

				firstDimensionComponentPointer, err := firstEntityLocalised.GetComponent("dimension")
				if err != nil {
					log.Panicln(err)
				}

				firstDimensionComponent := *firstDimensionComponentPointer
				firstWidth := firstDimensionComponent.GetData("width").(int)
				firstHeight := firstDimensionComponent.GetData("height").(int)

				secondPositionComponentPointer, err := secondEntityLocalised.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				secondPositionComponent := *secondPositionComponentPointer

				var secondX int = secondPositionComponent.GetData("x").(int)
				var secondY int = secondPositionComponent.GetData("y").(int)

				secondDimensionComponentPointer, err := secondEntityLocalised.GetComponent("dimension")
				if err != nil {
					log.Panicln(err)
				}
				secondDimensionComponent := *secondDimensionComponentPointer

				var secondWidth int = secondDimensionComponent.GetData("width").(int)
				var secondHeight int = secondDimensionComponent.GetData("height").(int)

				if firstX+firstWidth > secondX && firstX < secondX+secondWidth && firstY+firstHeight > secondY && firstY < secondY+secondHeight {
					log.Println(firstEntityLocalised.GetId() + " in collision with " + secondEntityLocalised.GetId())
				}
			}
		}
	}
}
