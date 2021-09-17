// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockfetcher

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// ChunkDataPackHandler is an autogenerated mock type for the ChunkDataPackHandler type
type ChunkDataPackHandler struct {
	mock.Mock
}

// HandleChunkDataPack provides a mock function with given fields: originID, chunkDataPack
func (_m *ChunkDataPackHandler) HandleChunkDataPack(originID flow.Identifier, chunkDataPack *flow.ChunkDataPack) {
	_m.Called(originID, chunkDataPack)
}

// NotifyChunkDataPackSealed provides a mock function with given fields: chunkID
func (_m *ChunkDataPackHandler) NotifyChunkDataPackSealed(chunkID flow.Identifier) {
	_m.Called(chunkID)
}