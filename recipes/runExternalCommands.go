// Run External Commands

/* Largely uses the exec function. */
package main

import(
	"os/exec"
	"fmt"	
)

/* Variables within the recipe that you can change. */
var (
        commandName = "echo"
	commandArg = "Hello, world!"
)

func main() {
	var err error
	
	/* Looks for executabe binary in dir given by $PATH. If / contained,
	   path is not consulted. Result could be absolute or relative to 
           current directory.  */
	path, err := exec.LookPath(commandName); if err != nil {
		fmt.Printf("\nERROR: %v\n", err)
	} else {
		fmt.Printf("\nBinary found. Path: %s\n", path)
	}

	/* exec.CMD is returned, it's a struct which contains info on the 
           command that will be run. Very useful and worth looking into. 
           The command args is variadic. */
	cmd := exec.Command(commandName, commandArg)

	/* Run the command and capture stdin and stdout, then return it. */
	cmdOutput, err := cmd.CombinedOutput(); if err != nil {
		fmt.Printf("\nERROR: %v\n", err)
	}

	fmt.Print("\nCommand Run: ", commandName,
		  "\nCommand argument(s): ", commandArg,
		  "\nCommand output: ", string(cmdOutput))

	return 
}

