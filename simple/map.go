package simple

//func Map[S, R any](src []S, fn func(S) R) (dst []R) {
//	dst = make([]R, len(src))
//	for i := 0; i < len(src); i++ {
//		dst[i] = fn(src[i])
//	}
//	return
//}
//
//func Filter[T any](src []T, fn func(T) bool) (dst []T) {
//	dst = make([]T, 0, len(src))
//	for i := 0; i < len(src); i++ {
//		if fn(src[i]) {
//			dst = append(dst, src[i])
//		}
//	}
//	return
//}
