# databricks-sdk-golang

This is a Golang SDK for [DataBricks REST API 2.0](https://docs.databricks.com/api/latest/index.html#) and [Azure DataBricks REST API 2.0](https://docs.azuredatabricks.net/api/latest/index.html).

## Usage

```go
import (
  databricks "github.com/polar-rams/databricks-sdk-golang"
  dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
)

opt := databricks.NewDBClientOption("", "", os.Getenv("DATABRICKS_HOST"), os.Getenv("DATABRICKS_TOKEN"))
c := dbAzure.NewDBClient(opt)

jobs, err := c.Jobs().List()
```

## Implementation Progress

| API  | Status |
| :--- | :---: |
| Account API | N/A |
| Clusters API | ✔ |
| Cluster Policies API | ✗ |
| DBFS API | ✔ |
| Global Init Scripts | ✗ |
| Groups API | ✔ |
| Instance Pools API | ✔ |
| IP Access List API | ✗ |
| Jobs API | ✔ |
| Libraries API | ✔ |
| MLflow** API | ✗ |
| Permissions API | ✗ |
| SCIM** API | ✗ |
| Secrets API | ✔ |
| Token API | ✔ |
| Token Management API | ✗ |
| Workspace API | ✔ |

** SCIM and MLflow are separate systems that are planned differently.

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
