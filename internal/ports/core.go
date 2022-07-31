package ports

type ArithemetictPort interface {
	Adddition(num, num1 int32) (int32, error)
	Subtraction(num, num1 int32) (int32, error)
	Multiplication(num, num1 int32) (int32, error)
	Division(num, num1 int32) (int32, error)
}
