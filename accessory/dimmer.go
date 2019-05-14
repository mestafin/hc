package accessory

import (
	"github.com/brutella/hc/service"
)

type Dimmer struct {
	*Accessory
	Dimmer *service.Dimmer
}

// NewDimmer returns a dimmer which implements model.Dimmer.
func NewDimmer(info Info) *Dimmer {
	acc := Dimmer{}
	acc.Accessory = New(info, TypeDimmer)
	acc.Dimmer = service.NewDimmer()
	acc.AddService(acc.Dimmer.Service)

	return &acc
}
