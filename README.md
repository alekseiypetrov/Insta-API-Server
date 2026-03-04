# Алгоритмы и методы распределенных систем

## Тема: Мини-соцсеть

Веб-приложение с логикой как у Instagram (но урезанная).

#### Функциональные требования
- архитектура: **микросервисная**
- фреймворк: Gin / Fiber (Go)
- авторизация: JWT (golang-jwt/jwt) + bcypt
- хранилище: MongoDB
- контейнеризация: Docker 
- логирование

Состоит из следующих микросервисов:

- User Service
- Post Service
- Feed Service

## User Service

Отвечает за:

- регистрацию    
- аутентификацию
- профиль
- подписки (follow / unfollow)
- список подписок

### Запросы

- POST /users
- POST /auth/login
- GET  /users/{id}
- PUT /users/me/avatar  [Authorization: Bearer **jwt**]
- POST /users/{id}/follow
- GET /users/{id}/following

## Post Service

Отвечает за:

- создание поста
- получение поста
- список постов пользователя
- лайк под постом

### Запросы

- POST /posts
- GET /posts/{id}
- GET /users/{id}/posts
- POST /posts/{id}/like 

## Feed Service

Отвечает за:

- формирование ленты пользователя
- агрегацию постов подписок

### Запросы

- GET /feed/me
Authorization: Bearer <jwt_token>

