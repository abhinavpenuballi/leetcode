package main

func predictPartyVictory(senate string) string {
	candidates, xor := []byte(senate), byte('R'^'D')

	for {
		for i := 0; i < len(candidates); i++ {
			if candidates[i] == '_' {
				continue
			}

			hasOpposition := false

			for j := 1; j < len(candidates); j++ {
				pick := (i + j) % len(candidates)
				if candidates[pick] == candidates[i]^xor {
					candidates[pick] = '_'
					hasOpposition = true
					break
				}
			}

			if !hasOpposition {
				if candidates[i] == 'R' {
					return "Radiant"
				} else {
					return "Dire"
				}
			}
		}
	}
}
