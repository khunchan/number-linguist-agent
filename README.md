# 🌍 NumberLinguist - Teneo Protocol Educational Agent

Educational agent that translates numbers (0-100) into 8 world languages.

## ✨ Features

Translates numbers 0-100 into 8 languages:
- 🇬🇧 English
- 🇪🇸 Spanish  
- 🇫🇷 French
- 🇩🇪 German
- 🇷🇺 Russian
- 🇨🇳 Chinese (Mandarin)
- 🇯🇵 Japanese
- 🇸🇦 Arabic

- ⚡ Instant responses (no AI latency)
- 💰 Minimal pricing: 0.00001 USDC per request
- 🔄 All fees return to owner

## 🎮 Commands

- `translate <number>` - Translate a number (0-100)
- `random` - Get a random number translation
- `help` - Show available commands

## 📱 NFT Details

- **NFT Token ID**: #735
- **Contract**: 0xd8493cc411D5d0da58dd7d6C0A22baEA9fbb3e5
- **Network**: PEAQ
- **Owner**: 0xD1FC8C4b5Df3390ccb7669C2f401186F23a9C770

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- PEAQ network access
- Teneo Protocol Agent NFT

### Installation

1. Clone the repository:
```bash
git clone https://github.com/khunchan/number-linguist-agent.git
cd number-linguist-agent
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment:
```bash
cp .env.example .env
# Edit .env with your credentials
```

4. Run the agent:
```bash
go run main.go
```

## ⚙️ Environment Variables
```env
PRIVATE_KEY=your_private_key_without_0x
NFT_TOKEN_ID=735
OWNER_ADDRESS=0xYourAddress
RATE_LIMIT_PER_MINUTE=0
```

⚠️ **SECURITY**: Never commit your `.env` file with real private keys!

## 🏗️ Architecture

- **Language**: Go
- **Framework**: Teneo Protocol Agent SDK
- **Type**: Command-Based Agent
- **Pricing**: Fixed rate (0.00001 USDC)
- **Categories**: Education, Language Learning

## 💬 Usage in Teneo Chatroom
```
@NumberLinguist translate 42
@NumberLinguist random
@NumberLinguist help
```

## 🛠️ Development

Built during Teneo Protocol Agent deployment on February 9, 2026.

### Technical Decisions

- No external APIs (instant responses)
- Dictionary-based translations (reliable)
- Composite number building (11-99 range)
- Minimal dependencies

## 📄 License

MIT

## 🔗 Links

- [Teneo Protocol](https://teneo-protocol.ai/)
- [Agent Console](https://agent-console.ai/)
- [Teneo SDK](https://github.com/TeneoProtocolAI/teneo-agent-sdk)
