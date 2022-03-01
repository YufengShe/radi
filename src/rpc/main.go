package rpc

import (
	"log"
	ccoperate "radi/rpc/kitex_gen/ccAPI/ccoperation"
)

func main() {
	svr := ccoperate.NewServer(new(CCOperationImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
