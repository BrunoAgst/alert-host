package services

import (
	"net/http"
)

type Host struct {
	URL string
}

func MonitorHosts(hosts []Host) {

	for _, host := range hosts {
		status, message := verifyHost(host)
		if !status {
			sendNotification(message)
		}
	}
}

func verifyHost(host Host) (bool, string) {
	res, _ := http.Get(host.URL)

	if res.StatusCode == 200 {
		return true, "host: " + host.URL + " page ok"
	}

	return false, "host: " + host.URL + " page down"
}
