package time

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getTime(w http.ResponseWriter, req *http.Request) {
	response := make(map[string]string)
	responseErr := make(map[string]string)
	queryParameter := req.URL.Query().Get("tz")
	timezoneList :=  strings.Split(queryParameter, ",")
	for _, timezone := range timezoneList {
		if timezone == "" {
			timezone = time.UTC.String()
		}
		time, err := getTimeByTimeZone(timezone)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			responseErr["error"] =  fmt.Sprintf("invalid timezone: [%v]", timezone)
			json.NewEncoder(w).Encode(responseErr)
			return
		}
		response[timezone] = time
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}