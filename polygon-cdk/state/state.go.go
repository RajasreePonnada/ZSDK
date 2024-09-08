package state

import (
    "github.com/ethereum/go-ethereum/core/state"
)

type CDKState struct {
    stateDB *state.StateDB
}

func NewCDKState(stateDB *state.StateDB) *CDKState {
    return &CDKState{stateDB: stateDB}
}

func (s *CDKState) Set(key, value []byte) {
    s.stateDB.SetState(common.BytesToAddress(key), common.BytesToHash(value))
}

func (s *CDKState) Get(key []byte) []byte {
    return s.stateDB.GetState(common.BytesToAddress(key), common.Hash{}).Bytes()
}