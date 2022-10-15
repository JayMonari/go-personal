package main

import (
	"bytes"
	"fmt"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const MinRead = bytes.MinRead

var ErrTooLarge = bytes.ErrTooLarge

// https://github.com/pingcap/tidb/blob/00791e7968ffad2de33d7af7f6f8a21580e2ab7e/ddl/cluster.go#L280
// https://github.com/ethereum/go-ethereum/blob/fb75f11e87420ec25ff72f7eeeb741fa8974e87e/trie/proof.go#L251-L258
// https://github.com/kubernetes/kubernetes/blob/61ca612cbb85efa13444a6d8ae517cd5e9c151a4/pkg/api/v1/endpoints/util.go#L177
// https://github.com/cockroachdb/cockroach/blob/1c377714b802ae9b354344a81f2611a96eb3255f/pkg/storage/pebble.go#L113
// https://github.com/hashicorp/vault/blob/6bede501766490007fb8d9323b9a37ef2883c07b/helper/dhutil/dhutil.go#L70
func ExampleCompare() {
	a4b := []byte("aaaab")
	a3b := []byte("aaab")
	a2 := []byte("aa")

	tellMe := func(a, b []byte) {
		switch bytes.Compare(a, b) {
		case -1:
			fmt.Printf("🤏 %q\thas a lexical score less than %q\n", a, b)
		case 0:
			fmt.Printf("🟰 %q\thas a lexical score that equals %q\n", a, b)
		case 1:
			fmt.Printf("💪 %q\thas a lexical score greater than %q\n", a, b)
		}
	}
	tellMe(a4b, a3b)
	tellMe(a2, a3b)
	tellMe(a4b, a2)
	tellMe(a2, a2)
	// Output:
	// 🤏 "aaaab"	has a lexical score less than "aaab"
	// 🤏 "aa"	has a lexical score less than "aaab"
	// 💪 "aaaab"	has a lexical score greater than "aa"
	// 🟰 "aa"	has a lexical score that equals "aa"
}

func ExampleContains() {
	data := []byte("My dentist tells me that chewing bricks is very bad for your teeth.")
	findOut := func(a, b []byte) {
		switch bytes.Contains(a, b) {
		case true:
			fmt.Printf("✅ %q contains %q\n", a, b)
		case false:
			fmt.Printf("❌ %q does NOT contain %q\n", a, b)
		}
	}
	findOut(data, []byte("my"))
	findOut(data, []byte("brushing"))
	findOut(data, []byte("wing"))
	findOut(data, []byte("hat"))
	findOut(data, []byte("My"))
	// Output:
	// ❌ "My dentist tells me that chewing bricks is very bad for your teeth." does NOT contain "my"
	// ❌ "My dentist tells me that chewing bricks is very bad for your teeth." does NOT contain "brushing"
	// ✅ "My dentist tells me that chewing bricks is very bad for your teeth." contains "wing"
	// ✅ "My dentist tells me that chewing bricks is very bad for your teeth." contains "hat"
	// ✅ "My dentist tells me that chewing bricks is very bad for your teeth." contains "My"
}

func ExampleContainsAny() {
	data := []byte("At that moment I was the most fearsome weasel in the entire swamp.")
	findOut := func(a []byte, chars string) {
		switch bytes.ContainsAny(a, chars) {
		case true:
			fmt.Printf("✅ %q contains one of these characters %q\n", a, chars)
		case false:
			fmt.Printf("❌ %q does NOT contain any of these characters %q\n", a, chars)
		}
	}

	findOut(data, "stuff")
	findOut(data, "x")
	findOut(data, "XXX")
	findOut(data, " ")
	// Output:
	// ✅ "At that moment I was the most fearsome weasel in the entire swamp." contains one of these characters "stuff"
	// ❌ "At that moment I was the most fearsome weasel in the entire swamp." does NOT contain any of these characters "x"
	// ❌ "At that moment I was the most fearsome weasel in the entire swamp." does NOT contain any of these characters "XXX"
	// ✅ "At that moment I was the most fearsome weasel in the entire swamp." contains one of these characters " "
}

func ExampleContainsRune() {
	data := []byte{0xf0, 0x9f, 0x98, 0x81}
	findOut := func(a []byte, r rune) {
		switch bytes.ContainsRune(a, r) {
		case true:
			fmt.Printf("✅ %q contains %q\n", a, r)
		case false:
			fmt.Printf("❌ %q does NOT contain %q\n", a, r)
		}
	}
	findOut(data, '😁')
	findOut(data, 0x1f601)
	findOut(data, '\U0001f601')
	findOut(data, '\u274c')
	// Output:
	// ✅ "😁" contains '😁'
	// ✅ "😁" contains '😁'
	// ✅ "😁" contains '😁'
	// ❌ "😁" does NOT contain '❌'
}

func ExampleCount() {
	data := []byte("It isn't true that my mattress is made of cotton candy.")
	emojis := []byte("🫶🙌👌👍🔥🙃😎😁🤗")
	howMany := func(a, b []byte) {
		fmt.Printf("There are %d of %q in %q\n", bytes.Count(a, b), b, a)
	}
	howMany(data, []byte("t"))
	howMany(data, []byte("is"))
	howMany(data, []byte(""))
	howMany(emojis, []byte{0xf0})
	howMany(emojis, []byte(""))
	// Output:
	// There are 9 of "t" in "It isn't true that my mattress is made of cotton candy."
	// There are 2 of "is" in "It isn't true that my mattress is made of cotton candy."
	// There are 56 of "" in "It isn't true that my mattress is made of cotton candy."
	// There are 9 of "\xf0" in "\U0001faf6🙌👌👍🔥🙃😎😁🤗"
	// There are 10 of "" in "\U0001faf6🙌👌👍🔥🙃😎😁🤗"
}

func ExampleCut() {
	data := []byte("Today we gathered moss for my uncle's wedding.")
	split := func(a, b []byte) {
		be, af, found := bytes.Cut(a, b)
		switch found {
		case true:
			fmt.Printf("✅ before: %q, after: %q, found: %t\n", be, af, found)
		case false:
			fmt.Printf("❌ could NOT find %q in %q:\n\tbefore: %q, after: %q\n", b, a, be, af)
		}
	}
	split(data, []byte("my"))
	split(data, []byte(" "))
	split(data, []byte("XXX"))
	// Output:
	// ✅ before: "Today we gathered moss for ", after: " uncle's wedding.", found: true
	// ✅ before: "Today", after: "we gathered moss for my uncle's wedding.", found: true
	// ❌ could NOT find "XXX" in "Today we gathered moss for my uncle's wedding.":
	// 	before: "Today we gathered moss for my uncle's wedding.", after: ""
}

func ExampleEqual() {
	data := []byte("Peter found road kill an excellent way to save money on dinner.")
	same := func(a, b []byte) {
		switch bytes.Equal(a, b) {
		case true:
			fmt.Printf("✅ %q is equal to %q\n", a, b)
		case false:
			fmt.Printf("❌ %q is NOT equal to %q\n", a, b)
		}
	}
	same([]byte(""), nil)
	same([]byte(""), []byte(""))
	same([]byte{}, []byte(""))
	same(data, []byte("Peter found road kill an excellent way to save money on dinner."))
	same(data, []byte("peter found road kill an excellent way to save money on dinner."))
	// Output:
	// ✅ "" is equal to ""
	// ✅ "" is equal to ""
	// ✅ "" is equal to ""
	// ❌ "Peter found road kill an excellent way to save money on dinner." is NOT equal to "Peter found road kill an excellent way to save\u00a0money on dinner."
	// ❌ "Peter found road kill an excellent way to save money on dinner." is NOT equal to "peter found road kill an excellent way to save money on dinner."
}

func ExampleEqualFold() {
	data := []byte("There's probably enough glass in my cupboard to build an undersea aquarium.")
	same := func(a, b []byte) {
		switch bytes.EqualFold(a, b) {
		case true:
			fmt.Printf("✅ %q is case-insensitively equal to %q\n", a, b)
		case false:
			fmt.Printf("❌ %q is NOT case-insensitively equal to %q\n", a, b)
		}
	}
	same([]byte(""), nil)
	same([]byte(""), []byte(""))
	same([]byte{}, nil)
	same([]byte("aaa"), []byte("AAA"))
	same(data, []byte("THERE'S pRoBabLY eNOuGh GlAsS iN my CuPBOaRd To BuILd AN uNdErSeA AqUaRiUM."))
	// Output:
	// ✅ "" is case-insensitively equal to ""
	// ✅ "" is case-insensitively equal to ""
	// ✅ "" is case-insensitively equal to ""
	// ✅ "aaa" is case-insensitively equal to "AAA"
	// ✅ "There's probably enough glass in my cupboard to build an undersea aquarium." is case-insensitively equal to "THERE'S pRoBabLY eNOuGh GlAsS iN my CuPBOaRd To BuILd AN uNdErSeA AqUaRiUM."
}

func ExampleFields() {
	// NOTE(jay): All acceptable field separators
	// 	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
	fields := func(a []byte) {
		fmt.Println("process fields:")
		for i, f := range bytes.Fields(a) {
			fmt.Printf("%d: %q\n", i, f)
		}
	}
	fields([]byte("I'm\tworried\nby\vthe\ffact\rthat my\u0085daughter\u00a0looks to the local\r\n\r\ncarpet\t\t\t\t\tseller                   as a role model."))
	fields([]byte("   \t \n \v \f \r \u0085 \u00a0        "))
	// Output:
	// process fields:
	// 0: "I'm"
	// 1: "worried"
	// 2: "by"
	// 3: "the"
	// 4: "fact"
	// 5: "that"
	// 6: "my"
	// 7: "daughter"
	// 8: "looks"
	// 9: "to"
	// 10: "the"
	// 11: "local"
	// 12: "carpet"
	// 13: "seller"
	// 14: "as"
	// 15: "a"
	// 16: "role"
	// 17: "model."
	// process fields:
}

func ExampleFieldsFunc() {
	dateSep := func(a []byte) {
		fmt.Println("process fields:")
		for i, f := range bytes.FieldsFunc(a, func(r rune) bool {
			return r == '-' || r == '/' || unicode.IsSpace(r)
		}) {
			fmt.Printf("%d: %q\n", i, f)
		}
	}
	dateSep([]byte("8/1/33"))
	dateSep([]byte("8-1-33"))
	dateSep([]byte("8-1-33 12-20-20    03/04/22 \t 06-/06-/07 "))
	// Output:
	// process fields:
	// 0: "8"
	// 1: "1"
	// 2: "33"
	// process fields:
	// 0: "8"
	// 1: "1"
	// 2: "33"
	// process fields:
	// 0: "8"
	// 1: "1"
	// 2: "33"
	// 3: "12"
	// 4: "20"
	// 5: "20"
	// 6: "03"
	// 7: "04"
	// 8: "22"
	// 9: "06"
	// 10: "06"
	// 11: "07"
}

func ExampleHasPrefix() {
	data := []byte("I want a giraffe, but I'm a turtle eating waffles.")
	prefixed := func(a, b []byte) {
		switch bytes.HasPrefix(a, b) {
		case true:
			fmt.Printf("✅ %q has %q as a prefix\n", a, b)
		case false:
			fmt.Printf("❌ %q does NOT have %q as a prefix\n", a, b)
		}
	}
	prefixed(data, []byte("I want"))
	prefixed(data, []byte("I want waffle fries"))
	prefixed(data, []byte("I want a giraffe, but I'm a turtle eating waffles."))
	// Output:
	// ✅ "I want a giraffe, but I'm a turtle eating waffles." has "I want" as a prefix
	// ❌ "I want a giraffe, but I'm a turtle eating waffles." does NOT have "I want waffle fries" as a prefix
	// ✅ "I want a giraffe, but I'm a turtle eating waffles." has "I want a giraffe, but I'm a turtle eating waffles." as a prefix
}

func ExampleHasSuffix() {
	data := []byte("If my calculator had a history, it would be more embarrassing than my browser history.")
	suffixed := func(a, b []byte) {
		switch bytes.HasSuffix(a, b) {
		case true:
			fmt.Printf("✅ %q has %q as a suffix\n", a, b)
		case false:
			fmt.Printf("❌ %q does NOT have %q as a suffix\n", a, b)
		}
	}
	suffixed(data, []byte("history."))
	suffixed(data, []byte("embarrassing my browser history."))
	suffixed(data, []byte("If my calculator had a history, it would be more embarrassing than my browser history."))
	// Output:
	// ✅ "If my calculator had a history, it would be more embarrassing than my browser history." has "history." as a suffix
	// ❌ "If my calculator had a history, it would be more embarrassing than my browser history." does NOT have "embarrassing my browser history." as a suffix
	// ✅ "If my calculator had a history, it would be more embarrassing than my browser history." has "If my calculator had a history, it would be more embarrassing than my browser history." as a suffix
}

func ExampleIndex() {
	data := []byte("He felt that dining on the bridge brought romance to his relationship with his cat.")
	idxOf := func(a, b []byte) {
		i := bytes.Index(a, b)
		if i == -1 {
			fmt.Printf("❌ %q could not be found in %q, index is %d\n", b, a, i)
			return
		}
		fmt.Printf("✅ %q starts at index: %d of %q\n\tyield: %q\n", b, i, a, a[i:])
	}
	idxOf(data, []byte("H"))
	idxOf(data, []byte("He"))
	idxOf(data, []byte("he"))
	idxOf(data, []byte("cat"))
	idxOf(data, []byte("dog"))
	// Output:
	// ✅ "H" starts at index: 0 of "He felt that dining on the bridge brought romance to his relationship with his cat."
	// 	yield: "He felt that dining on the bridge brought romance to his relationship with his cat."
	// ✅ "He" starts at index: 0 of "He felt that dining on the bridge brought romance to his relationship with his cat."
	// 	yield: "He felt that dining on the bridge brought romance to his relationship with his cat."
	// ✅ "he" starts at index: 24 of "He felt that dining on the bridge brought romance to his relationship with his cat."
	// 	yield: "he bridge brought romance to his relationship with his cat."
	// ✅ "cat" starts at index: 79 of "He felt that dining on the bridge brought romance to his relationship with his cat."
	// 	yield: "cat."
	// ❌ "dog" could not be found in "He felt that dining on the bridge brought romance to his relationship with his cat.", index is -1
}

func ExampleIndexAny() {
	data := []byte("Shakespeare was a famous 17th-century diesel mechanic.")
	findAny := func(a []byte, chars string) {
		i := bytes.IndexAny(a, chars)
		if i == -1 {
			fmt.Printf("❌ %q does NOT contain any of these characters %q\n", a, chars)
			return
		}
		fmt.Printf("✅ %q contains one of these characters %q at index: %d and is %q\n",
			a, chars, i, a[i])
	}
	findAny(data, "Smechanic")
	findAny(data, "")
	findAny(data, "zy18*/q")
	findAny(data, "zx8*/q")
	// Output:
	// ✅ "Shakespeare was a famous 17th-century diesel mechanic." contains one of these characters "Smechanic" at index: 0 and is 'S'
	// ❌ "Shakespeare was a famous 17th-century diesel mechanic." does NOT contain any of these characters ""
	// ✅ "Shakespeare was a famous 17th-century diesel mechanic." contains one of these characters "zy18*/q" at index: 25 and is '1'
	// ❌ "Shakespeare was a famous 17th-century diesel mechanic." does NOT contain any of these characters "zx8*/q"
}

func ExampleIndexByte() {
	data := []byte("It's not often you find a soggy banana on the street.")
	byteMe := func(a []byte, b byte) {
		i := bytes.IndexByte(a, b)
		if i == -1 {
			fmt.Printf("%q does not have byte: %q\n", a, b)
			return
		}
		fmt.Printf("%q is found at %d of %q and yields: %q\n", b, i, a, a[i:])
	}
	byteMe(data, 'b')
	byteMe(data, 'x')
	// Output:
	// 'b' is found at 32 of "It's not often you find a soggy banana on the street." and yields: "banana on the street."
	// "It's not often you find a soggy banana on the street." does not have byte: 'x'
}

func ExampleIndexFunc() {
	data := []byte("I caught my squirrel rustling through my gym bag.")
	nonPrintable := func(a []byte) {
		i := bytes.IndexFunc(a, func(r rune) bool {
			return !unicode.IsPrint(r)
		})
		if i == -1 {
			fmt.Printf("❌ %q does NOT contain any non-printable characters\n", a)
			return
		}
		fmt.Printf("✅ %q contains a non-printable character at %d and is %q\n", a, i, a[i])
	}
	data[20] = '\t'
	nonPrintable(data)
	data[20] = ' '
	nonPrintable(data)
	// Output:
	// ✅ "I caught my squirrel\trustling through my gym bag." contains a non-printable character at 20 and is '\t'
	// ❌ "I caught my squirrel rustling through my gym bag." does NOT contain any non-printable characters
}

func ExampleIndexRune() {
	data := []byte("Nobody questions who built the pyramids in Mexico🥑.")
	runed := func(a []byte, r rune) {
		i := bytes.IndexRune(a, r)
		if i == -1 {
			fmt.Printf("❌ %q does NOT contain the rune %q\n", a, r)
			return
		}
		fmt.Printf("✅ %q contains the rune at index: %d\n", a, i)
	}
	runed(data, '🥑')
	runed(data, '🌮')
	// Output:
	// ✅ "Nobody questions who built the pyramids in Mexico🥑." contains the rune at index: 49
	// ❌ "Nobody questions who built the pyramids in Mexico🥑." does NOT contain the rune '🌮'
}

func ExampleJoin() {
	data := [][]byte{
		{73},
		{97, 108, 119, 97, 121, 115},
		{100, 114, 101, 97, 109, 101, 100},
		{97, 98, 111, 117, 116},
		{98, 101, 105, 110, 103},
		{115, 116, 114, 97, 110, 100, 101, 100},
		{111, 110},
		{97},
		{100, 101, 115, 101, 114, 116},
		{105, 115, 108, 97, 110, 100},
		{117, 110, 116, 105, 108},
		{105, 116},
		{97, 99, 116, 117, 97, 108, 108, 121},
		{104, 97, 112, 112, 101, 110, 101, 100, 46},
	}
	fmt.Printf("joined bytes: %q\n", bytes.Join(data, []byte(" ")))

	oddWS := []byte("Nobody has\nencountered an explosive\r\ndaisy\t\tand lived to            tell the tale.")
	fmt.Printf("sanitized bytes: %q\n", bytes.Join(bytes.Fields(oddWS), []byte(" ")))

	// Output:
	// joined bytes: "I always dreamed about being stranded on a desert island until it actually happened."
	// sanitized bytes: "Nobody has encountered an explosive daisy and lived to tell the tale."
}

func ExampleLastIndex() {
	data := []byte("The fog was so dense even a laser decided it wasn't worth the effort.")
	lastIdx := func(a, b []byte) {
		i := bytes.LastIndex(a, b)
		if i == -1 {
			fmt.Printf("❌ %q could not be found in %q, index is %d\n", b, a, i)
			return
		}
		fmt.Printf("✅ last occurence of %q starts at index: %d of %q\n\tyield: %q\n", b, i, a, a[i:])
	}
	lastIdx(data, []byte("f"))
	lastIdx(data, []byte("XXX"))
	// Output:
	// ✅ last occurence of "f" starts at index: 64 of "The fog was so dense even a laser decided it wasn't worth the effort."
	// 	yield: "fort."
	// ❌ "XXX" could not be found in "The fog was so dense even a laser decided it wasn't worth the effort.", index is -1
}

func ExampleLastIndexAny() {
	data := []byte("Patricia found the meaning of life in a bowl of Cheerios.")
	findAny := func(a []byte, chars string) {
		i := bytes.LastIndexAny(a, chars)
		if i == -1 {
			fmt.Printf("❌ %q does NOT contain any of these characters %q\n", a, chars)
			return
		}
		fmt.Printf("✅ %q contains one of these characters %q last found at index: %d and is %q\n",
			a, chars, i, a[i])
	}
	findAny(data, "PCheerios")
	findAny(data, "")
	findAny(data, "zxy*/q")
	// Output:
	// ✅ "Patricia found the meaning of life in a bowl of Cheerios." contains one of these characters "PCheerios" last found at index: 55 and is 's'
	// ❌ "Patricia found the meaning of life in a bowl of Cheerios." does NOT contain any of these characters ""
	// ❌ "Patricia found the meaning of life in a bowl of Cheerios." does NOT contain any of these characters "zxy*/q"
}

func ExampleLastIndexByte() {
	data := []byte("Be careful with that butter knife.")
	byteMeLast := func(a []byte, b byte) {
		i := bytes.LastIndexByte(a, b)
		if i == -1 {
			fmt.Printf("%q does not have byte: %q\n", a, b)
			return
		}
		fmt.Printf("%q is found at last index: %d of %q and yields: %q\n", b, i, a, a[i:])
	}
	byteMeLast(data, 'f')
	byteMeLast(data, 'x')
	// Output:
	// 'f' is found at last index: 31 of "Be careful with that butter knife." and yields: "fe."
	// "Be careful with that butter knife." does not have byte: 'x'
}

func ExampleLastIndexFunc() {
	data := []byte("I ate a sock... because people on the Internet told me to.")
	lastPeriod := func(a []byte) {
		i := bytes.LastIndexFunc(a, func(r rune) bool {
			return r == '.'
		})
		if i == -1 {
			fmt.Printf("❌ %q does NOT contain a '.'\n", a)
			return
		}
		fmt.Printf("✅ %q contains a '.' at last index: %d and yields %q\n", a, i, a[i:])
	}
	data[len(data)-1] = ','
	lastPeriod(data)
	data[len(data)-1] = '.'
	lastPeriod(data)
	lastPeriod(bytes.ReplaceAll(data, []byte("."), []byte("")))
	// Output:
	// ✅ "I ate a sock... because people on the Internet told me to," contains a '.' at last index: 14 and yields ". because people on the Internet told me to,"
	// ✅ "I ate a sock... because people on the Internet told me to." contains a '.' at last index: 57 and yields "."
	// ❌ "I ate a sock because people on the Internet told me to" does NOT contain a '.'
}

func ExampleMap() {
	data := []byte("Th)e l!yric^s., of the213 so*ng() sound230ed li$$ke fi===n93gerna;ils on a cha_+}\\lkboa<]rd.")
	fmt.Printf("from:\t%q\nto:\t%q",
		data,
		bytes.Map(
			func(r rune) rune {
				switch {
				case !unicode.IsLetter(r) && !unicode.IsSpace(r):
					return -1
				case r > 'm':
					return unicode.ToUpper(r)
				default:
					return r
				}
			}, data),
	)
	// Output:
	// from:	"Th)e l!yric^s., of the213 so*ng() sound230ed li$$ke fi===n93gerna;ils on a cha_+}\\lkboa<]rd."
	// to:	"The lYRicS Of The SONg SOUNded like fiNgeRNailS ON a chalkbOaRd"
}

func ExampleRepeat() {
	data := []byte("The glacier came alive as the climbers hiked closer.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %q\n", r)
			fmt.Printf("%q\n", bytes.Repeat(data, 2))
		}
	}()
	fmt.Println(bytes.Repeat(data, 1<<62)) // XXX(jay): This panics
	// Output:
	// panic: "bytes: Repeat count causes overflow"
	// "The glacier came alive as the climbers hiked closer.The glacier came alive as the climbers hiked closer."
}

