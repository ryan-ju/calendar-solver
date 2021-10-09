package calendar

func GetShortIndexes(input []uint8) []ShortIndex {
	var result []ShortIndex
	var x uint8
	for i, v := range input {
		if i%2 == 0 {
			x = v
		} else {
			result = append(result, ShortIndex{
				X: x,
				Y: v,
			})
		}
	}
	return result
}
