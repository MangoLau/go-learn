package objpool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjpool(t *testing.T) {
	pool := NewObjPool(10)
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil { //尝试放置超出池大小的对象
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(1 * time.Millisecond); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("done")
}
