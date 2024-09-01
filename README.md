# Testing Code Overview

## Writing Test Code Functions

A test function starts with `Test` and takes `*testing.T` as the only parameter. For example:

```go
func Test[NameOfFunction](t *testing.T) {
    // TestRunner from Testing Generic Code
	TestRunner(t, ServiceTestModels)
}
```
# Test Run Setup
## 1. Command Line
a) To run your tests, use the command:
```
go test
```
This command shows whether test cases pass or fail. The output will look like this:
```
PASS
ok      bitbucket.org/lavalogic/magma-accounts/test     20.504s
PS C:\Lava_Logic\FloWMS\Magma\magma-accounts\test> 
```
b) To see detailed information about all running test cases, use:
```
go test -v
```
The output will look like this:
```
=== RUN   TestMagmaAccounts/Address_:_Successfully_Created
=== RUN   TestMagmaAccounts/Address_:_Successfully_Created#01
=== RUN   TestMagmaAccounts/Address_:_Successfully_Created#02
    --- PASS: TestMagmaAccounts/Accounts_:_Successfully_Created (0.05s)
    --- PASS: TestMagmaAccounts/Accounts_:_Successfully_Created#01 (0.06s)
    --- PASS: TestMagmaAccounts/Accounts_:_Successfully_Created#02 (0.12s)
PASS
ok      bitbucket.org/lavalogic/magma-accounts/test     20.443s
```
**Ex. For testing in the magma-accounts folder:**  
• If you are in the `/magma-accounts` directory, use:  
  `go test .\test\ -v`  

• If you are in the `/magma-accounts/test` directory, use:  
  `go test -v`


**Note:** The go test command only runs functions that start with `Test`FunName name and have *testing.T as the only parameter .

## 2. Running Directly from VS Code
To run test function directly from VS Code, ensure that the file ends with `_test.go` (e.g., `magma_account_test.go`) & inside file testing functions name start with `Test`[funname] . You should see a "Run" option directly on the testing function . 

# Test Timeout Setup

**Note :** By default, the `go test` command has a timeout of 30 seconds. To prevent tests from being prematurely terminated due to timeouts, you can manually set a custom timeout.

## 1) From the Command Line
To run tests with a custom timeout from the command line, use the `-timeout` flag to specify the duration. Here is an example command to set the timeout to 60 seconds:

```go
go test -timeout 60s -v
```
## 2) Directly from VS Code Settings

If you prefer running tests directly from VS Code, you can update the timeout settings in VS Code itself to prevent tests from timing out prematurely.

**Steps to Update Test Timeout in VS Code:**

1. **Go to VS Code Settings:**

2. **Search for `Go: Test Timeout`:**
   - In the search bar, type `Go: Test Timeout` to locate the test timeout setting for Go.

3. **Locate the Default Value:**
   - You will find the default value set to `30s`.

4. **Update the Timeout Value:**
   - Change the value to a duration that suits your requirements, such as `60s` or more.

By increasing the timeout value, you ensure that longer-running tests have enough time to complete without being forcefully terminated by the `TestRunner` function. 

# Test Generic Code Overview

## TableConfig Struct

```go
// Table configuration structure
type TableConfig struct {
    AllowedMethodType        []string
    UrlPath                  string
    DbModel                  interface{}
    ServiceName              string
    CustomRequestHandler     map[string]Handler
    StructHandler            StructHandlerFun
    FieldDataHandler         FieldDataHandlerFun
    CustomContextObject      GetContextObject
    CustomizeRequestBody     CustomizeRequestBody
    CustomInputTestCasesFile map[string]string
    QueryParameter           string // Default is "id"; e.g., 'uuid' for some endpoints
    ServiceTestModels        map[string]*TableConfig
    ChildTableIDRecorder     map[string][]interface{}
}

func (currentTableConfig *TableConfig) AddDataIntoChildTable(tableName string) (interface{}, error)
```

### TableConfig Attributes

- **AllowedMethodType**: Represents the types of HTTP operations that can be performed on the current table endpoint.

