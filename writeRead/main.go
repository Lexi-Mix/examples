package writeRead

import (
	"encoding/json"
	"fmt"
	"io"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	read bool
}

func (h *Human) String() string {
	return fmt.Sprintf("%s %d", h.Name, h.Age)
}

func (h *Human) Read(p []byte) (n int, err error) {
	if h.read {
		return 0, io.EOF
	}
	data := h.String()
	n = copy(p, data)
	h.read = true
	return n, nil
}

func (h *Human) Write(p []byte) (n int, err error) {
	h.read = false
	input := string(p)
	if _, err = fmt.Sscanf(input, "%s %d", &h.Name, &h.Age); err != nil {
		return 0, err
	}

	return len(p), nil
}

func (h *Human) PrettyJSON() error {
	indent, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(indent))

	return nil
}
