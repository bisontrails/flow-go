package badger

import (
	"github.com/dgraph-io/badger/v2"

	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/storage"
)

func InitAll(metrics module.CacheMetrics, db *badger.DB) *storage.All {
	headers := NewHeaders(metrics, db)
	guarantees := NewGuarantees(metrics, db)
	seals := NewSeals(metrics, db)
	index := NewIndex(metrics, db)
	results := NewExecutionResults(metrics, db)
	receipts := NewExecutionReceipts(metrics, db, results)
	payloads := NewPayloads(db, index, guarantees, seals, receipts)
	blocks := NewBlocks(db, headers, payloads)
	setups := NewEpochSetups(metrics, db)
	epochCommits := NewEpochCommits(metrics, db)
	statuses := NewEpochStatuses(metrics, db)

	chunkDataPacks := NewChunkDataPacks(db)
	commits := NewCommits(metrics, db)
	transactions := NewTransactions(metrics, db)
	transactionResults := NewTransactionResults(db)
	collections := NewCollections(db, transactions)

	return &storage.All{
		Headers:            headers,
		Guarantees:         guarantees,
		Seals:              seals,
		Index:              index,
		Payloads:           payloads,
		Blocks:             blocks,
		Setups:             setups,
		EpochCommits:       epochCommits,
		Statuses:           statuses,
		Results:            results,
		Receipts:           receipts,
		ChunkDataPacks:     chunkDataPacks,
		Commits:            commits,
		Transactions:       transactions,
		TransactionResults: transactionResults,
		Collections:        collections,
	}
}