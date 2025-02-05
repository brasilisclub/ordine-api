package services

func OrderExists(ordineId uint, productId uint) bool {
	_, err := GetOrderByIds(ordineId, productId)
	if err != nil {
		return false
	}
	return true
}
