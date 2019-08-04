package main

import (
	"grabvn-golang-bootcamp/protoc"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var client protoc.FeedbackServiceClient

func addFeedback(c *gin.Context) {
	var argument struct {
		Booking_code     string
		Passenger_id     int32
		Feedback_content string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	feedback := &protoc.FeedbackRequest{
		BookingCode:     argument.Booking_code,
		PassengerId:     argument.Passenger_id,
		FeedbackContent: argument.Feedback_content,
	}

	req := &protoc.CreateFeedbackRequest{Feedback: feedback}
	response, err := client.CreateFeedback(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, response.Msg)
}

func getPassengerFeedbacks(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	req := &protoc.GetFeedbacksByPassengerIdRequest{PassengerId: int32(id)}

	response, err := client.GetFeedbacksByPassengerId(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, response)

}

func getFeedback(c *gin.Context) {
	code := c.Params.ByName("code")
	if len(code) == 0 {
		c.String(400, "invalid param")
		return
	}

	req := &protoc.GetFeedbackByBookingCodeRequest{BookingCode: code}

	response, err := client.GetFeedbackByBookingCode(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, response)
}

func deletePassengerFeedbacks(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	req := &protoc.DeleteFeedbacksByPassengerIdRequest{PassengerId: int32(id)}

	response, err := client.DeleteFeedbacksByPassengerId(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, response)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = protoc.NewFeedbackServiceClient(conn)

	router := gin.Default()

	v1 := router.Group("api/feedback")
	{
		v1.POST("/", addFeedback)
		v1.GET("/passenger/:id", getPassengerFeedbacks)
		v1.GET("/bookingcode/:code", getFeedback)
		v1.DELETE("/:id", deletePassengerFeedbacks)
	}

	router.Run()
}
