// Copyright 2023 The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	_ "github.com/corazawaf/coraza-daddy"

	_ "embed"
    geo "github.com/corazawaf/coraza-geoip"
)

func init() {
    geo.RegisterDatabaseFromFile("geoip-database.mmdb", "city")
}

func main() {
	caddycmd.Main()
}
