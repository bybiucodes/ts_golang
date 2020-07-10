package HelloWorld

import (
	"errors"
	"fmt"
	"testing"
)

// 在某些时候，您可能希望使用`struct`来管理状态`state`，并公开一些方法，以使用户以您可以控制的方式更改状态。

/** Deposit
* 1. line:13 -> The compiler doesn't know what a Wallet is so let's tell it.  Now we've made our wallet, try and run the test again
* 2. wallet.Deposit/wallet.Balance undefined -> We need to define these methods.
* 3. Remember to only do enough to make the tests run. We need to make sure our test fails correctly with a clear error message.
* 4. We will need some kind of balance variable in our struct to store the state. Wallet add balance int.
# 5. In Go, when you call a function or a method the arguments are copied.
* When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
* A wallet, it is stored somewhere in memory. You can find out what the address of that bit of memory with &myVal
* We can test it from diff addresses.
* You can see that the addresses of the two balances are different. So when we change the value of the balance inside the code,
* we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.
* The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
*/

/** Deposit
 * Refactor
 * 1. To make Bitcoin you just use the syntax Bitcoin(999).
 * 2. By doing this we're making a new type and we can declare methods on them.
 * This can be very useful when you want to add some domain specific functionality on top of existing types.
 */

type Bitcoin int // 比特币

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Address of balance in test is %v \n", &w.balance)
	w.balance += amount
}

// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds // This is a positive change in itself because now our Withdraw function looks very clear.
	}
	w.balance -= amount
	// Again, it is very important to just write enough code to satisfy the compiler.
	// We correct our Withdraw method to return error and for now we have to return something so let's just return nil.
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance // w.balance
}

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	// Hopefully when returning an error of "oh no" you were thinking that we might iterate on that because it doesn't seem that useful to return.
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		// We want Withdraw to return an error if you try to take out more than you have and the balance should stay the same.
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.withdraw(Bitcoin(100))
		/*
		 * Err: wallet.Withdraw(Bitcoin(100)) used as value
		 * The wording is perhaps a little unclear, but our previous intent with Withdraw was just to call it, it will never return a value.
		 * To make this compile we will need to change it so it has a return type.
		 *
		 */

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{Bitcoin(20)}
		wallet.withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}

// ################################################################################################

// ################################################################################################
