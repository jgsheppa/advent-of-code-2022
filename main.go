package main

import (
	"fmt"
	"github.com/jgsheppa/advent-of-code-2022/common"
	"github.com/jgsheppa/advent-of-code-2022/daySix"
)

func main() {

	signal, err := common.ReadSignal("puzzle_input_day_6.txt")
	if err != nil {
		fmt.Printf("error while reading file: %v", err)
	}
	startOfPacket, err := daySix.FindStartOfMessageOrSignal(signal, 14)
	if err != nil {
		fmt.Printf("Received following error: %v", err)
	}
	fmt.Printf("FindStartOfMessage: %v \n", startOfPacket)
}
