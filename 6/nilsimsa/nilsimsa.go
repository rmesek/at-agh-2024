package nilsimsa

import (
	"fmt"
	"math/bits"
)

// Każda liczba całkowita między 0 a 255 występuje w tablicy
// hashConstants 1 raz
var hashConstants = [256]uint8{
	0x02, 0xD6, 0x9E, 0x6F, 0xF9, 0x1D, 0x04, 0xAB,
	0xD0, 0x22, 0x16, 0x1F, 0xD8, 0x73, 0xA1, 0xAC,
	0x3B, 0x70, 0x62, 0x96, 0x1E, 0x6E, 0x8F, 0x39,
	0x9D, 0x05, 0x14, 0x4A, 0xA6, 0xBE, 0xAE, 0x0E,
	0xCF, 0xB9, 0x9C, 0x9A, 0xC7, 0x68, 0x13, 0xE1,
	0x2D, 0xA4, 0xEB, 0x51, 0x8D, 0x64, 0x6B, 0x50,
	0x23, 0x80, 0x03, 0x41, 0xEC, 0xBB, 0x71, 0xCC,
	0x7A, 0x86, 0x7F, 0x98, 0xF2, 0x36, 0x5E, 0xEE,
	0x8E, 0xCE, 0x4F, 0xB8, 0x32, 0xB6, 0x5F, 0x59,
	0xDC, 0x1B, 0x31, 0x4C, 0x7B, 0xF0, 0x63, 0x01,
	0x6C, 0xBA, 0x07, 0xE8, 0x12, 0x77, 0x49, 0x3C,
	0xDA, 0x46, 0xFE, 0x2F, 0x79, 0x1C, 0x9B, 0x30,
	0xE3, 0x00, 0x06, 0x7E, 0x2E, 0x0F, 0x38, 0x33,
	0x21, 0xAD, 0xA5, 0x54, 0xCA, 0xA7, 0x29, 0xFC,
	0x5A, 0x47, 0x69, 0x7D, 0xC5, 0x95, 0xB5, 0xF4,
	0x0B, 0x90, 0xA3, 0x81, 0x6D, 0x25, 0x55, 0x35,
	0xF5, 0x75, 0x74, 0x0A, 0x26, 0xBF, 0x19, 0x5C,
	0x1A, 0xC6, 0xFF, 0x99, 0x5D, 0x84, 0xAA, 0x66,
	0x3E, 0xAF, 0x78, 0xB3, 0x20, 0x43, 0xC1, 0xED,
	0x24, 0xEA, 0xE6, 0x3F, 0x18, 0xF3, 0xA0, 0x42,
	0x57, 0x08, 0x53, 0x60, 0xC3, 0xC0, 0x83, 0x40,
	0x82, 0xD7, 0x09, 0xBD, 0x44, 0x2A, 0x67, 0xA8,
	0x93, 0xE0, 0xC2, 0x56, 0x9F, 0xD9, 0xDD, 0x85,
	0x15, 0xB4, 0x8A, 0x27, 0x28, 0x92, 0x76, 0xDE,
	0xEF, 0xF8, 0xB2, 0xB7, 0xC9, 0x3D, 0x45, 0x94,
	0x4B, 0x11, 0x0D, 0x65, 0xD5, 0x34, 0x8B, 0x91,
	0x0C, 0xFA, 0x87, 0xE9, 0x7C, 0x5B, 0xB1, 0x4D,
	0xE5, 0xD4, 0xCB, 0x10, 0xA2, 0x17, 0x89, 0xBC,
	0xDB, 0xB0, 0xE2, 0x97, 0x88, 0x52, 0xF7, 0x48,
	0xD3, 0x61, 0x2C, 0x3A, 0x2B, 0xD1, 0x8C, 0xFB,
	0xF1, 0xCD, 0xE4, 0x6A, 0xE7, 0xA9, 0xFD, 0xC4,
	0x37, 0xC8, 0xD2, 0xF6, 0xDF, 0x58, 0x72, 0x4E,
}

