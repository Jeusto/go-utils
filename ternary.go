package utils

func Ternary[T any](expression bool, ifCondition T, elseCondition T) T {
	if expression {
		return ifCondition
	}
	return elseCondition
}
