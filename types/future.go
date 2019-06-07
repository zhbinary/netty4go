//Created by zhbinary on 2018/10/18.
//Email: zhbinary@gmail.com
package types

import "time"

type Future interface {
	IsSuccess() bool
	Error() error
	AddListener(cb FutureCallback)
	RemoveListener(cb FutureCallback)
	Get() interface{}
	Get1(duration time.Duration) interface{}
	GetNow() interface{}
	Cancel() bool
	IsCancellable() bool
	IsCancelled() bool
	IsDone() bool
}

type FutureCallback func(future Future)
