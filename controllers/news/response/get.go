package response

import (
	admin_response "e-complaint-api/controllers/admin/response"
	category_response "e-complaint-api/controllers/category/response"
	file_response "e-complaint-api/controllers/news_file/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID         int                      `json:"id"`
	Admin      admin_response.Get       `json:"admin"`
	Category   category_response.Get    `json:"category"`
	Title      string                   `json:"title"`
	Content    string                   `json:"content"`
	TotalLikes int                      `json:"total_likes"`
	UpdatedAt  string                   `json:"update_at"`
	Files      []file_response.NewsFile `json:"files"`
}

func GetFromEntitiesToResponse(data *entities.News) *Get {
	var files []file_response.NewsFile
	for _, file := range data.Files {
		files = append(files, file_response.NewsFile{
			ID:     file.ID,
			NewsID: file.NewsID,
			Path:   file.Path,
		})
	}

	return &Get{
		ID:         data.ID,
		Admin:      *admin_response.GetFromEntitiesToResponse(&data.Admin),
		Category:   *category_response.GetFromEntitiesToResponse(&data.Category),
		Title:      data.Title,
		Content:    data.Content,
		TotalLikes: data.TotalLikes,
		UpdatedAt:  data.UpdatedAt.Format("2006-01-02 15:04:05"),
		Files:      files,
	}
}
