package check

import (
	"flag"
	"fmt"
	"io"
	"major/internal/models"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func (ch *CheckHandler) CliHandler() {
	// Определяем флаги
	stdinFlag := flag.Bool("stdin", false, "read configuration from standard input")
	silentFlag := flag.Bool("s", false, "silent mode: do not exit with error if issues found")
	flag.BoolVar(silentFlag, "silent", false, "silent mode (alternative)")
	flag.Parse()

	var inputData []byte
	var err error
	var source string

	if *stdinFlag {
		inputData, err = io.ReadAll(os.Stdin)
		if err != nil {
			printError("Error reading stdin: %v\n", err)
			os.Exit(1)
		}
		source = "STDIN"
	} else {
		if flag.NArg() < 1 {
			printUsage()
			os.Exit(1)
		}
		filePath := flag.Arg(0)
		inputData, err = os.ReadFile(filePath)
		if err != nil {
			printError("Error reading file %s: %v\n", filePath, err)
			os.Exit(1)
		}
		source = filePath
	}

	result, errGetResult := ch.service.CheckService.GetConfigSummary(
		&models.CheckRequest{
			Config: string(inputData),
		},
	)

	if errGetResult != nil {
		printError("Error check config file: %v\n", errGetResult)
	}

	// Выводим результат
	exitCode := printResults(source, result.Rules, *silentFlag)

	if exitCode != 0 && !*silentFlag {
		os.Exit(exitCode)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-s|--silent] [--stdin] <config-file>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	fmt.Fprintf(os.Stderr, "  -s, --silent    Don't exit with error if issues found\n")
	fmt.Fprintf(os.Stderr, "  --stdin         Read configuration from standard input\n")
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  %s config.yaml\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --stdin < config.json\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  cat config.yml | %s --stdin\n", os.Args[0])
}

func printError(format string, args ...interface{}) {
	red := color.New(color.FgRed, color.Bold)
	red.Fprintf(os.Stderr, "ERROR: ")
	fmt.Fprintf(os.Stderr, format, args...)
}

func printSuccess(format string, args ...interface{}) {
	green := color.New(color.FgGreen, color.Bold)
	green.Fprintf(os.Stdout, "✓ ")
	fmt.Fprintf(os.Stdout, format, args...)
}

func printWarning(format string, args ...interface{}) {
	yellow := color.New(color.FgYellow, color.Bold)
	yellow.Fprintf(os.Stdout, "⚠ ")
	fmt.Fprintf(os.Stdout, format, args...)
}

func printResults(source string, rules []*models.RuleFullData, silent bool) int {
	// Получаем ширину терминала
	termWidth := getTerminalWidth()
	separator := strings.Repeat("━", termWidth)
	thinSeparator := strings.Repeat("─", termWidth)

	// Заголовок
	fmt.Println()
	header := color.New(color.FgCyan, color.Bold)
	header.Printf("%s\n", separator)
	header.Printf("%*s\n", (termWidth+len("  Security Configuration Scanner (major)"))/2, "  Security Configuration Scanner (major)")
	header.Printf("%s\n", separator)
	fmt.Printf("  Source: %s\n", source)
	fmt.Printf("  Issues found: %d\n", len(rules))
	fmt.Println()

	if len(rules) == 0 {
		printSuccess("No security issues found!\n")
		fmt.Println()
		return 0
	}

	// Группируем по severities
	severityOrder := map[string]int{"HIGH": 0, "MEDIUM": 1, "LOW": 2}
	grouped := make(map[string][]*models.RuleFullData)
	for _, rule := range rules {
		severityName := rule.Issue.Severity.Name
		grouped[severityName] = append(grouped[severityName], rule)
	}

	// Сортируем severities
	severities := make([]string, 0, len(grouped))
	for s := range grouped {
		severities = append(severities, s)
	}
	sort.Slice(severities, func(i, j int) bool {
		return severityOrder[severities[i]] < severityOrder[severities[j]]
	})

	// Выводим результаты
	for _, severityName := range severities {
		rulesList := grouped[severityName]

		// Цвет для severities
		var severityColor *color.Color
		switch severityName {
		case "HIGH":
			severityColor = color.New(color.FgRed, color.BgWhite, color.Bold)
		case "MEDIUM":
			severityColor = color.New(color.FgYellow, color.Bold)
		case "LOW":
			severityColor = color.New(color.FgBlue, color.Bold)
		default:
			severityColor = color.New(color.FgWhite, color.Bold)
		}

		severityColor.Printf("\n[%s] %d issue(s)\n", severityName, len(rulesList))
		fmt.Println(thinSeparator)

		for i, rule := range rulesList {
			// Номер проблемы
			fmt.Printf("\n  %d. %s\n", i+1, rule.Name)

			// Описание
			fmt.Printf("     %s\n", color.YellowString("Description:"))
			// Форматируем описание с переносом по словам
			fmt.Printf("       %s\n", wrapText(rule.Issue.Description, termWidth-10))

			// Рекомендация
			fmt.Printf("     %s\n", color.GreenString("Recommendation:"))
			fmt.Printf("       %s\n", wrapText(rule.Issue.Recommendation, termWidth-10))

			// Выражение
			if rule.Expression != "" {
				fmt.Printf("     %s\n", color.HiBlackString("Expression:"))
				fmt.Printf("       %s\n", wrapText(rule.Expression, termWidth-10))
			}

			if i < len(rulesList)-1 {
				fmt.Println()
			}
		}
		fmt.Println()
	}

	fmt.Println(separator)
	if silent {
		printWarning("Silent mode enabled, exiting with code 0\n")
		fmt.Println()
		return 0
	}

	red := color.New(color.FgRed, color.Bold)
	red.Printf("\n✗ Validation failed! Found %d security issue(s)\n", len(rules))
	fmt.Println()
	return 1
}

func getTerminalWidth() int {
	// Если вывод не в терминал (например, перенаправлен в файл), используем 80 колонок
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return 80
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // значение по умолчанию при ошибке
	}
	return width
}

func wrapText(text string, maxWidth int) string {
	if maxWidth <= 0 {
		maxWidth = 70
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		if len(currentLine)+1+len(word) <= maxWidth {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	lines = append(lines, currentLine)

	return strings.Join(lines, "\n       ")
}
