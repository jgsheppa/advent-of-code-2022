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

		set[symbol[0]] = exists
		set[symbol[1]] = exists
		set[symbol[2]] = exists
		set[symbol[3]] = exists

		if len(set) == len(symbol) {
			break
		}

		firstMarkerAfterNSymbols++
	}

	return firstMarkerAfterNSymbols, nil
}
