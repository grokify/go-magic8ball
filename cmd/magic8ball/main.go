package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/grokify/go-magic8ball"
)

func main() {
	max := 1
	if len(os.Args) > 1 {
		maxInt, err := strconv.Atoi(os.Args[1])
		if err == nil && maxInt > 1 {
			max = maxInt
		}
	}

	for i := 1; i <= max; i++ {
		resp, err := magic8ball.Shake()
		if err != nil {
			panic(err)
		}
		respType, err := magic8ball.ResponseType(resp)
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
