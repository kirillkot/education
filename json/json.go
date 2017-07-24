package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Human struct {
	id          int32
	age         int32
	name        string
	phoneNumber int32
}

func (h *Human) UnmarshalJSON(b []byte) error {
	fmt.Println("I am here, in UnmarshalJson")
	return nil
}

func (h *Human) MarshalJSON() ([]byte, error) {
	fmt.Println("method, MarshalJSON")
	return nil, nil
}

func parseJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ParseJson")
	var man Human
	err := json.NewDecoder(r.Body).Decode(&man)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println(man)
}

func main() {
	TestCases := []struct {
		Name       string
		Body       string
		CErr       error
		StatusCode int
	}{
		{Name: "invalid_body(immutable field)", Body: `[{ids: [1], topic: "topic"}]`, StatusCode: 400},
	}

	for _, testCase := range TestCases {
		r, _ := http.NewRequest("PUT", "http://fake-url/", strings.NewReader(testCase.Body))
		w := httptest.NewRecorder()

		parseJson(w, r)
	}
}
