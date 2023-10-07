package betterjson

func decomment(data []byte) []byte {
	const (
		OUTSIDE = iota
		SINGLE_LINE
		MULTI_LINE
		MULTI_LINE_ENDING
		IN_STRING
	)

	state := OUTSIDE
	result := make([]byte, len(data))
	copy(result, data)

	stateHandlers := []func(int, []byte, []byte) int{
		// OUTSIDE
		func(i int, data, result []byte) int {
			if data[i] == '/' && i+1 < len(data) {
				if result[i+1] == '/' {
					state = SINGLE_LINE
					result[i] = ' '
					result[i+1] = ' '
					return i + 1 // skip the next character
				} else if data[i+1] == '*' {
					state = MULTI_LINE
					result[i] = ' '
					result[i+1] = ' '
					return i + 1 // skip the next character
				}
			} else if data[i] == '"' {
				state = IN_STRING
			}
			return i
		},
		// SINGLE_LINE
		func(i int, data, result []byte) int {
			if data[i] == '\n' {
				state = OUTSIDE
			} else {
				result[i] = ' '
			}
			return i
		},
		// MULTI_LINE
		func(i int, data, result []byte) int {
			if data[i] == '*' && i+1 < len(result) && data[i+1] == '/' {
				state = MULTI_LINE_ENDING
				result[i] = ' '
				result[i+1] = ' '
				return i + 1 // skip the next character
			} else if result[i] != '\n' {
				result[i] = ' '
			}
			return i
		},
		// MULTI_LINE_ENDING
		func(i int, data, result []byte) int {
			state = OUTSIDE
			return i
		},
		// IN_STRING
		func(i int, data, result []byte) int {
			if data[i] == '"' {
				state = OUTSIDE
			}
			return i
		},
	}

	for i := 0; i < len(result); i++ {
		i = stateHandlers[state](i, data, result)
	}

	return result
}
