package perfect

import "errors"

var ErrOnlyPositive = errors.New("only positive integers are allowed.")

type Classification string

const (
	ClassificationAbundant  Classification = "ClassificationAbundant"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationPerfect   Classification = "ClassificationPerfect"
)

// Classify returns the classification of a number according to Nicomuchus'
// classification scheme for positive integers. If n < 1 an error is returned.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	s := calculateAliquotSum(n)
	switch {
	case s > n:
		return ClassificationAbundant, nil
	case s < n:
		return ClassificationDeficient, nil
	case s == n:
		return ClassificationPerfect, nil
	default:
		panic("this should not be reachable")
	}
}

// calculateAliquotSum returns the aliquot sum of a positive integer.
func calculateAliquotSum(n int64) int64 {
	sum := int64(0)
	for factor := int64(1); factor <= n/2; factor++ {
		if n%factor == 0 {
			sum += factor
		}
	}
	return sum
}
