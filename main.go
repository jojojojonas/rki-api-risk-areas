package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiData struct {
	FieldName string `json:"objectIdFieldName"`
	UniqueIdField UniqueIdField `json:"uniqueIdField"`
	GlobalIdFieldName string `json:"globalIdFieldName"`
	ServerGens serverGens `json:"serverGens"`
	GeometryType string `json:"geometryType"`
	SpatialReference SpatialReference `json:"spatialReference"`
	Fields []Fields `json:"fields"`
	Features []Features `json:"features"`
}

type UniqueIdField struct {
	Name string `json:"name"`
	IsSystemMaintained bool `json:"isSystemMaintained"`
}

type serverGens struct {
	MinServerGen int `json:"minServerGen"`
	ServerGen int `json:"serverGen"`
}

type SpatialReference struct {
	Wkid int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}

type Fields struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Alias string `json:"alias"`
	SqlType string `json:"sqlType"`
	Length int `json:"length"`
	Domain string `json:"domain"`
	DefaultValue string `json:"defaultValue"`
}

type Features struct {
	Arributes Attributes `json:"attributes"`
}

type Attributes struct {
	State string `json:"LAN_ew_GEN"`
	Cases float64 `json:"cases7_bl_per_100k"`
}

type ReturnData struct {
	State string
	Cases float64
}

func main() {

	// Set gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Get request
	router.GET("/rki/corona", func (c *gin.Context) {

		// Get data from URL
		res, err := http.Get("https://services7.arcgis.com/mOBPykOjAyBO2ZKk/arcgis/rest/services/Coronaf%C3%A4lle_in_den_Bundesl%C3%A4ndern/FeatureServer/0/query?where=1%3D1&outFields=LAN_ew_GEN,cases7_bl_per_100k&returnGeometry=false&outSR=&f=json")
		if err != nil {
			fmt.Println("The following error occurred while accessing the Url: ", err)
		}

		// Create variable for decode data
		var data ApiData

		// Decode data
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			fmt.Println("Error beim Encoden: ", err)
		}

		// Create variable for return
		var returnData []ReturnData

		for _, value := range data.Features {
			if value.Arributes.Cases >= 50 {
				returnData = append(returnData, ReturnData{value.Arributes.State, value.Arributes.Cases})
			}
		}

		// Return JSON
		c.JSON(200, returnData)

	})

	router.Run(":80")

}