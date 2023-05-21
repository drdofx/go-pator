package handlers

import "strconv"

func StringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	return uint(u), err
}
