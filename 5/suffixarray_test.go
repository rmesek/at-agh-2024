package suffixarray_test

import (
	"slices"
	"testing"

	simplesuffixarray "github.com/MarcinCiura/AT-lab/5/suffixarray"
)

//"github.com/MarcinCiura/AT-lab/5/suffixtree"

func TestLookupAll(t *testing.T) {
	text := "ananas"
	data := []struct {
		in   string
		want []int
	}{
		{"ananas", []int{0}},
		{"nan", []int{1}},
		{"a", []int{0, 2, 4}},
	}
	sa := simplesuffixarray.New([]byte(text))
	for _, d := range data {
		if got := sa.LookupAll([]byte(d.in)); !slices.Equal(got, d.want) {
			t.Errorf("LookupAll(%v) == %v want %v", d.in, got, d.want)
		}
	}
}

func BenchmarkAhoCorasickBuild(b *testing.B) {

}

func BenchmarkSimpleSuffixArrayBuild(b *testing.B) {

}
func BenchmarkLibrarySuffixArrayBuild(b *testing.B) {

}

func BenchmarkAhoCorasickSearch(b *testing.B) {

}

func BenchmarkSimpleSuffixArraySearch(b *testing.B) {

}
func BenchmarkLibrarySuffixArraySearch(b *testing.B) {

}
