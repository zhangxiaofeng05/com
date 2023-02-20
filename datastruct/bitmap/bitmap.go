package bitmap

type BitMap struct {
	capacity uint64
	bits     []byte
}

func New(capacity uint64) *BitMap {
	return &BitMap{
		capacity: capacity,
		bits:     make([]byte, (capacity>>3)+1),
	}
}

func (b *BitMap) Set(num uint32) {
	//index = num/8
	arrayIndex := num >> 3
	position := num & 0x07
	b.bits[arrayIndex] |= 1 << position
}

func (b *BitMap) Contains(num uint32) bool {
	arrayIndex := num >> 3
	position := num & 0x07
	return (b.bits[arrayIndex] & (1 << position)) != 0
}

func (b *BitMap) Remove(num uint32) {
	arrayIndex := num >> 3
	position := num & 0x07
	b.bits[arrayIndex] &= ^(1 << position)
}
