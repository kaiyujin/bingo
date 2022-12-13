package request

type PostedGame struct {
	Names []string `json:"names" binding:"required,max=3"`
}
