package helper

func RemoveDuplicates(arr []string) []string {
	words_string := map[string]bool{}
	for i := range arr {
		words_string[arr[i]] = true
	}
	desired_output := []string{} // Keep all keys from the map into a slice.
	for key, _ := range words_string {
		desired_output = append(desired_output, key)
	}
	return desired_output
}
