package main

import (
	"fmt"
)

// ACCOUNT REPRESENTS A BANK ACCOUNT
type Account struct {
	Name          string
	AccountHolder string
	Age           float64
	AccountNumber string
	Pin           int
	Balance       float64
}

// NEWACCOUNT CREATES A NEW ACCOUNT
func NewAccount(name, accountHolder string, age float64, accountNumber string, pin int) *Account {
	return &Account{
		Name:          name,
		AccountHolder: accountHolder,
		Age:           age,
		AccountNumber: accountNumber,
		Pin:           pin,
		Balance:       100.0,
	}
}

// CHECKBALANCE CHECKS AND DISPLAYS THE ACCOUNT BALANCE
func (a *Account) CheckBalance() {
	var key int
	fmt.Println("Enter the Pin:")
	fmt.Scan(&key)

	if key == a.Pin {
		fmt.Printf("Mr. %s\n", a.AccountHolder)
		fmt.Printf("Your Current Balance is $%.2f\n", a.Balance)
	} else {
		fmt.Println("Your pin is Wrong. Try Again.")
	}
}

// WITHDRAW ALLOWS THE USER TO WITHDRAW FUNDS FROM THE ACCOUNT
func (a *Account) Withdraw() {
	var key int
	fmt.Println("Enter the Pin:")
	fmt.Scan(&key)

	if key == a.Pin {
		var withdrawAmount float64
		fmt.Println("Enter the amount you want to withdraw:")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > a.Balance {
			fmt.Println("Not Enough Balance. Welcome Again.")
			return
		}

		a.Balance -= withdrawAmount
		fmt.Printf("You have successfully withdrawn $%.2f\n", withdrawAmount)
		fmt.Printf("Your current Balance is $%.2f\n", a.Balance)
	} else {
		fmt.Println("Your pin is Wrong. Try Again.")
	}
}

// TRANSFER ALLOWS THE USER TO PERFORM TRANSACTION
func (a *Account) Transfer() {
	var method, key int
	var amount, targetAccountNumber float64

	fmt.Println("Select Method:")
	fmt.Println("1. JazzCash\t2. EasyPaisa\t3. Upaisa\t4. MasterCard")
	fmt.Scan(&method)

	fmt.Println("Enter the Account Number:")
	fmt.Scan(&targetAccountNumber)

	fmt.Println("Enter Your Pin:")
	fmt.Scan(&key)

	if key == a.Pin {
		fmt.Println("Enter the Amount:")
		fmt.Scan(&amount)

		if amount > a.Balance {
			fmt.Println("Sorry, Not Enough Balance")
			return
		}

		a.Balance -= amount
		fmt.Printf("You have successfully transferred $%.2f to %f\n", amount, targetAccountNumber)
	} else {
		fmt.Println("Your pin is Wrong. Try Again.")
	}
}

// DEPOSIT ALLOWS THE USER TO DEPOSIT FUNDS INTO THE ACCOUNT
func (a *Account) Deposit() {
	var depositAmount float64
	fmt.Println("How much amount you want to deposit?")
	fmt.Scan(&depositAmount)

	fmt.Printf("You have successfully deposited $%.2f to your %s Account.\n", depositAmount, a.Name)
	a.Balance += depositAmount
	fmt.Printf("Your new %s Account balance is $%.2f\n", a.Name, a.Balance)
}

// DISPLAYINFO DISPLAYS CUSTOMER INFORMATION
func (a *Account) DisplayInfo() {
	var key int
	fmt.Println("Enter your pin:")
	fmt.Scan(&key)

	if key == a.Pin {
		fmt.Printf("Bank Name: %s\n", a.Name)
		fmt.Printf("Account Holder Name: %s\n", a.AccountHolder)
		fmt.Printf("Your pin: %d\n", a.Pin)
		fmt.Printf("Your Account Number: %s\n", a.AccountNumber)
	} else {
		fmt.Println("Can't share our customer information with strangers.")
	}
}

func main() {
	var name, accountHolder, accountNumber string
	var age float64
	var pin int

	fmt.Println("* Fill This Form To Create and use your bank Account:")
	fmt.Print("Enter Your Bank Name(eg. Meezan/United/International): ")
	fmt.Scanln(&name)

	fmt.Println("Enter Account Holder Your Name:")
	fmt.Scanln(&accountHolder)

	fmt.Println("Enter Age:")
	fmt.Scanln(&age)

	fmt.Println("Enter Your Account Number(must be than 11/16 values):")
	fmt.Scanln(&accountNumber)

	fmt.Println("Create your Pin/Set your Pin:")
	fmt.Scanln(&pin)

	fmt.Println("Your Account has been Successfully Created.")
	fmt.Println("Enter any number to proceed.")
	var a int
	fmt.Scanln(&a)

	fmt.Printf("****** Welcome to your %s Bank/Account ******\n", name)
	fmt.Printf("Hi, %s. Welcome to your Account.\n", accountHolder)

	account := NewAccount(name, accountHolder, age, accountNumber, pin)

	var option int

	for {
		fmt.Println("Our Menu is here:")
		fmt.Println("1. Check Current Balance")
		fmt.Println("2. Withdraw Amount")
		fmt.Println("3. Perform Transaction")
		fmt.Println("4. Deposit Amount")
		fmt.Println("5. Customer Information")
		fmt.Println("6. Exit")

		fmt.Print("Select option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			account.CheckBalance()
		case 2:
			account.Withdraw()
		case 3:
			account.Transfer()
		case 4:
			account.Deposit()
		case 5:
			account.DisplayInfo()
		case 6:
			fmt.Println("Exiting. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}
