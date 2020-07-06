package state

type Inspector interface {
	Match(str string) (bool, string)
	Copy() Inspector
}
