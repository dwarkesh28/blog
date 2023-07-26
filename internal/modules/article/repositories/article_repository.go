package repositories

import (
	ArticalModels "blog/internal/modules/article/models2"
	articleModel "blog/internal/modules/article/models2"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (articleRepository *ArticleRepository) Find(id int) articleModel.Article {
	var article articleModel.Article
	articleRepository.DB.Joins("User").First(&article, id)
	return article
}

func (articleRepository *ArticleRepository) Create(article ArticalModels.Article) ArticalModels.Article {
	var newArticle articleModel.Article
	articleRepository.DB.Create(&article).Scan(&newArticle)
	return newArticle
}
