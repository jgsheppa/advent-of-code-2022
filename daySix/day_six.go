package daySix

func FindStartOfPacket(signal string) (int, error) {

	firstMarkerAfterNSymbols := 4
	var exists = struct{}{}

	for i := 0; i < len(signal)-3; i++ {
		set := make(map[byte]struct{})

		start := i
		end := i + 4
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

func FindStartOfMessage(signal string) (int, error) {

	firstMarkerAfterNSymbols := 14
	var exists = struct{}{}

	for i := 0; i < len(signal)-13; i++ {
		set := make(map[byte]struct{})

		start := i
		end := i + 14
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
