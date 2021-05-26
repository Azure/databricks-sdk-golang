package azure_test

import (
	"fmt"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/dbfs/httpmodels"
	"github.com/stretchr/testify/assert"
)

//Test the api: MKdirs, get-status
func TestAzureDBFSMkdirsAndStatus(t *testing.T) {
	path := "/dbfstest"
	mdirRequest := httpmodels.MkdirsReq{
		Path: path,
	}
	e := c.Dbfs().Mkdirs(mdirRequest)
	assert.Nil(t, e, "could not create directory")

	sRequest := httpmodels.GetStatusReq{
		Path: path,
	}
	sResp, e := c.Dbfs().GetStatus(sRequest)
	assert.Nil(t, e, "Couldn't get the status of the directory")
	assert.Equal(t, sResp.Path, path)
	assert.Equal(t, sResp.IsDir, true)

}

// Use create API to create a blank file
// Use List to check the directory that the fire size is not 0
// Addblock to add some data to the file
// Close the file and save it
func TestAzureDBFSCreateListAddblockClose(t *testing.T) {

	filepath := "/dbfstest/test.txt"
	dirpath := "/dbfstest/"
	createReq := httpmodels.CreateReq{
		Path:      filepath,
		Overwrite: true,
	}
	handler, e := c.Dbfs().Create(createReq)
	assert.Nil(t, e, "Cannot create file using the path. ")
	fmt.Sprintln("the handler is: ", handler.Handle)

	listRequest := httpmodels.ListReq{
		Path: dirpath,
	}
	listResponse, e := c.Dbfs().List(listRequest)
	assert.Nil(t, e, "Couldn't list the directory.")
	assert.NotEqual(t, len(*listResponse.Files), 0)

	addblockReq := httpmodels.AddBlockReq{
		Data:   "VGhpcyBpcyB0aGUgYWRkIGJsb2NrIHRlc3QuIA==", //'this is the addblock test'
		Handle: handler.Handle,
	}
	e = c.Dbfs().AddBlock(addblockReq)
	assert.Nil(t, e, "Could not add data to the file.")

	closeReq := httpmodels.CloseReq(handler)

	e = c.Dbfs().Close(closeReq)
	assert.Nil(t, e, "Could not close the file.")

}

//Put to upload new data and overwrite if there already an existing file.
//Read to get the data length and contents
func TestAzureDBFSPutRead(t *testing.T) {
	filepath := "/dbfstest/test.txt"
	putReq := httpmodels.PutReq{
		Path:      filepath,
		Contents:  "VGhpcyBpcyB0aGUgcHV0IHRlc3QuIA==", //'this is a put test'
		Overwrite: true,
	}

	e := c.Dbfs().Put(putReq)
	assert.Nil(t, e, "Could not put data to the file.")

	readReq := httpmodels.ReadReq{
		Path: filepath,
	}

	resp, e := c.Dbfs().Read(readReq)
	assert.Nil(t, e, "Could not read the file")
	assert.NotEqual(t, resp.BytesRead, 0)
	assert.Equal(t, resp.Data, putReq.Contents)
}

// Move the file from one dir to another
// Delete all the directories
func TestAzureDBFSMoveDelete(t *testing.T) {

	sourcePath := "/dbfstest"
	targetPath := "/dbfsnewtest"

	moveReq := httpmodels.MoveReq{
		SourcePath:      sourcePath,
		DestinationPath: targetPath,
	}
	e := c.Dbfs().Move(moveReq)
	assert.Nil(t, e, "Could not move file.")

	delRequest := httpmodels.DeleteReq{
		Path:      sourcePath,
		Recursive: true,
	}
	e = c.Dbfs().Delete(delRequest)
	assert.Nil(t, e, "Could not delete the source directory. ")

	delRequest = httpmodels.DeleteReq{
		Path:      targetPath,
		Recursive: true,
	}

	e = c.Dbfs().Delete(delRequest)
	assert.Nil(t, e, "Could not delete the target directory. ")

}
