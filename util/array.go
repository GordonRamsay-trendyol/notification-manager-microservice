package util

func ContainsStrList(arr []string, elements ...string) bool {
	if len(elements) == 0 || len(arr) == 0 {
		return false
	}

	for _, originalElement := range arr {
		for _, expectedElement := range elements {
			if originalElement == expectedElement {
				return true
			}
		}
	}

	return false
}
