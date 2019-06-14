//Created by zhbinary on 2019-06-14.
//Email: zhbinary@gmail.com
package concurrent

import (
	"github.com/zhbinary/heng/types"
	"testing"
)

func TestAbstractPromise_AddListener(t *testing.T) {
	promise := NewAbstractPromise()
	cnt := 0
	promise.AddListener(&types.FutureListenerAdapter{
		OperationCompleteCb: func(future types.Future) {
			cnt++
			t.Log(future)
		},
	})
	promise.AddListener(&types.FutureListenerAdapter{
		OperationCompleteCb: func(future types.Future) {
			cnt++
			t.Log(future)
		},
	})
	if cnt != 2 {
		t.Fatalf("cnt:%d", cnt)
	}
}
