package paytm

import (
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
	// "strings"
)

type Class struct {
	// Code string
	Fare int
	Seats string // seats status
	Fresh string // last updated x hrs ago
}

type Alternate struct {
	Src string
	Dst string
	Avail map[string]Class
}

type Train struct {
	Name string
	No string
	Dura string // duration
	Src string
	Dst string
	Dept string
	Arri string
	Avail map[string]Class // seats and fare by class 1A,2A,3A,SL
	Alt []Alternate
}

func checkErr(err error){
	if err != nil {
		panic(err.Error())
	}
}

func parseJson(body string) ([]Train) {

	trains := []Train{}

	gjson.Get(body, "body.trains").ForEach(func(_, v gjson.Result) bool {
		var s = new(Train)	
		s.Name = v.Get("trainName").String()
		s.No = v.Get("trainNumber").String()
		s.Dura = v.Get("duration").String()
		s.Avail = make(map[string]Class)
		s.Src = v.Get("source").String()
		s.Dst = v.Get("destination").String()

		date_d := v.Get("departure").String()
		s.Dept = date_d[8:10] +"." + date_d[11:16]
		date_a := v.Get("arrival").String()
		s.Arri = date_a[8:10] +"." + date_a[11:16]

		v.Get("availability").ForEach(func(_, v1 gjson.Result) bool {
			var a = new(Class)
			// a.Code = v1.Get("code").String()
			a.Fare = int(v1.Get("fare").Int())
			a.Seats = v1.Get("status").String()
			a.Fresh = v1.Get("time_of_availability").String()
			// s.Avail = append(s.Avail, *a)
			s.Avail[v1.Get("code").String()] = *a
			return true
		})

		v.Get("alternate_stations_data").ForEach(func(_, v1 gjson.Result) bool {
			var alt_ = new(Alternate)
			alt_.Src = v1.Get("new_source").String()
			alt_.Dst = v1.Get("new_destination").String()
			alt_.Avail = make(map[string]Class)
			v1.Get("availability").ForEach(func(_, v2 gjson.Result) bool {
				var a = new(Class)
				// a.Code = v2.Get("code").String()
				a.Fare = int(v2.Get("fare").Int())
				a.Seats = v2.Get("status").String()
				a.Fresh = v2.Get("time_of_availability").String()
				alt_.Avail[v2.Get("code").String()] = *a
				// alt_.Avail = append(alt_.Avail, *a)
				return true
			})
			s.Alt = append(s.Alt, *alt_)
			return true
		})

		trains = append(trains, *s)
		return true // keep iterating
	})
	
    return trains
}

func getRequest(src string, dst string, date string) string { // date yyyymmdd
	// return HardResponse

	res, err := http.Get("https://travel.paytm.com/api/trains/v3/search?source="+src+"&destination="+dst+"&departureDate="+date)
	checkErr(err)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	return string(body)
}

func Api(src string, dst string, date string) []Train {
	res := getRequest(src, dst, date)
	
	data := parseJson(res)
	// for _, t := range data {
	// 	fmt.Println(t)
	// }
	return data
}

// func main(){

// }