func ExampleReplace() {
	data := []byte("It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general.")
	etp := func(a, old, _new []byte, n int) {
		fmt.Printf("orig:\t%q\nnew:\t%q\n", a, bytes.Replace(a, old, _new, n))
		fmt.Println("-------------", "replaced", n, "-------------")
	}
	etp(data, []byte("difficult"), []byte("easy😎"), 1)
	etp(data, []byte("difficult"), []byte("easy😎"), 2)
	etp(data, []byte("tennis"), []byte("football🏈"), -1) // == bytes.ReplaceAll
	// Output:
	// orig:	"It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general."
	// new:	"It's much more easy😎 to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general."
	// ------------- replaced 1 -------------
	// orig:	"It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general."
	// new:	"It's much more easy😎 to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is easy😎 in general."
	// ------------- replaced 2 -------------
	// orig:	"It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general."
	// new:	"It's much more difficult to play football🏈 with a bowling ball than it is to bowl with a football🏈 ball, but football🏈 is difficult in general."
	// ------------- replaced -1 -------------
}

func ExampleReplaceAll() {
	data := []byte("It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general.")
	fmt.Printf("orig:\t%q\nnew:\t%q\n",
		data, bytes.ReplaceAll(data, []byte("tennis"), []byte("football🏈")),
	)
	// Output:
	// orig:	"It's much more difficult to play tennis with a bowling ball than it is to bowl with a tennis ball, but tennis is difficult in general."
	// new:	"It's much more difficult to play football🏈 with a bowling ball than it is to bowl with a football🏈 ball, but football🏈 is difficult in general."
}

