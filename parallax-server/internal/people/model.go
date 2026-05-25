package people

type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Occupation string `json:"occupation"`
	City       string `json:"city"`
}
