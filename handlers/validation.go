package handlers

import (
	"regexp"
	"strings"
)

func validPass(s string) bool {
	if len(s) < 5 || len(s) > 20 {
		return false
	}
	letters := "abcdefghijklmnopqrstuvwxyz"
	capitals := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := ".,`-=~!@#$%^&*()_+\\|/? []{}"
	if !strings.ContainsAny(s, specials) || !strings.ContainsAny(s, letters) || !strings.ContainsAny(s, capitals) || !strings.ContainsAny(s, numbers) {
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
	if len(s) < 5 || len(s) > 20 {
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

func validEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
