# **Danger!!! Never use in a production environment, only for entertainment**



## Fix some issues

### ThostFtdcUserApiDataType.h**(v6.5.1 finished)**

```c
#define THOST_FTDC_VTC_BankBankToFuture "102001"
#define THOST_FTDC_VTC_BankFutureToBank "102002"
#define THOST_FTDC_VTC_FutureBankToFuture "202001"
#define THOST_FTDC_VTC_FutureFutureToBank "202002"

#define THOST_FTDC_FTC_BankLaunchBankToBroker "102001"
#define THOST_FTDC_FTC_BrokerLaunchBankToBroker "202001"
#define THOST_FTDC_FTC_BankLaunchBrokerToBank "102002"
#define THOST_FTDC_FTC_BrokerLaunchBrokerToBank "202002"
```

### ThostFtdcMdApi.h**(v6.5.1 finished)**

```cpp
class CThostFtdcMdSpi
{
public:
    virtual ~CThostFtdcMdSpi(){};
    ...
```

### ThostFtdcTraderApi.h**(v6.5.1 finished)**

```cpp
class CThostFtdcTraderSpi
{
public:
    virtual ~CThostFtdcTraderSpi(){};
    ...
```
