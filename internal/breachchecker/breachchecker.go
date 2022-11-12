package breachchecker

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const pwnedPasswordsURL = "https://api.pwnedpasswords.com/range/"

func GetBreach(password string) (count int, err error) {
	hasher := sha1.New()
	_, err = io.WriteString(hasher, password)
	if err != nil {
		return 0, err
	}
	shaVal := hex.EncodeToString(hasher.Sum(nil))
	prefix := shaVal[:5]

	resp, err := http.Get(pwnedPasswordsURL + prefix)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Searching through the text should work
	// One way to solve this is to use this pattern
	// (rest of the hash value):(\d+)
	pResults := strings.Split(string(body), "\n")
	for _, result := range pResults {
		innerResult := strings.Split(result, ":")
		hashVal := innerResult[0]
		count := strings.TrimSpace(innerResult[1])
		if hashVal == strings.ToUpper(shaVal[5:]) {
			return strconv.Atoi(count)
		}
	}

	return 0, nil
}
