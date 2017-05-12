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

func parse_option(option string) (int, int) {
	var start, offset, pos int

	pos = strings.IndexAny(option, "+-")
	start,  _ = strconv.Atoi(option[0:pos])
	offset, _ = strconv.Atoi(option[pos:])
	if option[pos] == '-' {
		offset *= -1
	}
	return start, offset
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
  var start, offset, sync, tidx, eidx int

	start, offset = parse_option(os.Args[1])

  buf.ReadFrom(os.Stdin)

  for {
    line, _e = buf.ReadString('\n');
    if _e != nil {
      break
    }

    if len(syncstr) <= len(line) && syncstr == line[0:len(syncstr)] {
      tidx = len(syncstr)
      eidx = strings.Index(line, ">")
      sync, _e = strconv.Atoi(line[tidx:eidx])
      if sync >= start {
        sync = sync + offset
      }

      fmt.Printf("%s%d%s", syncstr, sync, line[eidx:])
    } else {
      fmt.Printf("%s", line)
    }
  }
}
