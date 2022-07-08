package objpool

import (
	"errors"
	"time"
)

type ReusableObj struct{}

type objPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(size int) *objPool {
	objPool := &objPool{}
	objPool.bufChan = make(chan *ReusableObj, size)
	for i := 0; i < size; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (p *objPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *objPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("over flow")
	}
}
