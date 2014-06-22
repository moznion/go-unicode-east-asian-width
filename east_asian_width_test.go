package eastasianwidth

import (
	"testing"
	"unicode"

	. "github.com/onsi/gomega"
)

func TestAmbiguousCharacterShouldBeUsedAsHalfwidth(t *testing.T) {
	RegisterTestingT(t)

	char := '‐'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeTrue())
	Expect(IsFullwidth(char)).To(BeFalse())
	Expect(IsHalfwidth(char)).To(BeTrue())
}

func TestAmbiguousCharacterShouldBeUsedAsFullwidth(t *testing.T) {
	RegisterTestingT(t)

	EastAsian = true

	char := '‐'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeTrue())
	Expect(IsFullwidth(char)).To(BeTrue())
	Expect(IsHalfwidth(char)).To(BeFalse())
}

func TestFullwidthCharacterShouldBeUsedAsIt(t *testing.T) {
	RegisterTestingT(t)

	var char rune

	char = '一'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeFalse())
	Expect(unicode.Is(EastAsianWide, char)).To(BeTrue())
	Expect(IsFullwidth(char)).To(BeTrue())

	char = '＂'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeFalse())
	Expect(unicode.Is(EastAsianFullwidth, char)).To(BeTrue())
	Expect(IsFullwidth(char)).To(BeTrue())
}

func TestHalfwidthCharacterShouldBeUsedAsIt(t *testing.T) {
	RegisterTestingT(t)

	var char rune

	char = '｡'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeFalse())
	Expect(unicode.Is(EastAsianHalfwidth, char)).To(BeTrue())
	Expect(IsHalfwidth(char)).To(BeTrue())

	char = '¬'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeFalse())
	Expect(unicode.Is(EastAsianNarrow, char)).To(BeTrue())
	Expect(IsHalfwidth(char)).To(BeTrue())

	char = 'ÿ'
	Expect(unicode.Is(EastAsianAmbiguous, char)).To(BeFalse())
	Expect(unicode.Is(EastAsianNeutral, char)).To(BeTrue())
	Expect(IsHalfwidth(char)).To(BeTrue())
}
