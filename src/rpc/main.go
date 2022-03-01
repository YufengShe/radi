package rpc

import (
	"log"
	ccapi "radi/rpc/kitex_gen/ccAPI/ccoperation"
)

func main() {
	svr := ccapi.NewServer(new(CCOperationImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
