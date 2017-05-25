package main

import (
	"go/ast"

	"github.com/THE108/enumlinter/checker"
	lg "github.com/THE108/enumlinter/logger"
	ph "github.com/THE108/enumlinter/position_helper"

	"golang.org/x/tools/go/loader"
)

func run(pkgpaths []string, typeToCheck string, logger lg.ILogger) error {
	prog, err := loadProgFromArgs(pkgpaths)
	if err != nil {
		return err
	}

	posHelper := ph.NewPositionHelper(prog.Fset)

	for _, pkgInfo := range prog.InitialPackages() {
		switchChecker := &checker.SwitchChecker{
			Pkg:       pkgInfo,
			PosHelper: posHelper,
			Logger:    logger,
		}

		assignmentChecker := &checker.AssignmentChecker{
			Pkg:         pkgInfo,
			PosHelper:   posHelper,
			Logger:      logger,
			TypeToCheck: typeToCheck,
		}

		for _, astfile := range pkgInfo.Files {
			ast.Inspect(astfile, func(node ast.Node) bool {
				switchChecker.Run(node)
				return assignmentChecker.Run(node)
			})
		}
	}
	return nil
}

func loadProgFromArgs(pkgpaths []string) (*loader.Program, error) {
	conf := &loader.Config{}
	if _, err := conf.FromArgs(pkgpaths, true); err != nil {
		return nil, err
	}

	return conf.Load()
}
