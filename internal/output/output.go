package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// ErrorInfo represents an error in responses
type ErrorInfo struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Hint    string   `json:"hint,omitempty"`
	Usage   []string `json:"usage,omitempty"`
}

// ErrorResponse is a standard error response
type ErrorResponse struct {
	Success bool       `json:"success"`
	Error   *ErrorInfo `json:"error"`
}

// SuccessResponse is a standard success response
type SuccessResponse struct {
	Success   bool   `json:"success"`
	Operation string `json:"operation,omitempty"`
	Message   string `json:"message,omitempty"`
}

// JSON outputs data as formatted JSON to stdout
func JSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// JSONString returns data as a formatted JSON string
func JSONString(data interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Human outputs human-readable text to stdout
func Human(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// HumanLn outputs human-readable text with newline
func HumanLn(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// Error outputs an error response
func Error(code, message string) error {
	resp := ErrorResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}
	return JSON(resp)
}

// ErrorWithHint outputs an error response with guidance for agents
func ErrorWithHint(code, message, hint string, usage ...string) error {
	resp := ErrorResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Hint:    hint,
			Usage:   usage,
		},
	}
	return JSON(resp)
}

// ErrorHuman outputs a human-readable error
func ErrorHuman(message string) {
	color.Red("Error: %s", message)
	fmt.Println()
}

// ErrorHumanWithHint outputs a human-readable error with guidance
func ErrorHumanWithHint(message, hint string, usage ...string) {
	color.Red("Error: %s", message)
	fmt.Println()
	if hint != "" {
		fmt.Printf("\n%s\n", hint)
	}
	if len(usage) > 0 {
		fmt.Println("\nExamples:")
		for _, u := range usage {
			fmt.Printf("  %s\n", u)
		}
	}
	fmt.Println()
}

// Success outputs a success response
func Success(operation, message string) error {
	resp := SuccessResponse{
		Success:   true,
		Operation: operation,
		Message:   message,
	}
	return JSON(resp)
}

// SuccessHuman outputs a human-readable success message
func SuccessHuman(message string) {
	color.Green("✓ %s", message)
	fmt.Println()
}

// Table outputs data in table format
func Table(headers []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("  ")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(rows)
	table.Render()
}

// TableWithColors outputs a table with colored headers
func TableWithColors(headers []string, rows [][]string) {
	// Color the headers
	coloredHeaders := make([]string, len(headers))
	for i, h := range headers {
		coloredHeaders[i] = color.New(color.Bold).Sprint(h)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(coloredHeaders)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("  ")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(rows)
	table.Render()
}

// Section outputs a section header
func Section(title string) {
	color.Cyan(title)
}

// Muted outputs muted/dimmed text
func Muted(format string, args ...interface{}) string {
	return color.New(color.Faint).Sprintf(format, args...)
}

// Bold outputs bold text
func Bold(format string, args ...interface{}) string {
	return color.New(color.Bold).Sprintf(format, args...)
}

// Yellow outputs yellow text
func Yellow(format string, args ...interface{}) string {
	return color.YellowString(format, args...)
}

// Green outputs green text
func Green(format string, args ...interface{}) string {
	return color.GreenString(format, args...)
}

// Red outputs red text
func Red(format string, args ...interface{}) string {
	return color.RedString(format, args...)
}

// Cyan outputs cyan text
func Cyan(format string, args ...interface{}) string {
	return color.CyanString(format, args...)
}

// KeyValue outputs a key-value pair for human output
func KeyValue(key, value string) {
	fmt.Printf("  %s: %s\n", color.New(color.Faint).Sprint(key), value)
}

// Divider outputs a divider line
func Divider() {
	fmt.Println(strings.Repeat("-", 40))
}
