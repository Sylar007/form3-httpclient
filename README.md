# **Home Exercise - Form3 Accounts API**
The goal of this exercise is to write a client library in Go to access Form3 fake account API.


## Name: Mohd Diah A.Karim
Most of my career in software development, i'm using C# for backend and Angular for frontend development. I'm only have a few months in commercial experience working with Go and couple of months learning the Go language.

In the home exercise, I'm creating a custom HTTP client to extends a common functionalities from standard library net/http. The custom HTTP client have a few generic methods that created based on the home exercise requirements such as:-

- Get (Fetch)
- Post (Create)
- Delete

Apart from that, i'm also added a few additional methods for setting-up some configuration for example SetHeaders, SetConnectionTimeout, SetResponseTimeout, SetMaxIdleConnections, SetHttpClient and Client Builder. Apart from that i'm using a singleton instance creation for this http client initialisation so that the connection can be re-used between http requests for efficient resource usage. 

After completed custom Http client library creation, i added  a few other fields in accounts model.
This file location is at `accounts/accountsModels.go` Then i proceed to create a unit test file for each functions. I started with drafting some scenarios follows by creating a functions to handle the operation to call the fake account API. 

Below are the functions created
- checkHealth - to get fake account service health (service is up/down)
- createAccount - to create a new organisation account 
- fetchAccount - to fetch a organisation account based on supplied Id 
- removeAccount - to delete a organisation account based on supplied Id 
- config- as a common/helper function

The Form3-API documentation is really helpful to understand the return json object, the expected error message and the response status code.
Here are the Test Scenario for each of the operations:-
- **createAccount**

**TestCreateAccount_WhenArrayOfNamesMoreThan4_ReturnValidationErrorMessage**
_When requesting to create an account to fake Account API, the test case for attribute names having more 5 items which will break the validation._
**TestCreateAccount_WhenMissingRequiredData_ReturnValidationErrorMessage**
_When requesting to create an account to fake Account API, the test case for attribute country is not supplied which will break the validation._
**TestCreateAccount_WhenValidData_ReturnAccountCreate**
_When requesting to create an account to fake Account API, the test case is all data supplied is valid which return account created._

- **fetchAccount**

**TestFetchAccount_WhenNonExistAccountId_ReturnMessageRecordNotExist**
_When requesting to fetch an account to fake Account API, the test case for account Id is not exist which will return message record does not exist._
**TestFetchAccount_WhenValidAccountId_ReturnAccountData**
_When requesting to create an account to fake Account API, the test case for account Id is exist  which will return the account data._
**TestFetchAccount_WhenInvalidFormatAccountId_ReturnMessageIdInvalidFormat**
_Not able to test due to some UUID must parse has been guarded._

- **deleteAccount**

**TestDeleteAccount_WhenNonExistAccountId_ReturnsNotFoundStatus**
_When requesting to delete an account to fake Account API, the test case for account Id is not exist which will return Not Found response status._
**TestDeleteAccount_WhenInvalidVersion_ReturnMessageInvalidVersion**
_When requesting to delete an account to fake Account API, the test case for version is not valid which will return message invalid version._
**TestDeleteAccount_WhenAccountIdExist_ReturnsNoContentStatus**
_When requesting to delete an account to fake Account API, the test case for account Id is exist which will return account data with response status no content._

How to run test 
- Run `docker compose up` from the terminal

What can be improved from this repo
- Implement async-await for calling API service from this http client
I haven't got a chance to implement it but gone through some documentation that it can be achieved by using channel. Looking forward to implement it for this repo improvment

- Implement retry mechanism for calling API service from this http client
In .NET(C#) usually we are using Polly for retry mechanism due to handling unstable newtork connection between http client to the API service. Due to lack of time and experience i'm not managed to implement it for now.  Looking forward to implement it later

- Handling concurrent request
In .NET(C#) we also using Polly bulk head or SemaphoreSlim for handling the overwhelm request to the API service. Looking forward find implementation/appoach in Go

*** The information and requirements from the Form3 Home Exercise page is easy to understand with a clear instructions 

