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

