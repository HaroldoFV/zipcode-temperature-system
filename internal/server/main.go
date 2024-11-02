package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"zipcode-temperature-system/configs"
	"zipcode-temperature-system/docs"
	_ "zipcode-temperature-system/docs"
	"zipcode-temperature-system/internal/dto"
	"zipcode-temperature-system/internal/service"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Started")
	defer log.Println("Request ended")

	cep := chi.URLParam(r, "cep")

	if !isValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	dir, _ := os.Getwd()
	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Erro ao carregar configurações:", err)
			panic(err)
		}
	}

	tempService := service.NewTemperatureService(config.WeatherApiKey)

	city, err := tempService.GetCity(cep)
	if err != nil {
		http.Error(w, "cannot find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := tempService.GetTemperature(city)
	if err != nil {
		http.Error(w, "error getting temperature", http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := dto.TemperatureResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func isValidCEP(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}

// Middleware to enable CORS
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Altere para um domínio específico se necessário
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Defina os métodos permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Defina os cabeçalhos permitidos

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// @title Zipcode Temperature API
// @version 1.0
// @description API para obter temperatura com base no CEP usando Swagger.
// @BasePath /

// @param cep path string true "CEP para buscar a temperatura"
// @Success 200 {object} dto.TemperatureResponse
// @Failure 404 {object} map[string]string "can not find zipcode"
// @Failure 422 {object} map[string]string "invalid zipcode"
// @Router /temperature/{cep} [get]
func main() {
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080" // Define localhost como padrão
	}
	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.Schemes = []string{"https"}

	r := chi.NewRouter()
	r.Use(cors)

	r.Get("/temperature/{cep}", GetTemperature)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