// hashBytes zwraca wartość funkcji haszującej Pearsona dla argumentów
// b0, b1, b2
func hashBytes(b0, b1, b2 uint8) uint8 {
	return hashConstants[hashConstants[hashConstants[b0]^b1]^b2]
}

// Histogram służy do obliczania funkcji haszującej łańcucha
type Histogram struct {
	counters           [256]int
	numTrigrams        int
	charsProcessed     int
	b0, b1, b2, b3, b4 uint8
}

// update zapamiętuje wynik haszowania trigramu (ba, bb, bc)
func (h *Histogram) update(ba, bb, bc uint8) {
	h.counters[hashBytes(ba, bb, bc)]++
	h.numTrigrams++
}

// ProcessRune dodaje 1 do niektórych liczników `h.counters`. Te
// liczniki są wynikami funkcji haszującej Pearsona dla trigramów
// złożonych z `r` i 2 spośród 4 runów poprzedzających `r`
func (h *Histogram) ProcessRune(r rune) {
	h.b0, h.b1, h.b2, h.b3, h.b4 = uint8(r), h.b0, h.b1, h.b2, h.b3
	h.charsProcessed++
	switch h.charsProcessed {
	default:
		h.update(h.b0^0x3F, h.b3, h.b4)
		h.update(h.b0^0x1F, h.b2, h.b4)
		h.update(h.b0^0x0F, h.b1, h.b4)
		fallthrough
	case 4:
		h.update(h.b0^0x07, h.b2, h.b3)
		h.update(h.b0^0x03, h.b1, h.b3)
		fallthrough
	case 3:
		h.update(h.b0^0x01, h.b1, h.b2)
	case 2:
	case 1:
	case 0:
		break
	}
}

// ProcessString dodaje 1 do niektórych liczników `h.counters`. Te
// liczniki są wynikami funkcji haszującej Pearsona dla wszystkich
// trigramów zawartych w kolejnych 5-gramach łańcucha `s`
func (h *Histogram) ProcessString(s string) {
	for _, r := range s {
		h.ProcessRune(r)
	}
}

// Fingerprint to maska 256 bitów
type Fingerprint [256 / 64]uint64

// SetBit ustawia `n`-ty bit maski `f`
func (f *Fingerprint) SetBit(n int) {
	f[n/64] |= uint64(1) << (n % 64)
}

// Fingerprint zwraca maskę 256 bitów. Ta maska to wynik funkcji
// haszującej tego łańcucha, który został przetworzony przez funkcję
// `h.ProcessString`
func (h *Histogram) Fingerprint() *Fingerprint {
	f := Fingerprint{}
	threshold := h.numTrigrams / 256
	for i, a := range h.counters {
		if a > threshold {
			f.SetBit(i)
		}
	}
	return &f
}

// Nilsimsa zwraca wynik funkcji haszującej łańcucha `s`
func Nilsimsa(s string) *Fingerprint {
	h := Histogram{}
	h.ProcessString(s)
	return h.Fingerprint()
}

// HammingDistance zwraca odległość Hamminga między maskami
// `f1` i `f2`, czyli liczbę tych bitów maski `f1`, które są
// różne od odpowiadających im bitów maski `f2`
func HammingDistance(f1, f2 *Fingerprint) int {
	r := 0
	for i := range f1 {
		r += bits.OnesCount64(f1[i] ^ f2[i])
	}
	return r
}

// Print wypisuje maskę `f` jako ciąg bitów
func (f *Fingerprint) Print() {
	for _, w := range f {
		fmt.Printf("%064b", w)
	}
	fmt.Println()
}

// Memo to struktura danych, w której są zapisywane argumenty
// i wartości funkcji Distance
type Memo struct {
	m map[string]int
}

// makeKey zwraca klucz mapy. Ten klucz jest inny dla każdej pary
// wycinków `s` i `t` pod warunkiem, że żaden z tych wycinków nie
// zawiera znaku '$'
func makeKey(s, t []byte) string {
	return string(s) + "$" + string(t)
}

