package clients

import (
	"context"
	"tasker-client-example/models"
	"tasker-client-example/pb"
)

type GrpcTest struct {
	GrpcClient pb.TaskServiceClient
	Ctx        context.Context
}

func (gt *GrpcTest) TestMethods(initialTask models.Task) (bool, error) {
	task := initialTask

	taskCreateResp, err := gt.GrpcClient.CreateTask(gt.Ctx, &pb.TaskCreateRequest{Title: task.Title, Text: task.Text})
	if err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Create", taskCreateResp.Success, "gRPC")

	taskGetResp, err := gt.GrpcClient.GetTask(gt.Ctx, &pb.TaskGetRequest{TaskId: taskCreateResp.TaskId})
	if err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Get", task.Title == taskGetResp.Title && task.Text == taskGetResp.Text, "gRPC")

	taskUpdateResp, err := gt.GrpcClient.UpdateTask(gt.Ctx, &pb.TaskUpdateRequest{
		TaskId:      taskCreateResp.TaskId,
		Title:       "Updated",
		Text:        "Updated",
		IsCompleted: taskGetResp.IsCompleted,
	})
	if err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Update", taskUpdateResp.Success, "gRPC")

	taskDeleteResp, err := gt.GrpcClient.DeleteTask(gt.Ctx, &pb.TaskDeleteRequest{TaskId: taskCreateResp.TaskId})
	if err != nil {
		return false, err
	}
	PrintMethodSucceed("Task Delete", taskDeleteResp.Success, "gRPC")

	return true, nil
}
