package types


type MinConsensusInfoEvent struct {
	Epoch            uint64 	 `json:"epoch"`
	ValidatorList    []string    `json:"validatorList"`
	EpochStartTime   uint64      `json:"epochTimeStart"`
	SlotTimeDuration uint64      `json:"slotTimeDuration"`
}
