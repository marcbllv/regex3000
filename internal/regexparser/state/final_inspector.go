package state

type FinalInspector struct {
	mustMatchEndOfString bool
}

func (inspector FinalInspector) Match(str string) (bool, string) {
	if inspector.mustMatchEndOfString && len(str) > 0 {
		return false, ""
	}
	return true, str
}

func (inspector FinalInspector) Copy() Inspector {
	return FinalInspector{}
}
