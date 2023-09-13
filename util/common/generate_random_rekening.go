package common

import "fmt"

func GenerateRandomRekeningNumber(userID string) string {
	// Check the length of the userID (assuming it's a UUID)
	if len(userID) < 10 {
		return fmt.Sprintf("10000%s", userID)
	} else if len(userID) < 100 {
		return fmt.Sprintf("1000%s", userID)
	} else {
		return fmt.Sprintf("100%s", userID)
	}
}
