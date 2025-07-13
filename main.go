package main

import (
	"fmt"

	"github.com/relax-space/gomod3B"
	gomod3C_v5 "github.com/relax-space/gomod3C/v5"
)

func main() {
	fmt.Printf("gomod3B.MultiB(2, 3)=%d\n", gomod3B.MultiB(2, 3))
	fmt.Printf("gomod3C_v5.AddC(2, 3)=%d\n", gomod3C_v5.AddC(2, 3))
}
