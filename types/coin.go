package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DARC  = "darc"  // 1 (base denom unit)
	MDARC = "mdarc" // 10^-3 (milli)
	UDARC = "udarc" // 10^-6 (micro)
	NDARC = "ndarc" // 10^-9 (nano)
	PDARC = "pdarc" // 10^-12 (pico)
	FDARC = "fdarc" // 10^-15 (femto)
	ADARC = "adarc" // 10^-18 (atto)

	KNSTL  = "knstl"  // 1 (base denom unit)
	MKNSTL = "mknstl" // 10^-3 (milli)
	UKNSTL = "uknstl" // 10^-6 (micro)
	NKNSTL = "nknstl" // 10^-9 (nano)
	PKNSTL = "pknstl" // 10^-12 (pico)
	FKNSTL = "fknstl" // 10^-15 (femto)
	AKNSTL = "aknstl" // 10^-18 (atto)

	DefaultBondDenom      = UDARC
	StakeDenom            = UDARC
	DefaultConsensusPower = 100000
)

func RegisterNativeCoinUnits() {
	_ = sdk.RegisterDenom(DARC, sdk.OneDec())
	_ = sdk.RegisterDenom(MDARC, sdk.NewDecWithPrec(1, 3))
	_ = sdk.RegisterDenom(UDARC, sdk.NewDecWithPrec(1, 6))
	// Base denom is darc
	//_ = sdk.RegisterDenom(NDARC, sdk.NewDecWithPrec(1, 9))
	//_ = sdk.RegisterDenom(PDARC, sdk.NewDecWithPrec(1, 12))
	//_ = sdk.RegisterDenom(FDARC, sdk.NewDecWithPrec(1, 15))
	//_ = sdk.RegisterDenom(ADARC, sdk.NewDecWithPrec(1, 18))
}
