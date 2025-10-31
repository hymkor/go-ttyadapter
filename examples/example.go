//go:build run

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nyaosorg/go-ttyadapter"
	"github.com/nyaosorg/go-ttyadapter/auto"
	"github.com/nyaosorg/go-ttyadapter/tty8"
)

func run(operations []string) error {
	var tty ttyadapter.Tty

	// If command-line arguments are given, use pseudo terminal for simulation.
	if len(operations) > 0 {
		// Append ESC to end to exit automatically.
		operations = append(operations, "\x1B")
		tty = &auto.Pilot{Text: operations}
	} else {
		tty = &tty8.Tty{}
	}

	if err := tty.Open(nil); err != nil {
		return err
	}
	defer tty.Close()

	for {
		key, err := tty.GetKey()
		if err != nil {
			return err
		}
		if key == "\x1B" {
			return nil
		}
		// Show typed key, converting ESC to literal name.
		fmt.Printf("<%s>\n", strings.ReplaceAll(key, "\x1B", "ESC"))
	}
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
