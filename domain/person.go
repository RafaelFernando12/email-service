package domain

type Person struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}
