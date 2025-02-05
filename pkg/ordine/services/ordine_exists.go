package services

func OrdineExists(ordineId string) error {
	_, err := GetOrdineById(ordineId)

	return err
}
