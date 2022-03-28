package types

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

const (
	TokenHeaderKey = "Authorization"
)

type TaskState string

type TaskWorkerState string

const (
	TaskStateWaiting  TaskState = "waiting"
	TaskStateDoing    TaskState = "doing"
	TaskStateFinished TaskState = "finished"
	TaskStateDropped  TaskState = "dropped"

	TaskWorkerStateDoing    TaskWorkerState = "doing"
	TaskWorkerStateFailed   TaskWorkerState = "failed"
	TaskWorkerStateFinished TaskWorkerState = "finished"
	TaskWorkerStateReverted TaskWorkerState = "reverted"

	TaskTypeKey         = "Task-type"
	TaskTypeLotusCommit = "lotus-commit2"
	TaskTypeLotusWPost  = "lotus-WPost"

	KeyAllowTasks string = "allow_tasks"
)

var TaskTypesMap = map[string]struct{}{
	TaskTypeLotusCommit: {},
	TaskTypeLotusWPost:  {},
}

type SectorCids struct {
	Unsealed cid.Cid
	Sealed   cid.Cid
}
type SectorRef struct {
	ID        abi.SectorID
	ProofType abi.RegisteredSealProof
}
type Commit2InputParam struct {
	SectorId   abi.SectorNumber `json:"sector_id"`
	MinerId    string           `json:"miner_id"`
	Commit1Out string           `json:"commit1_out"` // hex encoded
}
