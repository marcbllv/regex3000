package state

type FinalInspector struct{}

func (inspector FinalInspector) Match(str string) (bool, string) {
	return true, str
}

func (inspector FinalInspector) Copy() Inspector {
	return FinalInspector{}
}