func ExampleRunes() {
	data := []byte("😄👋😂💬🎉🤔👀")
	fmt.Printf("useless:\t%v\nmore info:\t%U\n", data, bytes.Runes(data))
	// Output:
	// useless:	[240 159 152 132 240 159 145 139 240 159 152 130 240 159 146 172 240 159 142 137 240 159 164 148 240 159 145 128]
	// more info:	[U+1F604 U+1F44B U+1F602 U+1F4AC U+1F389 U+1F914 U+1F440]
}

func ExampleSplit() {
	data := []byte("He🍌loved🍌eating🍌his🍌bananas🍌in🍌hot🍌dog🍌buns.")
	fmt.Printf("%q\n", bytes.Split(data, []byte("🍌")))
	fmt.Printf("%q\n", bytes.Split(data, nil))
	// Output:
	// ["He" "loved" "eating" "his" "bananas" "in" "hot" "dog" "buns."]
	// ["H" "e" "🍌" "l" "o" "v" "e" "d" "🍌" "e" "a" "t" "i" "n" "g" "🍌" "h" "i" "s" "🍌" "b" "a" "n" "a" "n" "a" "s" "🍌" "i" "n" "🍌" "h" "o" "t" "🍌" "d" "o" "g" "🍌" "b" "u" "n" "s" "."]
}

