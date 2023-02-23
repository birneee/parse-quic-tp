package internal

func InvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	inv := make(map[V]K)
	for k, v := range m {
		inv[v] = k
	}
	return inv

}
