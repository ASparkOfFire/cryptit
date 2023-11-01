package encrypt

func Nimbus(str string) string {
	var encryptedString string
	for _, c := range str {
		asciiCode := int(c)
		character := string(int32(asciiCode + 3))
		encryptedString += character
	}
	return encryptedString
}
