package utils

func inArray(elem string, arr []string ) bool {
	for _,arrayElem := range arr {
		if elem == arrayElem {
			return true
		}
	}

	return false
}
