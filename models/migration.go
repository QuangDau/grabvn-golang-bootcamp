package models

func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&Feedback{})
	db.Debug().AutoMigrate(&Feedback{})

	//Sample data
	feedback1 := Feedback{ID: 1, Booking_code: "KJAG12", Feedback_content: "GooD!!", Passenger_id: 1}
	feedback2 := Feedback{ID: 2, Booking_code: "KYSH43", Feedback_content: "BAD!!", Passenger_id: 1}
	feedback3 := Feedback{ID: 3, Booking_code: "NFXG08", Feedback_content: "AVERAGE!!", Passenger_id: 1}
	feedback4 := Feedback{ID: 4, Booking_code: "ZSDU99", Feedback_content: "NO COMMENT!!", Passenger_id: 2}
	feedback5 := Feedback{ID: 5, Booking_code: "PZHE59", Feedback_content: "QUITE OK!!", Passenger_id: 3}

	db.Create(&feedback1)
	db.Create(&feedback2)
	db.Create(&feedback3)
	db.Create(&feedback4)
	db.Create(&feedback5)
}
