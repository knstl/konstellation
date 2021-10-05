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

func NewFreeze(address, denom string, in, out bool) *Freeze {
	return &Freeze{
		Address: address,
		Denom:   denom,
		In:      in,
		Out:     out,
	}
}

/*
func (f Freeze) String() string {
	return fmt.Sprintf(`in=%t, out=%t`, f.In, f.Out)
}
*/

type Freezes []*Freeze

func (afs Freezes) String() string {
	out := fmt.Sprintf("%-44s|%-32s|%-32s\n", "Address", "Out-end-time", "In-end-time")
	for _, v := range afs {
		out += fmt.Sprintf("%-44s|%-32t|%-32t\n", v.Address, v.In, v.Out)
	}
	return strings.TrimSpace(out)
}
