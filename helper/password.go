package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	//"fmt"
	"time"
)

func EncryptHash(password string, salt []byte) string {
	if salt == nil {
		m := md5.New()
		m.Write([]byte(time.Now().String()))
		s := hex.EncodeToString(m.Sum(nil))
		salt = []byte(s[2:10])
	} else if len(salt) != 8 {
		panic("Encrypt_hash: salt != 8")
	}

	mac := hmac.New(sha256.New, salt)
	mac.Write([]byte(password))
	//s := fmt.Sprintf("%x", (mac.Sum(salt)))
	s := hex.EncodeToString(mac.Sum(nil))

	hasher := sha1.New()
	hasher.Write([]byte(s))

	//result := fmt.Sprintf("%x", (hasher.Sum(nil)))
	result := hex.EncodeToString(hasher.Sum(nil))

	p := string(salt) + result

	return p
}

func ValidateHash(hashed string, input_password string) bool {
	switch l := len(hashed); {
	case l == 48:
		salt := hashed[0:8]
		if hashed == EncryptHash(input_password, []byte(salt)) {
			return true
		} else {
			return false
		}
	case l == 32:
		//For discuz
		return false
	}
	return false
}

/*
func main() {
	hashed := EncryptHash("password", nil)
	fmt.Println(hashed)
	fmt.Println("----------------------------------------------")
	fmt.Println(ValidateHash(hashed, "password"))
}
*/
