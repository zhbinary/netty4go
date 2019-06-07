//Created by zhbinary on 2019-01-14.
//Email: zhbinary@gmail.com
package types

type Promise interface {
	Future
	SetSuccess(i interface{}) bool
	SetFailure(err error) bool
	SetUncancellable() bool
}
