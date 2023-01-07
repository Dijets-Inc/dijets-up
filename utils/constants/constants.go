// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.
package constants

import "path/filepath"

const (
	LogNameMain    = "main"
	LogNameControl = "control"
	LogNameTest    = "test"
)

var (
	LocalConfigDir   = filepath.Join("local", "default")
	LocalGenesisFile = "genesis.json"
)
