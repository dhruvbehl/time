package time

import "time"

func getTimeByTimeZone(timezone string) (string, error){
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).String(), nil
}