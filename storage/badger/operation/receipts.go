package operation

import (
	"github.com/dgraph-io/badger/v2"

	"github.com/onflow/flow-go/model/flow"
)

// InsertExecutionReceiptMeta inserts an execution receipt meta by ID.
func InsertExecutionReceiptMeta(receiptID flow.Identifier, meta *flow.ExecutionReceiptMeta) func(*badger.Txn) error {
	return insert(makePrefix(codeExecutionReceiptMeta, receiptID), meta)
}

// RetrieveExecutionReceipt retrieves a execution receipt meta by ID.
func RetrieveExecutionReceiptMeta(receiptID flow.Identifier, meta *flow.ExecutionReceiptMeta) func(*badger.Txn) error {
	return retrieve(makePrefix(codeExecutionReceiptMeta, receiptID), meta)
}

// IndexExecutionReceipt inserts an execution receipt ID keyed by block ID
func IndexExecutionReceipt(blockID flow.Identifier, receiptID flow.Identifier) func(*badger.Txn) error {
	return insert(makePrefix(codeOwnBlockReceipt, blockID), receiptID)
}

// LookupExecutionReceipt finds execution receipt ID by block
func LookupExecutionReceipt(blockID flow.Identifier, receiptID *flow.Identifier) func(*badger.Txn) error {
	return retrieve(makePrefix(codeOwnBlockReceipt, blockID), receiptID)
}

// Add2IndexAllBlockReceipts inserts an execution receipt ID keyed by block ID and execution ID
func Add2IndexAllBlockReceipts(blockID, receiptID flow.Identifier) func(*badger.Txn) error {
	return insert(makePrefix(codeAllBlockReceipts, blockID, receiptID), receiptID)
}

// LookupAllBlockReceipts finds execution receipt ID by block ID for all execution IDs
func LookupAllBlockReceipts(blockID flow.Identifier, receiptIDs *[]flow.Identifier) func(*badger.Txn) error {
	iterationFunc := receiptIterationFunc(receiptIDs)
	return traverse(makePrefix(codeAllBlockReceipts, blockID), iterationFunc)
}

// receiptIterationFunc returns an in iteration function which returns all receipt IDs found during traversal
func receiptIterationFunc(receiptIDs *[]flow.Identifier) func() (checkFunc, createFunc, handleFunc) {
	return func() (checkFunc, createFunc, handleFunc) {
		check := func(key []byte) bool {
			return true
		}
		var val flow.Identifier
		create := func() interface{} {
			return &val
		}
		handle := func() error {
			*receiptIDs = append(*receiptIDs, val)
			return nil
		}
		return check, create, handle
	}
}
