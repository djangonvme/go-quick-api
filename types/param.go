package types

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

// header 中 token key
const (
	AppName           = "task-dispatcher"
	ArgConfig         = "config"
	ArgConfigFilename = "config.ini"
	TokenHeaderKey    = "Authorization"
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

	TaskTypeKey         = "task_type"
	TaskTypeLotusCommit = "lotus-commit"
	TaskTypeLotusWPost  = "lotus-WPost"

	KeyAllowTasks string = "allow_tasks"
)

var TaskTypesMap = map[string]struct{}{
	TaskTypeLotusCommit: {},
	TaskTypeLotusWPost:  {},
}

/*var TaskStateMap = map[TaskState]struct{}{
	TaskStateWaiting: {},
	TaskStateDoing: {},
	TaskStateFinished: {},
	TaskStateDropped: {},
}
*/

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
