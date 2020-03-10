package pointer

import "testing"

func TestWallet(t *testing.T) {

	// wallet := Wallet{}
	// wallet.Deposit(Bitcoin(20))

	// want := Bitcoin(10)
	// got := wallet.Balance()

	// if want != got {
	// 	t.Errorf("want %s got %s", want, got)

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))

		want := Bitcoin(100)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(34.5))

		want := Bitcoin(65.5)

		assertNoError(t, err)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw with a nil error", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, startingBalance)
		assertErrors(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	t.Helper()
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}

func assertErrors(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %s, want %s", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didnt want one")
	}
}