- **UrlPath**: The endpoint path.

- **CustomRequestHandler**: Holds HTTP methods with Custom http handler functions of the endpoint.

- **StructHandler**: Callback function to mark internal struct attributes of the Model as `nil` to avoid empty record initialization during marshalling. Example:

  ```go
  type Accounts struct {
      Address *Address
  }

  func StructHandler(model interface{}) {
      // Mark the 'address' field as nil
      if acc, ok := model.(Accounts); ok {
          acc.Address = nil
      }
  }
``
- **FieldDataHandler**: Callback function to handle specific data value for fields of model.
  ```go
  FieldDataHandler: func(sf *reflect.StructField, v *reflect.Value, TableConfig *TableConfig) {
			switch sf.Name {
			case "ContactEmailAddress":
				gmail := GenerateRandomGmail()
				v.SetString(gmail)
			case "ContactNumber":
				num := GenerateStringOfDigit(10)
				v.SetString(num)
			}
		}
  ```

- **CustomContextObject**: Used to pass any specific type of context required by custom HTTP handlers.

- **CustomInputTestCasesFile**: Specifies JSON files for manual test cases. Example:

  ```go
  CustomInputTestCasesFile: map[string]string{
      http.MethodDelete: "test-data/accounts/delete-negative-tests.json",
      http.MethodPost:   "test-data/accounts/create-negative-tests.json",
  }
``
- **QueryParameter**: Specifies the query parameter if the endpoint uses something other than the default `id`. For example, `uuid` for pods service .

