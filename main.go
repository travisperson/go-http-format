package main

import (
	"bufio"
	"github.com/travisperson/go-http-format/http-formatter"
	"io"
	"os"
)

func main() {
	bufstdin := bufio.NewReader(os.Stdin)
	f := HttpFormatter.NewFormatter()

	for {
		if f.Format(bufstdin) == io.EOF {
			break
		}
	}
}
