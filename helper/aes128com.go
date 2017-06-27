package helper

import (
	//"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http/cookiejar"
)

func Aes128COMEncrypt(content string, constKey string) (string, error) {

	guidKey := GUID32BIT() //32位

	kakey := guidKey[:8]              //8位
	kbkey := guidKey[len(guidKey)-3:] //3位
	key := kakey + kbkey + constKey   //16位 , constKey = "23423" //5位

	if len(key) != 16 {
		return "", errors.New("Aes128 the key length must be 16 bytes.")
	}

	if crypted, err := AesCBCEncrypt([]byte(content), []byte(key)); err != nil {
		return "", err
	} else {
		data := kakey + string(crypted) + kbkey
		return data, err

	}
}

func Aes128COMDecrypt(crypted string, constKey string) (string, error) {
	if len(crypted) <= 11 {
		return "", errors.New("Aes128 the Ciphertext is too short!")
	}
	kakey := crypted[:8]              //前8位
	kbkey := crypted[len(crypted)-3:] //后3位
	key := kakey + kbkey + constKey   //16位 , constKey = "23423" //5位

	if len(key) != 16 {
		return "", errors.New("Aes128 the key length must be 16 bytes.")
	}

	crypteds := crypted[8 : len(crypted)-3] //截取密文

	if crypteds == "" {
		return "", errors.New("密文为空!")
	} else {

		// 对AES密文进行解密
		if decrypts, err := AesCBCDecrypt([]byte(crypteds), []byte(key)); err == nil {
			return string(decrypts), err
		} else {
			return "", err
		}
	}
}

func SetJsonCOMEncrypt(status int64, msg string, data interface{}) (string, error) {
	//{"status":1, //状态。（1成功， 0失败）"msg":"success", //描述"data":""}
	m := map[string]interface{}{}

	if status != 1 {
		status = 0
	}

	if msg == "" {
		if status == 1 {
			msg = "success"
		} else {
			msg = "failure"
		}
	}

	m["data"] = data
	m["msg"] = msg
	m["status"] = status
	if content, err := json.Marshal(m); err != nil {
		return "", err
	} else {
		if crypted, err := Aes128COMEncrypt(string(content), AesConstKey); err != nil {
			return "", err
		} else {
			return crypted, err
		}

	}

	return "", errors.New("SetJsonCOMEncrypt Error")

}

func GetJsonCOMDecrypt(status string, actionurl string, data interface{}, ckJar *cookiejar.Jar) (*simplejson.Json, error) {

	content, e := json.Marshal(data)
	if e != nil {
		return nil, e
	}

	crypted, e := Aes128COMEncrypt(string(content), AesConstKey)
	if e != nil {
		return nil, e
	}

	if ckJar == nil {
		ckJar, _ = cookiejar.New(nil)
	}

	if reps, e := SendPacket(status, actionurl, crypted, ckJar); e == nil {

		bodyByte, e := ioutil.ReadAll(reps.Body)
		if e != nil {
			return nil, e
		}

		bodyString := string(bodyByte)
		if reps.StatusCode == 200 {
			//fmt.Println("服务端响应成功")
			if bodyString == "" {
				return nil, errors.New("服务端响应正文为空!")
			} else {
				if s, err := Aes128COMDecrypt(bodyString, AesConstKey); err != nil {
					return nil, errors.New(fmt.Sprint("解密失败错误:", err.Error()))

				} else {
					return simplejson.NewJson([]byte(s))
				}

			}

		} else if reps.StatusCode == 401 {
			return nil, errors.New("服务端响应状态:尚未认证!")
		} else {
			return nil, errors.New(fmt.Sprint("服务端响应错误状态:", reps.StatusCode))
		}

	}

	return nil, errors.New("GetJsonCOMDecrypt Error")

}
