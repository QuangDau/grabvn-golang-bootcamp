package main

import (
	"grabvn-golang-bootcamp/models"
	"grabvn-golang-bootcamp/protoc"
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	models.AutoMigrations()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	protoc.RegisterFeedbackServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}

}

func (s *server) CreateFeedback(ctx context.Context, request *protoc.CreateFeedbackRequest) (*protoc.CreateFeedbackReply, error) {
	feedback := models.Feedback{
		Booking_code:     request.GetFeedback().GetBookingCode(),
		Passenger_id:     int(request.GetFeedback().GetPassengerId()),
		Feedback_content: request.GetFeedback().GetFeedbackContent(),
	}

	err := models.AddPassengerFeedback(feedback)

	if err != nil {
		return &protoc.CreateFeedbackReply{}, err
	}

	return &protoc.CreateFeedbackReply{Msg: "Successfully created feedback"}, nil
}

func (s *server) GetFeedbacksByPassengerId(ctx context.Context, request *protoc.GetFeedbacksByPassengerIdRequest) (*protoc.GetFeedbacksByPassengerIdReply, error) {
	id := int(request.GetPassengerId())

	feedbacks, err := models.GetFeedbacksByPassengerID(id)
	if err != nil {
		return &protoc.GetFeedbacksByPassengerIdReply{}, err
	}

	replyCollection := []*protoc.FeedbackResponse{}

	for _, feedback := range feedbacks {
		reply := &protoc.FeedbackResponse{
			Id:              int32(feedback.ID),
			BookingCode:     feedback.Booking_code,
			PassengerId:     int32(feedback.Passenger_id),
			FeedbackContent: feedback.Feedback_content,
		}
		replyCollection = append(replyCollection, reply)
	}

	return &protoc.GetFeedbacksByPassengerIdReply{Feedbacks: replyCollection}, nil
}

func (s *server) GetFeedbackByBookingCode(ctx context.Context, request *protoc.GetFeedbackByBookingCodeRequest) (*protoc.GetFeedbacksByBookingCodeReply, error) {
	bookingCode := request.GetBookingCode()

	feedback, err := models.GetFeedbackByBookingCode(bookingCode)
	if err != nil {
		return &protoc.GetFeedbacksByBookingCodeReply{}, err
	}

	reply := &protoc.FeedbackResponse{
		Id:              int32(feedback.ID),
		BookingCode:     feedback.Booking_code,
		PassengerId:     int32(feedback.Passenger_id),
		FeedbackContent: feedback.Feedback_content,
	}

	return &protoc.GetFeedbacksByBookingCodeReply{Feedback: reply}, nil
}

func (s *server) DeleteFeedbacksByPassengerId(ctx context.Context, request *protoc.DeleteFeedbacksByPassengerIdRequest) (*protoc.DeleteFeedbacksByPassengerIdReply, error) {
	id := int(request.GetPassengerId())

	err := models.DeleteFeedbacksByPassengerID(id)
	if err != nil {
		return &protoc.DeleteFeedbacksByPassengerIdReply{}, err
	}

	return &protoc.DeleteFeedbacksByPassengerIdReply{Msg: "successfully deleted feedbacks"}, nil
}
