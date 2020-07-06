package state

type EpsilonInspector struct{}

func (inspector EpsilonInspector) Match(str string) (bool, string) {
	return true, str
}

func (inspector EpsilonInspector) Copy() Inspector {
	return EpsilonInspector{}
}
