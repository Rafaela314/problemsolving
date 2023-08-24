package others

import (
	"bytes"
	"fmt"
	"strconv"
)

func CommonConversions() {

	// STRING TO BINARY
	number := 123

	binaryNumber := strconv.FormatInt(int64(number), 2)
	binaryNumber2 := fmt.Sprintf("x=%b", number)

	fmt.Printf("binaryNumber: %s, binaryNumber2: %s\n", binaryNumber, binaryNumber2) // 1111011

	// BINARY TO INT

	strBinary := "1111011"
	numberFromBinary, _ := strconv.ParseInt(strBinary, 2, 0)
	fmt.Printf("numberFromBinary: %d, \n", numberFromBinary) // 123

	// BYTE x STRING

	byteArray := []byte{97, 98, 99, 100, 101, 102}
	fmt.Println(string(byteArray)) // "abcdef"

	str := bytes.NewBuffer(byteArray).String()
	fmt.Println(string(str)) // "abcdef"

	var b byte
	b = 98
	fmt.Println(string(b)) // "b"

	// RUNE

	s := "sum"
	fmt.Println(string(s[1])) //"u"
	fmt.Println(rune('u'))    // 117

	// INT TO STRING

	i := 10
	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)
	fmt.Printf("%v, %v\n", s1, s2) // 10, 10

	// STRING TO INT

	t := "123456"

	d, err := strconv.Atoi(t)
	if err == nil {
		fmt.Printf("i=%d, type: %T\n", d, d) // i=123456, type: int
	}

	c, err := strconv.ParseInt(t, 10, 64)
	if err == nil {
		fmt.Printf("i=%d, type: %T\n", c, c) // i=123456, type: int64

	}

}
