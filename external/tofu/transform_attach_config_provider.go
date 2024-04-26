// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofu

import (
	"github.com/opentofu/opentofu/external/addrs"
	"github.com/opentofu/opentofu/external/configs"
)

// GraphNodeAttachProvider is an interface that must be implemented by nodes
// that want provider configurations attached.
type GraphNodeAttachProvider interface {
	// ProviderName with no module prefix. Example: "aws".
	ProviderAddr() addrs.AbsProviderConfig

	// Sets the configuration
	AttachProvider(*configs.Provider)
}