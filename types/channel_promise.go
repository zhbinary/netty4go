//Created by zhbinary on 2019-01-14.
//Email: zhbinary@gmail.com
package types

type ChannelPromise interface {
	Promise
	Channel() Channel
}
