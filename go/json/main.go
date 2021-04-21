package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"strconv"
	"time"
)

type MinimalEpochConsensusInfo struct {
	Epoch            uint64        `json:"epoch"`
	ValidatorList    [32]string    `json:"validatorList"`
	EpochStartTime   uint64        `json:"epochTimeStart"`
	SlotTimeDuration time.Duration `json:"slotTimeDuration"`
}

func newMinimalConsensusInfo(epoch uint64) *MinimalEpochConsensusInfo {
	validatorList := make([]string, 32)

	for idx := 0; idx < 32; idx++ {
		bs := []byte(strconv.Itoa(31415926))
		pubKey := common.Bytes2Hex(bs)
		validatorList[idx] = pubKey
	}

	var validatorList32 [32]string
	copy(validatorList32[:], validatorList)
	return &MinimalEpochConsensusInfo{
		Epoch:            epoch,
		ValidatorList:    validatorList32,
		EpochStartTime:   rand.Uint64(),
		SlotTimeDuration: time.Duration(6),
	}
}

func encode(v interface{}) []byte {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(v)
	return buffer.Bytes()
}

func decode(data []byte, v interface{}) {
	var buf bytes.Buffer
	buf.Write(data)
	json.NewDecoder(&buf).Decode(v)
}

func main() {
	consensusInfo0 := newMinimalConsensusInfo(1)
	consensusInfoEncoded0 := encode(consensusInfo0)
	fmt.Printf("consensusInfo0: %v\n len: %d\n", consensusInfoEncoded0, len(consensusInfoEncoded0))

	var consensusInfoDecoded0 *MinimalEpochConsensusInfo
	decode(consensusInfoEncoded0, &consensusInfoDecoded0)
	fmt.Printf("consensusInfoDecoded: %v", consensusInfoDecoded0)
}
