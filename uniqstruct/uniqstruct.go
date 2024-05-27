package uniqstruct

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

	return uniqSlice
}
