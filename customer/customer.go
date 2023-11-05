package customer

type Customer struct {
	Name string
	Age  int64
}

func (c Customer) GetAge() int64 {
	return c.Age
}
