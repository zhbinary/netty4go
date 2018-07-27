package example

import (
	"github.com/zhbinary/thor"
)

func Testserver() {
	server, err := thor.Bind("locahost:7777").Option().Handler().Sync()
	if(err != nil){

	}
}
