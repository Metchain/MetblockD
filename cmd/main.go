package main

import (
	"fmt"
	"github.com/Metchain/Metblock/utils/nodes"
)

func main() {
	fmt.Println(nodes.FindNeighbors("127.0.0.1", 5000, 0, 3, 5000, 5003))
}
