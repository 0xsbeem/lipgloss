package list

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestList(t *testing.T) {
	l := New().
		Item("Foo").
		Item("Bar").
		Item("Baz")

	expected := strings.TrimPrefix(`
• Foo
• Bar
• Baz`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestHide(t *testing.T) {
	l := New().
		Item("Foo").
		Item("Baz")

	expected := strings.TrimPrefix(`
• Foo
• Baz`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestListIntegers(t *testing.T) {
	l := New().
		Item("1").
		Item("2").
		Item("3")

	expected := strings.TrimPrefix(`
• 1
• 2
• 3`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestEnumerators(t *testing.T) {
	tests := []struct {
		enumeration Enumerator
		expected    string
	}{
		{
			enumeration: Alphabet,
			expected: `
A. Foo
B. Bar
C. Baz`,
		},
		{
			enumeration: Arabic,
			expected: `
1. Foo
2. Bar
3. Baz`,
		},
		{
			enumeration: Roman,
			expected: `
  I. Foo
 II. Bar
III. Baz`,
		},
		{
			enumeration: Bullet,
			expected: `
• Foo
• Bar
• Baz`,
		},
		{
			enumeration: Tree,
			expected: `
├─ Foo
├─ Bar
└─ Baz`,
		},
	}

	for _, test := range tests {
		expected := strings.TrimPrefix(test.expected, "\n")

		l := New().
			Enumerator(test.enumeration).
			Item("Foo").
			Item("Bar").
			Item("Baz")

		if l.String() != expected {
			t.Errorf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
		}
	}
}

func TestEnumeratorsTransform(t *testing.T) {
	tests := []struct {
		enumeration Enumerator
		style       lipgloss.Style
		expected    string
	}{
		{
			enumeration: Alphabet,
			style:       lipgloss.NewStyle().MarginRight(1).Transform(strings.ToLower),
			expected: `
a. Foo
b. Bar
c. Baz`,
		},
		{
			enumeration: Arabic,
			style: lipgloss.NewStyle().MarginRight(1).Transform(func(s string) string {
				return strings.Replace(s, ".", ")", 1)
			}),
			expected: `
1) Foo
2) Bar
3) Baz`,
		},
		{
			enumeration: Roman,
			style: lipgloss.NewStyle().Transform(func(s string) string {
				return "(" + strings.Replace(strings.ToLower(s), ".", "", 1) + ") "
			}),
			expected: `
  (i) Foo
 (ii) Bar
(iii) Baz`,
		},
		{
			enumeration: Bullet,
			style: lipgloss.NewStyle().Transform(func(s string) string {
				return "- " // this is better done by replacing the enumerator.
			}),
			expected: `
- Foo
- Bar
- Baz`,
		},
		{
			enumeration: Tree,
			style: lipgloss.NewStyle().MarginRight(1).Transform(func(s string) string {
				return strings.Replace(s, "─", "───", 1)
			}),
			expected: `
├─── Foo
├─── Bar
└─── Baz`,
		},
	}

	for _, test := range tests {
		expected := strings.TrimPrefix(test.expected, "\n")

		l := New().
			Enumerator(test.enumeration).
			EnumeratorStyle(test.style).
			Item("Foo").
			Item("Bar").
			Item("Baz")

		if l.String() != expected {
			t.Errorf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
		}
	}
}

func TestBullet(t *testing.T) {
	tests := []struct {
		enum Enumerator
		i    int
		exp  string
	}{
		{Alphabet, 0, "A"},
		{Alphabet, 25, "Z"},
		{Alphabet, 26, "AA"},
		{Alphabet, 51, "AZ"},
		{Alphabet, 52, "BA"},
		{Alphabet, 79, "CB"},
		{Alphabet, 701, "ZZ"},
		{Alphabet, 702, "AAA"},
		{Alphabet, 801, "ADV"},
		{Alphabet, 1000, "ALM"},
		{Roman, 0, "I"},
		{Roman, 25, "XXVI"},
		{Roman, 26, "XXVII"},
		{Roman, 50, "LI"},
		{Roman, 100, "CI"},
		{Roman, 701, "DCCII"},
		{Roman, 1000, "MI"},
	}

	for _, test := range tests {
		bullet := strings.TrimSuffix(test.enum(nil, test.i), ".")
		if bullet != test.exp {
			t.Errorf("expected: %s, got: %s\n", test.exp, bullet)
		}
	}
}

func TestData(t *testing.T) {
	data := NewStringData("Foo", "Bar", "Baz")
	filter := func(index int) bool {
		return index != 1
	}
	l := New().Data(NewFilter(data).Filter(filter))

	expected := strings.TrimPrefix(`
• Foo
• Baz`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestEnumeratorsAlign(t *testing.T) {
	fooList := strings.Split(strings.TrimSuffix(strings.Repeat("Foo ", 100), " "), " ")
	l := New().Enumerator(Roman)
	for _, f := range fooList {
		l.Item(f)
	}

	expected := strings.TrimPrefix(`
       I. Foo
      II. Foo
     III. Foo
      IV. Foo
       V. Foo
      VI. Foo
     VII. Foo
    VIII. Foo
      IX. Foo
       X. Foo
      XI. Foo
     XII. Foo
    XIII. Foo
     XIV. Foo
      XV. Foo
     XVI. Foo
    XVII. Foo
   XVIII. Foo
     XIX. Foo
      XX. Foo
     XXI. Foo
    XXII. Foo
   XXIII. Foo
    XXIV. Foo
     XXV. Foo
    XXVI. Foo
   XXVII. Foo
  XXVIII. Foo
    XXIX. Foo
     XXX. Foo
    XXXI. Foo
   XXXII. Foo
  XXXIII. Foo
   XXXIV. Foo
    XXXV. Foo
   XXXVI. Foo
  XXXVII. Foo
 XXXVIII. Foo
   XXXIX. Foo
      XL. Foo
     XLI. Foo
    XLII. Foo
   XLIII. Foo
    XLIV. Foo
     XLV. Foo
    XLVI. Foo
   XLVII. Foo
  XLVIII. Foo
    XLIX. Foo
       L. Foo
      LI. Foo
     LII. Foo
    LIII. Foo
     LIV. Foo
      LV. Foo
     LVI. Foo
    LVII. Foo
   LVIII. Foo
     LIX. Foo
      LX. Foo
     LXI. Foo
    LXII. Foo
   LXIII. Foo
    LXIV. Foo
     LXV. Foo
    LXVI. Foo
   LXVII. Foo
  LXVIII. Foo
    LXIX. Foo
     LXX. Foo
    LXXI. Foo
   LXXII. Foo
  LXXIII. Foo
   LXXIV. Foo
    LXXV. Foo
   LXXVI. Foo
  LXXVII. Foo
 LXXVIII. Foo
   LXXIX. Foo
    LXXX. Foo
   LXXXI. Foo
  LXXXII. Foo
 LXXXIII. Foo
  LXXXIV. Foo
   LXXXV. Foo
  LXXXVI. Foo
 LXXXVII. Foo
LXXXVIII. Foo
  LXXXIX. Foo
      XC. Foo
     XCI. Foo
    XCII. Foo
   XCIII. Foo
    XCIV. Foo
     XCV. Foo
    XCVI. Foo
   XCVII. Foo
  XCVIII. Foo
    XCIX. Foo
       C. Foo`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestIndent(t *testing.T) {
	l := New("foo", "bar", "baz").Enumerator(Arabic).Indent(2)

	expected := strings.TrimPrefix(`
  1. foo
  2. bar
  3. baz`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}

func TestOffset(t *testing.T) {
	l := New("foo", "bar", "baz", "qux", "quux").Enumerator(Arabic).Offset(2).Height(2)

	expected := strings.TrimPrefix(`
3. baz
4. qux`, "\n")

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}
func TestInvalidOffset(t *testing.T) {
	l := New("foo", "bar", "baz", "qux", "quux").
		Enumerator(Arabic).
		Offset(10).
		Height(2)

	expected := ""

	if l.String() != expected {
		t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n", expected, l.String())
	}
}