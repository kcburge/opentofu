// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package init

import (
	"testing"

	"github.com/opentofu/opentofu/external/backend/remote-state/inmem"
	"github.com/zclconf/go-cty/cty"
)

func TestDeprecateBackend(t *testing.T) {
	deprecateMessage := "deprecated backend"
	deprecatedBackend := deprecateBackend(
		inmem.New(),
		deprecateMessage,
	)

	_, diags := deprecatedBackend.PrepareConfig(cty.EmptyObjectVal)
	if len(diags) != 1 {
		t.Errorf("got %d diagnostics; want 1", len(diags))
		for _, diag := range diags {
			t.Errorf("- %s", diag)
		}
		return
	}

	desc := diags[0].Description()
	if desc.Summary != deprecateMessage {
		t.Fatalf("wrong message %q; want %q", desc.Summary, deprecateMessage)
	}
}