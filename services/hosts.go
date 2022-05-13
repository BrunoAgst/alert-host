package services

import (
	"host-monitoring/models"
	"net/http"
)

func MonitorHosts(hosts []models.Host) {

	for _, host := range hosts {
		status, message := verifyHost(host.URL)

		if !status {
			sendNotification(message)
		}
	}
}

func verifyHost(host string) (bool, string) {
	res, _ := http.Get(host)

	if res.StatusCode == 200 {
		return true, "host: " + host + " page ok"
	}
	return false, "host: " + host + " page down"

}
