package models

import (
	"errors"
)

type Feedback struct {
	ID               int    `gorm:"primary_key;auto_increment" json:"id"`
	Booking_code     string `gorm:"type:varchar(45);not null;unique" json:"booking_code"`
	Feedback_content string `gorm:"type:text;not null" json:"content"`
	Passenger_id     int    `gorm:"type:int;not null" json: "passenger_id"`
}

func AddPassengerFeedback(feedback Feedback) error {
	db := Connect()
	defer db.Close()

	if len(feedback.Booking_code) == 0 {
		err := errors.New("Feedback requires BookingCode")
		return err
	}

	if feedback.Passenger_id == 0 {
		err := errors.New("Feedback requires Passenger ID")
		return err
	}
	err := db.Create(&feedback).Error
	return err
}

func GetFeedbacksByPassengerID(id int) (feedbacks []Feedback, err error) {
	db := Connect()
	defer db.Close()

	result := []Feedback{}
	err = db.Where(&Feedback{Passenger_id: id}).Find(&result).Error

	return result, err
}

func GetFeedbackByBookingCode(code string) (feedback Feedback, err error) {
	db := Connect()
	defer db.Close()

	result := Feedback{Booking_code: code}
	err = db.Where(&result).First(&result).Error

	return result, err
}

func DeleteFeedbacksByPassengerID(id int) error {
	db := Connect()
	defer db.Close()

	err := db.Where("passenger_id = ?", id).Delete(Feedback{}).Error
	return err
}
