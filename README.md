go-ttyadapter
=============

This package provides an abstraction layer for reading keyboard input from terminal devices.

Features
--------

- Backends for common terminal implementations  
  - `tty8` → wrapper around [github.com/mattn/go-tty][go-tty]
  - `tty10` → wrapper around [golang.org/x/term][xterm]
  - `auto` → pseudo terminal that feeds programmed key sequences (useful for tests)

- Automatic handling of terminal modes  
  The library switches terminal modes (raw/cooked) internally as needed so callers do not have to manage mode changes manually.

- Test-friendly pseudo terminal for automated tests  
  Use `auto.Pilot` to simulate key sequences and verify interactive behavior in `go test` or CI.

- Designed for integration with readline-like tools  
  Intended to be a reusable terminal-input layer for libraries that implement line-editing or interactive selection UIs.

- Lightweight wrappers (not a pure-zero-dependency package)  
  This module provides small adapters over existing terminal libraries rather than reimplementing low-level terminal handling. Backends depend on [github.com/mattn/go-tty][go-tty] or [golang.org/x/term][xterm] as appropriate.

Example
-------

```examples/example.go
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
```

This example prints pressed keys enclosed in &lt; &gt; and exits when ESC is pressed.
The demo below shows two modes in sequence:

- Simulated input using auto.Pilot (numbers 1–5 are automatically typed)
- Manual input where actual keypresses (e.g. arrow keys, ESC) are captured from the terminal

![demo.gif](demo.gif)

[go-tty]: https://github.com/mattn/go-tty
[xterm]: https://pkg.go.dev/golang.org/x/term
