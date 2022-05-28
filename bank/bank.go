package bank

type Account struct {
	name string
	totalMoney int
}

func AccountCreate(name string){
	Account.name = name
}

