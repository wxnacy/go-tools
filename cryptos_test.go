package gotool

import "testing"

func TestAesEncryptDecrypt(t *testing.T) {
	key := Md5("wxnacy")
	text := "wxnacy"
	ciphertext, err := AesEncrypt(text, key)
	if err != nil {
		t.Error(err)
	}
	src, err := AesDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
	}
	if src != text {
		t.Errorf("%s != %s", text, src)
	}
}
