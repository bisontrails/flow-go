package hash

import (
	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/model/encoding"
	"github.com/dapperlabs/flow-go/model/flow"
)

// DefaultHasher is the default hasher used by Flow.
var DefaultHasher, _ = crypto.NewHasher(crypto.SHA3_256)

// SetTransactionHash computes a hash and saves it on the transaction.
func SetTransactionHash(tx *flow.Transaction) error {
	b, err := encoding.DefaultEncoder.EncodeTransaction(*tx)
	if err != nil {
		return err
	}

	hash := DefaultHasher.ComputeHash(b)

	tx.Hash = hash

	return nil
}

// SetChunkHash computes a hash and saves it on the chunk.
func SetChunkHash(c *flow.Chunk) error {
	b, err := encoding.DefaultEncoder.EncodeChunk(*c)
	if err != nil {
		return err
	}

	hash := DefaultHasher.ComputeHash(b)

	c.Hash = hash

	return nil
}
