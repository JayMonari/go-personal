package main

func GenerateDivTags(amount int) []string {
	tags := []string{}
	generateDivTagsFromPrefix(amount, amount, "", &tags)
	return tags
}

func generateDivTagsFromPrefix(nOpenTags, nCloseTags int, prefix string, result *[]string) {
	if nOpenTags > 0 {
		generateDivTagsFromPrefix(nOpenTags-1, nCloseTags, prefix+"<div>", result)
	}

	if nOpenTags < nCloseTags {
		generateDivTagsFromPrefix(nOpenTags, nCloseTags-1, prefix+"</div>", result)
	}

	if nCloseTags == 0 {
		*result = append(*result, prefix)
	}
}

// Test Case 1
// {
//   "numberOfTags": 3
// }
// Test Case 2
// {
//   "numberOfTags": 2
// }
// Test Case 3
// {
//   "numberOfTags": 1
// }
// Test Case 4
// {
//   "numberOfTags": 4
// }
// Test Case 5
// {
//   "numberOfTags": 5
// }
// Test Case 6
// {
//   "numberOfTags": 6
// }
