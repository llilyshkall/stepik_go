package main

func main() {
	groupCity := map[int][]string{
		10:   []string{"abc", "def", "g"}, // города с населением 10-99 тыс. человек
		100:  []string{"k", "m"},          // города с населением 100-999 тыс. человек
		1000: []string{"dsc"},             // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"k": 10,
		"M": 100,
	}
	for key, _ := range cityPopulation {
		exist := false
		for _, city := range groupCity[100] {
			if city == key {
				exist = true
			}
		}
		if !exist {
			delete(cityPopulation, key)
		}
	}
}
