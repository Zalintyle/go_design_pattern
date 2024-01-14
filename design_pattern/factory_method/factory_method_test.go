package factory_method

import (
	"fmt"
	"testing"
)

func TestMakeClothes(t *testing.T) {
	anta, _ := MakeClothes("ANTA")
	peak, _ := MakeClothes("PEAK")

	printDetails(anta)
	printDetails(peak)
}

func printDetails(c IClothes) {
	fmt.Printf("Clothes name: %s, Clothes size: %d\n", c.GetName(), c.GetSize())
}
