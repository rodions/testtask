# Реализация web сервиса для работы с заметками

Необходимо создать HTTP API для взаимодействия с сервисом заметок. 
Модель заметки описана ниже:

|Поле       |Тип    |Описание                                                        |
|-----------|:------|:---------------------------------------------------------------|
|id         |String|Строка в формате UUID                                            |
|title      |String|Заголовок заметки                                                |
|date_create|Int   |Время создания заметки в формате UNIX Time Stamp                 |
|date_update|Int   |Время последнего редактирования заметки в формате UNIX Time Stamp|

JSON-объект заметки будет выглядеть следующим образом:


    {
        "id": "10dcdccb-8876-4245-ac53-92900c6509bd",
        "title": "Текст заголовка",
        "text": "Текст заметки",
        "date_create": 1513759529,
        "date_update": 1517259529
    }

## Описание маршрутов

В рамках тестового задания нужно реализовать следующие маршруты:

1) **GET /notes/**  —  Используется для получения списка всех заметок.

Входные данные: Отсутствуют

Выходные данные: **200 Ok**

Массив объектов заметок (описание объекта заметки см. выше)



2) **POST /notes/**  — Используется для создания заметки

Входные данные:  body:


        {
        "title": "Тестовая заметка",
        "text": "Текст тестовой заметки"
        }

Выходные данные: **201 Created**


3) **GET /notes/&lt;note_id&gt;/** — Используется для получения информации о заметке по ее идентификатору

Входные данные: **note_id** — String в формате UUID - идентификатор заметки, отправляется в path

Выходные данные: **200 Ok**

Объект заметки (описание объекта заметки см. выше)


4) **PUT /notes/&lt;note_id&gt;/**  

Входные данные: **note_id** — String в формате UUID - идентификатор заметки, отправляется в path
body:


        {
        "title": "Тестовая заметка",
        "text": "Текст тестовой заметки"
        }


Выходные данные: **200 Ok**

Обновленный объект заметки (описание объекта заметки см. выше)
