/*
Copyright 2022 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cockroachdb

import (
	"context"

	"github.com/dapr/components-contrib/state"
)

// dbAccess is a private interface which enables unit testing of CockroachDB.
type dbAccess interface {
	Init(metadata state.Metadata) error
	Set(ctx context.Context, req *state.SetRequest) error
	BulkSet(ctx context.Context, req []state.SetRequest) error
	Get(ctx context.Context, req *state.GetRequest) (*state.GetResponse, error)
	Delete(ctx context.Context, req *state.DeleteRequest) error
	BulkDelete(ctx context.Context, req []state.DeleteRequest) error
	ExecuteMulti(req *state.TransactionalStateRequest) error
	Query(req *state.QueryRequest) (*state.QueryResponse, error)
	Ping(ctx context.Context) error
	Close() error
}
