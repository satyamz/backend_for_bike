package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//Ride : Struct to declare ride instance
type Ride struct {
	StartReading float64   `json:"start_reading"`
	EndReading   float64   `json:"end_reading"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}

//BillCalculate : Function to calculate bill
func BillCalculate(c *gin.Context) {
	var RideInstance Ride
	var TotalAmount float64
	var Extra float64
	var ExtraAmount float64

	c.Bind(&RideInstance)
	fmt.Println(RideInstance)
	Distance := RideInstance.EndReading - RideInstance.StartReading
	TimeTravelled := RideInstance.EndTime.Minute() - RideInstance.StartTime.Minute()
	/*
		loc, _ := time.LoadLocation("Asia/Kolkata")
		format := "Jan _2 2006 3:04:05 PM"
		timestamp := TimeTravelled
		t, err := time.ParseInLocation(format, timestamp, loc)
		fmt.Println(t, err)
	*/
	// Rs. 60/Hr | 12 Kms | Extra Rs4/Km.

	if TimeTravelled <= 60 {
		if Distance <= 10 {
			TotalAmount = 60
		} else {
			Extra = Distance - 10
			ExtraAmount = Extra * 4.0
			TotalAmount = 60 + ExtraAmount
		}
	} else {
		Extra = Distance - 10
		ExtraAmount = Extra * 4.0
		TotalAmount = 60 + ExtraAmount
	}

	/*
		if TimeTravelled.Minutes() < 31 && Distance <= 7 {
			TotalAmount = 29
		} else if (TimeTravelled.Minutes() < 60 && TimeTravelled.Minutes() > 31) && (Distance < 11) {
			TotalAmount = 50
		} else TimeTravelled.Minutes() > 60 {

		}
	*/
	c.JSON(200, gin.H{
		"Total Amount":               TotalAmount,
		"Start Time":                 RideInstance.StartTime.UTC(),
		"End Time":                   time.Now().UTC(),
		"Total Distance":             Distance,
		"Total Time(mins) travelled": TimeTravelled,
		"Extra Distance":             Extra,
		"Extra Amount":               ExtraAmount,
	})
	fmt.Println(Distance)
	fmt.Println(TimeTravelled)
	fmt.Println(TotalAmount)

}
func main() {
	r := gin.Default()
	r.POST("/bill", BillCalculate)
	r.Run(":9000")
}
