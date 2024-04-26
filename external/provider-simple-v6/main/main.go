// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/opentofu/opentofu/external/grpcwrap"
	plugin "github.com/opentofu/opentofu/external/plugin6"
	simple "github.com/opentofu/opentofu/external/provider-simple-v6"
	"github.com/opentofu/opentofu/external/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
