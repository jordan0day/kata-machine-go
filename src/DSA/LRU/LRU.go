package lru

func NewLRU[T comparable, U any](cap uint) *LRU[T, U] {
}

func (l *LRU[T, U]) Update(key T, val int) {

}

func (l *LRU[T, U]) Get(key T) (U, bool) {

}
