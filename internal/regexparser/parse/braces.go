package parse

type RangeQuantifier struct {
	Low  int
	High int
}

type RangeParser struct{}

func (rangeParser RangeParser) Parse(str []rune, pos int) (RangeQuantifier, int) {
	var currentPos int
	openBraceParser := RuneParser{r: '{'}
	lowerBoundParser := IntParser{}
	commaParser := RuneParser{r: ',', optional: true}
	higherBoundParser := IntParser{optional: true}
	closingBraceParser := RuneParser{r: '}'}

	_, currentPos = openBraceParser.Parse(str, pos)
	lowerBound, currentPos := lowerBoundParser.Parse(str, currentPos)
	_, currentPos = commaParser.Parse(str, currentPos)
	higherBound, currentPos := higherBoundParser.Parse(str, currentPos)
	_, currentPos = closingBraceParser.Parse(str, currentPos)

	if higherBound == -1 {
		higherBound = lowerBound
	}
	return RangeQuantifier{Low: lowerBound, High: higherBound}, currentPos
}

func ParseBraces(regex []rune, position int) (RangeQuantifier, int) {
	rangeParser := RangeParser{}
	rangeQuantifier, currentPosition := rangeParser.Parse(regex, position)
	return rangeQuantifier, currentPosition
}
