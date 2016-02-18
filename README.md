# go-smsaero

[![Coverage Status](https://coveralls.io/repos/github/ns3777k/go-smsaero/badge.svg?branch=master)](https://coveralls.io/github/ns3777k/go-smsaero?branch=master)
[![Build Status](https://travis-ci.org/ns3777k/go-smsaero.svg?branch=master)](https://travis-ci.org/ns3777k/go-smsaero)

API-клиент для [SMS Aero](http://smsaero.ru/).

Отличительные особенности сервиса и API:
- Пока вы не заключили договор, все смс будут проходить ручную модерацию :trollface:
- Независимо от того, возникла ли какая-то ошибка при обработке запроса (неверный логин/пароль и т.п.) сервис всегда вернет код ответа `200`
- JSON-ответы не описаны в принципе
- Даже если мы указываем, что хотим получить обратно JSON, он отдается с заголовком `Content-Type: text/plain`


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