// get sprawdza, czy funkcja Distance została wcześniej wywołana z
// argumentami `s` i `t`. Jeśli tak, to zwraca wynik poprzedniego
// wywołania funkcji Distance(s, t). Jeśli nie, to wykonuje funkcję
// Distance(s, t), a potem zapisuje jej argumenty i wynik w mapie
// `m.m`
func (m Memo) get(s, t []byte) int {
	k := makeKey(s, t)
	if v, ok := m.m[k]; ok {
		return v
	}
	m.m[k] = m.Distance(s, t)
	return m.m[k]
}

// put zapisuje w mapie `m` wartość `v`, która odpowiada argumentom
// `s` i `t`
func (m Memo) put(s, t []byte, v int) int {
	k := makeKey(s, t)
	m.m[k] = v
	return v
}

// Distance zwraca odległość edycyjną wycinków `s` i `t`
func (m Memo) Distance(s, t []byte) int {
	i, j := len(s), len(t)
	if len(s) == 0 {
		return m.put(s, t, len(t))
	} else if len(t) == 0 {
		return m.put(s, t, len(s))
	} else if s[i-1] == t[j-1] {
		return m.put(s, t, m.get(s[:i-1], t[:j-1]))
	} else { // if s[i-1] != t[j-1]
		return m.put(s, t, 1+min(
			m.get(s, t[:j-1]),       // Insert
			m.get(s[:i-1], t[:j-1]), // Substitute
			m.get(s[:i-1], t)))      // Delete
	}
}

// Allison zwraca odległość edycyjną wycinków `s` i `t`
func Allison(s, t []byte) int {
	m := Memo{m: map[string]int{}}
	return m.Distance(s, t)
}

// setNthBit zwraca maskę, w której bit na pozycji n jest równy 1
func setNthBit(n int) uint64 {
	return uint64(1) << n
}

// nthBit zwraca bit na pozycji n w masce m
func nthBit(m uint64, n int) uint64 {
	return (m >> n) & 1
}

// makeMask zwraca tablicę 256 masek; bity i-tej maski są równe 0
// na pozycjach równych wszystkim pozycjom znaku i we wzorcu pat
func makeMask(pat []byte) [256]uint64 {
	m := [256]uint64{}
	for c := 0; c < 256; c++ {
		m[c] = ^uint64(0) // Ustaw wszystkie bity maski m[c]
	}
	for j, c := range pat {
		m[c] &^= setNthBit(j) // Wyzeruj j-ty bit maski m[c]
	}
	// Dla 0 <= c < 256, 0 <= j < len(pat) zachodzi
	// (nthBit(m[c], j) == 0) == (pat[j] == c)
	return m
}

// FuzzyShiftOrL wywołuje funkcję `output(i, j)` dla każdej takiej
// pary indeksów (`i`, `j`), że odległość Levenshteina między
// wycinkiem `text[i:j]` a wzorcem `pat` wynosi co najwyżej 2
func FuzzyShiftOrL(pat, text []byte, output func(int, int)) {
	m := makeMask(pat)
	s0, s1, s2 := ^uint64(0), ^uint64(0), ^uint64(0)
	for j, c := range text {
		// Uwzględnij zamianę 1 znaku lub wstawienie 1 znaku
		s2 = ((s2 << 1) | m[c]) & (s1 << 1) & s1
		s1 = ((s1 << 1) | m[c]) & (s0 << 1) & s0
		s0 = (s0 << 1) | m[c]
		// Uwzględnij usunięcie 1 znaku
		s1 &= (s0 << 1)
		s2 &= (s1 << 1)
		if nthBit(s2, len(pat)-1) == 0 {
			for i := max((j - len(pat) - 2), 0); i <= min((j-len(pat)+2), len(text)); i++ {
				if Allison(text[i:j], pat) <= 2 {
					output(i, j) // Zwróć pozycję ostatniego znaku wycinka
				}
			}
		}
	}
}
