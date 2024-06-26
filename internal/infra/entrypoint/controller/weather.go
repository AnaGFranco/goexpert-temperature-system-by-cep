package controller

import (
	"github.com/gin-gonic/gin"
	"goexpert-temperature-system-by-cep/internal/application/service"
	"net/http"
)

type WeatherController struct {
	weatherService service.WeatherService
	zipCodeService service.ZipCodeService
}

func NewWeatherController(weatherService service.WeatherService, zipCodeService service.ZipCodeService) *WeatherController {
	return &WeatherController{
		weatherService: weatherService,
		zipCodeService: zipCodeService,
	}
}

func (h *WeatherController) GetWeather(c *gin.Context) {
	zipCode := c.Param("zipcode")

	if len(zipCode) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	location, err := h.zipCodeService.GetLocationByZipCode(zipCode)
	if location.Localidade == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	weather, err := h.weatherService.GetWeatherByLocation(location.Localidade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch weather data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"temp_C": weather.TempCelsius,
		"temp_F": weather.TempCelsius*1.8 + 32,
		"temp_K": weather.TempCelsius + 273.15,
	})
}
