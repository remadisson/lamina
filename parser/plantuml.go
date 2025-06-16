package parser

import (
	"fmt"
	"os"
	"strings"
)

func GeneratePlantUML(cfg *Config) string {
	var sb strings.Builder

	sb.WriteString("@startuml\n")
	content := `
!pragma layout smetana
skinparam lineType polyline
skinparam linetype ortho

skinparam backgroundColor #2A2B2C
skinparam shadowing false
skinparam componentStyle rectangle
skinparam roundcorner 9

skinparam defaultFontName "Segoe UI"
skinparam defaultFontSize 14
skinparam defaultTextAlignment center
skinparam DefaultFontColor #E0E0E0
skinparam dpi 150
skinparam rectangle {
	FontColor #E0E0E0
	BackgroundColor #2B2D31
	BorderColor #4E5056
	BorderThickness 1
}
skinparam component {
	FontColor #E0E0E0
	BackgroundColor #2B2D31
	BorderColor #4E5056
	BorderThickness 1
}
skinparam node {
	FontColor #E0E0E0
	BackgroundColor #2F3136
	BorderColor #4E5056
	BorderThickness 1
}
skinparam arrow {
	Color #5E81AC
	Thickness 2
	FontColor #D8DEE9
	FontSize 12
}`

	sb.WriteString(content + "\n")

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

	sb.WriteString("\n@enduml\n")
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
