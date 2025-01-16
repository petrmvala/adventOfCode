package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(f), "\n")[0]
	fmt.Println("Checksum:", SumDefrag(data))
}

func SumDefrag(input string) int {
	// println(input)
	// f := "[%s] blockid: %d, index: %d, revindex: %d, blklen: %d, lastblklen: %d, blktype: %s, lastlblktype: %s, fileId: %d, total: %d\n"
	in := []int{}
	strlen := 0
	for _, v := range strings.Split(input, "") {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		in = append(in, val)
		strlen += val
	}
	total := 0
	idx := 0
	for blk := 0; blk < len(in); blk++ {
		blklen := in[blk]
		// fmt.Printf(f, "block", blk, idx, strlen, blklen, -1, "N/A", "N/A", -1, total)
		if blk%2 == 0 { // current block is a file
			fileID := blk / 2
			for j := 0; j < blklen; j++ {
				total += idx * fileID
				idx++
				// fmt.Printf(f, "file", blk, idx, strlen, blklen, -1, "file", "N/A", fileID, total)
				if idx >= strlen {
					return total
				}
			}
			continue
		}
		// current block is empty space
		for blklen > 0 {
			// fmt.Printf(f, "empty", blk, idx, strlen, blklen, -1, "empty", "N/A", -1, total)
			lastBlockId := len(in) / 2
			lastBlockLen := &in[len(in)-1]
			for *lastBlockLen > 0 && blklen > 0 {
				// fmt.Printf(f, "lastFile", blk, idx, strlen, blklen, *lastBlockLen, "empty", "file", lastBlockId, total)
				total += idx * lastBlockId
				idx++
				strlen--
				blklen--
				*lastBlockLen--
				if idx >= strlen {
					return total
				}
			}
			if blklen == 0 {
				continue
			}
			// lastBlockLen == 0 --> get rid of the block
			in = in[:len(in)-1]
			lastBlockLen = &in[len(in)-1]
			// fmt.Printf(f, "trim", blk, idx, strlen, blklen, *lastBlockLen, "empty", "file", lastBlockId, total)
			// only last block of space is remaining
			if idx+*lastBlockLen >= strlen {
				return total
			}
			strlen -= *lastBlockLen
			in = in[:len(in)-1]
		}
	}
	return total
}
