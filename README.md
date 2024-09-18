# Building
Building - сервис для хранения данных об недвижимости

## Основные методы
Запись по 'POST' запросу из json 

Получения данных из БД с фильтрацией по городу, году и количеству этажей

## Установка 
```
git clone https://github.com/evilgooby/Buildings.git
```

## Использование:
### Пример запроса:
POST:
``` JSON
curl -X POST -H "Content-Type: application/json" -d '{
 "name":"ЖК Москва",
 "city": "Москва",
 "year": 2020,
  "floors": 9
}' http://localhost:8080/CreateBuilding
```
GET:
```
curl http://localhost:8080/GetBuildings?city=Москва&year=2022&floors=9
```

### Пример ответа:
POST:
``` JSON
{
"message": "Building created successfully"
}
``` 
GET:
``` JSON
[
  {
    "id": 7,
    "name": "danila",
    "city": "Москва",
    "year": 2022,
    "floors": 9
  }
]
```


