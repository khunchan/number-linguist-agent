package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Teneo-Protocol/teneo-agent-sdk/pkg/agent"
	"github.com/joho/godotenv"
)

// Словарь переводов чисел (основные значения)
var numberTranslations = map[int]map[string]string{
	0: {
		"English":  "zero",
		"Spanish":  "cero",
		"French":   "zéro",
		"German":   "null",
		"Russian":  "ноль",
		"Chinese":  "零 (líng)",
		"Japanese": "零 (rei)",
		"Arabic":   "صفر (sifr)",
	},
	1: {
		"English":  "one",
		"Spanish":  "uno",
		"French":   "un",
		"German":   "eins",
		"Russian":  "один",
		"Chinese":  "一 (yī)",
		"Japanese": "一 (ichi)",
		"Arabic":   "واحد (wahid)",
	},
	2: {
		"English":  "two",
		"Spanish":  "dos",
		"French":   "deux",
		"German":   "zwei",
		"Russian":  "два",
		"Chinese":  "二 (èr)",
		"Japanese": "二 (ni)",
		"Arabic":   "اثنان (ithnan)",
	},
	3: {
		"English":  "three",
		"Spanish":  "tres",
		"French":   "trois",
		"German":   "drei",
		"Russian":  "три",
		"Chinese":  "三 (sān)",
		"Japanese": "三 (san)",
		"Arabic":   "ثلاثة (thalatha)",
	},
	4: {
		"English":  "four",
		"Spanish":  "cuatro",
		"French":   "quatre",
		"German":   "vier",
		"Russian":  "четыре",
		"Chinese":  "四 (sì)",
		"Japanese": "四 (shi/yon)",
		"Arabic":   "أربعة (arba'a)",
	},
	5: {
		"English":  "five",
		"Spanish":  "cinco",
		"French":   "cinq",
		"German":   "fünf",
		"Russian":  "пять",
		"Chinese":  "五 (wǔ)",
		"Japanese": "五 (go)",
		"Arabic":   "خمسة (khamsa)",
	},
	6: {
		"English":  "six",
		"Spanish":  "seis",
		"French":   "six",
		"German":   "sechs",
		"Russian":  "шесть",
		"Chinese":  "六 (liù)",
		"Japanese": "六 (roku)",
		"Arabic":   "ستة (sitta)",
	},
	7: {
		"English":  "seven",
		"Spanish":  "siete",
		"French":   "sept",
		"German":   "sieben",
		"Russian":  "семь",
		"Chinese":  "七 (qī)",
		"Japanese": "七 (shichi/nana)",
		"Arabic":   "سبعة (sab'a)",
	},
	8: {
		"English":  "eight",
		"Spanish":  "ocho",
		"French":   "huit",
		"German":   "acht",
		"Russian":  "восемь",
		"Chinese":  "八 (bā)",
		"Japanese": "八 (hachi)",
		"Arabic":   "ثمانية (thamaniya)",
	},
	9: {
		"English":  "nine",
		"Spanish":  "nueve",
		"French":   "neuf",
		"German":   "neun",
		"Russian":  "девять",
		"Chinese":  "九 (jiǔ)",
		"Japanese": "九 (kyū)",
		"Arabic":   "تسعة (tis'a)",
	},
	10: {
		"English":  "ten",
		"Spanish":  "diez",
		"French":   "dix",
		"German":   "zehn",
		"Russian":  "десять",
		"Chinese":  "十 (shí)",
		"Japanese": "十 (jū)",
		"Arabic":   "عشرة (ashara)",
	},
	20: {
		"English":  "twenty",
		"Spanish":  "veinte",
		"French":   "vingt",
		"German":   "zwanzig",
		"Russian":  "двадцать",
		"Chinese":  "二十 (èrshí)",
		"Japanese": "二十 (nijū)",
		"Arabic":   "عشرون (ishrun)",
	},
	30: {
		"English":  "thirty",
		"Spanish":  "treinta",
		"French":   "trente",
		"German":   "dreißig",
		"Russian":  "тридцать",
		"Chinese":  "三十 (sānshí)",
		"Japanese": "三十 (sanjū)",
		"Arabic":   "ثلاثون (thalathun)",
	},
	40: {
		"English":  "forty",
		"Spanish":  "cuarenta",
		"French":   "quarante",
		"German":   "vierzig",
		"Russian":  "сорок",
		"Chinese":  "四十 (sìshí)",
		"Japanese": "四十 (yonjū)",
		"Arabic":   "أربعون (arba'un)",
	},
	50: {
		"English":  "fifty",
		"Spanish":  "cincuenta",
		"French":   "cinquante",
		"German":   "fünfzig",
		"Russian":  "пятьдесят",
		"Chinese":  "五十 (wǔshí)",
		"Japanese": "五十 (gojū)",
		"Arabic":   "خمسون (khamsun)",
	},
	60: {
		"English":  "sixty",
		"Spanish":  "sesenta",
		"French":   "soixante",
		"German":   "sechzig",
		"Russian":  "шестьдесят",
		"Chinese":  "六十 (liùshí)",
		"Japanese": "六十 (rokujū)",
		"Arabic":   "ستون (sittun)",
	},
	70: {
		"English":  "seventy",
		"Spanish":  "setenta",
		"French":   "soixante-dix",
		"German":   "siebzig",
		"Russian":  "семьдесят",
		"Chinese":  "七十 (qīshí)",
		"Japanese": "七十 (nanajū)",
		"Arabic":   "سبعون (sab'un)",
	},
	80: {
		"English":  "eighty",
		"Spanish":  "ochenta",
		"French":   "quatre-vingts",
		"German":   "achtzig",
		"Russian":  "восемьдесят",
		"Chinese":  "八十 (bāshí)",
		"Japanese": "八十 (hachijū)",
		"Arabic":   "ثمانون (thamanun)",
	},
	90: {
		"English":  "ninety",
		"Spanish":  "noventa",
		"French":   "quatre-vingt-dix",
		"German":   "neunzig",
		"Russian":  "девяносто",
		"Chinese":  "九十 (jiǔshí)",
		"Japanese": "九十 (kyūjū)",
		"Arabic":   "تسعون (tis'un)",
	},
	100: {
		"English":  "one hundred",
		"Spanish":  "cien",
		"French":   "cent",
		"German":   "hundert",
		"Russian":  "сто",
		"Chinese":  "百 (bǎi)",
		"Japanese": "百 (hyaku)",
		"Arabic":   "مئة (mi'a)",
	},
}

