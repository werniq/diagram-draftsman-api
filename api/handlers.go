package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/werniq/diagram-creating-api/Logger"
	"net/http"
)

type URIData struct {
	DiagramType string   `json:"type"`
	Options     *Options `json:"options"`
	Data        []*Data  `json:"data"`
}

type Data struct {
	Labels   []any     `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	BorderColor string `json:"borderColor"`
	BorderWidth string `json:"borderWidth"`
	Fill        bool   `json:"fill"`
	Data        []any  `json:"data"`
}

type Options struct {
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	BackgroundColor string `json:"backgroundColor"`
	Format          string `json:"format"`
}

// CreateDiagramWithRawData used for creating diagram with data, inputted manually by user
func CreateDiagramWithRawData(c *gin.Context) {
	uri, err := createDiagramUrl(c)
	if err != nil {
		Logger.Logger().Println("Error creating diagram URI: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uri": uri,
	})
}

func createDiagramUrl(c *gin.Context) (string, error) {
	var Data *URIData
	err := c.BindJSON(&Data)
	if err != nil {
		Logger.Logger().Println("error binding json: ", err)
		return "", err
	}

	uri := "https://quickchart.io/chart?chart={type:"
	uri += "'" + Data.DiagramType + "',"
	uri += "data:{labels:["

	for i := 0; i < len(Data.Data); i++ {
		v := Data.Data[i]
		uri += fmt.Sprintf("%d,", v.Labels[i])
	}
	uri += "],datasets:[{"

	for i, v := range Data.Data {
		uri += "'" + fmt.Sprintf("%v", v.Datasets[i].Label) + "',data:["
		for _, d := range v.Datasets[i].Data {
			uri += fmt.Sprintf("%v,", d)
		}
		uri += "]}"
		if Data.Data[i+1] != nil {
			uri += ","
		}
	}

	if Data.Options.Height != 0 && Data.Options.Width != 0 {
		uri += "?height=" + fmt.Sprintf("%d", Data.Options.Height)
		uri += "?width=" + fmt.Sprintf("%d", Data.Options.Width)
	}

	if Data.Options.BackgroundColor != "" {
		uri += "?backgroundColor=" + Data.Options.BackgroundColor
	}

	if Data.Options.Format != "" {
		uri += "?format=" + Data.Options.Format
	}

	return uri, nil
}

//func CreateDiagramWithDataFromFile(c *gin.Context) {
//
//}

// readFileData will be used to generate diagrams using file data
//func readFileData(c *gin.Context) (string, error) {
//	file, header, err := c.Request.FormFile("file")
//	if err != nil {
//		Logger.Logger().Println("error parsing file from request: ", err)
//		return "", err
//	}
//	var s
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//
//	}
//}
