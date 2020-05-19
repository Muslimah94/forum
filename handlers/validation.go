package handlers

import (
	"strings"
)

func validPass(s string) bool {
	if len(s) < 5 && len(s) > 20 {
		return false
	}
	letters := "abcdefghijklmnopqrstuvwxyz"
	capitals := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := ".,`-=~!@#$%^&*()_+\\|/? []{}"
	if !strings.ContainsAny(s, specials) && !strings.ContainsAny(s, letters) || !strings.ContainsAny(s, capitals) || !strings.ContainsAny(s, numbers) {
		return false
	}
	validstr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`-=~!@#$%^&*()_+\\|/? {}[]"
	for _, r := range s {
		if !strings.Contains(validstr, string(r)) {
			return false
		}
	}
	return true
}

func validNick(s string) bool {
	validstr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`-=~!@#$%^&*()_+\\|/? {}[]"
	for _, r := range s {
		if !strings.Contains(validstr, string(r)) {
			return false
		}
	}
	return true
}
