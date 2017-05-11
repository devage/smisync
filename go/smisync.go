// main of smisync
// 2013.03.21 ikpark@gmail.com

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func usage(cmd string) {
	fmt.Printf(
		"USAGE: %v <start>+/-<offset> < input.smi > output.smi\n",
		cmd)
}

func main() {

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(0)
	}

	const syncstr string = "<SYNC Start="
	var ( // buffer-related vars
		buf bytes.Buffer
		line string
		_e error
	)
	var start, offset, sync, pos, tidx, eidx int

	pos = strings.IndexAny(os.Args[1], "+-")
	start,  _e = strconv.Atoi(os.Args[1][0:pos])
	offset, _e = strconv.Atoi(os.Args[1][pos:])
	if os.Args[1][pos] == '-' {
		offset *= -1
	}

	buf.ReadFrom(os.Stdin)

	for {
		line, _e = buf.ReadString('\n');
		if _e != nil {
			break
		}

		if len(syncstr) <= len(line) &&
				syncstr == line[0:len(syncstr)] {
			fmt.Printf("%s", syncstr)

			tidx = len(syncstr)
			eidx = strings.Index(line, ">")
			sync, _e = strconv.Atoi(line[tidx:eidx])
			if sync >= start {
				fmt.Printf("%d", sync+offset)
			} else {
				fmt.Printf("%d", sync)
			}
			fmt.Printf("%s", line[eidx:])
		} else {
			fmt.Printf("%s", line)
		}
	}
}
