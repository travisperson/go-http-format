package HttpFormatter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"io"
	"io/ioutil"
	"net/http"
)

type Formatter struct {
	KeyColor   *color.Color
	ValueColor *color.Color
}

func NewFormatter() *Formatter {
	return &Formatter{
		KeyColor:   color.New(color.FgBlue, color.Bold),
		ValueColor: color.New(color.FgBlack, color.Bold),
	}
}

func (f *Formatter) Format(buf *bufio.Reader) error {
	head, err := buf.Peek(8)

	if err != nil {
		return err
	}

	if string(head[:8]) == "HTTP/1.1" {
		r, err := http.ReadResponse(buf, nil)

		if err != nil {
			return err
		}

		return f.printResponse(r)
	} else {
		r, err := http.ReadRequest(buf)

		if err != nil {
			return err
		}

		return f.printRequest(r)
	}

	return nil
}

func (f *Formatter) printBody(b io.ReadCloser) error {
	body, err := ioutil.ReadAll(b)

	if len(body) == 0 {
		return nil
	}

	if err != nil {
		return err
	}

	var v map[string]interface{}

	json.Unmarshal(body, &v)

	s, err := prettyjson.Marshal(v)

	if err != nil {
		return err
	}

	fmt.Println(string(s))
	fmt.Printf("\n")

	return nil
}

func (f *Formatter) printRequest(r *http.Request) error {
	f.KeyColor.Printf("%s ", r.Method)
	f.ValueColor.Printf("%s ", r.URL)
	f.KeyColor.Printf("%s\n", r.Proto)

	f.KeyColor.Printf("%s: ", "Host")
	f.ValueColor.Printf("%s\n", r.Host)

	for k, v := range r.Header {
		f.KeyColor.Printf("%s: ", k)
		f.ValueColor.Printf("%s\n", v[0])
	}

	fmt.Printf("\n")

	return f.printBody(r.Body)
}

func (f *Formatter) printResponse(r *http.Response) error {
	f.KeyColor.Printf("%s ", r.Proto)
	f.ValueColor.Printf("%s\n", r.Status)

	for k, v := range r.Header {
		f.KeyColor.Printf("%s: ", k)
		f.ValueColor.Printf("%s\n", v[0])
	}

	fmt.Printf("\n")

	return f.printBody(r.Body)
}
