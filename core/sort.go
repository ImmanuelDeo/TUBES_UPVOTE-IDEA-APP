package core

import "Tubes-IdeaManagerCLI/model"

//Handler sorting data dari vote terbesar
func SelectionSortByUpvotes(arr *[model.MaxIdeas]model.Idea, count int, ascending bool) {
	for i := 0; i < count-1; i++ {
		minOrMax := i
		for j := i + 1; j < count; j++ {
			if (ascending && arr[j].Upvotes < arr[minOrMax].Upvotes) ||
				(!ascending && arr[j].Upvotes > arr[minOrMax].Upvotes) {
				minOrMax = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minOrMax]
		arr[minOrMax] = temp
	}
}

//Handler sorting data dari tanggal terbaru
func InsertionSortByDate(arr *[model.MaxIdeas]model.Idea, count int, ascending bool) {
	for i := 1; i < count; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && ((ascending && arr[j].CreatedAt.After(key.CreatedAt)) || (!ascending && arr[j].CreatedAt.Before(key.CreatedAt))) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

//Handler sorting data dari ID terkecil abis input data baru
func SortIdeasByID(arr *[model.MaxIdeas]model.Idea, count int) {
	for i := 0; i < count-1; i++ {
		minIdx := i
		for j := i + 1; j < count; j++ {
			if arr[j].ID < arr[minIdx].ID {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}