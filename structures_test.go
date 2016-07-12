package EthereumAPI_test

import (
	"encoding/json"
	"testing"

	. "github.com/FactomProject/EthereumAPI"
)

func TestQuantityJSON(t *testing.T) {
	for i := 0; i < 100; i++ {
		q1 := NewQuantityFromInt(int64(i))
		j, err := json.Marshal(q1)
		if err != nil {
			t.Errorf("%v", err)
		}
		q2 := new(Quantity)
		err = json.Unmarshal(j, q2)
		if err != nil {
			t.Errorf("%v", err)
		}
		if int64(*q1) != int64(*q2) {
			t.Errorf("q1!=q2 - %v vs %v", q1, q2)
		}
	}
}
