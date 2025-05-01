package utils

import "strings"

func SshToHttps(sshUrl string) string {
	if strings.HasPrefix(sshUrl, "git@github.com:") {

		var sanitazedUrl string = strings.Replace(strings.TrimSpace(sshUrl), ".git")
		return "https://github.com/" + strings.TrimPrefix(sanitazedUrl, "git@github.com:")
	}
	return sshtUrl
}
