// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure_test

import (
	"fmt"
	"testing"
	"time"

	clusterHttpModels "github.com/Azure/databricks-sdk-golang/azure/clusters/httpmodels"
	clusterModels "github.com/Azure/databricks-sdk-golang/azure/clusters/models"
	"github.com/Azure/databricks-sdk-golang/azure/libraries/httpmodels"
	"github.com/Azure/databricks-sdk-golang/azure/libraries/models"
	"github.com/stretchr/testify/suite"
)

type ClusterId struct {
	suite.Suite
	clusterId string
}

/* Function called before any test method in the suite is called. Creates cluster and waits until it comes to RUNNING state */
func (suite *ClusterId) SetupSuite() {
	var clusterName = "acceptance_test_cluster_Libraries"
	sparkConf := map[string]string{
		"spark.speculation": "true",
	}

	req := clusterHttpModels.CreateReq{
		ClusterName:  clusterName,
		SparkVersion: "7.3.x-scala2.12",
		NodeTypeID:   "Standard_D3_v2",
		SparkConf:    sparkConf,
		NumWorkers:   3,
	}

	resp, _ := c.Clusters().Create(req)

	suite.clusterId = resp.ClusterID

	getReq := clusterHttpModels.GetReq{
		ClusterID: suite.clusterId,
	}

	for {
		resp, e := c.Clusters().Get(getReq)
		if resp.State == clusterModels.ClusterStateRunning || e != nil {
			break
		}
		fmt.Println(resp.State)
		time.Sleep(10 * time.Second)
	}
}

/* Test to kick start other tests of the suite */
func TestAzureLibraries(t *testing.T) {
	suite.Run(t, new(ClusterId))
}

/* Test the libraries are installed in the cluster and veryfying them by checking the status by listing cluster status */
func (suite *ClusterId) TestLibrariesInstall() {
	installReq := httpmodels.InstallReq{
		ClusterID: suite.clusterId,
		Libraries: &[]models.Library{
			models.Library{
				Jar: "dbfs:/mnt/libraries/library.jar",
				Egg: "dbfs:/mnt/libraries/library.egg",
				Whl: "dbfs:/mnt/libraries/mlflow-0.0.1.dev0-py2-none-any.whl",
				Pypi: models.PythonPyPiLibrary{
					Package: "simplejson",
					Repo:    "https://my-pypi-mirror.com",
				},
				Maven: models.MavenLibrary{},
				Cran:  models.RCranLibrary{},
			},
		},
	}
	error := c.Libraries().Install(installReq)
	suite.Equal(nil, error)

	resp, error := c.Libraries().AllClusterStatuses()
	suite.Equal(nil, error)
	var isClusterFound bool = false

	for _, value := range *resp.Statuses {
		if value.ClusterID == suite.clusterId {
			isClusterFound = true
			suite.NotEqual(0, len(*value.LibraryStatuses))
			break
		}
	}
	suite.Equal(isClusterFound, true)

	clusterStatusReq := httpmodels.ClusterStatusReq{
		ClusterID: suite.clusterId,
	}
	response, error := c.Libraries().ClusterStatus(clusterStatusReq)
	suite.Equal(nil, error)
	suite.NotEqual(0, len(*response.LibraryStatues))
}

/* Test the libraries are uninstalled in the cluster */
func (suite *ClusterId) TestLibrariesUninstall() {
	uninstallReq := httpmodels.UninstallReq{
		ClusterID: suite.clusterId,
		Libraries: &[]models.Library{
			models.Library{
				Jar: "dbfs:/mnt/libraries/library.jar",
				Egg: "dbfs:/mnt/libraries/library.egg",
				Whl: "dbfs:/mnt/libraries/mlflow-0.0.1.dev0-py2-none-any.whl",
				Pypi: models.PythonPyPiLibrary{
					Package: "simplejson",
					Repo:    "https://my-pypi-mirror.com",
				},
				Maven: models.MavenLibrary{},
				Cran:  models.RCranLibrary{},
			},
		},
	}
	error := c.Libraries().Uninstall(uninstallReq)
	suite.Equal(nil, error)
}

/* Delete the cluster after all the tests in the suite finish */
func (suite *ClusterId) TearDownSuite() {
	clusterDelReq := clusterHttpModels.DeleteReq{
		ClusterID: suite.clusterId,
	}
	error := c.Clusters().Delete(clusterDelReq)
	suite.Equal(nil, error)
}
