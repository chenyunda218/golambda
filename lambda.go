package golambda

func Map[I any, O any](slice []I, converter func(index int, m I) O) []O {
	var output []O
	for index, value := range slice {
		output = append(output, converter(index, value))
	}
	return output
}

func Filter[T any](slice []T, filter func(index int, s T) bool) []T {
	var output []T
	for index, value := range slice {
		if filter(index, value) {
			output = append(output, value)
		}
	}
	return output
}

func ForEach[T any](slice []T, operation func(index int, s *T)) {
	for index, value := range slice {
		operation(index, &value)
	}
}
