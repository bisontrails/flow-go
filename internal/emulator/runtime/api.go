package runtime

import (
	"math/big"

	crypto "github.com/dapperlabs/bamboo-node/pkg/crypto/oldcrypto"

	etypes "github.com/dapperlabs/bamboo-node/internal/emulator/types"
)

type EmulatorRuntimeAPI struct {
	registers *etypes.RegistersView
}

func NewEmulatorRuntimeAPI(registers *etypes.RegistersView) *EmulatorRuntimeAPI {
	return &EmulatorRuntimeAPI{registers}
}

func (i *EmulatorRuntimeAPI) GetValue(owner, controller, key []byte) ([]byte, error) {
	v, _ := i.registers.Get(fullKey(owner, controller, key))
	return v, nil
}

func (i *EmulatorRuntimeAPI) SetValue(owner, controller, key, value []byte) error {
	i.registers.Set(fullKey(owner, controller, key), value)
	return nil
}

func (i *EmulatorRuntimeAPI) CreateAccount(publicKey, code []byte) (id []byte, err error) {
	latestAccountID, _ := i.registers.Get(keyLatestAccount())

	accountIDInt := big.NewInt(0).SetBytes(latestAccountID)
	accountIDBytes := accountIDInt.Add(accountIDInt, big.NewInt(1)).Bytes()

	accountAddress := crypto.BytesToAddress(accountIDBytes)

	accountID := accountAddress.Bytes()

	i.registers.Set(fullKey(accountID, []byte{}, keyBalance()), big.NewInt(0).Bytes())
	i.registers.Set(fullKey(accountID, accountID, keyPublicKey()), publicKey)
	i.registers.Set(fullKey(accountID, accountID, keyCode()), code)

	i.registers.Set(keyLatestAccount(), accountID)

	address := crypto.BytesToAddress(accountID)

	return address.Bytes(), nil
}

func (i *EmulatorRuntimeAPI) GetAccount(address crypto.Address) *crypto.Account {
	accountID := address.Bytes()

	balanceBytes, exists := i.registers.Get(fullKey(accountID, []byte{}, keyBalance()))
	if !exists {
		return nil
	}

	balanceInt := big.NewInt(0).SetBytes(balanceBytes)

	publicKey, _ := i.registers.Get(fullKey(accountID, accountID, []byte("public_key")))
	code, _ := i.registers.Get(fullKey(accountID, accountID, keyCode()))

	return &crypto.Account{
		Address:    address,
		Balance:    balanceInt.Uint64(),
		Code:       code,
		PublicKeys: [][]byte{publicKey},
	}
}

func keyLatestAccount() crypto.Hash {
	return crypto.NewHash([]byte("latestAccount"))
}

func keyBalance() []byte {
	return []byte("balance")
}

func keyPublicKey() []byte {
	return []byte("public_key")
}

func keyCode() []byte {
	return []byte("code")
}

func fullKey(owner, controller, key []byte) crypto.Hash {
	fullKey := append(owner, controller...)
	fullKey = append(fullKey, key...)
	return crypto.NewHash(fullKey)
}
