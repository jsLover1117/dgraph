/*
 * Copyright 2020 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package worker

import (
	"context"
	"math"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"

	"github.com/dgraph-io/dgraph/protos/pb"
	"github.com/dgraph-io/dgraph/x"
)

// predicateSet is a map whose keys are predicates. It is meant to be used as a set.
type predicateSet map[string]struct{}

// Manifest records backup details, these are values used during restore.
// Since is the timestamp from which the next incremental backup should start (it's set
// to the readTs of the current backup).
// Groups are the IDs of the groups involved.
type Manifest struct {
	sync.Mutex
	//Type is the type of backup, either full or incremental.
	Type string `json:"type"`
	// Since is the timestamp at which this backup was taken. It's called Since
	// because it will become the timestamp from which to backup in the next
	// incremental backup.
	Since uint64 `json:"since"`
	// Groups is the map of valid groups to predicates at the time the backup was created.
	Groups map[uint32][]string `json:"groups"`
	// BackupId is a unique ID assigned to all the backups in the same series
	// (from the first full backup to the last incremental backup).
	BackupId string `json:"backup_id"`
	// BackupNum is a monotonically increasing number assigned to each backup in
	// a series. The full backup as BackupNum equal to one and each incremental
	// backup gets assigned the next available number. Used to verify the integrity
	// of the data during a restore.
	BackupNum uint64 `json:"backup_num"`
	// Version specifies the Dgraph version, the backup was taken on. For the backup taken on older
	// versions (<= 20.11), the predicates in Group map do not have namespace. Version will be zero
	// for older versions.
	Version int `json:"version"`
	// Path is the name of the backup directory to which this manifest belongs to.
	Path string `json:"path"`
	// Encrypted indicates whether this backup was encrypted or not.
	Encrypted bool `json:"encrypted"`
	// DropOperations lists the various DROP operations that took place since the last backup.
	// These are used during restore to redo those operations before applying the backup.
	DropOperations []*pb.DropOperation `json:"drop_operations"`
	// Compression keeps track of the compression that was used for the data.
	Compression string `json:"compression"`
}

type MasterManifest struct {
	Manifests []*Manifest
}

func (m *Manifest) getPredsInGroup(gid uint32) predicateSet {
	preds, ok := m.Groups[gid]
	if !ok {
		return nil
	}

	predSet := make(predicateSet)
	for _, pred := range preds {
		predSet[pred] = struct{}{}
	}
	return predSet
}

// GetCredentialsFromRequest extracts the credentials from a backup request.
func GetCredentialsFromRequest(req *pb.BackupRequest) *x.MinioCredentials {
	return &x.MinioCredentials{
		AccessKey:    req.GetAccessKey(),
		SecretKey:    req.GetSecretKey(),
		SessionToken: req.GetSessionToken(),
		Anonymous:    req.GetAnonymous(),
	}
}

func StoreExport(request *pb.ExportRequest, dir string, key x.Sensitive) error {
	db, err := badger.OpenManaged(badger.DefaultOptions(dir).
		WithSyncWrites(false).
		WithValueThreshold(1 << 10).
		WithNumVersionsToKeep(math.MaxInt32).
		WithEncryptionKey(key))

	if err != nil {
		return err
	}

	_, err = exportInternal(context.Background(), request, db, true)
	// It is important to close the db before sending err to ch. Else, we will see a memory
	// leak.
	db.Close()
	return errors.Wrapf(err, "cannot export data inside DB at %s", dir)
}
