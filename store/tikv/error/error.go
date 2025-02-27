// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package error

import (
	"github.com/pingcap/errors"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/pingcap/kvproto/pkg/pdpb"
	"github.com/pingcap/tidb/util/dbterror"
)

var (
	// ErrBodyMissing response body is missing error
	ErrBodyMissing = errors.New("response body is missing")
	// ErrTiDBShuttingDown is returned when TiDB is closing and send request to tikv fail, do not retry.
	ErrTiDBShuttingDown = errors.New("tidb server shutting down")
)

// MismatchClusterID represents the message that the cluster ID of the PD client does not match the PD.
const MismatchClusterID = "mismatch cluster id"

// error instances.
var (
	ErrTiKVServerTimeout           = dbterror.ClassTiKV.NewStd(CodeTiKVServerTimeout)
	ErrTiFlashServerTimeout        = dbterror.ClassTiKV.NewStd(CodeTiFlashServerTimeout)
	ErrResolveLockTimeout          = dbterror.ClassTiKV.NewStd(CodeResolveLockTimeout)
	ErrPDServerTimeout             = dbterror.ClassTiKV.NewStd(CodePDServerTimeout)
	ErrRegionUnavailable           = dbterror.ClassTiKV.NewStd(CodeRegionUnavailable)
	ErrTiKVServerBusy              = dbterror.ClassTiKV.NewStd(CodeTiKVServerBusy)
	ErrTiFlashServerBusy           = dbterror.ClassTiKV.NewStd(CodeTiFlashServerBusy)
	ErrTiKVStaleCommand            = dbterror.ClassTiKV.NewStd(CodeTiKVStaleCommand)
	ErrTiKVMaxTimestampNotSynced   = dbterror.ClassTiKV.NewStd(CodeTiKVMaxTimestampNotSynced)
	ErrGCTooEarly                  = dbterror.ClassTiKV.NewStd(CodeGCTooEarly)
	ErrQueryInterrupted            = dbterror.ClassTiKV.NewStd(CodeQueryInterrupted)
	ErrLockAcquireFailAndNoWaitSet = dbterror.ClassTiKV.NewStd(CodeLockAcquireFailAndNoWaitSet)
	ErrLockWaitTimeout             = dbterror.ClassTiKV.NewStd(CodeLockWaitTimeout)
	ErrTokenLimit                  = dbterror.ClassTiKV.NewStd(CodeTiKVStoreLimit)
	ErrUnknown                     = dbterror.ClassTiKV.NewStd(CodeUnknown)
)

// Registers error returned from TiKV.
var (
	_ = dbterror.ClassTiKV.NewStd(CodeDataOutOfRange)
	_ = dbterror.ClassTiKV.NewStd(CodeTruncatedWrongValue)
	_ = dbterror.ClassTiKV.NewStd(CodeDivisionByZero)
)

// ErrDeadlock wraps *kvrpcpb.Deadlock to implement the error interface.
// It also marks if the deadlock is retryable.
type ErrDeadlock struct {
	*kvrpcpb.Deadlock
	IsRetryable bool
}

func (d *ErrDeadlock) Error() string {
	return d.Deadlock.String()
}

// PDError wraps *pdpb.Error to implement the error interface.
type PDError struct {
	Err *pdpb.Error
}

func (d *PDError) Error() string {
	return d.Err.String()
}

// ErrKeyExist wraps *pdpb.AlreadyExist to implement the error interface.
type ErrKeyExist struct {
	*kvrpcpb.AlreadyExist
}

func (k *ErrKeyExist) Error() string {
	return k.AlreadyExist.String()
}

// IsErrKeyExist returns true if it is ErrKeyExist.
func IsErrKeyExist(err error) bool {
	_, ok := errors.Cause(err).(*ErrKeyExist)
	return ok
}

// ErrWriteConflict wraps *kvrpcpb.ErrWriteConflict to implement the error interface.
type ErrWriteConflict struct {
	*kvrpcpb.WriteConflict
}

func (k *ErrWriteConflict) Error() string {
	return k.WriteConflict.String()
}

// IsErrWriteConflict returns true if it is ErrWriteConflict.
func IsErrWriteConflict(err error) bool {
	_, ok := errors.Cause(err).(*ErrWriteConflict)
	return ok
}

//NewErrWriteConfictWithArgs generates an ErrWriteConflict with args.
func NewErrWriteConfictWithArgs(startTs, conflictTs, conflictCommitTs uint64, key []byte) *ErrWriteConflict {
	conflict := kvrpcpb.WriteConflict{
		StartTs:          startTs,
		ConflictTs:       conflictTs,
		Key:              key,
		ConflictCommitTs: conflictCommitTs,
	}
	return &ErrWriteConflict{WriteConflict: &conflict}
}

// ErrRetryable wraps *kvrpcpb.Retryable to implement the error interface.
type ErrRetryable struct {
	Retryable string
}

func (k *ErrRetryable) Error() string {
	return k.Retryable
}
