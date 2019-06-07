package main

import (
	"fmt"
	"github.com/zhbinary/heng/concurrent"
	"github.com/zhbinary/heng/types"
	"time"
)

func main() {
	promise := concurrent.NewDefaultPromise()
	promise.IsDone()

	go func() {
		time.Sleep(5 * time.Second)
		//promise.SetSuccess("Fuck")
		//promise.SetFailure(errors.New("Fuck "))
		promise.SetSuccess("Haha ")
	}()
	promise.SetUncancellable()
	promise.SetUncancellable()
	promise.AddListener(&MyListener{})
	fmt.Println(promise.IsDone())
	fmt.Println(promise.IsSuccess())
	result := promise.Get()
	fmt.Println(promise.IsDone())
	fmt.Println(promise.IsSuccess())
	fmt.Println(result)
}

type MyListener struct {
}

func (this *MyListener) OperationComplete(future types.Future) {
	fmt.Printf("Result:%d\n\n", future)
}
