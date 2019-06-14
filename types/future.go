//Created by zhbinary on 2018/10/18.
//Email: zhbinary@gmail.com
package types

import "time"

// Timeout:
// Completed: success,failed,canceled
type Future interface {
	Error() (err error)
	AddListener(cb FutureListener)
	RemoveListener(cb FutureListener)
	Get() (i interface{})
	Get0(duration time.Duration) (i interface{}, b bool)
	GetNow() (i interface{})
	Wait()
	Wait0(duration time.Duration) (b bool)
	IsSuccess() (b bool)
	IsCancelled() (b bool)
	IsDone() (b bool)
}

type FutureListener interface {
	OperationComplete(future Future)
}

type FutureListenerAdapter struct {
	OperationCompleteCb func(future Future)
}

func (this *FutureListenerAdapter) OperationComplete(future Future) {
	if this.OperationComplete != nil {
		this.OperationCompleteCb(future)
	}
}