func ExampleSplitAfter() {
	data := []byte("The🦃Guinea🦃fowl🦃flies🦃through🦃the🦃air🦃with🦃all🦃the🦃grace🦃of🦃a🦃turtle.")
	fmt.Printf("%q\n", bytes.SplitAfter(data, []byte("🦃")))
	fmt.Printf("%q\n", bytes.SplitAfter(data, []byte{}))
	// Output:
	// ["The🦃" "Guinea🦃" "fowl🦃" "flies🦃" "through🦃" "the🦃" "air🦃" "with🦃" "all🦃" "the🦃" "grace🦃" "of🦃" "a🦃" "turtle."]
	// ["T" "h" "e" "🦃" "G" "u" "i" "n" "e" "a" "🦃" "f" "o" "w" "l" "🦃" "f" "l" "i" "e" "s" "🦃" "t" "h" "r" "o" "u" "g" "h" "🦃" "t" "h" "e" "🦃" "a" "i" "r" "🦃" "w" "i" "t" "h" "🦃" "a" "l" "l" "🦃" "t" "h" "e" "🦃" "g" "r" "a" "c" "e" "🦃" "o" "f" "🦃" "a" "🦃" "t" "u" "r" "t" "l" "e" "."]
}

func ExampleSplitAfterN() {
	data := []byte("He🎭had🎭a🎭wall🎭full🎭of🎭masks🎭so🎭she🎭could🎭wear🎭a🎭different🎭face🎭every🎭day.")
	splits := func(a, b []byte, n int) {
		nsplits := bytes.SplitAfterN(a, b, n)
		fmt.Printf("len: %d, values: %q\n", len(nsplits), nsplits)
	}
	splits(data, []byte("🎭"), 0)  // useful for obfuscation?
	splits(data, []byte("🎭"), -1) // == bytes.SplitAfter
	splits(data, []byte("🎭"), 4)
	splits(data, []byte("🎭"), 20)
	// Output:
	// len: 0, values: []
	// len: 16, values: ["He🎭" "had🎭" "a🎭" "wall🎭" "full🎭" "of🎭" "masks🎭" "so🎭" "she🎭" "could🎭" "wear🎭" "a🎭" "different🎭" "face🎭" "every🎭" "day."]
	// len: 4, values: ["He🎭" "had🎭" "a🎭" "wall🎭full🎭of🎭masks🎭so🎭she🎭could🎭wear🎭a🎭different🎭face🎭every🎭day."]
	// len: 16, values: ["He🎭" "had🎭" "a🎭" "wall🎭" "full🎭" "of🎭" "masks🎭" "so🎭" "she🎭" "could🎭" "wear🎭" "a🎭" "different🎭" "face🎭" "every🎭" "day."]
}

