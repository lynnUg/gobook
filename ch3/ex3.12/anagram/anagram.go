package anagram


func IsAnagram(a, b string) bool {
     aMap := make(map[rune]int)
     for _,r := range a {
	aMap[r]++
     }

     bMap := make(map[rune]int)

     for _,r := range b {
	bMap[r]++
	if n,ok := aMap[r]; !ok || n < bMap[r] {
		return false
	}
     }

     return  true
}
