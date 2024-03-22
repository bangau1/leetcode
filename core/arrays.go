package core

func ArrayFill[T any](arr []T, val T) {
	n := len(arr)
	for i := 0; i < n; i++ {
		arr[i] = val
	}
}

func ArrayReverse[T any](arr []T) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
