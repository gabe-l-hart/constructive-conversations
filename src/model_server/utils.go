package main

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeFromSet(a string, set []string) []string {
	for i, b := range set {
		if a == b {
			set[i] = set[len(set)-1]
			set = set[:len(set)-1]
		}
	}
	return set
}

func addToSet(a string, set []string) []string {
	for _, b := range set {
		if a == b {
			return set
		}
	}
	set = append(set, a)
	return set
}
