// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/opentofu/opentofu/external/grpcwrap"
	"github.com/opentofu/opentofu/external/plugin"
	simple "github.com/opentofu/opentofu/external/provider-simple"
	"github.com/opentofu/opentofu/external/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