type NumberLinguistEduAgent struct{}

func (a *NumberLinguistEduAgent) ProcessTask(ctx context.Context, task string) (string, error) {
	log.Printf("📥 Received task: %s", task)

	// Очищаем входные данные
	task = strings.TrimSpace(task)
	task = strings.TrimPrefix(task, "/")
	taskLower := strings.ToLower(task)

	// Разбираем команду
	parts := strings.Fields(taskLower)
	if len(parts) == 0 {
		return "❌ No command provided. Try: translate <number>, random, or help", nil
	}

	command := parts[0]

	switch command {
	case "translate":
		if len(parts) < 2 {
			return "❌ Please provide a number! Example: translate 42", nil
		}

		// Парсим число
		numStr := parts[1]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return fmt.Sprintf("❌ '%s' is not a valid number!", numStr), nil
		}

		if num < 0 || num > 100 {
			return "❌ Please provide a number between 0 and 100", nil
		}

		return formatNumberTranslation(num), nil

	case "random":
		// Генерируем случайное число
		num := rand.Intn(101) // 0-100
		return fmt.Sprintf("🎲 Random number: %d\n\n%s", num, formatNumberTranslation(num)), nil

	case "help":
		return `📚 **NumberLinguist Help**

Commands:
• translate <number> - Translate number (0-100) to 8 languages
• random - Get random number translated
• help - Show this message

Example: translate 42

Supported languages: English, Spanish, French, German, Russian, Chinese, Japanese, Arabic`, nil

	default:
		// Попробуем распознать число без команды
		num, err := strconv.Atoi(command)
		if err == nil && num >= 0 && num <= 100 {
			return formatNumberTranslation(num), nil
		}

		return fmt.Sprintf("❓ Unknown command: '%s'. Try: translate <number>, random, or help", command), nil
	}
}

