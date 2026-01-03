package error

func ErrMapping(err error) bool {
	var (
		GeneralErrors  = GeneralErrors
		CustomerErrors = CustomerErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, CustomerErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
