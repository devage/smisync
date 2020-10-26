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
	fmt.Printf("USAGE: %v <start>+/-<offset> < input.smi > output.smi\n", cmd)
}

func parseOption(option string) (int, int) {

	pos := strings.IndexAny(option, "+-")
	start, _ := strconv.Atoi(option[0:pos])
	offset, _ := strconv.Atoi(option[pos:])
	if option[pos] == '-' {
		offset *= -1
	}
	return start, offset
}

func smiSync(line string, start int, offset int) string {

	const syncstr string = "<SYNC Start="
	syncedLine := line

	if len(syncstr) <= len(line) && syncstr == line[0:len(syncstr)] {
		tidx := len(syncstr)
		eidx := strings.Index(line, ">")
		sync, _ := strconv.Atoi(line[tidx:eidx])
		if sync >= start {
			sync = sync + offset
		}

		syncedLine = fmt.Sprintf("%s%d%s", syncstr, sync, line[eidx:])
	}

	return syncedLine
}

func main() {

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(0)
	}

	var buf bytes.Buffer

	start, offset := parseOption(os.Args[1])

	buf.ReadFrom(os.Stdin)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		syncedLine := smiSync(line, start, offset)
		fmt.Printf("%s", syncedLine)
	}
}
