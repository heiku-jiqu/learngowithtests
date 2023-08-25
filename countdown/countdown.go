package countdown

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, "Go!")
}
