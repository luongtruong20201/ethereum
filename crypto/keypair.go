package crypto

import (
	"ethereum/util"
	"strings"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
	address    []byte
	mnemonic   string
}

func GenerateNewKeyPair() *KeyPair {
	_, priv := secp256k1.GenerateKeyPair()
	keyPair, _ := NewKeyPairFromSec(priv)
	return keyPair
}

func NewKeyPairFromSec(secKey []byte) (*KeyPair, error) {
	pubKey, err := secp256k1.GeneratePubKey(secKey)
	if err != nil {
		return nil, err
	}
	return &KeyPair{
		PrivateKey: secKey,
		PublicKey:  pubKey,
	}, nil
}

func (k *KeyPair) Address() []byte {
	if k.address == nil {
		k.address = Sha3Bin(k.PublicKey[1:])[12:]
	}
	return k.address
}

func (k *KeyPair) Mnemonic() string {
	if k.mnemonic == "" {
		k.mnemonic = strings.Join(MnemonicEncode(util.Bytes2Hex(k.PrivateKey)), " ")
	}
	return k.mnemonic
}

func (k *KeyPair) AsStrings() (string, string, string, string) {
	return k.Mnemonic(), util.Bytes2Hex(k.Address()), util.Bytes2Hex(k.PrivateKey), util.Bytes2Hex(k.PublicKey)
}

func (k *KeyPair) RlpEncode() []byte {
	return k.RlpValue().Encode()
}

func (k *KeyPair) RlpValue() *util.Value {
	return util.NewValue(k.PrivateKey)
}
