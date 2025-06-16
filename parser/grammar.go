package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Config struct {
	Entries []*Entry `{ @@ }`
}

type Entry struct {
	Zone   *Zone   `  @@`
	Device *Device `| @@`
}

type Zone struct {
	Name        string  `"zone" @String`
	CIDR        string  `"{" "cidr" "=" @String`
	VLAN        int     `"vlan" "=" @Int`
	Description string  `"description" "=" @String`
	Parent      *string `( "parent" "=" @String )? "}"`
}

func (z Zone) String() string {
	parent := "<nil>"
	if z.Parent != nil {
		parent = *z.Parent
	}
	return fmt.Sprintf("Zone{Name: %s, CIDR: %s, VLAN: %d, Description: %s, Parent: %s}",
		z.Name, z.CIDR, z.VLAN, z.Description, parent)
}

type Device struct {
	Name string `"device" @String`
	IP   string `"{" "ip" "=" @String`
	MAC  string `"mac" "=" @String`
	Zone string `"zone" "=" @String "}"`
}

func unquoteToken(tok lexer.Token) lexer.Token {
	if tok.Type == 2 {
		raw := tok.Value
		if len(raw) >= 2 && raw[0] == '"' && raw[len(raw)-1] == '"' {
			tok.Value = raw[1 : len(raw)-1]
		}
	}
	return tok
}

var (
	laminaLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Comment", `(?:#|//)[^\n]*`},
		{"String", `"(\\"|[^"])*"`},
		{"Ident", `[a-zA-Z_][a-zA-Z0-9_-]*`},
		{"Int", `[0-9]+`},
		{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
		{"Whitespace", `[ \t\n\r]+`},
	})
	Parser = participle.MustBuild[Config](
		participle.Lexer(laminaLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.Unquote("String"),
		participle.UseLookahead(2),
	)
)
