package state

type Inspector interface {
	Match(str []rune, pos int) []int
	Copy() Inspector
}
