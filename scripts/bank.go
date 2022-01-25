package scripts

import "fmt"

// 不要使用共享数据来通信；使用通信 channel 来共享数据

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <- balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <- deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func BankMain() {
	quit := make(chan bool)
	go func() {
		for true {
			balance := Balance()
			fmt.Printf("Balance: %d\n", balance)
			if balance == 1000 {
				quit <- true
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go Deposit(100)
	}

	select {
	case <-quit:
	}
}
