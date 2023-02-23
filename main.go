package main

import (
	"bytes"
	"fmt"
	"github.com/birneee/parse-quic-tp/internal"
	"github.com/lucas-clemente/quic-go/logging"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

const PerspectiveNone = logging.Perspective(0)

func main() {
	app := &cli.App{
		Name: "parse-quic-tp",
		Action: func(c *cli.Context) error {
			in, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			r := bytes.NewReader(in)
			for r.Len() > 0 {
				entry, err := internal.ParseNextTransportParameter(r)
				if err != nil {
					return err
				}
				str, err := entry.String()
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", str)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
