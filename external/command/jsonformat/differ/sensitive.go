// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/opentofu/opentofu/external/command/jsonformat/computed"
	"github.com/opentofu/opentofu/external/command/jsonformat/computed/renderers"
	"github.com/opentofu/opentofu/external/command/jsonformat/structured"
	"github.com/opentofu/opentofu/external/command/jsonprovider"
	"github.com/opentofu/opentofu/external/plans"
)

type CreateSensitiveRenderer func(computed.Diff, bool, bool) computed.DiffRenderer

func checkForSensitiveType(change structured.Change, ctype cty.Type) (computed.Diff, bool) {
	return change.CheckForSensitive(
		func(value structured.Change) computed.Diff {
			return ComputeDiffForType(value, ctype)
		}, func(inner computed.Diff, beforeSensitive, afterSensitive bool, action plans.Action) computed.Diff {
			return computed.NewDiff(renderers.Sensitive(inner, beforeSensitive, afterSensitive), action, change.ReplacePaths.Matches())
		},
	)
}

func checkForSensitiveNestedAttribute(change structured.Change, attribute *jsonprovider.NestedType) (computed.Diff, bool) {
	return change.CheckForSensitive(
		func(value structured.Change) computed.Diff {
			return computeDiffForNestedAttribute(value, attribute)
		}, func(inner computed.Diff, beforeSensitive, afterSensitive bool, action plans.Action) computed.Diff {
			return computed.NewDiff(renderers.Sensitive(inner, beforeSensitive, afterSensitive), action, change.ReplacePaths.Matches())
		},
	)
}

func checkForSensitiveBlock(change structured.Change, block *jsonprovider.Block) (computed.Diff, bool) {
	return change.CheckForSensitive(
		func(value structured.Change) computed.Diff {
			return ComputeDiffForBlock(value, block)
		}, func(inner computed.Diff, beforeSensitive, afterSensitive bool, action plans.Action) computed.Diff {
			return computed.NewDiff(renderers.SensitiveBlock(inner, beforeSensitive, afterSensitive), action, change.ReplacePaths.Matches())
		},
	)
}