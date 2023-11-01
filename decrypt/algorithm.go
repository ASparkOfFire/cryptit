package decrypt

func Nimbus(str string) string {
	var decryptedString string
	for _, c := range str {
		asciiCode := int(c)
		character := string(int32(asciiCode - 3))
		decryptedString += character
	}
	return decryptedString
}
