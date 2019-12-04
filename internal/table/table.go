// Copyright 2019 Grabtaxi Holdings PTE LTE (GRAB), All rights reserved.
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file

package table

import (
	"io"
	"reflect"

	"github.com/grab/talaria/internal/presto"
)

// Table represents a table storage contract.
type Table interface {
	io.Closer
	Name() string
	Schema() (map[string]reflect.Type, error)
	GetSplits(desiredColumns []string, outputConstraint *presto.PrestoThriftTupleDomain, maxSplitCount int) ([]Split, error)
	GetRows(splitID []byte, columns []string, maxBytes int64) (*PageResult, error)
}

// Appender represents an appender of data to the table.
type Appender interface {
	Append(payload []byte) error
}

// Split represents a split
type Split struct {
	Key   []byte   // The key of the split (SplitID).
	Addrs []string // The list of hosts
}

// PageResult represents a result
type PageResult struct {
	Columns   []presto.Column // The list of columns returned
	NextToken []byte          // The next token if the result is incomplete
}
