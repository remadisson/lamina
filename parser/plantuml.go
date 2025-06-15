package parser

import (
	"fmt"
	"os"
	"strings"
)

func GeneratePlantUML(cfg *Config) string {
	var sb strings.Builder
	sb.WriteString("@startuml\n")
	sb.WriteString("skinparam componentStyle rectangle\n")

	// Zonen als Komponenten
	for _, entry := range cfg.Entries {
		if entry.Zone != nil {
			z := entry.Zone
			sb.WriteString(fmt.Sprintf("component \"%s\\n%s\" as %s\n", z.Name, z.CIDR, z.Name))
			if z.Parent != nil && *z.Parent != "root" {
				sb.WriteString(fmt.Sprintf("%s --> %s\n", z.Name, *z.Parent))
			}
		}
	}

	for _, entry := range cfg.Entries {
		if entry.Device != nil {
			d := entry.Device
			sb.WriteString(fmt.Sprintf("node \"%s\\n%s\" as %s\n", d.Name, d.IP, d.Name))
			sb.WriteString(fmt.Sprintf("%s --> %s\n", d.Name, d.Zone))
		}
	}

	sb.WriteString("@enduml\n")
	return sb.String()
}

func WriteToFile(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
