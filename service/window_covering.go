// THIS FILE IS AUTO-GENERATED
// Modified by mestafin

package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeWindowCovering = "8C"

type WindowCovering struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState
	
	// Modified by mestafin
	HoldPosition *characteristic.HoldPosition
	CurrentVerticalTiltAngle *characteristic.CurrentVerticalTiltAngle
	//XX *characteristic.TargetVerticalTiltAngle
	TargetVerticalTiltAngle *characteristic.TargetVerticalTiltAngle
}

func NewWindowCovering() *WindowCovering {
	svc := WindowCovering{}
	svc.Service = New(TypeWindowCovering)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)
	
	// Modified by mestafin
	svc.HoldPosition = characteristic.NewHoldPosition()
	svc.AddCharacteristic(svc.HoldPosition.Characteristic)

	svc.CurrentVerticalTiltAngle = characteristic.NewCurrentVerticalTiltAngle()
	svc.AddCharacteristic(svc.CurrentVerticalTiltAngle.Characteristic)

	//svc.XX = characteristic.NewTargetVerticalTiltAngle()
	svc.TargetVerticalTiltAngle = characteristic.NewTargetVerticalTiltAngle()
	svc.AddCharacteristic(svc.TargetVerticalTiltAngle.Characteristic)

	return &svc
}
