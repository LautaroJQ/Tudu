package dtos

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type EditTaskRequest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}

type TaskResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  string `json:"done"`
}
