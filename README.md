# 🌍 NumberLinguist Agent

Educational agent that translates numbers (0-100) into 8 world languages!

## 📋 NFT Details
- **Token ID:** #735
- **Contract:** 0xd8493cc411D5d0da58dd7d6C0A22baEA9fbb3e5
- **Network:** Peaq Network

## 🚀 Quick Start

### 1. Настройте .env файл

```bash
# Откройте .env и заполните:
PRIVATE_KEY=ваш_приватный_ключ_без_0x
NFT_TOKEN_ID=735
OWNER_ADDRESS=0xВаш_Адрес_Кошелька
```

### 2. Установите зависимости

```bash
go mod tidy
```

### 3. Запустите агента

```bash
go run main.go
```

Вы увидите:
```
🚀 NumberLinguist Agent Started!
📚 Commands: translate <number>, random, help
🌍 Supporting 8 world languages
🔖 NFT Token ID: 735
```

## 💬 Использование в чате

Откройте [dashboard.teneo.pro](https://dashboard.teneo.pro) и используйте команды:

```
@NumberLinguist translate 42
@NumberLinguist random
@NumberLinguist 7
@NumberLinguist help
```

## 🎯 Для выполнения задания (100 запросов)

Просто отправляйте числа от 0 до 100:

```
@NumberLinguist 1
@NumberLinguist 2
@NumberLinguist 3
... (до 100)
```

Или используйте команды:
```
@NumberLinguist translate 1
@NumberLinguist translate 2
@NumberLinguist random
... (100 раз)
```

## 🌍 Поддерживаемые языки

1. English
2. Spanish
3. French
4. German
5. Russian
6. Chinese
7. Japanese
8. Arabic

## 📚 Команды

- `translate <number>` - Перевести число на 8 языков
- `random` - Случайное число с переводом
- `help` - Показать справку

## ✅ Чеклист запуска

- [ ] Заполнить .env (PRIVATE_KEY, OWNER_ADDRESS)
- [ ] Убедиться что NFT_TOKEN_ID=735
- [ ] Запустить `go mod tidy`
- [ ] Запустить `go run main.go`
- [ ] Проверить что агент онлайн
- [ ] Протестировать команды в чате
- [ ] Выполнить 100 запросов для награды

## 💰 Награда

После 100 запросов вы получите **300,000 баллов**!

Каждый запрос стоит ~$0.00001, итого ~$0.001 за все 100.

---

Built with [Teneo Agent SDK](https://github.com/TeneoProtocolAI/teneo-agent-sdk)
