# \DefaultAPI

All URIs are relative to *http://localhost/api/plugins/grafana-slo-app/resources*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SloGet**](DefaultAPI.md#V1SloGet) | **Get** /v1/slo | 
[**V1SloIdDelete**](DefaultAPI.md#V1SloIdDelete) | **Delete** /v1/slo/{id} | 
[**V1SloIdGet**](DefaultAPI.md#V1SloIdGet) | **Get** /v1/slo/{id} | 
[**V1SloIdPut**](DefaultAPI.md#V1SloIdPut) | **Put** /v1/slo/{id} | 
[**V1SloPost**](DefaultAPI.md#V1SloPost) | **Post** /v1/slo | 



## V1SloGet

> ApiSLOListResponse V1SloGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/slo-openapi-client/go/slo"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1SloGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SloGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SloGet`: ApiSLOListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1SloGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SloGetRequest struct via the builder pattern


### Return type

[**ApiSLOListResponse**](ApiSLOListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SloIdDelete

> V1SloIdDelete(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/slo-openapi-client/go/slo"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V1SloIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SloIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SloIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SloIdGet

> SloV00Slo V1SloIdGet(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/slo-openapi-client/go/slo"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1SloIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SloIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SloIdGet`: SloV00Slo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1SloIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SloIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SloV00Slo**](SloV00Slo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SloIdPut

> V1SloIdPut(ctx, id).SloV00Slo(sloV00Slo).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/slo-openapi-client/go/slo"
)

func main() {
	id := "id_example" // string | 
	sloV00Slo := *openapiclient.NewSloV00Slo("Description_example", "Name_example", []openapiclient.SloV00Objective{*openapiclient.NewSloV00Objective(float64(123), "Window_example")}, *openapiclient.NewSloV00Query("Type_example"), "Uuid_example") // SloV00Slo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V1SloIdPut(context.Background(), id).SloV00Slo(sloV00Slo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SloIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SloIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sloV00Slo** | [**SloV00Slo**](SloV00Slo.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SloPost

> ApiSLOCreateResponse V1SloPost(ctx).SloV00Slo(sloV00Slo).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/slo-openapi-client/go/slo"
)

func main() {
	sloV00Slo := *openapiclient.NewSloV00Slo("Description_example", "Name_example", []openapiclient.SloV00Objective{*openapiclient.NewSloV00Objective(float64(123), "Window_example")}, *openapiclient.NewSloV00Query("Type_example"), "Uuid_example") // SloV00Slo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1SloPost(context.Background()).SloV00Slo(sloV00Slo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SloPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SloPost`: ApiSLOCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1SloPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SloPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sloV00Slo** | [**SloV00Slo**](SloV00Slo.md) |  | 

### Return type

[**ApiSLOCreateResponse**](ApiSLOCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

