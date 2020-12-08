package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filename = flag.String("f", "input-data", "path of file to load")
)

func splitFunc(c rune) bool {
	return c == '\n'
}

func readDataArr(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	strData := string(b)
	return strings.FieldsFunc(strData, splitFunc)
}

type Cmd struct {
	accum int
	step  int
	index int
	val   int
	t     string
}

func parseCmd(command string, i int) Cmd {
	cmdFields := strings.Fields(command)
	val, _ := strconv.Atoi(cmdFields[1])
	switch f := cmdFields[0]; f {
	case "nop":
		return Cmd{0, 1, i, val, "nop"}
	case "acc":
		return Cmd{val, 1, i, val, "acc"}
	case "jmp":
		return Cmd{0, val, i, 1, "jmp"}
	default:
		fmt.Println("Bad command", command)
	}
	return Cmd{}

}

func nextCmd(commands []string, cmdLoop []Cmd) (bool, Cmd) {
	lastCmd := cmdLoop[len(cmdLoop)-1]
	nextIndex := lastCmd.index + lastCmd.step
	nextC := parseCmd(commands[nextIndex], nextIndex)
	for _, c := range cmdLoop {
		if c.index == nextC.index {
			return false, Cmd{}
		}
		if c.index == len(commands) {
			return false, Cmd{}
		}
	}
	return true, nextC
}

func findLoop(instructions []string, start int) []Cmd {
	var loop []Cmd

	nextCommand := parseCmd(instructions[start], start)
	newCommand := true
	for newCommand == true {
		loop = append(loop, nextCommand)
		newCommand, nextCommand = nextCmd(instructions, loop)
	}
	return loop
}

func completeLoop(cmds []string, loop []Cmd, targetIndex int) (bool, []Cmd) {
	lc := loop[len(loop)-1]
	startInd := lc.index + lc.step
	nextCommand := parseCmd(cmds[startInd], startInd)
	newCommand := true
	for nextCommand.index != targetIndex {
		loop = append(loop, nextCommand)
                if nextCommand.index + nextCommand.step >= len(cmds) {
                  return true, loop
                }
		newCommand, nextCommand = nextCmd(cmds, loop)
		if newCommand == false {
			return false, loop
		}
                fmt.Println("hi", nextCommand.index, targetIndex)
	}
	return true, loop
}

func modifyLoopToReachEnd(badLoop []Cmd, inputLines []string, targetIndex int) []Cmd {
	var loop []Cmd

	for _, c := range badLoop {
		modifiedLoop := loop
		loop = append(loop, c)
                fmt.Println(len(loop), len(modifiedLoop))
		if c.t == "nop" || c.t == "jmp" {
			c.step, c.val = c.val, c.step
			fmt.Println("new swap", c)
		}
		modifiedLoop = append(modifiedLoop, c)
		fmt.Println("new loop", modifiedLoop)
		validLoop, finalLoop := completeLoop(inputLines, modifiedLoop, targetIndex)
		if validLoop {
			return finalLoop
		}
		// reset modifiedLoop since that didn't work
	}
	return loop
}

func accVal(cmdLoop []Cmd) int {
	sum := 0
	for _, c := range cmdLoop {
		sum += c.accum
	}
	return sum
}

func main() {
	flag.Parse()
	lineData := readDataArr(*filename)

	commandLoop := findLoop(lineData, 0)

	fmt.Println("loop sum is ", accVal(commandLoop))

	goodLoop := modifyLoopToReachEnd(commandLoop, lineData, len(lineData))
        fmt.Println(commandLoop)
        fmt.Println(goodLoop)
	fmt.Println("command sum is ", accVal(goodLoop))
}
