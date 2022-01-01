package e2

type PaymentStrategy interface {
	Pay(*PaymentContext)
}
