// echo -e "a\nb\nc" | ./myXargs ls -la
// ./myFind -f -ext 'txt' . | ./myXargs ./myWc -l

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var pipeArgs []string
	for scanner.Scan() {
		str := scanner.Text()
		pipeArgs = append(pipeArgs, str)
	}

	var command string
	var args []string
	cmdAndArgs := os.Args[1:]
	for i, arg := range cmdAndArgs {
		if i == 0 {
			command = arg
		} else {
			args = append(args, arg)
		}
	}

	args = append(args, pipeArgs...)

	cmd := exec.Command(command, args...)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdout))
		log.Fatal(err)
	}
	str := string(stdout)
	fmt.Printf("%s", str)

}
