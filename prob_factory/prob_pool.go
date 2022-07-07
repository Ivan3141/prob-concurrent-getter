package prob_factory

import "sync"

type Pooler interface {
	Get() interface{}
	Put(interface{})
}

// Recycle用于数据的重复利用
type Recycler interface {
	Recycle() error
}

type RecyclerPool struct {
	pool *sync.Pool
}

func (r *RecyclerPool) Get() interface{} {
	return r.pool.Get()
}

func (r *RecyclerPool) Put(i interface{}) error {
	err := i.(Recycler).Recycle()
	if err != nil {
		return err
	}
	r.pool.Put(i)
	return nil
}

func NewRecyclerPool(fn func() Recycler) *RecyclerPool {
	return &RecyclerPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return fn()
			},
		},
	}
}
