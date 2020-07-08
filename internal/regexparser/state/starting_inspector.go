package state

type StartingInspector struct{}

func (inspector StartingInspector) Match(str string) (bool, string) {
	return true, str
}

func (inspector StartingInspector) Copy() Inspector {
	return StartingInspector{}
}
