# goctp

Golang bindings for CTP.



## Fix some issues



### ThostFtdcUserApiDataType.h

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



### Library file configuration

```bash
mv ./libthostmduserapi.so /usr/lib/libthostmduserapi.so
mv ./libthosttraderapi.so /usr/lib/libthosttraderapi.so

vim /etc/ld.so.conf
/usr/lib/
ldconfig

# test
gcc -lthostmduserapi --verbose
gcc -lthosttraderapi --verbose
```
