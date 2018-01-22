package main

import (
	App "github.com/velrino/RedFull/app"
)

func main() {
	App.DatabaseInit()
	App.Routes()
}