//Created by zhbinary on 2019-01-29.
//Email: zhbinary@gmail.com
package concurrent

import (
	"github.com/zhbinary/heng/types"
	"sync"
	"sync/atomic"
	"time"
)

const (
	PromiseStatusIncompletedNew = iota
	PromiseStatusCompletedSucceed
	PromiseStatusCompletedFailed
	PromiseStatusCompletedCanceled
)

type AbstractPromise struct {
	result    interface{}
	status    int32
	callbacks sync.Map
	once      sync.Once
	done      chan struct{}
}

func NewAbstractPromise() *AbstractPromise {
	return &AbstractPromise{done: make(chan struct{}), status: PromiseStatusIncompletedNew}
}

func (this *AbstractPromise) IsSuccess() bool {
	return this.status == PromiseStatusCompletedSucceed
}

func (this *AbstractPromise) Error() (err error) {
	if err, ok := this.result.(error); ok {
		return err
	}
	return
}

func (this *AbstractPromise) AddListener(cb types.FutureListener) {
	if cb != nil {
		this.callbacks.Store(cb, cb)
	}
}

func (this *AbstractPromise) RemoveListener(cb types.FutureListener) {
	if cb != nil {
		this.callbacks.Delete(cb)
	}
}

func (this *AbstractPromise) Get() interface{} {
	//if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedNew, PromiseStatusIncompletedWaiting) {
	//}
	this.await()
	return this.result
}

func (this *AbstractPromise) Get0(duration time.Duration) (i interface{}, b bool) {
	//if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedNew, PromiseStatusIncompletedWaiting) {
	//}
	return this.result, this.await0(duration)
}

func (this *AbstractPromise) GetNow() interface{} {
	return this.result
}

func (this *AbstractPromise) Wait() {
	this.await()
}

func (this *AbstractPromise) Wait0(duration time.Duration) (b bool) {
	return this.await0(duration)
}

func (this *AbstractPromise) Cancel() {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedNew, PromiseStatusCompletedCanceled) {
		this.notifyAllWaiters()
	}
}

func (this *AbstractPromise) IsCancellable() bool {
	return this.status == PromiseStatusIncompletedNew
}

func (this *AbstractPromise) IsCancelled() bool {
	return this.status == PromiseStatusCompletedCanceled
}

func (this *AbstractPromise) IsDone() bool {
	return this.status == PromiseStatusCompletedSucceed || this.status == PromiseStatusCompletedFailed || this.status == PromiseStatusCompletedCanceled
}

func (this *AbstractPromise) await() {
	<-this.done
}

func (this *AbstractPromise) await0(duration time.Duration) (b bool) {
	timer := time.NewTimer(duration)
	select {
	case <-this.done:
		if !timer.Stop() {
			<-timer.C
		}
		return true
	case <-timer.C:
	}
	return false
}

func (this *AbstractPromise) SetSuccess() {
	this.SetSuccess0(nil)
}

func (this *AbstractPromise) SetSuccess0(i interface{}) {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedNew, PromiseStatusCompletedSucceed) {
		this.result = i
		this.notifyAllWaiters()
		this.notifyListeners()
	}
}

func (this *AbstractPromise) SetFailure(err error) {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedNew, PromiseStatusCompletedFailed) {
		this.result = err
		this.notifyAllWaiters()
		this.notifyListeners()
	}
}

func (this *AbstractPromise) notifyAllWaiters() {
	select {
	case <-this.done:
	default:
		close(this.done)
	}
}

func (this *AbstractPromise) notifyListeners() {
	this.callbacks.Range(func(key, value interface{}) bool {
		if l, ok := value.(types.FutureListener); ok {
			l.OperationComplete(this)
		}
		return true
	})
}
