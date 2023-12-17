package gohelp

import (
	"fmt"
)

// Index return index of item if exists. -1 if not
func Index[T comparable](value T, slice []T) int {
	for i, a := range slice {
		if a == value {
			return i
		}
	}
	return -1
}

// ExistsInArray Check if item exists in slice
func ExistsInArray[T comparable](value T, slice []T) bool {
	for _, a := range slice {
		if a == value {
			return true
		}
	}
	return false
}

// Implode join values via separator
func Implode[T comparable](values []T, sep string) string {
	var result string
	for _, value := range values {
		result += fmt.Sprintf("%v", value) + sep
	}
	return result[:len(result)-len(sep)]
}

// ArrayIntersect shows duplicated elements in both slices
func ArrayIntersect[T comparable](target []T, destination []T) (intersect []T, removedTargets int32) {
	for _, tItem := range target {
		targetFound := false
		for _, dItem := range destination {
			if tItem == dItem {
				intersect = append(intersect, tItem)
				targetFound = true
				break
			}
		}
		if targetFound == false {
			removedTargets += 1
		}
	}
	return
}

// AppendUnique append unique value to array
func AppendUnique[T comparable](slice []T, values ...T) []T {
	for _, v := range values {
		var exists bool
		for _, s := range slice {
			if s == v {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, v)
		}
	}
	return slice
}

// Unique create unique slice between left and right slices
func Unique[T comparable](left []T, right []T) []T {
	return AppendUnique[T]([]T{}, append(left, right...)...)
}

// UniqueLeft create unique slice from left and right slices
func UniqueLeft[T comparable](left []T, right []T) []T {
	var leftUnique = make([]T, 0)
	for _, leftValue := range left {
		var found bool
		for _, rightValue := range right {
			if leftValue == rightValue {
				found = true
				break
			}
		}
		if !found {
			leftUnique = append(leftUnique, leftValue)
		}
	}
	return AppendUnique[T]([]T{}, leftUnique...)
}

// Filter argument source via callback
func Filter[S any](s []S, callback func(S) bool) []S {
	var j int
	var r = make([]S, len(s))
	for i := range s {
		if callback(s[i]) {
			r[j] = s[i]
			j++
		}
	}
	return r[:j]
}

// Map argument source and result destination
func Map[S, D any](s []S, callback func(S) D) []D {
	r := make([]D, len(s))
	for i, v := range s {
		r[i] = callback(v)
	}
	return r
}

// Reduce implementation
func Reduce[S, D any](init D, s []S, callback func(D, S) D) D {
	for _, v := range s {
		init = callback(init, v)
	}
	return init
}

// Each iterate thought []S and apply on each item collback
func Each[S any](s []S, callback func(S)) {
	for i := range s {
		callback(s[i])
	}
	return
}
