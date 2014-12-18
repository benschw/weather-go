package weather

import (
	"github.com/benschw/weather-go/openweather"
	"github.com/benschw/weather-go/weather/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var _ = log.Printf

type WeatherClient interface {
	FindForLocation(city string, state string) (openweather.Conditions, error)
}

type WeatherService struct {
	Bind          string
	Db            gorm.DB
	WeatherClient WeatherClient
}

func NewWeatherService(bind string, dbStr string) (*WeatherService, error) {
	s := &WeatherService{}

	db, err := DbOpen(dbStr)
	if err != nil {
		return s, err
	}

	s.Db = db
	s.Bind = bind
	s.WeatherClient = &openweather.WeatherClient{}

	return s, nil
}

func (s *WeatherService) MigrateDb() error {

	s.Db.AutoMigrate(api.Location{})
	return nil
}

func (s *WeatherService) Run() error {

	// route handlers
	resource := &LocationResource{
		Db:            s.Db,
		WeatherClient: s.WeatherClient,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/location", resource.Add).Methods("POST")
	r.HandleFunc("/location", resource.FindAll).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Find).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Save).Methods("PUT")
	r.HandleFunc("/location/{id}", resource.Delete).Methods("DELETE")

	http.Handle("/", r)

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}

func DbOpen(dbStr string) (gorm.DB, error) {
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		return db, err
	}
	db.SingularTable(true)
	return db, nil
}
