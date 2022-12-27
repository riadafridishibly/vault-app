package secret

import "testing"

func TestVerify(t *testing.T) {
	const pass = "hello world"
	hashed := HashPassword([]byte(pass))
	if ok := Verify(hashed, pass); !ok {
		t.Error("Verifiation failed")
	}
}
