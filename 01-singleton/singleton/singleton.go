package singleton

var p *person

func GetInstance() *person {
	if p == nil {
		p = &person{}
	}

	return p
}

type person struct {
	name string
	age  int
}

func (p *person) SetName(n string) {
	p.name = n
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) IncremetAge() {
	p.age++
}

func (p *person) GetAge() int {
	return p.age
}
