package letter

const testVersion = 1

func ConcurrentFrequency(in []string) FreqMap {
	out := FreqMap{}
	channel := make(chan FreqMap, len(in))
	defer close(channel)

	for _, inp := range in {
		go func(w string) { channel <- Frequency(w) }(inp)
	}

	for range in {
		for key, value := range <-channel {
			out[key] += value
		}
	}
	return out
}