const HardResponse = `{
    "error": null,
    "status": {
        "result": "success",
        "message": {
            "title": "Successful",
            "message": "A successful operation has made"
        }
    },
    "body": {
        "trains": [
            {
                "departure": "2020-02-14T21:13:00+00:00",
                "arrival": "2020-02-15T14:35:00+00:00",
                "trainName": "MAHANANDA EXP",
                "trainNumber": "15484",
                "source": "DDU",
                "destination": "NJP",
                "source_name": "Deen Dayal Upadhyaya Jn",
                "destination_name": "New Jalpaiguri",
                "duration": "17:22",
                "runningOn": {
                    "sun": "Y",
                    "mon": "Y",
                    "tue": "Y",
                    "wed": "Y",
                    "thu": "Y",
                    "fri": "Y",
                    "sat": "Y"
                },
                "classes": [
                    "SL",
                    "3A",
                    "2A"
                ],
                "message_enabled": "false",
                "message_text": "",
                "train_type": "O",
                "alternate_stations": false,
                "tatkal_text": "View Tatkal Seats",
                "tatkal_enabled": false,
                "availability": [
                    {
                        "code": "3A",
                        "name": "AC 3 Tier",
                        "status": "AVAILABLE-0001",
                        "fare": 1040,
                        "time_of_availability": "5 hrs ago",
                        "time_diff": 322,
                        "colour": "GREEN",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "SL",
                        "name": "Sleeper Class",
                        "status": "RLWL7/WL7",
                        "fare": 385,
                        "time_of_availability": "2 hrs ago",
                        "time_diff": 108,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "2A",
                        "name": "AC 2 Tier",
                        "status": "RLWL1/WL1",
                        "fare": 1485,
                        "time_of_availability": "13 days ago",
                        "time_diff": 19293,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    }
                ],
                "availability_count": 1,
                "least_price": 1040,
                "cache_time": 322,
                "stationCodeMatch": 1
            },
            {
                "departure": "2020-02-14T11:58:00+00:00",
                "arrival": "2020-02-15T00:20:00+00:00",
                "trainName": "NDLS SCL PSK EXP",
                "trainNumber": "15602",
                "source": "DDU",
                "destination": "NJP",
                "source_name": "Deen Dayal Upadhyaya Jn",
                "destination_name": "New Jalpaiguri",
                "duration": "12:22",
                "runningOn": {
                    "sun": "N",
                    "mon": "N",
                    "tue": "N",
                    "wed": "N",
                    "thu": "N",
                    "fri": "Y",
                    "sat": "N"
                },
                "classes": [
                    "SL",
                    "3A",
                    "2A",
                    "1A"
                ],
                "message_enabled": "false",
                "message_text": "",
                "train_type": "O",
                "alternate_stations": false,
                "tatkal_text": "View Tatkal Seats",
                "tatkal_enabled": false,
                "availability": [
                    {
                        "code": "SL",
                        "name": "Sleeper Class",
                        "status": "RLWL8/WL7",
                        "fare": 395,
                        "time_of_availability": "2 hrs ago",
                        "time_diff": 107,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "3A",
                        "name": "AC 3 Tier",
                        "status": "RLWL7/WL7",
                        "fare": 1045,
                        "time_of_availability": "7 hrs ago",
                        "time_diff": 413,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "2A",
                        "name": "AC 2 Tier",
                        "status": "RLWL3/WL3",
                        "fare": 1500,
                        "time_of_availability": "13 hrs ago",
                        "time_diff": 774,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "1A",
                        "name": "First Class AC",
                        "status": "BOOKING CLOSED",
                        "fare": 2490,
                        "time_of_availability": "13 hrs ago",
                        "time_diff": 774,
                        "colour": "RED",
                        "tip_text": "",
                        "quota": "GN"
                    }
                ],
                "availability_count": 0,
                "least_price": 395,
                "cache_time": 107,
                "stationCodeMatch": 1
            },
            {
                "departure": "2020-02-14T18:23:00+00:00",
                "arrival": "2020-02-15T08:30:00+00:00",
                "trainName": "NORTH EAST EXP",
                "trainNumber": "12506",
                "source": "DDU",
                "destination": "NJP",
                "source_name": "Deen Dayal Upadhyaya Jn",
                "destination_name": "New Jalpaiguri",
                "duration": "14:07",
                "runningOn": {
                    "sun": "Y",
                    "mon": "Y",
                    "tue": "Y",
                    "wed": "Y",
                    "thu": "Y",
                    "fri": "Y",
                    "sat": "Y"
                },
                "classes": [
                    "SL",
                    "3A",
                    "2A"
                ],
                "message_enabled": "false",
                "message_text": "",
                "train_type": "O",
                "alternate_stations": true,
                "tatkal_text": "View Tatkal Seats",
                "tatkal_enabled": false,
                "availability": [
                    {
                        "code": "SL",
                        "name": "Sleeper Class",
                        "status": "RLWL9/WL9",
                        "fare": 420,
                        "time_of_availability": "a minute ago",
                        "time_diff": 1,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN",
                        "pnr_prediction": {
                            "value": 49,
                            "color": "#DA9800"
                        }
                    },
                    {
                        "code": "3A",
                        "name": "AC 3 Tier",
                        "status": "RLWL14/WL12",
                        "fare": 1075,
                        "time_of_availability": "a day ago",
                        "time_diff": 1756,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "2A",
                        "name": "AC 2 Tier",
                        "status": "RLWL3/WL3",
                        "fare": 1515,
                        "time_of_availability": "2 days ago",
                        "time_diff": 2840,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    }
                ],
                "alternate_stations_data": [
                    {
                        "original_source": "DDU",
                        "original_destination": "NJP",
                        "new_source": "ANVT",
                        "new_destination": "NCB",
                        "new_source_name": "Anand Vihar Terminus Delhi",
                        "new_destination_name": "New Cooch Behar",
                        "original_departure_date": "20200214",
                        "new_departure_date": "20200214",
                        "original_departure": "2020-02-14T18:23:00+00:00",
                        "original_arrival": "2020-02-15T08:30:00+00:00",
                        "new_departure": "2020-02-14T06:45:00+00:00",
                        "new_arrival": "2020-02-15T10:50:00+00:00",
                        "new_duration": "28:05",
                        "original_duration": "14:07",
                        "expand_text": "Ticket will be booked from <b>\"Anand Vihar Terminus Delhi\"</b> to <b>\"New Cooch Behar\"</b> but you board train at <b>\"Deen Dayal Upadhyaya Jn\"</b> and get off train at <b>\"New Jalpaiguri\"</b>",
                        "button_text": "Book for ₹690",
                        "availability": [
                            {
                                "code": "SL",
                                "name": "Sleeper Class",
                                "status": "AVAILABLE-0154",
                                "fare": 690,
                                "time_of_availability": "6 hrs ago",
                                "time_diff": 378,
                                "colour": "GREEN",
                                "tip_text": "",
                                "quota": "GN"
                            }
                        ]
                    }
                ],
                "alternate_stations_title": "Guaranteed Seat Assistance",
                "alternate_stations_text": "Available seats on the same train from other stations overlapping on your route.",
                "availability_count": 0,
                "least_price": 420,
                "cache_time": 1,
                "stationCodeMatch": 1
            },
            {
                "departure": "2020-02-14T11:48:00+00:00",
                "arrival": "2020-02-15T05:00:00+00:00",
                "trainName": "BRAHMPUTRA MAIL",
                "trainNumber": "15956",
                "source": "DDU",
                "destination": "NJP",
                "source_name": "Deen Dayal Upadhyaya Jn",
                "destination_name": "New Jalpaiguri",
                "duration": "17:12",
                "runningOn": {
                    "sun": "Y",
                    "mon": "Y",
                    "tue": "Y",
                    "wed": "Y",
                    "thu": "Y",
                    "fri": "Y",
                    "sat": "Y"
                },
                "classes": [
                    "SL",
                    "3A",
                    "2A"
                ],
                "message_enabled": "false",
                "message_text": "",
                "train_type": "O",
                "alternate_stations": false,
                "tatkal_text": "View Tatkal Seats",
                "tatkal_enabled": false,
                "availability": [
                    {
                        "code": "SL",
                        "name": "Sleeper Class",
                        "status": "RLWL8/WL8",
                        "fare": 435,
                        "time_of_availability": "5 hrs ago",
                        "time_diff": 322,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    },
                    {
                        "code": "3A",
                        "name": "AC 3 Tier",
                        "status": "RLWL1/WL1",
                        "fare": 1170,
                        "time_of_availability": "10 days ago",
                        "time_diff": 14791,
                        "colour": "ORANGE",
                        "tip_text": "",
                        "quota": "GN"
                    }
                ],
                "availability_count": 0,
                "least_price": 435,
                "cache_time": 322,
                "stationCodeMatch": 1
            },
            {
                "departure": "2020-02-14T01:33:00+00:00",
                "arrival": "2020-02-14T13:20:00+00:00",
                "trainName": "DBRT RAJDHANI",
                "trainNumber": "12424",
                "source": "DDU",
                "destination": "NJP",
                "source_name": "Deen Dayal Upadhyaya Jn",
                "destination_name": "New Jalpaiguri",
                "duration": "11:47",
                "runningOn": {
                    "sun": "Y",
                    "mon": "Y",
                    "tue": "Y",
                    "wed": "Y",
                    "thu": "Y",
                    "fri": "Y",
                    "sat": "Y"
                },
                "classes": [
                    "3A",
                    "2A",
                    "1A"
                ],
                "message_enabled": "false",
                "message_text": "",
                "train_type": "R",
                "alternate_stations": false,
                "tatkal_text": "View Tatkal Seats",
                "tatkal_enabled": false,
                "availability": [
                    {
                        "code": "1A",
                        "name": "First Class AC",
                        "status": "BOOKING CLOSED",
                        "fare": 2925,
                        "time_of_availability": "a month ago",
                        "time_diff": 53471,
                        "colour": "RED",
                        "tip_text": "",
                        "quota": "GN"
                    }
                ],
                "availability_count": 0,
                "least_price": 2925,
                "cache_time": 53471,
                "stationCodeMatch": 1
            }
        ],
        "filters": [
            {
                "id": "nonac",
                "title": "Non AC",
                "filters": [
                    {
                        "label": "Sleeper Class",
                        "values": [
                            "SL"
                        ]
                    }
                ]
            },
            {
                "id": "ac",
                "title": "AC",
                "filters": [
                    {
                        "label": "AC 3 Tier",
                        "values": [
                            "3A"
                        ]
                    },
                    {
                        "label": "AC 2 Tier",
                        "values": [
                            "2A"
                        ]
                    },
                    {
                        "label": "All Other ACs",
                        "values": [
                            "1A"
                        ]
                    }
                ]
            }
        ],
        "serverId": "DM07AP33MS3",
        "quota": [
            "GN",
            "LD"
        ],
        "search_source": "DDU",
        "search_destination": "NJP",
        "search_source_name": "Deen Dayal Upadhyaya Jn",
        "search_destination_name": "New Jalpaiguri",
        "tip_enabled": true,
        "tip_text": "Click here to know what RLWL, RAC, PQWL etc. mean",
        "promotional_messages": {
            "title": "Offers",
            "autoscroll": true,
            "messages": [
                {
                    "type": "text",
                    "message": "Pay through UPI and get Instant Refund on Cancellation"
                }
            ]
        }
    },
    "meta": {
        "quota": {
            "GN": "General",
            "TQ": "Tatkal",
            "LD": "Ladies"
        },
        "classes": {
            "1A": "First Class AC",
            "EC": "Executive Class",
            "2A": "AC 2 Tier",
            "FC": "First Class",
            "3A": "AC 3 Tier",
            "3E": "AC 3 Tier Economy",
            "CC": "AC Chair Car",
            "EA": "Anubhuti",
            "SL": "Sleeper Class",
            "2S": "Second Sitting"
        },
        "requestid": "4c776513-0575-42f6-b99c-30bdc27ab202"
    },
    "code": 200
}`