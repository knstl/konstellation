package types

import (
	"fmt"
	"strings"
)

const (
	FreezeIn    = "in"
	FreezeOut   = "out"
	FreezeInOut = "in-out"
)

type Freeze struct {
	In  bool `json:"in"`
	Out bool `json:"out"`
}

func NewFreeze(in, out bool) *Freeze {
	return &Freeze{In: in, Out: out}
}

func (f Freeze) String() string {
	return fmt.Sprintf(`in=%t, out=%t`, f.In, f.Out)
}

type AddressFreeze struct {
	Address string `json:"address"`
	In      bool   `json:"in"`
	Out     bool   `json:"out"`
}

func NewAddressFreeze(address string, in, out bool) *AddressFreeze {
	return &AddressFreeze{Address: address, In: in, Out: out}
}

func (af AddressFreeze) String() string {
	return fmt.Sprintf(`%s: in=%t, out=%t`, af.Address, af.In, af.Out)
}

type AddressFreezes []*AddressFreeze

func (afs AddressFreezes) String() string {
	out := fmt.Sprintf("%-44s|%-32s|%-32s\n", "Address", "Out-end-time", "In-end-time")
	for _, v := range afs {
		out += fmt.Sprintf("%-44s|%-32t|%-32t\n", v.Address, v.In, v.Out)
	}
	return strings.TrimSpace(out)
}
