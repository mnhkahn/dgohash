package runtime_test

import (
	. "runtime"
	"testing"
)

func BenchmarkMemHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, g := range goldenMurmur3 {
			StringHash(g.in, 0)
		}
	}
}

type _Golden struct {
	out uint32
	in  string
}

var goldenMurmur3 = []_Golden{
	{0x00000000, ""},
	{0x3c2569b2, "a"},
	{0x9bbfd75f, "ab"},
	{0xb3dd93fa, "abc"},
	{0x43ed676a, "abcd"},
	{0xe89b9af6, "abcde"},
	{0x6181c085, "abcdef"},
	{0x883c9b06, "abcdefg"},
	{0x49ddccc4, "abcdefgh"},
	{0x421406f0, "abcdefghi"},
	{0x88927791, "abcdefghij"},
	{0x91e056d3, "Discard medicine more than two years old."},
	{0xc4d1cdf9, "He who has a shady past knows that nice guys finish last."},
	{0x92a09da9, "I wouldn't marry him with a ten foot pole."},
	{0xba22e6c4, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0xb3ba11cb, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0x941ada4d, "Nepal premier won't resign."},
	{0x03f1f7b4, "For every action there is an equal and opposite government program."},
	{0x03946117, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0x91e89ce1, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0xdc39bd00, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0xe898a1fa, "size:  a.out:  bad magic"},
	{0xcb5affb4, "The major problem is with sendmail.  -Mark Horton"},
	{0xc84510d4, "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	{0xd4466554, "If the enemy is within range, then so are you."},
	{0xe718d618, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0xa6fb1684, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0x65cb8d60, "C is as portable as Stonehedge!!"},
	{0x164935d1, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0x33e03966, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0x04944630, "How can you write a big system without C++?  -Paul Glick"},
}
