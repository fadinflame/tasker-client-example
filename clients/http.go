package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	pbWrapper "google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"tasker-client-example/models"
	"tasker-client-example/pb"
)

type HttpTest struct {
	Client  http.Client
	BaseURL string
}

func (hc *HttpTest) TestMethods(initialTask models.Task) (bool, error) {
	task := initialTask

	// CREATE TASK
	createJson, err := json.Marshal(task)
	if err != nil {
		return false, nil
	}

	resp, err := hc.Client.Post(fmt.Sprintf(hc.BaseURL, "create-task"), "application/json", bytes.NewBuffer(createJson))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	createResp := pb.TaskCreateResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Create", createResp.Success, "HTTP")

	// GET TASK
	taskPath := fmt.Sprintf("task/%d", createResp.TaskId)
	resp, err = hc.Client.Get(fmt.Sprintf(hc.BaseURL, taskPath))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	getResp := pb.TaskGetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Get", getResp.Title == task.Title, "HTTP")

	// UPDATE TASK
	updateJson, err := json.Marshal(models.Task{
		Id:          getResp.TaskId,
		Title:       "Updated",
		Text:        "Updated",
		IsCompleted: getResp.IsCompleted,
	})
	if err != nil {
		return false, nil
	}

	resp, err = hc.Client.Post(fmt.Sprintf(hc.BaseURL, taskPath), "application/json", bytes.NewBuffer(updateJson))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	updateResp := pb.TaskUpdateResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Update", updateResp.Success, "HTTP")

	// DELETE TASK
	delRequest, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf(hc.BaseURL, taskPath), nil)
	resp, err = hc.Client.Do(delRequest)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	deleteResp := pbWrapper.BoolValue{}
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Delete", deleteResp.Value, "HTTP")

	return true, nil
}
