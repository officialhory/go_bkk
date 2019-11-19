package main

type server struct{}

type bkkAPIResponse struct {
	Data struct {
		Entry struct {
			StopTimes []struct {
				PredictedArrivalTime int `json:"predictedArrivalTime"`
			} `json:"stopTimes"`
		} `json:"entry"`
	} `json:"data"`
}
