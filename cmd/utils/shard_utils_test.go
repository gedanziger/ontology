/*
 * Copyright (C) 2019 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils_test

import (
	"fmt"
	"github.com/ontio/ontology/cmd/utils"
	"testing"
)

func TestBuildShardCommandArgs(t *testing.T) {
	args := make(map[string]string)
	args["a"] = "1"
	args["b"] = "2"

	shardID := uint64(2)
	parentShardID := uint64(1)
	parentPort := uint64(10000)
	cmdArgs, err := utils.BuildShardCommandArgs(args, shardID, parentShardID, parentPort)
	if err != nil {
		t.Fatalf("failed to build shard cmd args: %s", err)
	}

	if !isExits(cmdArgs, "--a=1") {
		t.Fatalf("arg 'a' not exist in %v", cmdArgs)
	}
	if !isExits(cmdArgs, "--b=2") {
		t.Fatalf("arg 'a' not exist in %v", cmdArgs)
	}
	if !isExits(cmdArgs, fmt.Sprintf("--%s=%d", utils.ShardIDFlag.GetName(), shardID)) {
		t.Fatalf("arg 'a' not exist in %v", cmdArgs)
	}
	if !isExits(cmdArgs, fmt.Sprintf("--%s=%d", utils.ParentShardIDFlag.GetName(), parentShardID)) {
		t.Fatalf("arg 'a' not exist in %v", cmdArgs)
	}
	if !isExits(cmdArgs, fmt.Sprintf("--%s=%d", utils.ParentShardPortFlag.GetName(), parentPort)) {
		t.Fatalf("arg 'a' not exist in %v", cmdArgs)
	}
}

func isExits(args []string, arg string) bool {
	for _, s := range args {
		if s == arg {
			return true
		}
	}
	return false
}
