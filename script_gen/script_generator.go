package script_gen

import (
	"errors"
	"fmt"

	"github.com/jbmccuddin/script-kitty/store"
)

func GenerateScript(config store.Config, argv []string) (string, error) {
	if len(argv) <= 1 {
		msg := fmt.Sprintf("Invalid cmd: %s", argv)
		return "", errors.New(msg)
	}
	cmd := argv[1]
	var script string
	for _, a := range config.Projects["global"].Aliases {
		if a.Alias == cmd {
			script = script + a.Script
		}
	}
	return script, nil
}
