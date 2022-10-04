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
				var firstX int = firstEntity.GetComponent("position").GetData("x").(int)
				var firstY int = firstEntity.GetComponent("position").GetData("y").(int)
				var firstWidth int = firstEntity.GetComponent("dimension").GetData("width").(int)
				var firstHeight int = firstEntity.GetComponent("dimension").GetData("height").(int)
				var secondX int = secondEntity.GetComponent("position").GetData("x").(int)
				var secondY int = secondEntity.GetComponent("position").GetData("y").(int)
				var secondWidth int = secondEntity.GetComponent("dimension").GetData("width").(int)
				var secondHeight int = secondEntity.GetComponent("dimension").GetData("height").(int)
				if firstX+firstWidth > secondX && firstX < secondX+secondWidth && firstY+firstHeight > secondY && firstY < secondY+secondHeight {
					log.Println(firstEntity.GetId() + " in collision with " + secondEntity.GetId())
				}
			}
		}
	}
}
