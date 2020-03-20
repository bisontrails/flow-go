package flow

type ChunkBody struct {
	CollectionIndex uint

	// execution info
	StartState      StateCommitment // start state when starting executing this chunk
	EventCollection Identifier      // Events generated by executing results

	// Computation consumption info
	TotalComputationUsed uint64 // total amount of computation used by running all txs in this chunk
	NumberOfTransactions uint64 // number of transactions inside the collection
}

type Chunk struct {
	ChunkBody

	Index uint64 // chunk index inside the ER (starts from zero)
	// EndState inferred from next chunk or from the ER
	EndState StateCommitment
}

// ID returns a unique id for this entity
func (ch *Chunk) ID() Identifier {
	return MakeID(ch.ChunkBody)
}

// Checksum provides a cryptographic commitment for a chunk content
func (ch *Chunk) Checksum() Identifier {
	return MakeID(ch)
}

// RegisterTouch captures the register value before an update or read
type RegisterTouch struct {
	RegisterID RegisterID
	Value      RegisterValue
	Proof      StorageProof
}

// ChunkDataPack holds all register touches (any read, or write)
// note that we have to capture a read proof for each write before updating the registers
type ChunkDataPack struct {
	ChunkID         Identifier
	StartState      StateCommitment
	RegisterTouches []RegisterTouch
}

// ID returns the unique identifier for the concrete view, which is the ID of
// the chunk the view is for.
func (c *ChunkDataPack) ID() Identifier {
	return c.ChunkID
}

// Checksum returns the checksum of the chunk data pack.
func (c *ChunkDataPack) Checksum() Identifier {
	return MakeID(c)
}

type ChunkHeader struct {
	ChunkID     Identifier
	StartState  StateCommitment
	RegisterIDs []RegisterID
}

// ChunkState represents the state registers used by a particular chunk.
type ChunkState struct {
	ChunkID   Identifier
	Registers Ledger
}

// ID returns the unique identifier for the concrete view, which is the ID of
// the chunk the view is for.
func (c *ChunkState) ID() Identifier {
	return c.ChunkID
}

// Checksum returns the checksum of the chunk state.
func (c *ChunkState) Checksum() Identifier {
	return MakeID(c)
}

// Note that this is the basic version of the List, we need to substitute it with something like Merkel tree at some point
type ChunkList []*Chunk

func (cl ChunkList) Fingerprint() Identifier {
	return MerkleRoot(GetIDs(cl)...)
}

func (cl *ChunkList) Insert(ch *Chunk) {
	*cl = append(*cl, ch)
}

func (cl ChunkList) Items() []*Chunk {
	return cl
}

// ByChecksum returns an entity from the list by entity fingerprint
func (cl ChunkList) ByChecksum(cs Identifier) (*Chunk, bool) {
	for _, ch := range cl {
		if ch.Checksum() == cs {
			return ch, true
		}
	}
	return nil, false
}

// ByIndex returns an entity from the list by index
// Todo return a boolean for out of bound
// https://github.com/dapperlabs/flow-go/issues/2796
func (cl ChunkList) ByIndex(i uint64) *Chunk {
	return cl[i]
}

// ByIndexWithProof returns an entity from the list by index and proof of membership
func (cl ChunkList) ByIndexWithProof(i uint64) (*Chunk, Proof) {
	return cl[i], nil
}

// Len returns the number of Chunks in the list. It is also part of the sort
// interface that makes ChunkList sortable
func (cl ChunkList) Len() int {
	return len(cl)
}

// Less returns true if element i in the ChunkList is less than j based on its chunk ID.
// Otherwise it returns true.
// It satisfies the sort.Interface making the ChunkList sortable.
func (cl ChunkList) Less(i, j int) bool {
	return cl[i].ID().String() < cl[j].ID().String()
}

// Swap swaps the element i and j in the ChunkList.
// It satisfies the sort.Interface making the ChunkList sortable.
func (cl ChunkList) Swap(i, j int) {
	cl[j], cl[i] = cl[i], cl[j]
}
