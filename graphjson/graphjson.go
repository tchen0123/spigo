// functions for writing graphml

package graphjson

import (
	"fmt"
	"os"
)

var Enabled bool
var file *os.File
var edgeid int // unique id for each edge
var comma bool

// write the header to the file
func Setup(arch string) {
	if Enabled == false {
		return
	}
	file, _ = os.Create("spigo.json")
	Write(fmt.Sprintf("{\n  \"arch\":\"%v\",\n  \"version\":\"spigo-0.3\",\n  \"args\":\"%v\",\n  \"graph\":[", arch, os.Args))
	comma = false
}

func Write(str string) {
	file.WriteString(str)
}

func commaNewline() string {
	if comma {
		return ",\n"
	} else {
		comma = true
		return "\n"
	}
}

func WriteNode(name, service string) {
	if Enabled == false {
		return
	}
	// node id should be unique and service indicates service type
	Write(fmt.Sprintf("%v    { \"node\":\"%v\", \"service\":\"%v\" }", commaNewline(), name, service))
}

func Edge(from, to string) string {
	if Enabled == false {
		return ""
	}
	edgeid = edgeid + 1
	return fmt.Sprintf("%v    { \"edge\":\"e%v\", \"source\":\"%v\", \"target\":\"%v\" }", commaNewline(), edgeid, from, to)
}

func WriteEdge(from, to string) {
	if Enabled == false {
		return
	}
	Write(Edge(from, to))
}

func Close() {
	if Enabled == false {
		return
	}
	Write("\n  ]\n}\n")
	file.Close()
}
