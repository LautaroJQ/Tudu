package dtos

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type EditTaskRequest struct {
	Title string `json:"title"`
}

type TaskResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  string `json:"done"`
}
