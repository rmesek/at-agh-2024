package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcinCiura/AT-lab/5/suffixtree"
)

var files = []string{
	"mDNAs/dziobak.txt",
	"mDNAs/jeż.txt",
	"mDNAs/kot_domowy.txt",
	"mDNAs/walen.txt",
}

func main() {
	text := []byte{}
	lens := []int{}
	for i, fn := range files {
		f, err := os.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}
		text = append(text, f...)
		text = append(text, byte('0'+i))
		lens = append(lens, len(text))

		index := suffixtree.New(text)

		bits := make([]uint, index.NumNodes())
		dfs1(index, index.Root(), bits, lens)
		dfs2(index, index.Root(), "", bits, (1<<len(lens))-1)
	}
}

// Zwróci wyżej bitowy OR dla node
func dfs1(index *suffixtree.Index, node int, bits []uint, lens []int) uint {
	if index.IsLeaf(node) {
		return uint(1) << contains(lens, index.SuffixStart(node))
	}
	bits[node] = uint(0)
	for _, n := range index.Edges(node) {
		bits[node] |= dfs1(index, n, bits, lens)
	}
	return bits[node]
}

func dfs2(index *suffixtree.Index, node int, path string, bits []uint, mask uint) {
	if bits[node] != mask {
		return
	}

	if node != index.Root() {
		path += index.EdgeLabel(node)
	}
	fmt.Println(len(path), path)

	for _, n := range index.Edges(node) {
		dfs2(index, n, path, bits, mask)
	}
}

func contains(lens []int, pos int) int {
	for i, l := range lens {
		if pos < l {
			return i
		}
	}
	return -1
}
