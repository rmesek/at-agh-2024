package matching

import (
	"os"
	"slices"
	"testing"
)

func TestPreprocess(t *testing.T) {
	data := []string{
		"aaaaaaa",
		"pies",
		"dźwiedź",
		"owocowo",
		"indianin",
		"nienapełnienie",
	}
	for _, in := range data {
		got := Preprocess([]byte(in))
		want := SimplePreprocess([]byte(in))
		if !slices.Equal(got, want) {
			t.Errorf(`Preprocess(%#v) == %#v want %#v`,
				in, got, want)
		}
	}
}

func indices(pat, text []byte) []int {
	r := []int{}
	for i := 0; i+len(pat) <= len(text); i++ {
		if slices.Equal(text[i:i+len(pat)], pat) {
			r = append(r, i)
		}
	}
	return r
}

type TestData struct {
	Text string
	Pats []string
}

var tests = []TestData{
	{"aaaaaaa", []string{"aa", "bb"}},
	{"pies", []string{"pies", "pi", "bb", "pies1"}},
	{"dźwiedź", []string{"dźwiedź", "dzwiedz"}},
	{"owocowo", []string{"owo", "cowo", "owocowo", "oc"}},
	{"indianin", []string{"i", "in", "bbb"}},
	{"nienapełnienie", []string{"nie"}},
}

func TestNaive(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			Naive(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`Naive(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func TestBackwardNaive(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			BackwardNaive(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`BackwardNaive(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func TestBoyerMoore(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			BoyerMoore(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`BoyerMoore(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func TestKMP(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			KMP(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`KMP(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func TestKarpRabin(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			KarpRabin(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`KarpRabin(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func TestShiftOr(t *testing.T) {
	for _, test := range tests {
		for _, pat := range test.Pats {
			// fmt.Println(test, pat)
			pat := []byte(pat)
			text := []byte(test.Text)
			got := []int{}
			ShiftOr(pat, text, func(n int) { got = append(got, n) })
			want := indices(pat, text)
			if !slices.Equal(got, want) {
				// Zgłoś błąd, korzystając z funkcji `t.Errorf`
				t.Errorf(`ShiftOr(%#v, %#v) == %#v want %#v`,
					pat, text, got, want)
			}
		}
	}
}

func BenchmarkShortNaive(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkLongNaive(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkShortBackwardNaive(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkLongBackwardNaive(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkShortBoyerMoore(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkLongBoyerMoore(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkShortKMP(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkLongKMP(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkShortKarpRabin(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkLongKarpRabin(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkShortShiftOr(b *testing.B) {
	pat := []byte("dniem")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func BenchmarkLongShiftOr(b *testing.B) {
	pat := []byte("gospodarze")
	text, err := os.ReadFile("pan-tadeusz.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}
