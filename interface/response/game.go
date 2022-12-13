package response

type Game struct {
	Id            string
	CalledNumbers []int8
	Users         []User
}

type User struct {
	Id      string
	Name    string
	Numbers []int8
}
