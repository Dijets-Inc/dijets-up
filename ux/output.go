// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.
package ux

import (
	"fmt"

	"github.com/lasthyphen/dijetsnodego/utils/logging"
)

func Print(log logging.Logger, msg string, args ...interface{}) {
	fmtMsg := fmt.Sprintf(msg, args...)
	fmt.Println(fmtMsg)
	log.Info(fmtMsg)
}
