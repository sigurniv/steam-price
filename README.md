### Установка
```sh
$ go get -u github.com/sigurniv/steam-price
```

### Запуск
```sh
# Запуск всех сервисов на kubernetes кластере
$ cd $GOPATH/src/github.com/sigurniv/steam-price
$ make deploy

# Запуск отдельного сервиса на kubernetes кластере
$ cd $GOPATH/src/github.com/sigurniv/steam-price/service/currency_rate
$ make deploy
```

### Структура
`steam-price/service/currency_rate` - сервис курсов валют

`steam-price/service/steam_game` - сервис информации по игре steam

`steam-price/service/steam_game_price` - сервис получения игры в заданной валюте

### Тестирование
#### Сервис курсов валют
Получение списка валют
  ```sh     
curl -X GET \
  http://192.168.99.100:31001/currency/list \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```

Добавление пары
```sh    
curl -X POST \
  http://192.168.99.100:31001/currency/pair \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 6fb2538a-c2d7-bb41-9b0e-e99615f0a9fb' \
  -d '{
	"pair": "RUB_USD"
}'
```

Получение пары валют
```sh
curl -X GET \
  http://192.168.99.100:31001/currency/pair/RUB_USD/rate \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```

#### Cервис информации по игре steam
Поиск игр по имени
```sh
curl -X GET \
  http://192.168.99.100:31002/game/search/Counter \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```

Получение информации по игре
```sh
curl -X GET \
  http://192.168.99.100:31002/game/10 \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```

#### Cервис получения игры в заданной валюте
Получение информации по игре в заданной валюте
```sh
curl -X GET \
  http://192.168.99.100:31003/game/10/RUB \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```