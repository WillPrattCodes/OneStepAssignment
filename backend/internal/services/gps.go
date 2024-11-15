package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//load env
func init() {
	godotenv.Load()
}

//was considering moving these structs to a separate file but it's only used here so I left it here

//lastest device point struct
type DevicePoint struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Speed float64 `json:"speed"`
}
//struct for each device in json response
type Device struct {
	DeviceID          string      `json:"device_id"`
	DisplayName       string      `json:"display_name"`
	ActiveState       string      `json:"active_state"`
	Online            bool        `json:"online"`
	LatestDevicePoint DevicePoint `json:"latest_device_point"`
}
//struct for overall json response
type GPSResponse struct {
	ResultList []Device `json:"result_list"`
}

//service func to fetch data from gps api
func FetchGPSData(db *sql.DB) ([]Device, error) {
	gpsKey := os.Getenv("GPS_KEY")

	//request url
	url := fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=%s", gpsKey)

	//get req
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GPS data: %w", err)
	}
	defer resp.Body.Close()

	//check resp status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	//read resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resp body: %w", err)
	}

	//parse json response into struct
	var gpsResponse GPSResponse
	if err := json.Unmarshal(body, &gpsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}
	//resturn parsed device list
	return gpsResponse.ResultList, nil
}
