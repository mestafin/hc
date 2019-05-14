// THIS FILE IS AUTO-GENERATED
// mestafin added
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeDimmer = "249"

type Dimmer struct {
	*Service

	On *characteristic.On
	Brightness *characteristic.Brightness
}

func NewDimmer() *Dimmer {
	svc := Dimmer{}
	svc.Service = New(TypeDimmer)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	// mestafin added
	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	return &svc
}
