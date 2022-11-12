package breachchecker

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetBreach(password string) (count int, err error) {
	hasher := sha1.New()
	pBytes := []byte(password)
	hasher.Write(pBytes)
	shaVal := hex.EncodeToString(hasher.Sum(nil))
	prefix := shaVal[:5]

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

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
