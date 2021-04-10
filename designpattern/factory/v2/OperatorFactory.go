package v2

type OperatorFactory struct {
}

func NewOperatorFactory() *OperatorFactory {
	return &OperatorFactory{}
}

func (p *OperatorFactory) CreateOperate(operatename string) Operator {
	switch operatename {
	case "+":
		return &AddOperator{}
	case "-":
		return &AddOperator{}
	default:
		panic("Invalid operator")
		return nil
	}
}
