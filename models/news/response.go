package news

import "time"

type NewsResponse struct {
	ID        uint  		`json:"id"`	
	CreatedAt time.Time  	`json:"createdAt"`		
	UpdatedAt time.Time 	`json:"updatedAt"`
	Title string 			`json:"title"`
	Author string 			`json:"author"`
	Content string 			`json:"content"`
}

func (newsResponse *NewsResponse) MapFromDatabase(news News) {
	newsResponse.ID = news.ID
	newsResponse.CreatedAt = news.CreatedAt
	newsResponse.UpdatedAt = news.UpdatedAt
	newsResponse.Title = news.Title
	newsResponse.Author = news.Author
	newsResponse.Content = news.Content
}


func MapFromDatabaseList(listNewsDB []News) []NewsResponse {
	var result []NewsResponse
	var newsResponse NewsResponse
	for _, news := range listNewsDB {
		newsResponse.ID = news.ID
		newsResponse.CreatedAt = news.CreatedAt
		newsResponse.UpdatedAt = news.UpdatedAt
		newsResponse.Title = news.Title
		newsResponse.Author = news.Author
		newsResponse.Content = news.Content
		result = append(result, newsResponse)
	}
	return result
}