package scripts

import "fmt"

// Vehicle 交通工具
type Vehicle interface {
	GetSpeed() float32
	SetSpeed(speed float32)
	Run()
}

type Vehicle0 struct {
	speed float32
	Vehicle
}

func (v *Vehicle0) GetSpeed() float32 {
	return v.speed
}

func (v *Vehicle0) SetSpeed(speed float32) {
	v.speed = speed
}

func (v *Vehicle0) String() string {
	return fmt.Sprintf("%#v", *v)
}

func (v *Vehicle0) Run() {
	// 抛出异常
	panic("not impl")
}


// Car 汽车
type Car struct {
	Vehicle0
}

func (c *Car) Run() {
	fmt.Println(fmt.Sprintf("%T running, speed=%f", c, c.GetSpeed()))
}

// Train 火车
type Train struct {
	Vehicle0
}

func (t *Train) Run() {
	fmt.Println(fmt.Sprintf("%T running, speed=%f", t, t.GetSpeed()))
}