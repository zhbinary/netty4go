//Created by zhbinary on 2019-01-29.
//Email: zhbinary@gmail.com
package concurrent

import (
	"context"
	"github.com/zhbinary/heng/types"
	"sync"
	"sync/atomic"
	"time"
)

const (
	PromiseStatusIncompletedCancelable = iota
	PromiseStatusIncompletedUncancelable
	PromiseStatusIncompletedWaiting
	PromiseStatusCompletedSucceed
	PromiseStatusCompletedFailed
	PromiseStatusCompletedCanceled
)

type AbstractPromise struct {
	result    interface{}
	status    int32
	done      chan interface{}
	cancel    context.CancelFunc
	callbacks sync.Map
	once      sync.Once
}

func NewAbstractPromise() *AbstractPromise {
	return &AbstractPromise{done: make(chan interface{}), status: PromiseStatusIncompletedCancelable}
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

func (this *AbstractPromise) AddListener(listener types.FutureCallback) {
	this.callbacks.Store(listener, listener)
}

func (this *AbstractPromise) RemoveListener(listener types.FutureCallback) {
	this.callbacks.Delete(listener)
}

func (this *AbstractPromise) Get() interface{} {
	this.await()
	return this.result
}

func (this *AbstractPromise) Get1(duration time.Duration) interface{} {
	this.await0(duration)
	return this.result
}

func (this *AbstractPromise) GetNow() interface{} {
	return this.result
}

func (this *AbstractPromise) Cancel() (b bool) {
	if this.cancel != nil && atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedCancelable, PromiseStatusCompletedCanceled) {
		this.cancel()
		this.notifyAll()
		return true
	}
	return
}

func (this *AbstractPromise) SetUncancellable() bool {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedCancelable, PromiseStatusIncompletedUncancelable) {
		return true
	}
	return !this.IsDone() || !this.IsCancelled()
}

func (this *AbstractPromise) IsCancellable() bool {
	return this.status == PromiseStatusIncompletedCancelable
}

func (this *AbstractPromise) IsCancelled() bool {
	return this.status == PromiseStatusCompletedCanceled
}

func (this *AbstractPromise) IsDone() bool {
	return this.status == PromiseStatusCompletedSucceed || this.status == PromiseStatusCompletedFailed || this.status == PromiseStatusCompletedCanceled
}

func (this *AbstractPromise) await() {
	this.once.Do(func() {
		this.await0(-1)
	})
}

func (this *AbstractPromise) await0(duration time.Duration) {
	if this.IsDone() {
		return
	}

	var ctx context.Context
	if duration == -1 {
		ctx, this.cancel = context.WithCancel(context.TODO())
	} else {
		ctx, this.cancel = context.WithTimeout(context.TODO(), duration)
	}

	<-ctx.Done()
	this.nofityListeners()
}

func (this *AbstractPromise) SetSuccess(i interface{}) (b bool) {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedCancelable, PromiseStatusCompletedSucceed) ||
		atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedUncancelable, PromiseStatusCompletedSucceed) {
		this.result = i
		this.notifyAll()
		this.nofityListeners()
		return true
	}
	return
}

func (this *AbstractPromise) SetFailure(err error) (b bool) {
	if atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedCancelable, PromiseStatusCompletedFailed) ||
		atomic.CompareAndSwapInt32(&this.status, PromiseStatusIncompletedUncancelable, PromiseStatusCompletedFailed) {
		this.result = err
		this.notifyAll()
		this.nofityListeners()
		return true
	}
	return
}

func (this *AbstractPromise) notifyAll() {
	if this.cancel != nil {
		this.cancel()
	}
}

func (this *AbstractPromise) nofityListeners() {
	this.callbacks.Range(func(key, value interface{}) bool {
		if cb, ok := value.(types.FutureCallback); ok {
			cb(this)
		}
		return true
	})
}
