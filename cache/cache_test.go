package cache

import (
	"testing"
)

func TestCache_Encode_Decode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"
	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}

	_, err = decode(string(bytes))
	if err != nil {
		t.Error(err)
	}

}
