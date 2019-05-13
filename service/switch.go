// THIS FILE IS AUTO-GENERATED
// mestafin added
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSwitch = "49"

type Switch struct {
	*Service

	On *characteristic.On
	Brightness *characteristic.Brightness
}

func NewSwitch() *Switch {
	svc := Switch{}
	svc.Service = New(TypeSwitch)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	// mestafin added
	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	return &svc
}
