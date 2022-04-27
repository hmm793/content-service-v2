package main

import (
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/delivery/asset_delivery"
	"content-service-v2/app/internal/delivery/content_delivery"
	"content-service-v2/app/internal/delivery/content_type_delivery"
	"content-service-v2/app/internal/repository/psql"
	"content-service-v2/app/internal/usecase/content_type_usecase"
	"content-service-v2/app/internal/usecase/content_usecase"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Content Service")
	
	// Environment Variables
	err := godotenv.Load("app/environments/app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}

	// Koneksi Ke Database
	db, Sql, err := config.ConnectPostgres()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	// Seeder
	// err = seeders.ContentTypeDBSeed(db)
	// if err != nil {
	// 	log.Fatal("Content Type Seeder Failed")
	// }
	// Seeder

	// Content Type
	contentTypeRepository := psql.NewRepositoryContentType(db)
	contentTypeService := content_type_usecase.NewServiceContentType(contentTypeRepository)
	contentTypeHandler := content_type_delivery.NewContentTypeHandler(contentTypeService)

	// Content
	contentRepository := psql.NewRepositoryContent(db)
	contentService := content_usecase.NewServiceContent(contentRepository)
	contentHandler := content_delivery.NewContentHandler(contentService)

	// Asset
	assetHandler := asset_delivery.NewAssetHandler()

	// Router
	router := gin.Default()
	router.Static("/api/v1/images", os.Getenv("tempPath"))
	router.Static("/api/v1/asset", os.Getenv("assetPath"))
	api := router.Group("/api/v1")
	api.GET("/content_types", contentTypeHandler.GetContentTypes)
	api.POST("/create_content", contentHandler.CreateContent)
	api.PUT("/update_content/:id", contentHandler.UpdateContent)               
	api.DELETE("/delete_content/:id", contentHandler.DeleteContent)
	api.GET("/get_content/:id", contentHandler.FindById)
	api.GET("/contents", contentHandler.FindAllContent)
	api.POST("/upload_images", assetHandler.UploadAsset)
	router.Run(":"+os.Getenv("PORT_DEV"))
}
