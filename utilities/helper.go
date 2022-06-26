package utilities

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func OrAssignment(a, b string, rest ...string) string {
	if a != "" {
		return a
	}
	if b != "" {
		return b
	}
	for _, v := range rest {
		if v != "" {
			return v
		}
	}
	return ""
}

func GetDate(dateOfBirth string) (time.Time, error) {
	day := dateOfBirth[0:2]
	month := dateOfBirth[3:5]
	year := dateOfBirth[6:len(dateOfBirth)]

	y, err := strconv.Atoi(year)
	if err != nil {
		return time.Time{}, err
	}

	m, _ := strconv.Atoi(month)
	if err != nil {
		return time.Time{}, err
	}

	d, _ := strconv.Atoi(day)
	if err != nil {
		return time.Time{}, err
	}

	date := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)

	return date, nil

}

// Calculate provides an age from the given time to the present time.
func AgeCalculate(givenTime time.Time) int {
	presentTime := time.Now()

	switch givenLocation := givenTime.Location(); givenLocation {
	case time.UTC, nil:
		presentTime = presentTime.UTC()
	default:
		presentTime = presentTime.In(givenLocation)
	}

	givenYear := givenTime.Year()
	presentYear := presentTime.Year()

	age := presentYear - givenYear

	givenYearIsLeapYear := isLeapYear(givenYear)
	presentYearIsLeapYear := isLeapYear(presentYear)

	givenYearDay := givenTime.YearDay()
	presentYearDay := presentTime.YearDay()

	if givenYearIsLeapYear && !presentYearIsLeapYear && givenYearDay >= 60 {
		givenYearDay--
	} else if presentYearIsLeapYear && !givenYearIsLeapYear && presentYearDay >= 60 {
		givenYearDay++
	}

	if presentYearDay < givenYearDay {
		age--
	}

	return age
}

func isLeapYear(givenYear int) bool {
	if givenYear%400 == 0 {
		return true
	} else if givenYear%100 == 0 {
		return false
	} else if givenYear%4 == 0 {
		return true
	}

	return false
}
func IoReaderFromBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(requestBody), nil
}

func UpdateQueryParams(req *http.Request, qs url.Values) {
	q := req.URL.Query()
	for k, values := range qs {
		for _, v := range values {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
}

func UpdateHeaders(req *http.Request, headers http.Header) {
	for k, values := range headers {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
}

func ContainsInStringArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
