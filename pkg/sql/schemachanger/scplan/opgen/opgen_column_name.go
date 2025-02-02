// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package opgen

import (
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scpb"
)

func init() {
	opRegistry.register(
		(*scpb.ColumnName)(nil),
		add(
			to(scpb.Status_PUBLIC,
				minPhase(scop.PreCommitPhase),
				emit(func(this *scpb.ColumnName) scop.Op {
					return &scop.SetColumnName{
						TableID:  this.TableID,
						ColumnID: this.ColumnID,
						Name:     this.Name,
					}
				}),
			),
		),
		drop(
			to(scpb.Status_ABSENT,
				minPhase(scop.PostCommitPhase),
				emit(func(this *scpb.ColumnName) scop.Op {
					return &scop.SetColumnName{
						TableID:  this.TableID,
						ColumnID: this.ColumnID,
						Name:     this.Name,
					}
				}),
			),
		),
	)
}
