package depcalldemoe

import (
	"fmt"

	"github.com/csguojin/depcalldemof"
	"github.com/csguojin/depcalldemog"
)

func MyFunc1() {
	fmt.Println("depcalldemoe.MyFunc1")
	depcalldemog.MyFunc1()
	depcalldemof.MyFunc1()
}
