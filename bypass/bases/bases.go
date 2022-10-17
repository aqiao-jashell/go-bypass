package bypass

import (
	"encoding/base64"
	// "fmt"
)


func Base_en(str string) string{
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	// fmt.Println(encoded)
	return encoded
}

func Base_de(str string) string{
	strbytes := str
	decoded, _ := base64.StdEncoding.DecodeString(strbytes)
	decodestr := string(decoded)
	// fmt.Println(decodestr)
	return decodestr
}

// func main() {
// 	Encode("qwe")
// }
