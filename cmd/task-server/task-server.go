package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Task struct {
	ID          int64      `json:"id,omitempty"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      uint8      `json:"status"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CreateTaskResponse struct {
	ID     int64   `json:"id"`
	Status uint8   `json:"status"`
	Error  *string `json:"error,omitempty"`
}

var Config = struct {
	BindAddr   string
	ConfigPath string
}{}

func init() {
	flag.StringVar(&Config.ConfigPath, "config-path", "./configs", "path to the config file")
}

func InitConfig() {
	viper.AddConfigPath(Config.ConfigPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		viper.SetDefault("bindaddr", ":8080")
		log.Print(err.Error())
	}

	Config.BindAddr = viper.GetString("bindaddr")
}

func main() {

	flag.Parse()

	var buf []byte
	task_buf := Task{}
	task_response := CreateTaskResponse{}
	var id uuid.UUID

	InitConfig()

	http.HandleFunc("POST /task", func(w http.ResponseWriter, r *http.Request) {

		err := json.NewDecoder(r.Body).Decode(&task_buf)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			println(err.Error())
		}

		id = uuid.New()

		task_buf.ID = int64(id.ID())

		task_response.ID = task_buf.ID
		task_response.Status = task_buf.Status

		task_buf.CreatedAt = time.Now()

		buf, err = json.Marshal(task_response)
		if err != nil {
			log.Print("marshal error\n")
		}

		w.Write(buf)
	})

	http.HandleFunc("GET /task", func(w http.ResponseWriter, r *http.Request) {
		if task_buf.ID == 0 {
			w.Write([]byte("don't have any tasks"))
		} else {
			buf, err := json.Marshal(task_buf)
			if err != nil {
				log.Print("marshal error\n")
			}

			w.Write(buf)
		}
	})

	http.HandleFunc("PATCH /task", func(w http.ResponseWriter, r *http.Request) {

	})

	err := http.ListenAndServe(Config.BindAddr, nil)
	if err != nil {
		println("ListenAndServe error")
		os.Exit(1)
	}

}
