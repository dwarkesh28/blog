package services

import (
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	"blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticle() ArticleResponse.Articles
	GetStoriesArticle() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user responses.User) (ArticleResponse.Article, error)
}
