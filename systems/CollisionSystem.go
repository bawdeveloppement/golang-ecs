package systems

import (
	"hermannvincent/deliveryservice/ecs"
	"log"
)

type CollisionSystem struct {
	ecs.ISystem

	Id    string
	World ecs.IWorld
}

func (c *CollisionSystem) GetId() string {
	return c.Id
}

func (c *CollisionSystem) Update() {
	for _, firstEntity := range c.World.GetEntities() {
		for _, secondEntity := range c.World.GetEntities() {
			if firstEntity.GetId() != secondEntity.GetId() {

				firstPositionComponentPointer, err := firstEntity.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				firstPositionComponent := *firstPositionComponentPointer

				var firstX int = firstPositionComponent.GetData("x").(int)
				var firstY int = firstPositionComponent.GetData("y").(int)

				firstDimensionComponentPointer, err := firstEntity.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				firstDimensionComponent := *firstDimensionComponentPointer

				var firstWidth int = firstDimensionComponent.GetData("width").(int)
				var firstHeight int = firstDimensionComponent.GetData("height").(int)

				secondPositionComponentPointer, err := secondEntity.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				secondPositionComponent := *secondPositionComponentPointer

				var secondX int = secondPositionComponent.GetData("x").(int)
				var secondY int = secondPositionComponent.GetData("y").(int)

				secondDimensionComponentPointer, err := secondEntity.GetComponent("position")
				if err != nil {
					log.Panicln(err)
				}
				secondDimensionComponent := *secondDimensionComponentPointer

				var secondWidth int = secondDimensionComponent.GetData("width").(int)
				var secondHeight int = secondDimensionComponent.GetData("height").(int)

				if firstX+firstWidth > secondX && firstX < secondX+secondWidth && firstY+firstHeight > secondY && firstY < secondY+secondHeight {
					log.Println(firstEntity.GetId() + " in collision with " + secondEntity.GetId())
				}
			}
		}
	}
}
