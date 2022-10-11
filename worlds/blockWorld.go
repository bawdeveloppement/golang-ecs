package worlds

import (
	"hermannvincent/forgottenkingdom/ecs"
	"hermannvincent/forgottenkingdom/utils"
)

type IBlockWorld interface {
	ecs.IWorld

	AddBlock(*utils.IBlock) error
	GetBlock(string) *utils.IBlock
	GetBlocks() []*utils.IBlock
	GetBlockByPosition(utils.Vec3f) *utils.IBlock
}
