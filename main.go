package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/protocol/alerthash"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

type Network struct {
	ChainId string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

type BlockEvent struct {
	BlockHash string  `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Network   Network `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
}

type Finding struct {
	Protocol    string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Severity    int64  `protobuf:"varint,2,opt,name=severity,proto3" json:"severity,omitempty"`
	Type        int64  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	AlertId     string `protobuf:"bytes,5,opt,name=alertId,proto3" json:"alertId,omitempty"`
	Name        string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

type BotInfo alerthash.BotInfo

type Alert struct {
	BlockEvent BlockEvent
	Finding    Finding
	BotInfo    BotInfo
}

const schema = `
	type Network struct {
		ChainId String
	}

	type BlockEvent struct {
		BlockHash String
		Network Network
	}

	type Finding struct {
		Protocol String
		Severity Int
		Type Int
		AlertId String
		Name String
		Description String
	}

	type BotInfo struct {
		BotImage String
		BotID String
	}

	type Alert struct {
		BlockEvent BlockEvent
		Finding Finding
		BotInfo BotInfo
	}
`

func main() {
	typeSystem, err := ipld.LoadSchemaBytes([]byte(schema))
	if err != nil {
		panic(err)
	}
	schemaType := typeSystem.TypeByName("Alert")

	node := bindnode.Wrap(&Alert{
		// TODO: Fill with anything
		Finding: Finding{
			AlertId: "SOME-PROTOCOL-ALERT-1",
		},
	}, schemaType)

	b, err := ipld.Encode(node, dagcbor.Encode)
	if err != nil {
		panic(err)
	}

	// Example output: 0x2389c487f81c6b9dab3d07ff3780c1cb9cc08fa418ba79551cd46083adb0df89
	fmt.Println(crypto.Keccak256Hash(b).Hex())
}
