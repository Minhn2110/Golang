package main

import (
	"libraries/cmd"
	"libraries/util"
)

func main() {
	// util.Flag()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
	util.Logger()
	util.Viper()

}
