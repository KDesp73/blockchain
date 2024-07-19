package wallet

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/encryption"
	"blockchain/internal/transaction"
	"crypto/rsa"
	"fmt"
)

type Wallet struct {
	Name string
	PublicKey *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewWallet(name string) *Wallet {
	wallet := &Wallet{Name: name}

	public, private, err := encryption.GenerateRSAKeys()

    if err != nil {
        fmt.Println("Error generating RSA keys: ", err)
        return nil
    }

	wallet.PublicKey = public
	wallet.privateKey = private

    return wallet
}

func (w *Wallet) ToString() string {
	publicKey, err := encryption.FormatPublicKey(w.PublicKey)
	
	if err != nil {
		fmt.Println("Could not format public key: ", err)
		return ""
	}

	return fmt.Sprintf("name: %s\npublic key: %s\n", w.Name, publicKey)
}

func (w* Wallet) SendMoney(bc *blockchain.Blockchain, amount int, payeePublicKey *rsa.PublicKey) {
	t := transaction.Transaction {
		Amount: amount, 
		Payer: w.PublicKey,
		Payee: payeePublicKey,
	}

	signature, err := encryption.SignData(w.privateKey, []byte(t.ToString()))

	if err != nil {
		fmt.Println("Could not sign the transaction: ", err)
		return
	}

	bc.MineBlock(t, w.PublicKey, string(signature))
}
