package azure_test

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
	"github.com/stretchr/testify/assert"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func TestAzureWorkspaceList(t *testing.T) {
	// list and assert
	const rootPath = "/"
	listReq := httpmodels.ListReq{
		Path: rootPath,
	}
	workSpaceList, e := c.Workspace().List(listReq)
	assert.Nil(t, e, "could not list workspaces")
	assert.NotEqual(t, len(*workSpaceList.Objects), 0, "No workspaces found in root path")
}

func TestAzureWorkspaceImportAndDelete(t *testing.T) {
	// import and assert
	var scalaLanguage models.Language = models.LanguageScala
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importReqScala := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  scalaLanguage,
		Format:    sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}
	assert.Nil(t, c.Workspace().Import(importReqScala), fmt.Sprintf("could not import request: %s", importReqScala.Language))

	// delete and assert
	deleteReq := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteReq), fmt.Sprintf("could not delete the imported resquest: %s", importReqScala.Language))
}

func TestAzureWorkspaceExport(t *testing.T) {
	// import and assert
	var pythonLanguage models.Language = models.LanguagePython
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importReqPython := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  pythonLanguage,
		Format:    sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}
	assert.Nil(t, c.Workspace().Import(importReqPython), fmt.Sprintf("could not import request for test: %s", importReqPython.Language))

	// Export and assert
	exportReq := httpmodels.ExportReq{
		Path:   samplePath,
		Format: sourceFormat,
	}
	exportRes, e := c.Workspace().Export(exportReq)
	assert.Nil(t, e, fmt.Sprintf("could not export notebook: %s", importReqPython.Path))
	strImportReq, e := base64.StdEncoding.DecodeString(importReqPython.Content)
	assert.Nil(t, e, fmt.Sprintf("could not decode import request content: %s", importReqPython.Content))
	strExportRes, e := base64.StdEncoding.DecodeString(exportRes.Content)
	assert.Nil(t, e, fmt.Sprintf("could not decode export response content: %s", exportRes.Content))
	assert.Equal(t, true, strings.Contains(string(strExportRes), string(strImportReq)), fmt.Sprintf("Imported content: %s\t Exported content: %s. Are not equal.", importReqPython.Content, exportRes.Content))

	// delete and assert
	deleteReq := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteReq), fmt.Sprintf("could not delete the imported resquest: %s", importReqPython.Language))
}

func TestAzureWorkspaceGetStatus(t *testing.T) {
	// import and assert
	var scalaLanguage models.Language = models.LanguageScala
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importReqScala := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  scalaLanguage,
		Format:    sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}
	assert.Nil(t, c.Workspace().Import(importReqScala), fmt.Sprintf("could not import request: %s", importReqScala.Language))

	// getstatus and assert
	statusReq := httpmodels.GetStatusReq{
		Path: samplePath,
	}
	getStatusRes, e := c.Workspace().GetStatus(statusReq)
	assert.Nil(t, e, fmt.Sprintf("could not get status from workspace: %s", importReqScala.Path))
	assert.Equal(t, samplePath, getStatusRes.Path, fmt.Sprintf("Import path: %s and Status path: %s. Not equal.", samplePath, getStatusRes.Path))

	// delete and assert
	deleteReq := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteReq), fmt.Sprintf("could not delete the imported resquest: %s", importReqScala.Language))
}

func TestAzureWorkspaceMkdirs(t *testing.T) {
	// import and assert
	samplePath := fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	mkdirsReq := httpmodels.MkdirsReq{
		Path: samplePath,
	}

	assert.Nil(t, c.Workspace().Mkdirs(mkdirsReq), fmt.Sprintf("could not mkdir: %s", samplePath))

	// delete and assert
	deleteReq := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteReq), fmt.Sprintf("could not delete the imported resquest: %s", samplePath))
}
