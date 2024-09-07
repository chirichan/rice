package rice

import (
	"sort"
	"time"
)

type Numbers interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | int | uint
}

// RemoveItem 移除 slice 中的一个元素
func RemoveItem[T any](items []T, index int) []T {
	if index >= len(items) {
		return []T{}
	}
	return append(items[:index], items[index+1:]...)
}

// RemoveItemNoOrder 移除 slice 中的一个元素，无序
func RemoveItemNoOrder[T any](items []T, index int) []T {
	if index >= len(items) {
		return []T{}
	}
	items[index] = items[len(items)-1]
	return items[:len(items)-1]
}

// DeDuplicate slice 去重
func DeDuplicate[T comparable](items []T) []T {
	m := make(map[T]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	s := make([]T, 0)
	for k := range m {
		s = append(s, k)
	}
	return s
}

// DeDuplicateInPlace slice 就地去重
func DeDuplicateInPlace[T Numbers](items []T) []T {
	// if there are 0 or 1 items we return the slice itself.
	if len(items) < 2 {
		return items
	}

	// make the slice ascending sorted.
	sort.SliceStable(items, func(i, j int) bool { return items[i] < items[j] })

	uniqPointer := 0

	for i := 1; i < len(items); i++ {
		// compare a current item with the item under the unique pointer.
		// if they are not the same, write the item next to the right of the unique pointer.
		if items[uniqPointer] != items[i] {
			uniqPointer++
			items[uniqPointer] = items[i]
		}
	}

	return items[:uniqPointer+1]
}

// Difference 取 items1 中有，而 items2 中没有的
func Difference[T comparable](items1, items2 []T) []T {

	mb := make(map[T]struct{}, len(items2))

	for _, x := range items2 {
		mb[x] = struct{}{}
	}

	var diff []T

	for _, x := range items1 {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

// DifferenceBoth 取 slice1, slice2 的差集
func DifferenceBoth[T comparable](items1, items2 []T) []T {
	var diff []T

	// Loop two times, first to find items1 strings not in items2,
	// second loop to find items2 strings not in items1
	for i := 0; i < 2; i++ {
		for _, s1 := range items1 {
			found := false
			for _, s2 := range items2 {
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
			items1, items2 = items2, items1
		}
	}

	return diff
}

// Intersection 两个 slice 的交集
func Intersection[T comparable](s1, s2 []T) (inter []T) {
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
	// Remove dups from slice.
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

// IsTimeExistIntersection 两个时间段是否有交集 false 没有交集，true 有交集
func IsTimeExistIntersection(startTime, endTime time.Time, anotherStartTime, anotherEndTime time.Time) bool {
	if anotherStartTime.After(endTime) || anotherEndTime.Before(startTime) {
		return false
	} else {
		return true
	}
}

// IsTimestampExistIntersection 两个时间段是否有交集 false 没有交集，true 有交集
func IsTimestampExistIntersection(startTime, endTime int64, anotherStartTime, anotherEndTime int64) bool {
	if endTime < anotherStartTime || startTime > anotherEndTime {
		return false
	} else {
		return true
	}
}

// MaxNumber booleans, numbers, strings, pointers, channels, arrays
func MaxNumber[T Numbers](e ...T) T {
	sort.Slice(e, func(i, j int) bool { return e[i] < e[j] })
	return e[len(e)-1]
}

func MinNumber[T Numbers](e ...T) T {
	sort.Slice(e, func(i, j int) bool { return e[i] < e[j] })
	return e[0]
}

// NotIn e 不在 s 中吗？ true 不在， false 在
func NotIn[T comparable](e T, items []T) bool {
	for _, item := range items {
		if e == item {
			return false
		}
	}
	return true
}

// In e 在 items 中吗？ true 在，false 不在
func In[T comparable](e T, items []T) bool {
	for _, item := range items {
		if e == item {
			return true
		}
	}
	return false
}

// Reverse 反转 slice
func Reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

// Pagination 切片分页
func Pagination[T any](page, pageSize int, s []T) []T {
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
// func Filter[T any](objs []T, filter func(obj T) bool) []T {
// 	res := make([]T, 0, len(objs))
// 	for i := range objs {
// 		ok := filter(objs[i])
// 		if ok {
// 			res = append(res, objs[i])
// 		}
// 	}
// 	return res
// }

// Map one slice
// func Map[T any, K any](objs []T, mapper func(obj T) ([]K, bool)) []K {
// 	res := make([]K, 0, len(objs))
// 	for i := range objs {
// 		others, ok := mapper(objs[i])
// 		if ok {
// 			res = append(res, others...)
// 		}
// 	}
// 	return res
// }

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
