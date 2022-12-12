package rice

import (
	"sort"
	"time"
)

type Numbers interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | int | uint
}

// SliceRemoveIndex 移除 slice 中的一个元素
func SliceRemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// SliceRemoveIndexUnOrder 移除 slice 中的一个元素（无序，但效率高）
func SliceRemoveIndexUnOrder[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// RemoveDuplicates slice 去重
func RemoveDuplicates[T comparable](s1 []T) []T {

	m1 := make(map[T]struct{})

	for _, v := range s1 {
		m1[v] = struct{}{}
	}

	s2 := make([]T, 0)

	for k := range m1 {
		s2 = append(s2, k)
	}

	return s2
}

// RemoveDuplicatesInPlace slice 就地去重
func RemoveDuplicatesInPlace(userIDs []int64) []int64 {
	// if there are 0 or 1 items we return the slice itself.
	if len(userIDs) < 2 {
		return userIDs
	}

	// make the slice ascending sorted.
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	uniqPointer := 0

	for i := 1; i < len(userIDs); i++ {
		// compare a current item with the item under the unique pointer.
		// if they are not the same, write the item next to the right of the unique pointer.
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// SliceDifference 取 a 中有，而 b 中没有的
func SliceDifference[T comparable](a, b []T) []T {

	mb := make(map[T]struct{}, len(b))

	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []T

	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

// SliceDifferenceBoth 取 slice1, slice2 的差集
func SliceDifferenceBoth[T comparable](slice1, slice2 []T) []T {
	var diff []T

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

// SliceIntersection 两个 slice 的交集
func SliceIntersection[T comparable](s1, s2 []T) (inter []T) {
	hash := make(map[T]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

// Remove dups from slice.
func removeDups[T comparable](elements []T) (nodups []T) {
	encountered := make(map[T]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

// TimeExistIntersection 两个时间段是否有交集 false 没有交集，true 有交集
func TimeExistIntersection(startTime, endTime time.Time, anotherStartTime, anotherEndTime time.Time) bool {

	if anotherStartTime.After(endTime) || anotherEndTime.Before(startTime) {
		return false
	} else {
		return true
	}
}

// TimestampExistIntersection 两个时间段是否有交集 false 没有交集，true 有交集
func TimestampExistIntersection(startTime, endTime int64, anotherStartTime, anotherEndTime int64) bool {

	if endTime < anotherStartTime || startTime > anotherEndTime {
		return false
	} else {
		return true
	}
}

// SliceIn e 是否在 s 中
func SliceIn[T comparable](e T, s []T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// MaxNumber booleans, numbers, strings, pointers, channels, arrays
func MaxNumber[T Numbers](n ...T) T {

	sort.Slice(n, func(i, j int) bool { return n[i] < n[j] })

	return n[len(n)-1]
}

func MinNumber[T Numbers](n ...T) T {

	sort.Slice(n, func(i, j int) bool { return n[i] < n[j] })

	return n[0]
}

// NotIn e 不在 s 中吗？ true 不在， false 在
func NotIn[T comparable](e T, s []T) bool {

	for _, id := range s {
		if e == id {
			return false
		}
	}
	return true
}

// In e 在 s 中吗？ true 在，false 不在
func In[T comparable](e T, s []T) bool {

	for _, id := range s {
		if e == id {
			return true
		}
	}
	return false
}

// SliceReverse 反转 slice
func SliceReverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Pageination 切片分页
func Pageination[T any](page, pageSize int, s []T) []T {
	if page <= 0 {
		page = 1
	}
	if len(s) >= pageSize*(page-1) {
		if pageSize*page <= len(s) {
			s = s[pageSize*(page-1) : pageSize*page]
		} else if page-1 == 0 {
			s = s[:]
		} else if pageSize*page > len(s) {
			s = s[pageSize*(page-1):]
		}
		return s
	} else {
		return s
	}
}

// Filter filter one slice
//func Filter[T any](objs []T, filter func(obj T) bool) []T {
//	res := make([]T, 0, len(objs))
//	for i := range objs {
//		ok := filter(objs[i])
//		if ok {
//			res = append(res, objs[i])
//		}
//	}
//	return res
//}

// Map one slice
//func Map[T any, K any](objs []T, mapper func(obj T) ([]K, bool)) []K {
//	res := make([]K, 0, len(objs))
//	for i := range objs {
//		others, ok := mapper(objs[i])
//		if ok {
//			res = append(res, others...)
//		}
//	}
//	return res
//}

// First make return first for slice
func First[T any](objs []T) (T, bool) {
	if len(objs) > 0 {
		return objs[0], true
	}
	return *new(T), false
}

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type SliceIterator[T any] struct {
	Elements []T
	value    T
	index    int
}

// NewSliceIterator Create an iterator over the slice xs
func NewSliceIterator[T any](xs []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: xs,
	}
}

// Next Move to next value in collection
func (iter *SliceIterator[T]) Next() bool {
	if iter.index < len(iter.Elements) {
		iter.value = iter.Elements[iter.index]
		iter.index += 1
		return true
	}

	return false
}

// Value Get current element
func (iter *SliceIterator[T]) Value() T {
	return iter.value
}

type mapIterator[T any] struct {
	source Iterator[T]
	mapper func(T) T
}

// Next advance to next element
func (iter *mapIterator[T]) Next() bool {
	return iter.source.Next()
}

func (iter *mapIterator[T]) Value() T {
	value := iter.source.Value()
	return iter.mapper(value)
}

func Map[T any](iter Iterator[T], f func(T) T) Iterator[T] {
	return &mapIterator[T]{
		iter, f,
	}
}

type filterIterator[T any] struct {
	source Iterator[T]
	pred   func(T) bool
}

func (iter *filterIterator[T]) Next() bool {
	for iter.source.Next() {
		if iter.pred(iter.source.Value()) {
			return true
		}
	}
	return false
}

func (iter *filterIterator[T]) Value() T {
	return iter.source.Value()
}

func Filter[T any](iter Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIterator[T]{
		iter, pred,
	}
}

func Collect[T any](iter Iterator[T]) []T {
	var xs []T

	for iter.Next() {
		xs = append(xs, iter.Value())
	}

	return xs
}

type Reducer[T, V any] func(accum T, value V) T

// Reduce values iterated over to a single value
func Reduce[T, V any](iter Iterator[V], f Reducer[T, V]) T {
	var accum T
	for iter.Next() {
		accum = f(accum, iter.Value())
	}
	return accum
}