func ExampleSplitN() {
	data := []byte("Purple💜is💜the💜best💜city💜in💜the💜forest.")
	splits := func(a, b []byte, n int) {
		nsplits := bytes.SplitN(a, b, n)
		fmt.Printf("len: %d, values: %q\n", len(nsplits), nsplits)
	}
	splits(data, []byte("💜"), 0)  // useful for obfuscation?
	splits(data, []byte("💜"), -1) // == bytes.SplitAfter
	splits(data, []byte("💜"), 3)
	splits(data, []byte("💜"), 10)
	// Output:
	// len: 0, values: []
	// len: 8, values: ["Purple" "is" "the" "best" "city" "in" "the" "forest."]
	// len: 3, values: ["Purple" "is" "the💜best💜city💜in💜the💜forest."]
	// len: 8, values: ["Purple" "is" "the" "best" "city" "in" "the" "forest."]
}

func ExampleTitle() {
	// XXX(jay): Don't use this! DEPRECATED
	// bytes.Title()
	src := []byte("The rain pelted the windshield as the darkness engulfed us.")
	dst := make([]byte, len(src))
	// Use this in it's place.
	nDst, nSrc, err := cases.Title(language.English, cases.NoLower).Transform(dst, src, true)
	fmt.Printf("orig:\t%q\nresult:\t%q\nnDst: %d\tnSrc: %d\terr: %v\n",
		src, dst, nDst, nSrc, err)
	// Output:
	// orig:	"The rain pelted the windshield as the darkness engulfed us."
	// result:	"The Rain Pelted The Windshield As The Darkness Engulfed Us."
	// nDst: 59	nSrc: 59	err: <nil>
}

