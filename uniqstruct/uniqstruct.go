package uniqstruct

import "sort"

type ID int
type Structure struct {
	ID
}

func Uniqstruct(s []Structure) []Structure {

	var uniqSlice []Structure

	for _, item := range s {
		var isFound = false

		for _, uniqItem := range uniqSlice {
			if item.ID == uniqItem.ID {
				isFound = true
				break
			}
		}

		if !isFound {
			uniqSlice = append(uniqSlice, item)
		}
	}

	sort.Slice(uniqSlice, func(i, j int) bool {
		return uniqSlice[i].ID < uniqSlice[j].ID
	})

	return uniqSlice
}
