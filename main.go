package main

import (
	"fmt"
	"os"

	"github.com/jbmccuddin/script-kitty/builtin_cmds"
	"github.com/jbmccuddin/script-kitty/builtin_cmds/activate"
	"github.com/jbmccuddin/script-kitty/builtin_cmds/create"
	"github.com/jbmccuddin/script-kitty/script_gen"
	"github.com/jbmccuddin/script-kitty/store"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		printHelp()
		return
	}

	cmd := args[1]
	var config store.Config
	var err error
	config, err = store.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if builtin_cmds.ACTIVATE == cmd {
		if len(args) <= 2 {
			printCmdHelp(cmd)
		}
	} else if builtin_cmds.CREATE == cmd {
		if len(args) <= 2 {
			printCmdHelp(cmd)
		}
		switch args[2] {
		case create.ALIAS:
			create.CreateAlias(config)
			return
		}
	} else if config.ContainsAlias(cmd) {
		var result string
		result, err = script_gen.GenerateScript(config, args)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(result)
	} else {
		printHelp()
	}
}

func printCmdHelp(cmd string) {
	switch cmd {
	case builtin_cmds.ACTIVATE:
		activate.PrintActivateHelp()
	case builtin_cmds.CREATE:
		create.PrintCreateHelp()
	}
}

func printHelp() {
	fmt.Println("HELP:")
	for _, cmd := range builtin_cmds.BUILTIN_CMDS {
		printCmdHelp(cmd)
	}
}
