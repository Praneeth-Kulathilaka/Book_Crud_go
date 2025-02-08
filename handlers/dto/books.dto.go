package dto

type BookDTO struct {
	Author []string `json:"author_name"`
	Title string `json:"title"`
}

// type ByIdDTO struct {
// 	Title string `json:"title"`
// 	Author []string `json:"authors"`
// 	Pages string `json:"number_of_pages"`
// }