package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) issueTime(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issuePrice:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(3 * time.Second)
		sendTimeRequest(input)
	}()
}

func sendTimeRequest(request models.Request) {
	var timeEntry string
	if rand.Intn(10)%10 > 3 {
		minDate := time.Date(2023, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		maxDate := time.Date(2024, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		delta := maxDate - minDate

		sec := rand.Int63n(delta) + minDate
		timeEntry = time.Unix(sec, 0).Format(time.DateTime)

		fmt.Println("calculated")
	}
	answer := models.TimeRequest{
		AccessToken: 123,
		Time:        timeEntry,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/tickets/%d/update_time_entry/", request.TicketId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
