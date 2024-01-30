package util

const (
	USD = "USD"
	CNY = "CNY"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case USD, CNY:
		return true
	}
	return false
}
