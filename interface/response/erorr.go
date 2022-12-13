package response

type Error struct {
	Code    string
	Message string
	Details []ErrorDetail
}

type ErrorDetail struct {
	Target  string
	Message string
}
