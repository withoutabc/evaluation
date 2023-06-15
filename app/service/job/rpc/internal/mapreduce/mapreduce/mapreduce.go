package mapreduce

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"rpc/app/common/consts/errs"
	"rpc/app/service/job/rpc/internal/mapreduce/model"
	"rpc/app/service/job/rpc/pb"
	"rpc/utils/time"
	"strconv"
)

type (
	Model interface {
		WordCount(jobName, inputPath, outputPath string) int32
		JoinWord(jobName, outputPath string, inputPath []string) int32
		JobList(Id string) ([]*pb.Jobs, int32)
		JoinData(jobName string) int32
	}
	DefaultModel struct {
		Ssh *ssh.Client
	}
)

func (d *DefaultModel) WordCount(jobName, inputPath, outputPath string) int32 {
	//开启mapreduce
	session, err := d.Ssh.NewSession()
	if err != nil {
		log.Println(err)
		return errs.InternalServer
	}
	defer session.Close()
	// input path output path
	if err = session.Start("hadoop jar " +
		streamingJarPath +
		Reduce +
		JobName + "\"" + jobName + "\"" +
		" -input " + inputPath +
		" -output " + outputPath +
		WordCountMapper +
		WordCountReducer +
		WordCountMapperFile +
		WordCountReducerFile +
		" &"); err != nil {
		return errs.InternalServer
	}
	return errs.No
}

func (d *DefaultModel) JoinData(jobName string) int32 {
	return errs.No
}

func (d *DefaultModel) JoinWord(jobName, outputPath string, inputPath []string) int32 {
	session, err := d.Ssh.NewSession()
	if err != nil {
		log.Println(err)
		return errs.InternalServer
	}
	defer session.Close()
	var inputPaths string
	for _, path := range inputPath {
		inputPaths = fmt.Sprintf(" %s -input %s", inputPaths, path)
	}
	// input path output path
	if err = session.Start("hadoop jar " +
		streamingJarPath +
		"-D " + jobName +
		inputPaths +
		" -output " + outputPath +
		WordCountMapper +
		WordCountReducer +
		WordCountMapperFile +
		WordCountReducer +
		" &"); err != nil {
		return errs.InternalServer
	}
	return errs.No
}

func (d *DefaultModel) JobList(Id string) ([]*pb.Jobs, int32) {
	var urlPrefix string
	if Id != "0" {
		urlPrefix = url + "/" + Id
	} else {
		urlPrefix = url
	}
	resp, err := http.Get(urlPrefix)
	if err != nil {
		log.Println("获取作业列表失败：", err)
		return nil, errs.InternalServer
	}
	defer resp.Body.Close()
	// 解析JSON响应
	var yarnResponse model.YarnApplications
	json.NewDecoder(resp.Body).Decode(&yarnResponse)
	var jobs []*pb.Jobs
	for _, app := range yarnResponse.Apps.App {
		job := &pb.Jobs{
			Id:           app.Id,
			Name:         app.Name,
			Status:       app.State,
			FinalStatus:  app.FinalStatus,
			Progress:     app.Progress,
			TrackingUrl:  app.TrackingUrl,
			StartedTime:  time.TimTransfer(app.StartedTime),
			LaunchTime:   time.TimTransfer(app.LaunchTime),
			FinishedTime: time.TimTransfer(app.FinishedTime),
			ElapsedTime:  strconv.Itoa(app.ElapsedTime),
		}
		jobs = append(jobs, job)
	}
	return jobs, errs.No
}

func NewModel(ssh *ssh.Client) *DefaultModel {
	return &DefaultModel{
		Ssh: ssh,
	}
}
