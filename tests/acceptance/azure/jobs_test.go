package azure_test

import (
	"log"
	"testing"
	"time"

	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"
	workspaceHTTPModels "github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	workspaceModels "github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
	"github.com/stretchr/testify/assert"
)

const testJobNotebookPath = "/ScalaExampleNotebook"

var testJobClusterSpec = &models.NewCluster{
	SparkVersion: "7.3.x-scala2.12",
	NodeTypeID:   "Standard_D3_v2",
	NumWorkers:   1,
}

func beforeTestAzureJobsJobs(t *testing.T) {
	importReq := workspaceHTTPModels.ImportReq{
		Content:   "MSsx\n",
		Path:      testJobNotebookPath,
		Language:  workspaceModels.LanguageScala,
		Overwrite: true,
		Format:    workspaceModels.ExportFormatSource,
	}
	err := c.Workspace().Import(importReq)
	assert.Nil(t, err)
}

func TestAzureJobsJobs(t *testing.T) {
	beforeTestAzureJobsJobs(t)

	const jobName = "my-test-job-name"

	// Create
	createReq := httpmodels.CreateReq{
		Name:       jobName,
		NewCluster: testJobClusterSpec,
		NotebookTask: &models.NotebookTask{
			NotebookPath: testJobNotebookPath,
		},
		MaxRetries: 1,
	}
	createResp, err := c.Jobs().Create(createReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	// List
	listResp, err := c.Jobs().List()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, 1, len(*listResp.Jobs))

	// Get
	jobID := createResp.JobID
	getReq := httpmodels.GetReq{
		JobID: jobID,
	}
	getResp, err := c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, jobName, getResp.Settings.Name)
	assert.Equal(t, int32(1), getResp.Settings.MaxRetries)

	// Update
	updateReq := httpmodels.UpdateReq{
		JobID: jobID,
		NewSettings: &models.JobSettings{
			MaxRetries: 2,
		},
	}
	err = c.Jobs().Update(updateReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	getResp, err = c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), getResp.Settings.MaxRetries)

	// Reset
	resetReq := httpmodels.ResetReq{
		JobID: jobID,
		NewSettings: &models.JobSettings{
			Name:       jobName,
			NewCluster: testJobClusterSpec,
			NotebookTask: &models.NotebookTask{
				NotebookPath: testJobNotebookPath,
			},
			MaxRetries: 1,
		},
	}
	err = c.Jobs().Reset(resetReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	getResp, err = c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), getResp.Settings.MaxRetries)

	// Delete
	deleteReq := httpmodels.DeleteReq{
		JobID: jobID,
	}
	err = c.Jobs().Delete(deleteReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	_, err = c.Jobs().Get(getReq)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func beforeTestAzureJobsRuns(t *testing.T) int64 {
	beforeTestAzureJobsJobs(t)

	createJobReq := httpmodels.CreateReq{
		Name:       "my-test-job-name",
		NewCluster: testJobClusterSpec,
		NotebookTask: &models.NotebookTask{
			NotebookPath: testJobNotebookPath,
		},
		MaxRetries: 1,
	}
	createJobResp, err := c.Jobs().Create(createJobReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	return createJobResp.JobID
}

func afterTestAzureJobsRuns(t *testing.T, jobID int64) {
	deleteJobReq := httpmodels.DeleteReq{
		JobID: jobID,
	}
	err := c.Jobs().Delete(deleteJobReq)
	assert.Nil(t, err)
}

func TestAzureJobsRuns(t *testing.T) {
	jobID := beforeTestAzureJobsRuns(t)

	// Run Now
	runNowReq := httpmodels.RunNowReq{
		JobID: jobID,
	}
	runNowResp, err := c.Jobs().RunNow(runNowReq)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, runNowResp.NumberInJob, int64(1))
	assert.GreaterOrEqual(t, runNowResp.RunID, int64(1))
	time.Sleep(1 * time.Second)

	// Runs List
	runsListReq := httpmodels.RunsListReq{
		JobID:         jobID,
		ActiveOnly:    false,
		CompletedOnly: false,
		Offset:        0,
		Limit:         10,
	}
	runsListResp, err := c.Jobs().RunsList(runsListReq)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*runsListResp.Runs))

	// Runs Get
	runID := runNowResp.RunID
	runsGetReq := httpmodels.RunsGetReq{
		RunID: runID,
	}
	runsGetResp, err := c.Jobs().RunsGet(runsGetReq)
	assert.Nil(t, err)
	assert.Equal(t, jobID, runsGetResp.JobID)

	// wait for the run to complete
	timeoutCounter := 0
	for {
		timeoutCounter++

		runID := runNowResp.RunID
		runsGetReq := httpmodels.RunsGetReq{
			RunID: runID,
		}
		runsGetResp, err := c.Jobs().RunsGet(runsGetReq)
		log.Printf("Waiting for the run to complete, current state: %+v", runsGetResp.State)

		assert.Nil(t, err)
		if runsGetResp.State.LifeCycleState == models.RunLifeCycleStateTerminated {
			break
		}

		time.Sleep(10 * time.Second)
		assert.LessOrEqual(t, timeoutCounter, 60) // total of 600 seconds
	}

	// Runs Export
	runsExportReq := httpmodels.RunsExportReq{
		RunID:         runID,
		ViewsToExport: models.ViewsToExportCode,
	}
	runsExportResp, err := c.Jobs().RunsExport(runsExportReq)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(*runsExportResp.Views), 1)
	time.Sleep(1 * time.Second)

	// Runs Get Output
	runsGetOutputReq := httpmodels.RunsGetOutputReq{
		RunID: runID,
	}
	runsGetOutputResp, err := c.Jobs().RunsGetOutput(runsGetOutputReq)
	assert.Nil(t, err)
	assert.Equal(t, runsGetOutputResp.Metadata.RunID, runID)
	time.Sleep(15 * time.Second)

	afterTestAzureJobsRuns(t, jobID)

	// Runs Submit
	runsSubmitReq := httpmodels.RunsSubmitReq{
		RunName:    "my-test-run-name",
		NewCluster: testJobClusterSpec,
		NotebookTask: &models.NotebookTask{
			NotebookPath: testJobNotebookPath,
		},
	}
	runsSubmitResp, err := c.Jobs().RunsSubmit(runsSubmitReq)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, runsSubmitResp.RunID, int64(1))

	// Runs Cancel
	runsCancelReq := httpmodels.RunsCancelReq{
		RunID: int64(runsSubmitResp.RunID),
	}
	err = c.Jobs().RunsCancel(runsCancelReq)
	assert.Nil(t, err)
	time.Sleep(10 * time.Second)

	// Runs Delete
	runsDeleteReq := httpmodels.RunsDeleteReq{
		RunID: runID,
	}
	err = c.Jobs().RunsDelete(runsDeleteReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
}
