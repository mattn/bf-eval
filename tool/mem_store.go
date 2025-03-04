package tool

type MemStore []uint8

func (m MemStore) Get(ptr int) uint8 {
	return m[ptr+2048]
}

func (m MemStore) Set(ptr int, v uint8) {
	m[ptr+2048] = v
}

func (m MemStore) AddTo(ptr int, v uint8) {
	m[ptr+2048] += v
}

func (m MemStore) SubFrom(ptr int, v uint8) {
	m[ptr+2048] -= v
}

func NewMemStore() MemStore {
	return make(MemStore, 4096)
}