// Форматирует перевод числа на разные языки
func formatNumberTranslation(num int) string {
	translations, exists := numberTranslations[num]

	if !exists {
		// Для чисел без точного словаря пытаемся построить из составных
		translations = buildCompositeTranslation(num)
	}

	result := fmt.Sprintf("🔢 **Number: %d**\n\n", num)

	// Список языков в порядке отображения
	languages := []string{"English", "Spanish", "French", "German", "Russian", "Chinese", "Japanese", "Arabic"}

	for _, lang := range languages {
		if trans, ok := translations[lang]; ok {
			result += fmt.Sprintf("🌍 %s: **%s**\n", lang, trans)
		}
	}

	return result
}

// Строит составной перевод для чисел без точного словаря
func buildCompositeTranslation(num int) map[string]string {
	if num >= 11 && num <= 19 {
		// Числа 11-19 (особые случаи)
		return map[string]string{
			"English":  numberWord(num, "en"),
			"Spanish":  numberWord(num, "es"),
			"French":   numberWord(num, "fr"),
			"German":   numberWord(num, "de"),
			"Russian":  numberWord(num, "ru"),
			"Chinese":  numberWord(num, "zh"),
			"Japanese": numberWord(num, "ja"),
			"Arabic":   numberWord(num, "ar"),
		}
	}

	if num >= 21 && num <= 99 {
		// Составные числа 21-99
		tens := (num / 10) * 10
		ones := num % 10

		tensTranslations := numberTranslations[tens]
		onesTranslations := numberTranslations[ones]

		return map[string]string{
			"English":  fmt.Sprintf("%s-%s", tensTranslations["English"], onesTranslations["English"]),
			"Spanish":  compositeSpanish(tens, ones),
			"French":   compositeFrench(tens, ones),
			"German":   compositeGerman(tens, ones),
			"Russian":  compositeRussian(tens, ones),
			"Chinese":  fmt.Sprintf("%s%s", tensTranslations["Chinese"], onesTranslations["Chinese"]),
			"Japanese": fmt.Sprintf("%s%s", tensTranslations["Japanese"], onesTranslations["Japanese"]),
			"Arabic":   compositeArabic(tens, ones),
		}
	}

	// Fallback для остальных чисел
	return map[string]string{
		"English":  fmt.Sprintf("%d", num),
		"Spanish":  fmt.Sprintf("%d", num),
		"French":   fmt.Sprintf("%d", num),
		"German":   fmt.Sprintf("%d", num),
		"Russian":  fmt.Sprintf("%d", num),
		"Chinese":  fmt.Sprintf("%d", num),
		"Japanese": fmt.Sprintf("%d", num),
		"Arabic":   fmt.Sprintf("%d", num),
	}
}

// Вспомогательные функции для составных чисел

func compositeSpanish(tens, ones int) string {
	if tens == 20 {
		return "veinti" + numberTranslations[ones]["Spanish"]
	}
	return numberTranslations[tens]["Spanish"] + " y " + numberTranslations[ones]["Spanish"]
}

func compositeFrench(tens, ones int) string {
	if tens == 70 || tens == 90 {
		base := tens - 10
		return numberTranslations[base]["French"] + "-" + numberTranslations[10+ones]["French"]
	}
	return numberTranslations[tens]["French"] + "-" + numberTranslations[ones]["French"]
}

func compositeGerman(tens, ones int) string {
	return numberTranslations[ones]["German"] + "und" + numberTranslations[tens]["German"]
}

func compositeRussian(tens, ones int) string {
	return numberTranslations[tens]["Russian"] + " " + numberTranslations[ones]["Russian"]
}

