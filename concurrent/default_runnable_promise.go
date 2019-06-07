//Created by zhbinary on 2019-01-29.
//Email: zhbinary@gmail.com
package concurrent

import "github.com/zhbinary/heng/types"

type DefaultRunnablePromise struct {
	*AbstractPromise
	task types.Runnable
}

func (this *DefaultRunnablePromise) Run() {
	if this.task != nil {
		this.task()
	}
}

func NewDefaultRunnablePromise(Task types.Runnable) types.RunnablePromise {
	return &DefaultRunnablePromise{AbstractPromise: NewAbstractPromise(), task: Task}
}
