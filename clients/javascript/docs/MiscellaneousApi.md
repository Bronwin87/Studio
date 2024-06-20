# StudioAiEmpowerLabs.MiscellaneousApi

All URIs are relative to *https://studio.aiempowerlabs.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getConfiguration**](MiscellaneousApi.md#getConfiguration) | **GET** /info | 



## getConfiguration

> [{String: [String]}] getConfiguration()



Get configuration information

### Example

```javascript
import StudioAiEmpowerLabs from 'studio_ai_empower_labs';

let apiInstance = new StudioAiEmpowerLabs.MiscellaneousApi();
apiInstance.getConfiguration((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

**[{String: [String]}]**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json
