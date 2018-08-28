package main

import (
	"github.com/btcsuite/btcutil"
	"github.com/leekchan/accounting"
)

// FormatBTC formats money
func FormatBTC(amount uint64) string {
	f := btcutil.Amount(amount)
	return accounting.FormatNumberFloat64(f.ToBTC(), 8, ",", ".")
}