func ExampleToLower() {
	data := []byte("PEopLE wHo InsIst on pICKInG tHEIr tEEtH wItH tHEIr ELBows ArE so AnnoyInG!")
	fmt.Printf("before:\t%q\nafter:\t%q\n", data, bytes.ToLower(data))
	// Output:
	// before:	"PEopLE wHo InsIst on pICKInG tHEIr tEEtH wItH tHEIr ELBows ArE so AnnoyInG!"
	// after:	"people who insist on picking their teeth with their elbows are so annoying!"
}

func ExampleToLowerSpecial() {
	fmt.Printf("%q\n", bytes.ToLowerSpecial(unicode.TurkishCase, []byte("Genç ÇOCUK, saDEce teSTTen çıKMak IÇin kolunu kırmakla suÇlandı.")))
	// Output:
	// "genç çocuk, sadece testten çıkmak ıçin kolunu kırmakla suçlandı."
}

func ExampleToTitle() {
	// NOTE(jay): Compare with bytes.ToUpper and bytes.Title
	data := []byte("He found a leprechaun in his walnut shell.")
	data2 := []byte("он нашел лепрекона в скорлупе грецкого ореха.")
	comp := func(a []byte) {
		fmt.Printf("before:\t%q\nafter:\t%q\n", a, bytes.ToTitle(a))
	}
	comp(data)
	comp(data2)
	// Output:
	// before:	"He found a leprechaun in his walnut shell."
	// after:	"HE FOUND A LEPRECHAUN IN HIS WALNUT SHELL."
	// before:	"он нашел лепрекона в скорлупе грецкого ореха."
	// after:	"ОН НАШЕЛ ЛЕПРЕКОНА В СКОРЛУПЕ ГРЕЦКОГО ОРЕХА."
}

