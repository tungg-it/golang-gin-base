package main

import (
	"tungit/cmd"
)

// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @description Example: "Bearer {token}".
func main() {
	cmd.InitApi()
}
