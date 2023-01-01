package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvestCalculation struct {
	Jenis_Kelamin  string  `json:"jenis_kelamin"`
	Usia           int8    `json:"usia"`
	Perokok        string  `json:"perokok"`
	Nominal        float64 `json:"nominal"`
	Lama_Investasi float64 `json:"lama_investasi"`
}

type Total struct {
	Awal  float64 `json:"awal"`
	Bunga float64 `json:"bunga"`
	Akhir float64 `json:"akhir"`
}

func Controller(e echo.Context) error {
	cal := &InvestCalculation{}

	if err := e.Bind(cal); err != nil {
		return err
	}

	var data = map[int]*Total{}

	for i := 1; i <= int(cal.Lama_Investasi); i++ {
		awal := cal.Nominal
		bunga := Perokok(cal) - awal
		akhir := Usia(cal)

		a := akhir - bunga
		b := bunga
		c := awal + bunga

		data[i] = &Total{
			Awal:  a,
			Bunga: b,
			Akhir: c,
		}

	}

	return e.JSON(SuccessResponseWithData(data))
}

func Usia(cal *InvestCalculation) float64 {

	if cal.Usia <= 30 {

		cal.Nominal = cal.Nominal + (0.01 * cal.Nominal)

		return cal.Nominal

	} else if cal.Usia <= 50 {

		cal.Nominal = cal.Nominal + (0.005 * cal.Nominal)

		return cal.Nominal

	} else if cal.Usia > 50 {

		cal.Nominal = cal.Nominal + (0 * cal.Nominal)

		return cal.Nominal

	}

	return 1
}

func Perokok(cal *InvestCalculation) float64 {

	if cal.Perokok == "Ya" && cal.Jenis_Kelamin == "Pria" {

		cal.Nominal = cal.Nominal + (0.01 * cal.Nominal)

		return cal.Nominal

	} else if cal.Perokok == "Bukan" && cal.Jenis_Kelamin == "Pria" {

		cal.Nominal = cal.Nominal + (0.02 * cal.Nominal)

		return cal.Nominal

	} else if cal.Perokok == "Ya" && cal.Jenis_Kelamin == "Wanita" {

		cal.Nominal = cal.Nominal + (0.02 * cal.Nominal)

		return cal.Nominal

	} else if cal.Perokok == "Bukan" && cal.Jenis_Kelamin == "Wanita" {

		cal.Nominal = cal.Nominal + (0.03 * cal.Nominal)

		return cal.Nominal

	}

	return 1
}

var Success string = "200"

type SuccessResponseSpec struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseWithData(data interface{}) (int, *SuccessResponseSpec) {
	return http.StatusOK, &SuccessResponseSpec{
		Code:    Success,
		Message: "success",
		Data:    data,
	}
}

func main() {
	e := echo.New()

	e.POST("", Controller)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
