package main

import (
	"time"

	"github.com/vdonkey/accelerator/v5/main/v2binding"
)

func main() {
	v2binding.StartAPIInstance(".")
	for {
		time.Sleep(time.Minute)
	}
}
