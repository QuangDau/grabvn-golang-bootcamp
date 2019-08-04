# grabvn-golang-bootcamp
Assignment 3 : Create CRUD API for Passenger feedback using gRPC

Usage: POST:"http:localhost:8080/api/feedback/" add a passenger feedback :{
                                                                             passenger_id : int //required
                                                                             booking_code :string// required
                                                                             feedback_content: string //required
                                                                          }
                                                                          
                                                                          
       GET:"http:localhost:8080/api/feedback/passenger/:id" get all feedbacks belong to a passenger. 
            Example: GET "http:localhost:8080/api/feedback/passenger/1" returns {
                                                                                    "feedbacks": [
                                                                                        {
                                                                                            "id": 1,
                                                                                            "passenger_id": 1,
                                                                                            "booking_code": "KJAG12",
                                                                                            "feedback_content": "GooD!!"
                                                                                        },
                                                                                        {
                                                                                            "id": 2,
                                                                                            "passenger_id": 1,
                                                                                            "booking_code": "KYSH43",
                                                                                            "feedback_content": "BAD!!"
                                                                                        },
                                                                                        {
                                                                                            "id": 3,
                                                                                            "passenger_id": 1,
                                                                                            "booking_code": "NFXG08",
                                                                                            "feedback_content": "AVERAGE!!"
                                                                                        }
                                                                                    ]
                                                                                }
                                                                                
                                                                                
       GET:"http:localhost:8080/api/feedback/bookingcode/:code" get feedback belongs to a booking code. 
            Example: GET:"http:localhost:8080/api/feedback/bookingcode/NFXG08" returns  {
                                                                                            "id": 3,
                                                                                            "passenger_id": 1,
                                                                                            "booking_code": "NFXG08",
                                                                                            "feedback_content": "AVERAGE!!"
                                                                                        }
                                                                                        
       DELETE:"http:localhost:8080/api/feedback/:id" delete all feedbacks belong to a passenger. 
            Example: DELETE:"http:localhost:8080/api/feedback/1"
