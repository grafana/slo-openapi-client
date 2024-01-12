# Objective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float64** |  | 
**Window** | **string** |  | 

## Methods

### NewObjective

`func NewObjective(value float64, window string, ) *Objective`

NewObjective instantiates a new Objective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectiveWithDefaults

`func NewObjectiveWithDefaults() *Objective`

NewObjectiveWithDefaults instantiates a new Objective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Objective) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Objective) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Objective) SetValue(v float64)`

SetValue sets Value field to given value.


### GetWindow

`func (o *Objective) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *Objective) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *Objective) SetWindow(v string)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


