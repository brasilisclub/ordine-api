package services

func UserExists(userName string) bool {

	_, err := getUserFromDb(userName)
	if err != nil {
		return false
	}
	return true
}