func compositeArabic(tens, ones int) string {
	return numberTranslations[tens]["Arabic"] + " و " + numberTranslations[ones]["Arabic"]
}

func numberWord(num int, lang string) string {
	// Упрощенная версия для чисел 11-19
	words := map[string]map[int]string{
		"en": {
			11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
			16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
		},
		"es": {
			11: "once", 12: "doce", 13: "trece", 14: "catorce", 15: "quince",
			16: "dieciséis", 17: "diecisiete", 18: "dieciocho", 19: "diecinueve",
		},
		"fr": {
			11: "onze", 12: "douze", 13: "treize", 14: "quatorze", 15: "quinze",
			16: "seize", 17: "dix-sept", 18: "dix-huit", 19: "dix-neuf",
		},
		"de": {
			11: "elf", 12: "zwölf", 13: "dreizehn", 14: "vierzehn", 15: "fünfzehn",
			16: "sechzehn", 17: "siebzehn", 18: "achtzehn", 19: "neunzehn",
		},
		"ru": {
			11: "одиннадцать", 12: "двенадцать", 13: "тринадцать", 14: "четырнадцать", 15: "пятнадцать",
			16: "шестнадцать", 17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать",
		},
		"zh": {
			11: "十一 (shíyī)", 12: "十二 (shí'èr)", 13: "十三 (shísān)", 14: "十四 (shísì)", 15: "十五 (shíwǔ)",
			16: "十六 (shíliù)", 17: "十七 (shíqī)", 18: "十八 (shíbā)", 19: "十九 (shíjiǔ)",
		},
		"ja": {
			11: "十一 (jūichi)", 12: "十二 (jūni)", 13: "十三 (jūsan)", 14: "十四 (jūshi)", 15: "十五 (jūgo)",
			16: "十六 (jūroku)", 17: "十七 (jūshichi)", 18: "十八 (jūhachi)", 19: "十九 (jūkyū)",
		},
		"ar": {
			11: "أحد عشر (ahad ashar)", 12: "اثنا عشر (ithna ashar)", 13: "ثلاثة عشر (thalatha ashar)",
			14: "أربعة عشر (arba'a ashar)", 15: "خمسة عشر (khamsa ashar)", 16: "ستة عشر (sitta ashar)",
			17: "سبعة عشر (sab'a ashar)", 18: "ثمانية عشر (thamaniya ashar)", 19: "تسعة عشر (tis'a ashar)",
		},
	}

	if word, ok := words[lang][num]; ok {
		return word
	}
	return fmt.Sprintf("%d", num)
}

func main() {
	// Инициализация random seed
	rand.Seed(time.Now().UnixNano())

	// Загружаем переменные окружения
	godotenv.Load()

	// Конфигурация агента
	config := agent.DefaultConfig()

	config.Name = "NumberLinguist"
	config.Description = "Educational agent that translates numbers (0-100) into multiple world languages"
	config.Capabilities = []string{"number_translation", "language_learning"}

	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.NFTTokenID = os.Getenv("NFT_TOKEN_ID")
	config.OwnerAddress = os.Getenv("OWNER_ADDRESS")

	// Оптимизация
	config.RateLimitPerMinute = 0
	config.MaxConcurrentTasks = 100
	config.TaskTimeout = 30

	// Создаём агента
	enhancedAgent, err := agent.NewEnhancedAgent(&agent.EnhancedAgentConfig{
		Config:       config,
		AgentHandler: &NumberLinguistEduAgent{},
	})

	if err != nil {
		log.Fatalf("❌ Failed to create agent: %v", err)
	}

	log.Println("🚀 NumberLinguist Agent Started!")
	log.Println("📚 Commands: translate <number>, random, help")
	log.Println("🌍 Supporting 8 world languages")
	log.Printf("🔖 NFT Token ID: %s", os.Getenv("NFT_TOKEN_ID"))

	// Запуск
	enhancedAgent.Run()
}
