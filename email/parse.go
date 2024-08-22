package email

import "strings"

func GetEmailDomain(email string) string {
	atIndex := strings.LastIndexByte(email, '@')
	if atIndex >= 0 && atIndex < len(email)-1 {
		return strings.ToLower(email[atIndex+1:])
	}
	return ""
}
