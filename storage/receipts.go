// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package storage

import (
	"github.com/onflow/flow-go/model/flow"
)

type ExecutionReceipts interface {

	// Store stores an execution receipt.
	Store(result *flow.ExecutionReceipt) error

	// ByID retrieves an execution receipt by its ID.
	ByID(resultID flow.Identifier) (*flow.ExecutionReceipt, error)

	// Index indexes an execution receipt by block ID.
	Index(blockID flow.Identifier, resultID flow.Identifier) error

	// Add2IndexAllBlockReceipts adds the receipt to the index of all receipts for the block
	Add2IndexAllBlockReceipts(receipt *flow.ExecutionReceipt) error

	// ByBlockID retrieves an execution receipt by block ID.
	ByBlockID(blockID flow.Identifier) (*flow.ExecutionReceipt, error)

	// GetAllBlockReceipts retrieves all execution receipts for a block ID
	GetAllBlockReceipts(blockID flow.Identifier) ([]*flow.ExecutionReceipt, error)
}
