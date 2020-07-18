package cmd

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ridhoperdana/articletest"
	"github.com/ridhoperdana/articletest/internal/delivery"
	"github.com/ridhoperdana/articletest/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	repo articletest.ArticleRepository

	RootCMD = &cobra.Command{
		Use: "articletest",
	}

	port = ""

	httpServerCMD = &cobra.Command{
		Use:   "http",
		Short: "Run HTTP Server",
		Long:  "Run HTTP Server for articletest",
		Run: func(cmd *cobra.Command, args []string) {
			e := echo.New()
			service := articletest.NewArticleService(repo)
			delivery.RegisterHTTPPath(e, service)
			port = os.Getenv("PORT")
			if port == "" {
				port = "6969"
			}
			logrus.Infof("starting HTTP server with port %v", port)
			if err := e.Start(":" + port); err != nil {
				logrus.Fatal(err)
			}
		},
	}
)

func initApp() {
	mongoConnectTimeout, err := strconv.ParseInt(os.Getenv("MONGO_CONNECT_TIMEOUT"), 10, 64)
	if err != nil {
		logrus.Fatal("MONGO_CONNECT_TIMEOUT is not well-set", err)
	}

	mongoSelectionTimeout, err := strconv.ParseInt(os.Getenv("MONGO_SERVER_SELECTION_TIMEOUT"), 10, 64)
	if err != nil {
		logrus.Fatal("MONGO_SERVER_SELECTION_TIMEOUT is not well-set", err)
	}

	mongoClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGO_URI")).
			SetConnectTimeout(time.Duration(mongoConnectTimeout)*time.Second).
			SetServerSelectionTimeout(time.Duration(mongoSelectionTimeout)*time.Second),
	)
	if err != nil {
		logrus.Fatal("Mongo connection failed: ", err.Error())
	}

	if mongoClient == nil {
		logrus.Fatal("Mongo client nil")
	}

	repo = repository.NewMongoRepository(mongoClient.Database("articletest"))
}

func init() {
	cobra.OnInitialize(initApp)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	RootCMD.AddCommand(httpServerCMD)
	httpServerCMD.Flags().StringVarP(&port, "port", "p", "",
		"Port number for the App. Default is 8080.")
}
