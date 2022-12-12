package util

import (
	"regexp"
	"strings"
)

var r, _ = regexp.Compile("[A-Za-z]")

type Util struct {
}

func NewUtil() Util {
	return Util{}
}

func (u Util) uppercaseAndRemoveSpaces(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, " ", ""))
}

func (u Util) cleanInput(text string) string {
	var upppercaseLine string
	upppercaseLine = u.uppercaseAndRemoveSpaces(text)
	upppercaseLine = strings.Join(r.FindAllString(upppercaseLine, len(upppercaseLine)), "")
	return upppercaseLine
}

// Encryptions Using Caeser to a char
func (u Util) caesarEncryption(pt rune, key int) rune {
	pt = pt - 65
	return ((pt + rune(key-65)) % 26) + 65
}

// Decryptions Using Caeser to a char
func (u Util) caesarDecryption(st rune, key int) rune {
	st = st - 65
	result := (st - rune(key-65)) % 26

	if result >= 0 {
		return result + 65
	} else {
		return (26 + result) + 65
	}
}

// Encryption Function
func (u Util) VigenereEncryption(plainText string, key string) string {
	var cleanPlainText, cleanKey string = u.cleanInput(plainText), u.cleanInput(key)
	var newResult string

	for i := 0; i < len(cleanPlainText); i++ {
		c := u.caesarEncryption(rune(cleanPlainText[i]), int(cleanKey[i%(len(cleanKey))]))
		newResult = newResult + string(c)
	}
	return newResult
}

// Decryption Function
func (u Util) VigenereDecryption(cipherText string, key string) string {
	var cleanCipherText, cleanKey string = u.cleanInput(cipherText), u.cleanInput(key)
	var newResult string

	for i := 0; i < len(cleanCipherText); i++ {
		c := u.caesarDecryption(rune(cleanCipherText[i]), int(cleanKey[i%(len(cleanKey))]))
		newResult = newResult + string(c)
	}
	return newResult
}
