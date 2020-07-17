package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ridhoperdana/articletest"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionNameArticle = "article"

type mongoRepository struct {
	DB   *mongo.Database
	coll *mongo.Collection
}

func (m mongoRepository) Createarticles(article articletest.Article) (articletest.Article, error) {
	id := uuid.NewV4()
	article.Id = id.String()
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	_, err := m.coll.InsertOne(context.TODO(), article)
	if err != nil {
		return articletest.Article{}, err
	}
	return article, nil
}

func (m mongoRepository) Deletearticlebyid(articleId string) error {
	result, err := m.coll.DeleteOne(context.TODO(), bson.M{
		"id": articleId,
	})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return articletest.ErrorNotFound{
			Message: "article " + articleId + "not found",
		}
	}
	return nil
}

func (m mongoRepository) Listarticles(num int32, cursor string) (articletest.Articles, string, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(num)).SetSort(bson.M{
		"created_at": -1,
	})
	filter := bson.M{}
	if cursor != "" {
		decoded, err := articletest.Decode(cursor)
		if err != nil {
			return nil, "", err
		}

		timeCursor, err := time.Parse("2006-01-02 15:04:05 +0000", decoded[:len(decoded)-4])
		if err != nil {
			return nil, "", err
		}

		filter["created_at"] = bson.M{
			"$lt": timeCursor,
		}
	}

	var results []articletest.Article
	cur, err := m.coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, "", err
	}

	for cur.Next(context.TODO()) {
		var article articletest.Article
		if err := cur.Decode(&article); err != nil {
			logrus.Error(err)
			continue
		}
		results = append(results, article)
	}

	if err := cur.Err(); err != nil {
		return nil, "", err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, "", err
	}

	if len(results) == 0 {
		return make(articletest.Articles, 0), "", nil
	}

	nextCursor := articletest.Encode(results[len(results)-1].CreatedAt.String())

	return results, nextCursor, nil
}

func (m mongoRepository) Showarticlebyid(articleId string) (articletest.Article, error) {
	res := m.coll.FindOne(context.TODO(), bson.M{
		"id": articleId,
	})

	var result articletest.Article
	if err := res.Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return articletest.Article{}, articletest.ErrorNotFound{
				Message: "article " + articleId + " not found",
			}
		}
		return articletest.Article{}, err
	}

	return result, nil
}

func (m mongoRepository) Updatearticle(articleId string, article articletest.Article) (articletest.Article, error) {
	article.UpdatedAt = time.Now()
	res, err := m.coll.UpdateOne(context.TODO(), bson.M{
		"id": articleId,
	}, bson.M{
		"$set": bson.M{
			"title":      article.Title,
			"updated_at": article.UpdatedAt,
			"author":     article.Author,
			"publisher":  article.Publisher,
			"tag":        article.Tag,
		},
	})
	if err != nil {
		return articletest.Article{}, err
	}

	if res.ModifiedCount == 0 {
		return articletest.Article{}, articletest.ErrorNotFound{
			Message: "article " + articleId + " not found",
		}
	}

	return article, nil
}

func NewMongoRepository(db *mongo.Database) articletest.ArticleRepository {
	return &mongoRepository{
		DB:   db,
		coll: db.Collection(CollectionNameArticle),
	}
}
