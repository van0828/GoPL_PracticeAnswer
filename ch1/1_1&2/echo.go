package __1_2

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for index, arg := range os.Args[1:] {
		fmt.Println(fmt.Sprintf("index: %d, arg: %s", index, arg))
	}
}
