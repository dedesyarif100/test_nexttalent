package country

type Country struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Timezone struct {
	// {
	// 	"year": 2023,
	// 	"month": 9,
	// 	"day": 2,
	// 	"hour": 12,
	// 	"minute": 12,
	// 	"seconds": 11,
	// 	"milliSeconds": 121,
	// 	"dateTime": "2023-09-02T12:12:11.1214976",
	// 	"date": "09/02/2023",
	// 	"time": "12:12",
	// 	"timeZone": "Asia/Jakarta",
	// 	"dayOfWeek": "Saturday",
	// 	"dstActive": false
	// }
	Year         int
	Month        int
	Day          int
	Hour         int
	Minute       int
	Seconds      int
	MilliSeconds int
	Datetime     string
	Date         string
	Time         string
	TimeZone     string
	DayOfWeek    string
	DstActive    bool
}
