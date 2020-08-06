package paych

import "github.com/filecoin-project/specs-actors/actors/builtin"

// Maximum number of lanes in a channel.
const LaneLimit = 256

const SettleDelay = builtin.EpochsInHour * 12

// Maximum size of a secret that can be submitted with a payment channel update (in bytes).
const MaxSecretSize = 256
