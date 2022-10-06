package ecs

type IItem interface {
	GetId() string
	GetName() string
	GetDescription() string
	Use(*IWorld) string
}

const (
	SLOT_LEFTHAND  = "SLOT_LEFTHAND"
	SLOT_RIGHTHAND = "SLOT_RIGHTHAND"
	SLOT_CHEST     = "SLOT_CHEST"
	SLOT_FOOT      = "SLOT_FOOT"
)

type Item struct {
	Id          string
	Name        string
	Description string
	SlotType    string
}

func (item *Item) GetId() string {
	return item.Id
}

func (item *Item) GetName() string {
	return item.Name
}

func (item *Item) GetDescription() string {
	return item.Description
}
