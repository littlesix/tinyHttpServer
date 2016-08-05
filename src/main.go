package main

import (
	"pandora"
)

func main()  {
	pandora.Init("cfg.conf")
	pandora.Start()
}