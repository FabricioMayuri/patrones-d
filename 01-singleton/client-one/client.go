package client_one

import "github.com/FabricioMayuri/patrones-d/tree/master/01-singleton/singleton/singleton.go"

func IncremetAge() {
	p := singleton.GetInstance()
	p.IncremetAge()
}
