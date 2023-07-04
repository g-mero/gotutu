package main

import (
	"github.com/g-mero/gotutu/handle"
	"github.com/g-mero/gotutu/routes"
	"log"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
}

func main() {
	handle.InitStorages()
	routes.InitRouter()
}