func ExampleToTitleSpecial() {
	fmt.Printf("%q\n", bytes.ToTitleSpecial(unicode.TurkishCase, []byte("Genç ÇOCUK, saDEce teSTTen çıKMak IÇin kolunu kırmakla suÇlandı.")))
	// Output:
	// "GENÇ ÇOCUK, SADECE TESTTEN ÇIKMAK IÇİN KOLUNU KIRMAKLA SUÇLANDI."
}

func ExampleToUpper() {
	// NOTE(jay): Compare with bytes.ToTitle and bytes.Title
	data := []byte("He found a leprechaun in his walnut shell.")
	data2 := []byte("он нашел лепрекона в скорлупе грецкого ореха.")
	comp := func(a []byte) {
		fmt.Printf("before:\t%q\nafter:\t%q\n", a, bytes.ToUpper(a))
	}
	comp(data)
	comp(data2)
	// Output:
	// before:	"He found a leprechaun in his walnut shell."
	// after:	"HE FOUND A LEPRECHAUN IN HIS WALNUT SHELL."
	// before:	"он нашел лепрекона в скорлупе грецкого ореха."
	// after:	"ОН НАШЕЛ ЛЕПРЕКОНА В СКОРЛУПЕ ГРЕЦКОГО ОРЕХА."
}

func ExampleToUpperSpecial() {
	fmt.Printf("%q\n", bytes.ToUpperSpecial(unicode.TurkishCase, []byte("Genç ÇOCUK, saDEce teSTTen çıKMak IÇin kolunu kırmakla suÇlandı.")))
	// Output:
	// "GENÇ ÇOCUK, SADECE TESTTEN ÇIKMAK IÇİN KOLUNU KIRMAKLA SUÇLANDI."
}

func ExampleToValidUTF8() {
	// NOTE(jay): Bad bytes taken from
	// 	https://www.cl.cam.ac.uk/~mgk25/ucs/examples/UTF-8-test.txt
	data := []byte{
		0x54,
		0xc0, 0xaf, // Bad bytes
		0x6f, 0x64, 0x61,
		0xff, // Bad byte
		0x79, 0x20, 0x69,
		0xfe, // Bad byte
		0x73, 0x20,
		0xc1, 0xbf, // Bad bytes
		0x74, 0x68, 0x65, 0x20,
		0xf8, 0x87, 0xbf, 0xbf, // Bad bytes
		0x64, 0x61, 0x79, 0x20, 0x49, 0x27, 0x6c, 0x6c, 0x20, 0x66,
		0xc0, 0x80, 0xfc, 0x80, 0x80, 0x80, 0x80, // Bad bytes
		0x69, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x20, 0x6b, 0x6e, 0x6f, 0x77, 0x20,
		0xed, 0xa0, 0x80, // Bad bytes
		0x77, 0x68, 0x61, 0x74, 0x20, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x20, 0x74,
		0xed, 0xa0, 0x80, 0xed, 0xb0, 0x80, // Bad bytes
		0x61, 0x73, 0x74, 0x65, 0x73, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	}
	fmt.Printf("%q\n", bytes.ToValidUTF8(data, []byte("")))
	fmt.Printf("%q\n", bytes.ToValidUTF8(data, []byte("🥴")))
	// Output:
	// "Today is the day I'll finally know what brick tastes like."
	// "T🥴oda🥴y i🥴s 🥴the 🥴day I'll f🥴inally know 🥴what brick t🥴astes like."
}

