package services

import "strconv"

func ProductExists(productId uint) bool {
	_, err := GetProductById(strconv.FormatUint(uint64(productId), 10))

	if err != nil {
		return false
	}
	return true
}
