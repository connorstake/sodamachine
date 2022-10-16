package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Deposit  int    `json:"deposit"`
	Role     string `json:"role"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) DecreaseFunds(amount int) error {
	if amount > user.Deposit {
		return fmt.Errorf("there is not enough money deposited fulfill this order. cost: %v, deposited: %v, username: %s", amount, user.Deposit, user.Username)

	}
	user.Deposit = user.Deposit - amount
	return nil
}

func (user *User) GetChange() ([]int, error) {
	denoms := []int{100, 50, 20, 10, 5}
	var change []int
	for _, denom := range denoms {
		if user.Deposit%denom >= 0 {
			numCoins := (user.Deposit - (user.Deposit % denom)) / denom
			for i := 0; i < numCoins; i++ {
				change = append(change, denom)
				user.DecreaseFunds(denom)
			}
		}
	}
	return change, nil
}

func (user *User) DepositFunds(amount int) error {
	tm := createAcceptedTokenMap()

	if _, ok := tm[amount]; !ok {
		return fmt.Errorf("deposited funds are not in correct denomination")
	}

	user.Deposit += amount
	return nil
}

func (user *User) ResetDeposit() error {
	user.Deposit = 0
	return nil
}

func createAcceptedTokenMap() map[int]bool {
	acceptedTokens := make(map[int]bool)
	acceptedTokens[5] = true
	acceptedTokens[10] = true
	acceptedTokens[20] = true
	acceptedTokens[50] = true
	acceptedTokens[100] = true

	return acceptedTokens
}
