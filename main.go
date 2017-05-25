package main

import (
	"log"
	"os"
	"strings"

	"github.com/THE108/enumlinter/logger"

	"golang.org/x/tools/go/loader"
)

const (
	typeToCheckFlagName = "-type"
)

func parseArgs() (typeToCheck string, pkgpaths []string, ok bool) {
	if len(os.Args) < 3 {
		return
	}

	if os.Args[1] == typeToCheckFlagName {
		typeToCheck = os.Args[2]
		pkgpaths = os.Args[3:]
		ok = true
		return
	}

	if prefix := typeToCheckFlagName + "="; strings.HasPrefix(os.Args[1], prefix) {
		typeToCheck = strings.TrimPrefix(os.Args[1], prefix)
		pkgpaths = os.Args[2:]
		ok = true
		return
	}
	return
}

func main() {
	typeToCheck, pkgpaths, ok := parseArgs()
	if !ok {
		log.Fatalf("Usage: enumlinter -type <enumerated type> <args>\n%s",
			"<enumerated type> is an enumerated type to check."+loader.FromArgsUsage)
	}

	if err := run(pkgpaths, typeToCheck, logger.NewLogger(os.Stdout)); err != nil {
		log.Fatalf("Can't load prog from args: %s", err.Error())
	}
}
