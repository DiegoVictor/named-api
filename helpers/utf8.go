package helpers

func ToUtf8(encoded []byte) string {
	buffer := make([]rune, len(encoded))
	for i, b := range encoded {
		buffer[i] = rune(b)
	}
	return string(buffer)
}
