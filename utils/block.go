package utils

type IBlock interface {
	GetId() string
	GetPosition() Vec3f
	GetRotation() Vec3f
	GetDimension() Vec3f
	GetTexture() string
}

type Block struct {
	Id        string
	Position  Vec3f
	Rotation  Vec3f
	Dimension Vec3f
	Texture   string
}

func (b *Block) GetId() string {
	return b.Id
}

func (b *Block) GetPosition() Vec3f {
	return b.Position
}

func (b *Block) GetRotation() Vec3f {
	return b.Rotation
}

func (b *Block) GetDimension() Vec3f {
	return b.Dimension
}

func (b *Block) GetTexture() string {
	return b.Texture
}
