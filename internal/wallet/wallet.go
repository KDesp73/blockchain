package wallet

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rsa"
	"fmt"
)

type Wallet struct {
	name string
	PublicKey *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}
func NewWallet(name string) *Wallet {
	wallet := &Wallet{name: name} // Initialize wallet with a new instance of Wallet

	public, private, err := encryption.GenerateRSAKeys()

    if err != nil {
        fmt.Println("Error generating RSA keys: ", err)
        return nil
    }

	wallet.PublicKey = public
	wallet.PrivateKey = private

    return wallet
}

func (w *Wallet) ToString() string {
	publicKey, err := encryption.FormatPublicKey(w.PublicKey)
	
	if err != nil {
		fmt.Println("Could not format public key: ", err)
		return ""
	}

	return fmt.Sprintf("name: %s\npublic key: %s\n", w.name, publicKey)
}

func (w* Wallet) SendMoney(bc *blockchain.Blockchain, amount int, payeePublicKey *rsa.PublicKey) {
	t := transaction.Transaction {
		Amount: amount, 
		Payer: w.PublicKey,
		Payee: payeePublicKey,
	}

	signature, err := encryption.SignData(w.PrivateKey, []byte(t.ToString()))

	if err != nil {
		fmt.Println("Could not sign the transaction: ", err)
		return
	}

	bc.MineBlock(t, w.PublicKey, string(signature))
}
