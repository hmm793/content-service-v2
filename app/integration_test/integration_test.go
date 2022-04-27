package main

import (
	"bytes"
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/delivery/asset_delivery"
	"content-service-v2/app/internal/delivery/content_delivery"
	"content-service-v2/app/internal/delivery/content_type_delivery"
	"content-service-v2/app/internal/repository/psql"
	"content-service-v2/app/internal/usecase/content_type_usecase"
	"content-service-v2/app/internal/usecase/content_usecase"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)


func init_app() http.Handler {
	fmt.Println("Content Service Test")
	
	// Environment Variables
	err := godotenv.Load(`C:\Users\DRAGON\Documents\Golang\go\src\content-service-v2-2\app\environments\app.env`)
	if err != nil {
		log.Fatal("Error loading app.env file")
	}

	// Seeder
	// err = seeders.ContentTypeDBSeed(db)
	// if err != nil {
	// 	log.Fatal("Content Type Seeder Failed")
	// }
	// Seeder

	// Koneksi Ke Database
	db, _, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	// defer Sql.Close()

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
	return router
}

func TestGetAllContentType(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader("")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/content_types", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	bodyResult, _ := io.ReadAll(response.Body)
	var responseBody interface{}
	json.Unmarshal(bodyResult, &responseBody)
	assert.NotEqual(t, 0, len(responseBody.([]interface{})))
}

func TestCreateAsset(t *testing.T) {
	router := init_app()
	path := `C:\Users\DRAGON\Documents\Golang\go\src\content-service-v2-2\app\integration_test\img_avatar.png`

    body := new(bytes.Buffer)
    
	writer := multipart.NewWriter(body)
	
	err := writer.WriteField("userId","123123131")
    
	part, err := writer.CreateFormFile("fileName", path)
    
	assert.NoError(t, err)
    sample, err := os.Open(path)
    assert.NoError(t, err)
	
    _, err = io.Copy(part, sample)
    assert.NoError(t, err)
    assert.NoError(t, writer.Close())

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/upload_images", body)

	request.Header.Add("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	
	bodyResult, _ := io.ReadAll(response.Body)
	
	var responseBody interface{}
	json.Unmarshal(bodyResult, &responseBody)
	
	assert.NotEqual(t, nil, (responseBody.(map[string]interface{}))["image"])
}

func TestCreateContent(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`
	{
		"data" : [
			{"image":"/api/v1/images/123123131/4a412dbd1b0d4cb8a1dd81f6b8f00d64-img_avatar-2022-04-22.png", "description":"this is a description 5"},
			{"image":"/api/v1/images/123123131/c97f663ce05b4cfd9919989be6643ba7-img_avatar-2022-04-22.png", "title":"this is title 2"}
		],
		"name" : "Banner Depan",
		"isPublished" : 1,
		"contentTypeId" : 1,
		"createdByName" : "john",
		"createdById" : 123123131
	}
	`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/create_content", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, float64(1), responseBody["contentTypeId"])
	assert.Equal(t, float64(123123131), responseBody["createdById"])
	assert.Equal(t, "john", responseBody["createdByName"])
	assert.NotEqual(t, nil, responseBody["data"])
	assert.NotEqual(t, nil, responseBody["id"])
	assert.Equal(t, float64(1), responseBody["isPublished"])
	assert.Equal(t, "Banner Depan", responseBody["name"])
}

func TestCreateContentWithEmptyField(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/create_content", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.NotEqual(t, nil, responseBody["message"])
}

func TestCreateContentWithWrongContentTypeId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{"data" : "[{'image' :'gambar 1', 'title':'content 1'},{'image' :'gambar 2', 'title':'content 2'} ]","isPublished" : true,"contentTypeId" : 100,"createdByName" : "john","createdById" : 283478293}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/contents", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Create Content Failed", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 400, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetContentWithContentTypeId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents/byContentTypeId/1", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get Content By Content Type Id Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetContentWithWrongContentTypeId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents/byContentTypeId/1000", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get Content By Content Type Id Failed", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 400, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"])
}

func TestGetContentWithId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents/byId/1", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get Content By Id Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetContentWithWrongId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents/byId/150", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get Content By Id Failed", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 400, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"])
}

func TestUpdateContent(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{"data" : "[{'image' :'gambar 1', 'title':'content 1'},{'image' :'gambar 2', 'title':'content 2'} ]","is_published" : true,"content_type_id" : 1,"updated_by_name" : "john","updated_by" : 283478293}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/contents/5", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println("balas",response)

	assert.Equal(t, "Update Content Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetAllContent(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get All Content Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetAllContentWithLimit(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents?limit=5", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get All Content Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestGetAllContentWithLimitAndOffset(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/contents?limit=5&offset=1", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Get All Content Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestDeleteContent(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/contents/byId/1", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Delete Content Success", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 200, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"])
}
func TestDeleteContentWithWrongId(t *testing.T){
	router := init_app()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/contents/byId/150", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, "Delete Content Failed", responseBody["meta"].(map[string]interface{})["message"])
	assert.Equal(t, 400, int((responseBody["meta"].(map[string]interface{})["code"]).(float64)))
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"])
}