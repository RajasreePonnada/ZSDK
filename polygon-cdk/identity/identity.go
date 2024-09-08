// identity.go
package identity

import (
    "encoding/json"
    "github.com/ethereum/go-ethereum/common"
    "github.com/yourusername/identity-management/polygon-cdk/state"
)

type Identity struct {
    Address  string            `json:"address"`
    DataHash string            `json:"dataHash"`
    Metadata map[string]string `json:"metadata"`
}

func CreateIdentity(s *state.CDKState, identity Identity) error {
    data, err := json.Marshal(identity)
    if err != nil {
        return err
    }
    s.Set(common.HexToHash(identity.Address), data)
    return nil
}

func GetIdentity(s *state.CDKState, address string) (Identity, error) {
    data := s.Get(common.HexToHash(address))
    var identity Identity
    err := json.Unmarshal(data, &identity)
    return identity, err
}

// Add UpdateIdentity and DeleteIdentity functions as needed