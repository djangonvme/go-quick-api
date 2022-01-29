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

	TaskWorkerStateFailed   TaskWorkerState = "failed"
	TaskWorkerStateFinished TaskWorkerState = "finished"
	TaskWorkerStateReverted TaskWorkerState = "reverted"

	TaskTypeKey         = "Task-type"
	TaskTypeLotusCommit = "lotus-commit"
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
type CommitInputParam struct {
	Sector SectorRef
	Ticket abi.SealRandomness
	Seed   abi.InteractiveSealRandomness
	Pieces []abi.PieceInfo
	Cids   SectorCids
}
