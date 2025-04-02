# ApiReportListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | [**[]ReportV1Report**](ReportV1Report.md) |  | 

## Methods

### NewApiReportListResponse

`func NewApiReportListResponse(reports []ReportV1Report, ) *ApiReportListResponse`

NewApiReportListResponse instantiates a new ApiReportListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportListResponseWithDefaults

`func NewApiReportListResponseWithDefaults() *ApiReportListResponse`

NewApiReportListResponseWithDefaults instantiates a new ApiReportListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *ApiReportListResponse) GetReports() []ReportV1Report`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *ApiReportListResponse) GetReportsOk() (*[]ReportV1Report, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *ApiReportListResponse) SetReports(v []ReportV1Report)`

SetReports sets Reports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


