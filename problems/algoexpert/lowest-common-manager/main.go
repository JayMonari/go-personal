package main

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return getOrgInfo(org, reportOne, reportTwo).lowestCommonManager
}

type OrgInfo struct {
	lowestCommonManager *OrgChart
	nImportantReports   int
}

func getOrgInfo(manager, reportOne, reportTwo *OrgChart) OrgInfo {
	nReports := 0
	for _, r := range manager.DirectReports {
		orgInfo := getOrgInfo(r, reportOne, reportTwo)
		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}
		nReports += orgInfo.nImportantReports
	}
	if manager == reportOne || manager == reportTwo {
		nReports++
	}

	var lowestCommonManager *OrgChart
	if nReports == 2 {
		lowestCommonManager = manager
	}
	return OrgInfo{
		lowestCommonManager: lowestCommonManager,
		nImportantReports:   nReports,
	}
}

// Test Case 1
// {
//   "topManager": "A",
//   "reportOne": "E",
//   "reportTwo": "I",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C"], "id": "A", "name": "A"},
//       {"directReports": ["D", "E"], "id": "B", "name": "B"},
//       {"directReports": ["F", "G"], "id": "C", "name": "C"},
//       {"directReports": ["H", "I"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": [], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": [], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"}
//     ]
//   }
// }
// Test Case 2
// {
//   "topManager": "A",
//   "reportOne": "A",
//   "reportTwo": "B",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 3
// {
//   "topManager": "A",
//   "reportOne": "B",
//   "reportTwo": "F",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 4
// {
//   "topManager": "A",
//   "reportOne": "G",
//   "reportTwo": "M",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 5
// {
//   "topManager": "A",
//   "reportOne": "U",
//   "reportTwo": "S",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 6
// {
//   "topManager": "A",
//   "reportOne": "Z",
//   "reportTwo": "M",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 7
// {
//   "topManager": "A",
//   "reportOne": "O",
//   "reportTwo": "I",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 8
// {
//   "topManager": "A",
//   "reportOne": "T",
//   "reportTwo": "Z",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 9
// {
//   "topManager": "A",
//   "reportOne": "T",
//   "reportTwo": "V",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 10
// {
//   "topManager": "A",
//   "reportOne": "T",
//   "reportTwo": "H",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 11
// {
//   "topManager": "A",
//   "reportOne": "W",
//   "reportTwo": "V",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 12
// {
//   "topManager": "A",
//   "reportOne": "Z",
//   "reportTwo": "B",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 13
// {
//   "topManager": "A",
//   "reportOne": "Q",
//   "reportTwo": "W",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 14
// {
//   "topManager": "A",
//   "reportOne": "A",
//   "reportTwo": "Z",
//   "orgChart": {
//     "nodes": [
//       {"directReports": ["B", "C", "D", "E", "F"], "id": "A", "name": "A"},
//       {"directReports": ["G", "H", "I"], "id": "B", "name": "B"},
//       {"directReports": ["J"], "id": "C", "name": "C"},
//       {"directReports": ["K", "L"], "id": "D", "name": "D"},
//       {"directReports": [], "id": "E", "name": "E"},
//       {"directReports": ["M", "N"], "id": "F", "name": "F"},
//       {"directReports": [], "id": "G", "name": "G"},
//       {"directReports": ["O", "P", "Q", "R"], "id": "H", "name": "H"},
//       {"directReports": [], "id": "I", "name": "I"},
//       {"directReports": [], "id": "J", "name": "J"},
//       {"directReports": ["S"], "id": "K", "name": "K"},
//       {"directReports": [], "id": "L", "name": "L"},
//       {"directReports": [], "id": "M", "name": "M"},
//       {"directReports": [], "id": "N", "name": "N"},
//       {"directReports": [], "id": "O", "name": "O"},
//       {"directReports": ["T", "U"], "id": "P", "name": "P"},
//       {"directReports": [], "id": "Q", "name": "Q"},
//       {"directReports": ["V"], "id": "R", "name": "R"},
//       {"directReports": [], "id": "S", "name": "S"},
//       {"directReports": [], "id": "T", "name": "T"},
//       {"directReports": [], "id": "U", "name": "U"},
//       {"directReports": ["W", "X", "Y"], "id": "V", "name": "V"},
//       {"directReports": [], "id": "W", "name": "W"},
//       {"directReports": ["Z"], "id": "X", "name": "X"},
//       {"directReports": [], "id": "Y", "name": "Y"},
//       {"directReports": [], "id": "Z", "name": "Z"}
//     ]
//   }
// }
