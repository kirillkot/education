package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

// Human struct for testing override methods
type Human struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

// UnmarshalJSON ovveride method for struct Human
func (h *Human) UnmarshalJSON(data []byte) error {
	type Alias Human
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return nil
}

// MarshalJSON override method for struct Human
func (h *Human) MarshalJSON() ([]byte, error) {
	type Alias Human
	return json.Marshal(&struct {
		*Alias
		ID int `json:"id"`
	}{
		Alias: (*Alias)(h),
		ID:    rand.Int(),
	})
}

func main() {
	Tom := &Human{
		Age:  20,
		ID:   1,
		Name: "Tom",
	}

	// Override MarshalJson with change params struct
	_ = json.NewEncoder(os.Stdout).Encode(&Tom)

	buf, err := json.Marshal(&Tom)
	if err != nil {
		panic(err)
	}

	var TomUnmarshal Human
	e := json.Unmarshal(buf, &TomUnmarshal)
	if e != nil {
		panic(e)
	}
}
