//Isomorphic String

func isomorphicString(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc := t[i]

		// Check s -> t mapping
		if mapped, exists := sToT[sc]; exists {
			if mapped != tc {
				return false
			}
		}

		// Check t -> s mapping
		if mapped, exists := tToS[tc]; exists {
			if mapped != sc {
				return false
			}
		}

		sToT[sc] = tc
		tToS[tc] = sc
	}

	return true
}