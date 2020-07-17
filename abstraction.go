package articletest

type ArticleService interface {
	Createarticles(article Article) (Article, error)
	Deletearticlebyid(articleId string) error
	Listarticles(num int32, cursor string) (Articles, string, error)
	Showarticlebyid(articleId string) (Article, error)
	Updatearticle(articleId string, article Article) (Article, error)
}

type ArticleRepository interface {
	Createarticles(article Article) (Article, error)
	Deletearticlebyid(articleId string) error
	Listarticles(num int32, cursor string) (Articles, string, error)
	Showarticlebyid(articleId string) (Article, error)
	Updatearticle(articleId string, article Article) (Article, error)
}
