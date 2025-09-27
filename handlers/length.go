package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Value		float64
	From		string
	To			string
	Result		float64
	Converted 	bool
}

func convertLength (value float64, from, to string) float64 {
	unitMap := map[string]float64 {
		"mm": 0.001,
		"cm": 0.01,
		"m": 1,
		"km": 1000,
		"in": 0.0254,
		"ft": 0.3048,
		"yd": 0.9144,
		"mi": 1609.34,
	}

	// valueInMeter := value * unitMap[from]
	return value * unitMap[from] / unitMap[to]
}

func LengthHandler(templateWeb *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Print("test")

		data := Data{}
// 
		if req.Method == http.MethodPost {

			if err := req.ParseForm(); err != nil {
				http.Error(w, "form error", http.StatusBadRequest)
				return
			}

			// req.ParseForm()
			valueStr := req.FormValue("value")
			from := req.FormValue("from")
			to := req.FormValue("to")

			value, _ := strconv.ParseFloat(valueStr, 64)
			result := convertLength(value, from, to)

			data = Data{
				Value: value,
				From: from,
				To: to,
				Result: result,
				Converted: true,
			}
		// templateWeb.Execute(w, map[string]interface{} {
		// 	"Result": result,
		// })

		} 
		
		if err := templateWeb.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}