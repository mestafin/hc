// THIS FILE IS AUTO-GENERATED
// modified by mestafin

package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeAccessoryInformation = "3E"

type AccessoryInformation struct {
	*Service

	Identify         *characteristic.Identify
	Manufacturer     *characteristic.Manufacturer
	Model            *characteristic.Model
	Name             *characteristic.Name
	SerialNumber     *characteristic.SerialNumber
	FirmwareRevision *characteristic.FirmwareRevision
	
	// added by mestafin
	HardwareRevision *characteristic.HardwareRevision
	Version *characteristic.Version
}

func NewAccessoryInformation() *AccessoryInformation {
	svc := AccessoryInformation{}
	svc.Service = New(TypeAccessoryInformation)

	svc.Identify = characteristic.NewIdentify()
	svc.AddCharacteristic(svc.Identify.Characteristic)

	svc.Manufacturer = characteristic.NewManufacturer()
	svc.AddCharacteristic(svc.Manufacturer.Characteristic)

	svc.Model = characteristic.NewModel()
	svc.AddCharacteristic(svc.Model.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.SerialNumber = characteristic.NewSerialNumber()
	svc.AddCharacteristic(svc.SerialNumber.Characteristic)

	svc.FirmwareRevision = characteristic.NewFirmwareRevision()
	svc.AddCharacteristic(svc.FirmwareRevision.Characteristic)
	
	// added by mestafin
	svc.HardwareRevision = characteristic.NewHardwareRevision()
	svc.AddCharacteristic(svc.HardwareRevision.Characteristic)
	
	svc.Version = characteristic.NewVersion()
	svc.AddCharacteristic(svc.Version.Characteristic)
	

	return &svc
}
