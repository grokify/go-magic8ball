package main

import (
	"fmt"

	"github.com/grokify/go-magic8ball/magic"
)

func main() {
	for i := 0; i < 100; i++ {
		resp, err := magic.Shake()
		if err != nil {
			panic(err)
		}
		respType, err := magic.ResponseType(resp)
		if err != nil {
			panic(err)
		}
		outlook := "positive"
		if respType == 0 {
			outlook = "neutral"
		} else if respType < 0 {
			outlook = "negative"
		}
		fmt.Printf("%s (%s)\n", resp, outlook)
	}
}
