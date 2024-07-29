package main

import (
	"fmt"
	"gopl.io/ch2/conv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		d, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mm: %v\n", err)
			os.Exit(1)
		}

		miles := conv.Miles(d)
		meters := conv.MilesToMeters(miles)

		fmt.Printf("%v = %v", miles, meters)
	}
}
