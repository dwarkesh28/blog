package repositories

import ArticalModels "blog/internal/modules/article/models2"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticalModels.Article
	Find(id int) ArticalModels.Article
	Create(article ArticalModels.Article) ArticalModels.Article
}

