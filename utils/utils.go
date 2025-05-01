package utils

import "strings"

func SshToHttps(sshUrl string) string {
	var sanitazedUrl string = strings.ReplaceAll(strings.TrimSpace(sshUrl), ".git", "")
	if strings.HasPrefix(sshUrl, "git@github.com:") {
		return "https://github.com/" + strings.TrimPrefix(sanitazedUrl, "git@github.com:")
	}

	return sanitazedUrl
}
