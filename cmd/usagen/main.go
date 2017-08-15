//go:generate usagen -o z_usage.go usagen

// The usagen tool generates usage documentation for a given command. It does so
// by executing the command with the "-help" flag and storing the output as a
// package doc comment.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func usage() {
	const use = `
Generate usage documentation for a given command.

Usage:

	usagen [OPTION]... CMD

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// output specifies the output path.
		output string
	)
	flag.StringVar(&output, "o", "", "output path")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	cmdName := flag.Arg(0)

	// Generate usage documentation.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatalf("unable to create file; %v", err)
		}
		defer f.Close()
		w = f
	}
	if err := genUsage(w, cmdName); err != nil {
		log.Fatalf("%+v", err)
	}
}

// genUsage generates usage documentation for a given command.
func genUsage(w io.Writer, cmdName string) error {
	// Capture output from running:
	//    foo -help
	buf := new(bytes.Buffer)
	cmd := exec.Command(cmdName, "-help")
	cmd.Stderr = buf
	cmd.Run()

	// Add DO NOT EDIT notice.
	out := new(bytes.Buffer)
	fmt.Fprintf(out, "// Generated by `usagen %s`; DO NOT EDIT.\n\n", cmdName)

	// Prefix each line with "//    ".
	lines := strings.Split(buf.String(), "\n")
	for i, line := range lines {
		prefix := "//    "
		if len(line) == 0 {
			if i == len(lines)-1 {
				break
			}
			prefix = "//"
		}
		fmt.Fprintf(out, "%s%s\n", prefix, line)
	}
	out.WriteString("package main\n")

	// Write usage info to w.
	if _, err := w.Write(out.Bytes()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