- **CustomizeRequestBody**: Allows additional JSON data to be added to the request body. Example:

  ```go
  CustomizeRequestBody: func(bodyMap map[string]interface{}, tc *TestCase) {
      additionalBody := `{
          "schedule": {
              "recipients": "recipient1@flowms.com,recipient2@flowms.com",
              "scheduleFrequency": "hourly",
              "second": 30,
              "minute": 39
          }
      }`

      // Merge additional JSON object into body
      tc.MergeJsonStringIntoRequestBody(bodyMap, additionalBody)
  }
``
- **ChildTableIDRecorder**: Holds inserted records in the child table, which are used to remove records later from child table.

- **AddDataIntoChildTable()**: Inserts data into a child table based on the table name/endpoint name and returns the target ID for use as a foreign key.

## Generic Important Functions

- **Testing HTTP Operations**:
- ```go
  func TestPostEndpoint(t *testing.T, tableConfig *TableConfig)
  func TestPatchEndpoint(t *testing.T, tableConfig *TableConfig)
  func TestDeleteEndpoint(t *testing.T, tableConfig *TableConfig)
  func TestGetEndPoint(t *testing.T, tableConfig *TableConfig)
  ```

- **TestInsertIntoDatabase()**
  
  Inserts a record into the database based on `TableConfig` and returns the inserted target IDs. This function helps in setting up the database with necessary test data.

- **TestDeleteFromDatabase()**
  
  Deletes records from the database based on a list of target IDs and `TableConfig`. This function cleans up records after tests are executed, ensuring that the database remains clean.

- **TestHttpRequestHandler()**
  
  Performs CRUD operations on an endpoint and tests common parameters of test cases.
  
- **DeleteChildTableDataFromTableConfig()**
  
  Deletes data from child tables based on the `ChildTableIDRecorder` map in `TableConfig`. This function ensures that any related data in child tables is removed .

## Random Data Generator Functions

- **TestGenerateDataForFields()**
  
  Generates random data for a model based on field data types. It includes the `FieldDataHandlerFun` callback to handle specific data requirements for different fields. This function is used to create diverse test data to thoroughly evaluate the model.

- **TestGetNotNullFields()**
  
  Returns a list of fields that are required based on the model's `GetCreationRules` function. This helps in identifying which fields must be provided with values to create a valid record.

- **TestGenerateRandomTestCases()**
  
  Generates positive test cases for an endpoint based on a specified number. This function is used to create a set of valid test cases that help in ensuring the endpoint's functionality under various scenarios.

- **TestGenerateNNFNegativeTestCases()**
  
  Generates negative test cases for not-null fields (NNF). This function is designed to test the endpoint's response when required fields are missing, ensuring that the system correctly handles such scenarios.

- **TestGenerateRNFNegativeTestCases()**
  
  Generates negative test cases for record not found (RNF). This function tests how the endpoint responds when trying to access or manipulate records that do not exist, verifying that the system appropriately handles such cases.

 - **Common other functions**
    ```go
    GenerateRandomString()
    GenerateRandomStringByChar(noOfChar int)
    GenerateRandomGmail()
    GenerateRandomDate() 
    GenerateStringOfDigit(noOfChar int)
    GenerateRandomInt()
    GenerateIntBetweenRange(start int, end int)
    GenerateUniqueUint()
    ```

## Common Issues and Panic Reasons

1. **Missing `StructHandler` or `FieldDataHandler` Callback Functions**

   **Issue:** If you do not provide `StructHandler` or `FieldDataHandler` callback functions in `TableConfig`, the generic code may panic during execution.

   **Solution:** Always specify these callback functions in `TableConfig` even if you do not intend to handle specific data. This helps avoid unexpected panics.

2. **Repeated DbModel Type in TableConfig**

    **Pani Log :** ```go 
    <orm.RegisterModel> model  repeat register, must be unique
       ```

    **Issue :** When configuring `TableConfig`, ensure that each `DbModel` type is unique. If you accidentally register the same `DbModel` type multiple times, it will lead to registration conflicts during ORM setup. Ex :
    ```go
        var ServiceTestModels = map[string]*TableConfig{
      	"accounts-address": &TableConfig{
              DbModel :(Account)
              },
	    "accounts&TableConfig{
            DbModel :(Accounts)
      }
   ```
    **Solution:** To resolve this issue, make sure that each `DbModel` type is only registered once in the `ServiceTestModels` map. This ensures that there are no conflicts during ORM registration.

3. **Custom Handler called gRPC Issues :**
   
   **Issue :** When using custom handlers in your TableConfig, if these handlers make calls to gRPC services, it may cause panics in the generic code. This happens because the generic code  not be designed to handle gRPC calls, leading to unexpected behavior or errors.


4. **Unmarked Struct Attributes of Model to Nil:**

     **Panic Log :**
   
     ```go
     panic: reflect.Set: value of type map[string]interface {} is not assignable to type *accountspb.AccountAddress [recovered]
        panic: reflect.Set: value of type map[string]interface {} is not assignable to type *accountspb.AccountAddress

    goroutine 6 [running]:
          C:/Program Files/Go/src/reflect/value.go:2325 +0xe6
    bitbucket.org/lavalogic/pyro/dbutils.SetFieldValue({0xd48100?, 0xc0001391e0?}, {0xc7f13a, 0x7}, {0xc909e0, 0xc00022cc30})
        c:/Lava_Logic/FloWMS/Magma/pyro/dbutils/model.go:227 +0x19f
    bitbucket.org/lavalogic/pyro/dbutils.MapToStruct(0xd48100?, {0xd48100, 0xc0001391e0})
        c:/Lava_Logic/FloWMS/Magma/pyro/dbutils/model.go:329 +0x125
    bitbucket.org/lavalogic/pyro.MapSliceToModelSlice({0xc000078718, 0x1, 0xd48100?}, {0xd48100, 0xc0001388f0})
        c:/Lava_Logic/FloWMS/Magma/pyro/http_crud.go:286 +0x3e5
    bitbucket.org/lavalogic/pyro.HandleCRUD({0xe78780, 0xc00022cba0}, {0x0, 0x0}, {0xd48100, 0xc0001388f0}, 0xc000200120)
        c:/Lava_Logic/FloWMS/Magma/pyro/http_crud.go:100 +0x1285
     ```
   **Issue :**  If internal struct attributes are not marked as `nil` in the `StructHandler` function, it can lead to panics in the generic code. This is because the generic code expects certain fields to be explicitly set to `nil` .


 



