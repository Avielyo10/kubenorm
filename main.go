package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	// Version is the version of the application
	Version = "0.1.1"
	flagSet = flag.NewFlagSet("kubenorm", flag.ExitOnError)

	// Flags
	nano    = flagSet.Bool("n", false, "nano")
	micro   = flagSet.Bool("u", false, "micro")
	milli   = flagSet.Bool("m", false, "milli")
	kilo    = flagSet.Bool("k", false, "kilo")
	mega    = flagSet.Bool("M", false, "mega")
	giga    = flagSet.Bool("g", false, "giga")
	tera    = flagSet.Bool("t", false, "tera")
	peta    = flagSet.Bool("p", false, "peta")
	exa     = flagSet.Bool("e", false, "exa")
	version = flagSet.Bool("v", false, "version")

	helpMsg    = "Usage: kubenorm [options] [value]\nOptions:\n  -n nano\n  -u micro\n  -m milli\n  -k kilo\n  -M mega\n  -g giga\n  -t tera\n  -p peta\n  -e exa\n  -v version\n"
	versionMsg = fmt.Sprintf("kubenorm version: v%s\n", Version)
)

func main() {
	flagSet.Parse(os.Args[1:])
	flagSet.Usage = func() {
		fmt.Println(helpMsg)
	}

	var rvalue string
	if *version {
		rvalue = versionMsg
	} else if flagSet.NFlag() > 0 {
		rvalue = handleScale()
	} else {
		rvalue = helpMsg
	}
	fmt.Println(rvalue)
}

func handleScale() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input: ", err)
	}
	input = strings.TrimSuffix(input, "\n")

	var scale resource.Scale
	switch {
	case *nano:
		scale = resource.Nano
	case *micro:
		scale = resource.Micro
	case *milli:
		scale = resource.Milli
	case *kilo:
		scale = resource.Kilo
	case *mega:
		scale = resource.Mega
	case *giga:
		scale = resource.Giga
	case *tera:
		scale = resource.Tera
	case *peta:
		scale = resource.Peta
	case *exa:
		scale = resource.Exa
	default:
		scale = resource.Mega
	}

	quantity := resource.MustParse(input)
	return fmt.Sprintln(quantity.ScaledValue(scale))
}
