package main

import (
	App "./app"
)

func main() {
	App.DatabaseInit()
	App.Routes()
}