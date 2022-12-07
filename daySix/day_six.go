package daySix

func FindStartOfMessageOrSignal(signal string, length int) (int, error) {

	firstMarkerAfterNSymbols := length
	var exists = struct{}{}

	for i := 0; i < len(signal)-length-1; i++ {
		set := make(map[byte]struct{})

		start := i
		end := i + length
		var symbol string
		symbol = signal[start:end]

		for _, ch := range symbol {
			set[byte(ch)] = exists
		}

		if len(set) == len(symbol) {
			break
		}

		firstMarkerAfterNSymbols++
	}

	return firstMarkerAfterNSymbols, nil
}
