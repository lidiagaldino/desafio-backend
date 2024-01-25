package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/lidiagaldino/desafio-backend/cmd/server/config"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
	"github.com/lidiagaldino/desafio-backend/internal/infra/database"
	"github.com/lidiagaldino/desafio-backend/internal/infra/sns"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web"
	"github.com/spf13/viper"
)

func loadConfiguration() (config config.Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	return
}

func main() { 
	config, err := loadConfiguration()
	if err != nil {
		log.Fatal(err)
  }
	client, err := database.Connect(config.DB_URL)
	

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_KEY_ID,
			config.AWS_SECRET_ACCESS_KEY,
			"",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	sns := sns.NewSNS(sess)

	productRepository := database.NewProductRepository(client)
	categoryRepository := database.NewCategoryRepository(client)

	deleteProductUsecase := usecase.NewDeleteProductUsecase(productRepository)
	createProductUsecase := usecase.NewCreateProductUsecase(productRepository, sns, config.AWS_SNS_TOPIC_ARN, categoryRepository)
	updateProductUsecase := usecase.NewUpdateProductUsecase(productRepository, sns, config.AWS_SNS_TOPIC_ARN, categoryRepository)
	getProductUsecase := usecase.NewFindProductByIDUsecase(productRepository)
	getProductsUsecase := usecase.NewFindAllProductsUsecase(productRepository)

	productUsecases := usecase.NewProductUsecases(
		*getProductUsecase,
		*getProductsUsecase,
		*createProductUsecase,
		*updateProductUsecase,
		*deleteProductUsecase,
	)

	
	deleteCategoryUsecase := usecase.NewDeleteCategoryUsecase(categoryRepository)
	createCategoryUsecase := usecase.NewCreateCategoryUsecase(categoryRepository, sns, config.AWS_SNS_TOPIC_ARN)
	updateCategoryUsecase := usecase.NewUpdateCategoryUsecase(categoryRepository, sns, config.AWS_SNS_TOPIC_ARN)
	getCategoryUsecase := usecase.NewFindCategoryByIDUsecase(categoryRepository)
	getCategoriesUsecase := usecase.NewFindAllCategoriesUsecase(categoryRepository)
	categoryUsecases := usecase.NewCategoryUsecases(
    *getCategoryUsecase,
    *getCategoriesUsecase,
    *createCategoryUsecase,
    *updateCategoryUsecase,
		*deleteCategoryUsecase,
  )
	web.Initialize(productUsecases, categoryUsecases)
 }
