package client_one

import "singleton"

func IncremetAge() {
	p := singleton.GetInstance()
	p.IncremetAge()
}
