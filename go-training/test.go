package main

import "training/basic"

func main() {
	// Test Data types
	d := basic.DataTypes{}
	d.TestPassMapIntoFunction()
	d.TestPassSliceIntoFunction()
	d.TestChannelType()

	// Test type conversions
	t := basic.TypeConversions{}
	t.DataLoss()
	t.Overflow()
	t.ConversionError()
	t.TypeAssertion()
}
