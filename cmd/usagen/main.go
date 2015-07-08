//go:generate usagen usagen

// usagen generates usage documentation for a given command. It does so by
// executing the command with the "--help" flag and storing the output as a
// package doc comment.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// flagOut specifies the output path.
var flagOut string

const use = `
Usage: usagen [OPTION]... CMD
Generate usage documentation for a given command.

Flags:`

func init() {
	flag.StringVar(&flagOut, "o", "z_usage.go", "Output path.")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, use[1:])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := genUsage(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// genUsage generates usage documentation for a given command.
func genUsage(cmdName string) error {
	// Capture output from running:
	//    foo --help
	buf := new(bytes.Buffer)
	cmd := exec.Command(cmdName, "--help")
	cmd.Stderr = buf
	cmd.Run()

	// Add DO NOT EDIT notice.
	out := new(bytes.Buffer)
	fmt.Fprintf(out, "// generated by `usagen %s`; DO NOT EDIT\n\n", cmdName)

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

	// Store usage info in z_usage.go.
	return ioutil.WriteFile(flagOut, out.Bytes(), 0644)
}