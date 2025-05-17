package core

import "Tubes-IdeaManagerCLI/model"

func SequentialSearch(arr [model.MaxIdeas]model.Idea, count int, keyword string) int {
	for i := 0; i < count; i++ {
		if arr[i].Title == keyword {
			return i
		}
	}
	return -1
}

func BinarySearch(arr [model.MaxIdeas]model.Idea, count int, keyword string) int {
	low := 0
	high := count - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid].Title == keyword {
			return mid
		} else if arr[mid].Title < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}