package azure

import databricks "github.com/polar-rams/databricks-sdk-golang"

// DBClient is the client for Azure implements DBClient
type DBClient struct {
	Option databricks.DBClientOption
}

func NewDBClient(option *databricks.DBClientOption) *DBClient {
	return &DBClient{
		Option: *option,
	}
}

// Clusters returns an instance of ClustersAPI
func (c DBClient) Clusters() ClustersAPI {
	var clustersAPI ClustersAPI
	return clustersAPI.init(c)
}

// Dbfs returns an instance of DbfsAPI
func (c DBClient) Dbfs() DbfsAPI {
	var dbfsAPI DbfsAPI
	return dbfsAPI.init(c)
}

// Groups returns an instance of GroupAPI
func (c DBClient) Groups() GroupsAPI {
	var groupsAPI GroupsAPI
	return groupsAPI.init(c)
}

// InstancePools returns an instance of InstancePoolsAPI
func (c DBClient) InstancePools() InstancePoolsAPI {
	var instancePoolsAPI InstancePoolsAPI
	return instancePoolsAPI.init(c)
}

// Jobs returns an instance of JobsAPI
func (c DBClient) Jobs() JobsAPI {
	var jobsAPI JobsAPI
	return jobsAPI.init(c)
}

// Libraries returns an instance of LibrariesAPI
func (c DBClient) Libraries() LibrariesAPI {
	var libraries LibrariesAPI
	return libraries.init(c)
}

// Secrets returns an instance of SecretsAPI
func (c DBClient) Secrets() SecretsAPI {
	var secretsAPI SecretsAPI
	return secretsAPI.init(c)
}

// Token returns an instance of TokensAPI
func (c DBClient) Token() TokenAPI {
	var tokenAPI TokenAPI
	return tokenAPI.init(c)
}

// Workspace returns an instance of WorkspaceAPI
func (c DBClient) Workspace() WorkspaceAPI {
	var workspaceAPI WorkspaceAPI
	return workspaceAPI.init(c)
}

func (c *DBClient) performQuery(method, path string, data interface{}, headers map[string]string) ([]byte, error) {
	return databricks.PerformQuery(c.Option, method, path, data, headers)
}
