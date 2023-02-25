package core

type Storage interface {
	Put(*Block) error
}

type MemoryStore struct {
}

func NewMemoryStote() *MemoryStore {
	return &MemoryStore{}
}

func (ms *MemoryStore) Put(b *Block) error {
	return nil
}
