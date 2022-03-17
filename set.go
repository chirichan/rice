package rice

import "time"

// SliceRemoveIndex 移除 slice 中的一个元素
func SliceRemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// SliceRemoveIndexUnOrder 移除 slice 中的一个元素（无序，但效率高）
func SliceRemoveIndexUnOrder[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// SliceSet slice 去重
func SliceSet[T comparable](s1 []T) []T {

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

//Remove dups from slice.
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

// IsHaveIntersectionTimestamp 两个时间段是否有交集 false 没有交集，true 有交集
func TimestampExistIntersection(startTime, endTime int64, anotherStartTime, anotherEndTime int64) bool {

	if endTime < anotherStartTime || startTime > anotherEndTime {
		return false
	} else {
		return true
	}
}
