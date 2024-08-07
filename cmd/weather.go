package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

type Timelines struct {
	Hourly []Hourly `json:"hourly"`
}

type Hourly struct {
	Time   string `json:"time"`
	Values Values `json:"values"`
}

type Values struct {
	CloudBase                float64 `json:"cloudBase"`
	CloudCeiling             float64 `json:"cloudCeiling"`
	CloudCover               int     `json:"cloudCover"`
	DewPoint                 float64 `json:"dewPoint"`
	Evapotranspiration       float64 `json:"evapotranspiration"`
	FreezingRainIntensity    float64 `json:"freezingRainIntensity"`
	Humidity                 float64 `json:"humidity"`
	IceAccumulation          float64 `json:"iceAccumulation"`
	IceAccumulationLwe       float64 `json:"iceAccumulationLwe"`
	PrecipitationProbability float64 `json:"precipitationProbability"`
	PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
	RainAccumulation         float64 `json:"rainAccumulation"`
	RainAccumulationLwe      float64 `json:"rainAccumulationLwe"`
	RainIntensity            float64 `json:"rainIntensity"`
	SleetAccumulation        float64 `json:"sleetAccumulation"`
	SleetAccumulationLwe     float64 `json:"sleetAccumulationLwe"`
	SleetIntensity           float64 `json:"sleetIntensity"`
	SnowAccumulation         float64 `json:"snowAccumulation"`
	SnowAccumulationLwe      float64 `json:"snowAccumulationLwe"`
	SnowIntensity            float64 `json:"snowIntensity"`
	Temperature              float64 `json:"temperature"`
	TemperatureApparent      float64 `json:"temperatureApparent"`
	UvHealthConcern          int     `json:"uvHealthConcern"`
	UvIndex                  int     `json:"uvIndex"`
	Visibility               float64 `json:"visibility"`
	WeatherCode              int     `json:"weatherCode"`
	WindDirection            float64 `json:"windDirection"`
	WindGust                 float64 `json:"windGust"`
	WindSpeed                float64 `json:"windSpeed"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}

type Response struct { // The response needs to be filled into this struct
	Timelines Timelines `json:"timelines"`
	Location  Location  `json:"location"`
}

var verbose bool
var WeatherAPIKey string = "uzxuk8FTWTE1M9mTctZ5kFFY2PV8oqca"

func GetWeatherData(location string) {
	var weather Response
	// Long, Lat, err := GetLongAndLatFromLocation(location)

	response, err := http.Get(fmt.Sprintf("https://api.tomorrow.io/v4/weather/forecast?location=%s&timesteps=1h&apikey=%s", url.QueryEscape(location), WeatherAPIKey))

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	if responseData, err := io.ReadAll(response.Body); err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(responseData, &weather)
	}
	fmt.Printf("The current temperature in %s is %.2f°C.\n", weather.Location.Name, weather.Timelines.Hourly.Temperature-273.15)
}

func init() {
	rootCmd.AddCommand(GetWeather)

	GetWeather.Flags().BoolVarP(&verbose, "verbose", "v", false, "Shows some extra informations")

}

var GetWeather = &cobra.Command{
	Use:   "weather",
	Short: "Displays the weather of a location",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		location := args[0]
		GetWeatherData(location)
	},
}
