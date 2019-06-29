//Created by zhbinary on 2019-01-14.
//Email: zhbinary@gmail.com
package types

type Promise interface {
	Future
	SetSuccess()
	SetSuccess0(i interface{})
	SetFailure(e error)
	Cancel()
}
