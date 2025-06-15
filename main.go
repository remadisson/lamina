package main

import (
	"fmt"
	"os"
	"subra/parser"
)

func main() {
	data, err := os.ReadFile("lamina.dsl")
	if err != nil {
		panic(err)
	}

	cfg, err := parser.Parser.ParseString("", string(data))
	if err != nil {
		panic(err)
	}

	var root = "\"root\""
	for _, entry := range cfg.Entries {
		if entry.Zone != nil {
			if (entry.Zone.Parent) == nil {
				entry.Zone.Parent = &root
			}
			fmt.Println(entry.Zone)
		}
		if entry.Device != nil {
			fmt.Printf("Device: %s (%s) -> Zone: %s\n", entry.Device.Name, entry.Device.IP, entry.Device.Zone)
		}
	}
}
