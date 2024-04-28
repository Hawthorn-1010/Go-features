package shallowanddeepcopy

type Per struct {
	Name    string
	Age     int
	HouseId [3]int
}

type Per1 struct {
	Name    string
	Age     int
	HouseId [3]int
	Tags    map[string]string
}

type Man struct {
	Name    string            `json:"name"`
	HouseId []int             `json:"houseId"`
	Tags    map[string]string `json:"tags"`
}
