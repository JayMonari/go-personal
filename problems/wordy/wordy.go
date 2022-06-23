package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Operators maps the words used for operators with their mathmatical
// equivalent.
var Operators = map[string]int{
	"plus":       '+',
	"minus":      '-',
	"multiplied": '*',
	"divided":    '/',
}

// Answer returns a result to a spoken mathmatical question. The question must
// use basic arithmetic operators and must be infix notation else -1 and false
// is returned.
func Answer(question string) (int, bool) {
	if strings.Contains(question, "cubed") {
		return -1, false
	}

	regDigs := regexp.MustCompile("[-0-9]+")
	regOps := regexp.MustCompile("(plus|minus|divided|multiplied)")
	operands := findOpers(regDigs, question)
	operators := findOpers(regOps, question)
	if len(operands)-len(operators) != 1 {
		return -1, false
	} else if len(operands) == 0 {
		return -1, false
	}

	res, err := evaluate(operands.createIterator(), operators.createIterator())
	if err != nil {
		return -1, false
	}
	return res, true
}

// evaluate returns a result based on the passed in operands and operators. An
// error is returned if the order of indices is not operand, operator, operand,
// operator....
func evaluate(operandIt, operatorIt iterator) (int, error) {
	prevPair := operandIt.next()
	result := prevPair[1]
	for operandIt.hasNext() {
		currPair := operandIt.next()
		operatorPair := operatorIt.next()
		if prevPair[0] > operatorPair[0] || currPair[0] < operatorPair[0] {
			return -1, fmt.Errorf("unexpected notation used")
		}
		switch operatorPair[1] {
		case '+':
			result += currPair[1]
		case '-':
			result -= currPair[1]
		case '*':
			result *= currPair[1]
		case '/':
			result /= currPair[1]
		}
		prevPair = currPair
	}
	return result, nil
}

// findOpers returns all operators or operands in a string with thier starting
// index.
func findOpers(re *regexp.Regexp, question string) IdxOperCollection {
	reOps := re.FindAllString(question, -1)
	idxs := re.FindAllStringIndex(question, -1)
	ops := make(IdxOperCollection, len(reOps))
	for i, o := range reOps {
		n := convertToInt(o)
		ops[i] = [2]int{idxs[i][0], n}
	}
	return ops
}

// convertToInt ...
func convertToInt(x string) int {
	op, ok := Operators[x]
	switch {
	case ok:
		return op
	default:
		n, _ := strconv.Atoi(x)
		return n
	}
}
