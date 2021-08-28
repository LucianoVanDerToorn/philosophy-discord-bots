package util

func RemoveEmptyFromStringArray(arr []string) []string {
	var arrNew []string
	for _, s := range arr {
		if s != "" {
			arrNew = append(arrNew, s)
		}
	}
	return arrNew
}
