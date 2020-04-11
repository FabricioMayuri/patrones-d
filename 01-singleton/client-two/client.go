package client_two

import "../singleton"

func IncremetAge() {
	p := singleton.GetInstance()
	p.IncremetAge()
}
