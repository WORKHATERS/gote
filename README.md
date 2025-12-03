# gote — Telegram Bot API для Go

**gote** — это современная, минималистичная и удобная библиотека для работы с [Telegram Bot API](https://core.telegram.org/bots/api) на языке **Go**.
Она предоставляет полный доступ ко всем возможностям Telegram: сообщения, медиа, клавиатуры, команды, callback-запросы, webhooks, inline-режим и многое другое.

---

## Установка

```bash
go get github.com/MOXHATKA/gote
```

---

## Быстрый старт

   ```go
   package main

   import (
   	"context"

   	"github.com/MOXHATKA/gote/pkg/core"
   	"github.com/MOXHATKA/gote/pkg/types"
   	"github.com/MOXHATKA/gote/pkg/updater"
   )

   func main() {
   	ctx, cancel := context.WithCancel(context.Background())
   	defer cancel()

   	bot := core.NewBot(ctx, "ВАШ_ТОКЕН_БОТА")

   	poller := updater.NewPoller(bot)
   	updates := poller.Start()

   	for u := range updates {
   		if u.Message != nil {
   			bot.SendMessage(ctx, types.SendMessage{
   				ChatId: u.Message.Chat.Id,
   				Text:   u.Message.Text,
   			})
   		}
   	}
   }
   ```

---

## Архитектура

Библиотека разделена на пакеты:

| Пакет          | Назначение                                                                                           |
| -------------- | ---------------------------------------------------------------------------------------------------- |
| `pkg/core`     | Основной объект `Bot`, методы API, отправка запросов, логирование.                                   |
| `pkg/updater`  | Механизм получения обновлений (polling или webhook).                                                 |
| `pkg/types`    | Типы данных, соответствующие Telegram Bot API (сообщения, медиа, чаты, пользователи, кнопки и т.д.). |
| `internal/env` | Утилита для загрузки конфигурации из `.env` файлов.                                                  |

---

## Основные возможности

- Отправка и редактирование сообщений
- Работа с inline и reply клавиатурами
- Callback-запросы и inline-режим
- Отправка фото, видео, документов и медиа-групп
- Webhook и long polling режимы
- Управление командами, чатами, пользователями
- Встроенная система логирования
- Полная типизация всех объектов Telegram API

---

## Принцип работы

1. **Создание объекта бота:**

   ```go
   bot := core.NewBot(ctx, token)
   ```

2. **Запуск получения обновлений:**

   ```go
   poller := updater.NewPoller(bot)
   updates := poller.Start()
   ```

3. **Обработка обновлений:**

   ```go
   for update := range updates {
       if update.Message != nil {
           bot.SendMessage(ctx, ...)
       }
       if update.CallbackQuery != nil {
           bot.AnswerCallbackQuery(ctx, ...)
       }
       // и так далее
   }
   ```

---

## Преимущества gote

* **Минимализм:** чистый и понятный API без избыточных абстракций
* **Гибкость:** легко интегрируется в любые проекты
* **Полный контроль:** доступ к каждому полю Telegram API
* **Расширяемость:** настраиваемый HTTP-клиент, логгер

---

## Примеры

Примеры доступны в папке [`examples/`](examples/)

---

## Лицензия

MIT License © 2025 MOXHATKA
