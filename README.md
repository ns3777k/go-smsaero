# go-smsaero

API-клиент для [SMS Aero](http://smsaero.ru/).

Все что могу сказать, это одно из худших API, которое я видел :trollface:

## Пример

```go
package main

import (
    "github.com/ns3777k/go-smsaero/smsaero"
    "log"
    "time"
)

func main() {
    client := smsaero.NewClient(nil, "my@email", "my_md5_hash")
    smsID, err := client.Send(79999999999, "проверка API", "NEWS")
    if err != nil {
        log.Fatalln(err)
    }
    
    for {
        status, err := client.GetStatus(smsID)
        if err != nil {
            log.Println(err)
            goto sleep
        }

        log.Println(status)
sleep:
        time.Sleep(time.Second * 15)
    }
}
```
