package util

const (
	Pending  = "PENDING"
	Shipping = "SHIPPING"
	Failed   = "FAILED"
	Complete = "COMPLETE"
)

func IsSupportedOrderStatus(orderStatus string) bool {
	switch orderStatus {
	case Pending, Shipping, Failed, Complete:
		return true
	}
	return false
}
