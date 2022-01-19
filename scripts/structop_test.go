package scripts

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestVehicle(t *testing.T) {
	v := &Vehicle0{speed: 0.0}
	assert.Panics(t, func() {
		v.Run()
	})
}

func TestCar(t *testing.T) {
	v := &Car{}
	v.SetSpeed(120)
	v.Run()
}

func TestTrain(t *testing.T) {
	v := &Train{}
	v.SetSpeed(380)
	v.Run()
}

func TestOverride(t *testing.T) {
	v := &Train{}
	v.speed = 380.0
	v.Run()
	vP := unsafe.Pointer(v)
	v2 := (*Vehicle0)(vP)
	assert.NotNil(t, v2)
	assert.Panics(t, func(){v2.Run()})
}
