# MarginAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPair** | **string** | Currency pair | [optional] 
**AccountType** | **string** | Account type, risk - risk rate account, mmr - maintenance margin rate account, inactive - market not activated | [optional] 
**Leverage** | **string** | User current market leverage multiple | [optional] 
**Locked** | **bool** | Whether account is locked | [optional] 
**Risk** | **string** | Leveraged Account Current Risk Rate (Returned when the Account is a Risk Rate Account) | [optional] 
**Mmr** | **string** | Leveraged Account Current Maintenance Margin Rate (returned when the Account is a Maintenance Margin Rate Account) | [optional] 
**Base** | [**MarginAccountCurrency**](MarginAccountCurrency.md) |  | [optional] 
**Quote** | [**MarginAccountCurrency**](MarginAccountCurrency.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


