/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package pingpong

import (
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc"
)

func Topology() []nwo.Topology {
	// Create an empty FSC topology
	topology := fsc.NewTopology()

	// Add the initiator fsc node
	topology.AddNodeByName("initiator").SetExecutable(
		"github.com/hyperledger-labs/fabric-smart-client/integration/generic/pingpong/cmd/initiator",
	)
	// Add the responder fsc node
	topology.AddNodeByName("responder").RegisterResponder(
		&Responder{}, &Initiator{},
	)
	return []nwo.Topology{topology}
}
