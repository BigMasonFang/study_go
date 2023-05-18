// file: fedex/fedex.go
package main

// this package must be main because it is build independently

type shipper struct{}

func (s shipper) Name() string {
	return "Federal Express (Fedex)"
}
func (s shipper) Currency() string {
	return "USD"
}
func (s shipper) CalculateRate(weight float32) float32 {
	cost := weight * 1.8
	tax := cost * .10
	return cost + tax
}

var Shipper shipper
