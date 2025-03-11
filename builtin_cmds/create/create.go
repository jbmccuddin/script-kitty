package create

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jbmccuddin/script-kitty/store"
)

func PrintCreateHelp() {
	fmt.Println("Create cmd help msg")
}

func CreateAlias(config store.Config) error {
	reader := bufio.NewReader(os.Stdin)
	var aliasName string

	for {
		fmt.Print("Enter a single word: ")
		aliasName, _ = reader.ReadString('\n')
		aliasName = strings.TrimSpace(aliasName) // Trim leading & trailing spaces

		if strings.Contains(aliasName, " ") {
			fmt.Println("Please enter a single-word alias.")
			continue
		}

		break
	}
	script, _ := reader.ReadString('\n')
	script = strings.TrimSpace(script)
	config.AddGlobalAlias(aliasName, script)
	return nil
}
