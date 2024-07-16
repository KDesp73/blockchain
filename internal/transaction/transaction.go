package transaction

import (
	"blockchain/internal/encryption"
	"crypto/rsa"
	"fmt"
	"strconv"
)

type Transaction struct {
	Amount int
	Payer *rsa.PublicKey
	Payee *rsa.PublicKey
}

func (t Transaction) ToString() string {
	payerKey, err := encryption.FormatPublicKey(t.Payer);
	
	if err != nil {
		fmt.Println("Could not format Payer's public key: ", err)
		return ""
	}

	payeeKey, err := encryption.FormatPublicKey(t.Payee);

	if err != nil {
		fmt.Println("Could not format Payee's public key: ", err)
		return ""
	}

	return strconv.Itoa(t.Amount) + payerKey + payeeKey
}
