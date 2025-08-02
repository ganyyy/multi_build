package main

import (
	"B/common"
	"SB/comp"
)

func main() {
	common.InitServer("SB" + comp.GenRandomSuffix())
}