func ExampleTrim() {
	data := []byte("🍞🥜🥜🍇🍞👵Peanut butter and jelly caused the elderly lady to think about her past.🍞🥜🥜🍓🍓🍞👵")
	fmt.Printf("trimmed: %q", bytes.Trim(data, "👵🍇🍞🥜🍓"))
	// Output:
	// trimmed: "Peanut butter and jelly caused the elderly lady to think about her past."
}

func ExampleTrimFunc() {
	data := []byte(`
/###%%%#####%%%#####'#####'#####'#####'###'#####'#####'#####%%%#####%%%####\
*###%%%#####%%%#############################################%%%#####%%%####*
*###%%%### The secret ingredient to his wonderful life was crime. ##%%%####*
*####%#######%###############################################%#######%#####*
\###'#'#####'#'#############################################'#'#####'#'####/
`[1:])
	fmt.Printf("cleaned: %q\n", bytes.TrimFunc(data, func(r rune) bool {
		return r == '\n' || r == ' ' || r == '#' || r == '\\' ||
			r == '*' || r == '\'' || r == '/' || r == '%'
	}))
	// Output:
	// cleaned: "The secret ingredient to his wonderful life was crime."
}

func ExampleTrimLeft() {
	data := []byte("🕵️‍♂️📔🎤🗣The mysterious diary records the voice.")
	fmt.Printf("trimmed: %q\n", bytes.TrimLeft(data, "📔🕵️‍🗣♂️🎤"))
	// Output:
	// trimmed: "The mysterious diary records the voice."
}

func ExampleTrimLeftFunc() {
	data := []byte(`
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%             _       _                  %
%  _ __   ___| |_ ___| |__   ___  _ __   %
% | '_ \ / _ \ __/ __| '_ \ / _ \| '_ \  %
% | |_) |  __/ |_\__ \ | | | (_) | |_) | %
% | .__/ \___|\__|___/_| |_|\___/| .__/  %
% |_|                            |_|     %
%                                        %
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
The pet shop stocks everything you need to keep your anaconda happy.`)
	fmt.Printf("trimmed: %q\n", bytes.TrimLeftFunc(data, func(r rune) bool {
		return unicode.In(r, unicode.Punct, unicode.White_Space) || r == '|'
	}))
	// Output:
	// trimmed: "The pet shop stocks everything you need to keep your anaconda happy."
}

func ExampleTrimPrefix() {
	data := []byte("He had reached the point where he was paranoid about being paranoid.")
	fmt.Printf("trimmed: %q\n", bytes.TrimPrefix(data, []byte("He had reached the point where")))
	fmt.Printf("trimmed: %q\n", bytes.TrimPrefix(data, []byte("He had reached")))
	fmt.Printf("trimmed: %q\n", bytes.TrimPrefix(data, []byte("XXX")))
	// Output:
	// trimmed: " he was paranoid about being paranoid."
	// trimmed: " the point where he was paranoid about being paranoid."
	// trimmed: "He had reached the point where he was paranoid about being paranoid."
}

func ExampleTrimRight() {
	fmt.Printf("trimmed: %q", bytes.TrimRight(
		[]byte("Imagine his surprise when he discovered that the safe was full of pudding.🫃🍮🍮🍮🍮🍮🍮🍮🍮🍮"),
		"🍮🫃🧷"))
	// Output:
	// trimmed: "Imagine his surprise when he discovered that the safe was full of pudding."
}

func ExampleTrimRightFunc() {
	data := []byte(`Harrold felt confident that nobody would ever suspect his spy pigeon.
			-- Narrator 04/20/69`)
	fmt.Printf("trimmed: %q\n", bytes.TrimRightFunc(data, func(r rune) bool {
		return unicode.In(r, unicode.Letter, unicode.Digit, unicode.White_Space) ||
			r == '-' || r == '/'
	}))
	// Output:
	// trimmed: "Harrold felt confident that nobody would ever suspect his spy pigeon."
}

func ExampleTrimSpace() {
	data := []byte("\t \n \v \f \r \u0085 \u00a0   You bite up because of your lower jaw.😬   \u00a0 \u0085 \r \f \v \n \t")
	fmt.Printf("with space:\t%q\nwithout:\t%q\n", data, bytes.TrimSpace(data))
	// Output:
	// with space:	"\t \n \v \f \r \u0085 \u00a0   You bite up because of your lower jaw.😬   \u00a0 \u0085 \r \f \v \n \t"
	// without:	"You bite up because of your lower jaw.😬"
}

func ExampleTrimSuffix() {
	data := []byte("He put heat on the wound to see what would grow.")
	fmt.Printf("trimmed: %q\n", bytes.TrimSuffix(data, []byte("what would grow.")))
	fmt.Printf("trimmed: %q\n", bytes.TrimSuffix(data, []byte("grow.")))
	fmt.Printf("trimmed: %q\n", bytes.TrimSuffix(data, []byte("He put")))
	// Output:
	// trimmed: "He put heat on the wound to see "
	// trimmed: "He put heat on the wound to see what would "
	// trimmed: "He put heat on the wound to see what would grow."
}
