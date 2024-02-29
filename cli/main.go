package cli

import (
	"fmt"

	"github.com/JubaerHossain/clid/cli/create"
	"github.com/spf13/cobra"
)

var asciiArt = `
________  ___       ___  ________     
|\   ____\|\  \     |\  \|\   ___ \    
\ \  \___|\ \  \    \ \  \ \  \_|\ \   
 \ \  \    \ \  \    \ \  \ \  \ \\ \  
  \ \  \____\ \  \____\ \  \ \  \_\\ \ 
   \ \_______\ \_______\ \__\ \_______\
    \|_______|\|_______|\|__|\|_______|
  
  A CLI tool to building restful API with Go
`

// Create a Cobra command for clid
var command = &cobra.Command{
	Use:   "clid",
	Short: asciiArt,
}

func AddCommand(command *cobra.Command) {
	command.AddCommand(command)
}

func init() {
	command.AddCommand(create.Create)
}

// run is the main function to execute the clid command
func Run() {
	err := command.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
