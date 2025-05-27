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