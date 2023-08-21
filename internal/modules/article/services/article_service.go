package services

import (
	articleModel "blog/internal/modules/article/models2"
	ArticleRepository "blog/internal/modules/article/repositories"
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	"blog/internal/modules/user/responses"
	"errors"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticle() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticle() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article

	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user responses.User, folderPath string) (ArticleResponse.Article, error) {
	var article articleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID
	article.Image = folderPath

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("errors in creating the article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}
