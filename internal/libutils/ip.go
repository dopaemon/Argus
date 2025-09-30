package libutils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetHostIP() (string, error) {
	services := []string{
		"https://api.ipify.org",
		"https://ifconfig.me",
		"https://checkip.amazonaws.com",
	}

	for _, url := range services {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		ip := strings.TrimSpace(string(body))
		if ip != "" {
			return ip, nil
		}
	}

	return "", errors.New("cannot get public IP")
}
