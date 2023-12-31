package golambda

func Map[I any, O any](slice []I, converter func(index int, m I) O) []O {
	var output []O = make([]O, 0)
	for index, value := range slice {
		output = append(output, converter(index, value))
	}
	return output
}

func Filter[T any](slice []T, filter func(index int, s T) bool) []T {
	var output []T = make([]T, 0)
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

func NewMayBe[T any](d *T) MayBe[T] {
	return MayBe[T]{Data: d}
}

type MayBe[T any] struct {
	Data *T `json:"data"`
}

func (m MayBe[T]) Just(f func(data T)) MayBe[T] {
	if m.Data != nil {
		f(*m.Data)
	}
	return m
}

func (m MayBe[T]) Nothing(f func()) MayBe[T] {
	if m.Data == nil {
		f()
	}
	return m
}

func (m MayBe[T]) Value() *T {
	return m.Data
}

func Reference[T any](t T) *T {
	return &t
}

func Convert[A, B any](data A, converter func(A) B) B {
	return converter(data)
}

func ConvertMayBe[A, B any](m MayBe[A], converter func(*A) *B) MayBe[B] {
	if m.Data == nil {
		return MayBe[B]{}
	}
	return MayBe[B]{Data: converter(m.Data)}
}
