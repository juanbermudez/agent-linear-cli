package display

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// Priority icons for different priority levels
var PriorityIcons = map[int]string{
	0: "---",   // No priority
	1: "⚠⚠⚠",  // Urgent
	2: "▄▆█",   // High
	3: "▄▆ ",   // Medium
	4: "▄  ",   // Low
}

// PriorityNames for human-readable output
var PriorityNames = map[int]string{
	0: "No priority",
	1: "Urgent",
	2: "High",
	3: "Medium",
	4: "Low",
}

// PriorityIcon returns the icon for a priority level
func PriorityIcon(priority int) string {
	if icon, ok := PriorityIcons[priority]; ok {
		return icon
	}
	return "---"
}

// PriorityName returns the name for a priority level
func PriorityName(priority int) string {
	if name, ok := PriorityNames[priority]; ok {
		return name
	}
	return "Unknown"
}

// TimeAgo returns a human-readable relative time string
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case diff < 7*24*time.Hour:
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	case diff < 30*24*time.Hour:
		weeks := int(diff.Hours() / (24 * 7))
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	case diff < 365*24*time.Hour:
		months := int(diff.Hours() / (24 * 30))
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	default:
		years := int(diff.Hours() / (24 * 365))
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

// TimeAgoShort returns a short relative time string
func TimeAgoShort(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "now"
	case diff < time.Hour:
		return fmt.Sprintf("%dm", int(diff.Minutes()))
	case diff < 24*time.Hour:
		return fmt.Sprintf("%dh", int(diff.Hours()))
	case diff < 7*24*time.Hour:
		return fmt.Sprintf("%dd", int(diff.Hours()/24))
	case diff < 30*24*time.Hour:
		return fmt.Sprintf("%dw", int(diff.Hours()/(24*7)))
	case diff < 365*24*time.Hour:
		return fmt.Sprintf("%dmo", int(diff.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%dy", int(diff.Hours()/(24*365)))
	}
}

// FormatDate formats a time as a date string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats a time as a datetime string
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

// FormatISO formats a time as ISO 8601
func FormatISO(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseISO parses an ISO 8601 string
func ParseISO(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// Truncate truncates a string to maxLen, adding "..." if truncated
// Handles unicode properly
func Truncate(s string, maxLen int) string {
	if maxLen <= 3 {
		return s
	}

	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}

	runes := []rune(s)
	return string(runes[:maxLen-3]) + "..."
}

// TruncateMiddle truncates a string in the middle, preserving start and end
func TruncateMiddle(s string, maxLen int) string {
	if maxLen <= 5 {
		return Truncate(s, maxLen)
	}

	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}

	half := (maxLen - 3) / 2
	return string(runes[:half]) + "..." + string(runes[len(runes)-half:])
}

// Pad pads a string to the specified width
// Handles unicode properly
func Pad(s string, width int) string {
	runeCount := utf8.RuneCountInString(s)
	if runeCount >= width {
		return s
	}
	return s + strings.Repeat(" ", width-runeCount)
}

// PadLeft pads a string on the left to the specified width
func PadLeft(s string, width int) string {
	runeCount := utf8.RuneCountInString(s)
	if runeCount >= width {
		return s
	}
	return strings.Repeat(" ", width-runeCount) + s
}

// Initials returns the initials from a name (max 2 chars)
func Initials(name string) string {
	if name == "" {
		return "--"
	}

	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "--"
	}

	if len(parts) == 1 {
		runes := []rune(parts[0])
		if len(runes) >= 2 {
			return strings.ToUpper(string(runes[:2]))
		}
		return strings.ToUpper(parts[0])
	}

	// First letter of first and last name
	first := []rune(parts[0])
	last := []rune(parts[len(parts)-1])
	return strings.ToUpper(string(first[0]) + string(last[0]))
}

// BoolToYesNo converts a boolean to "Yes" or "No"
func BoolToYesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// BoolToCheckmark converts a boolean to "✓" or ""
func BoolToCheckmark(b bool) string {
	if b {
		return "✓"
	}
	return ""
}

// StatusIcon returns a status indicator icon
func StatusIcon(status string) string {
	switch strings.ToLower(status) {
	case "triage":
		return "◇"
	case "backlog":
		return "◌"
	case "unstarted", "todo":
		return "○"
	case "started", "in progress":
		return "◐"
	case "completed", "done":
		return "●"
	case "canceled", "cancelled":
		return "⊘"
	default:
		return "○"
	}
}

// HealthIcon returns a health indicator
func HealthIcon(health string) string {
	switch strings.ToLower(health) {
	case "ontrack", "on track":
		return "🟢"
	case "atrisk", "at risk":
		return "🟡"
	case "offtrack", "off track":
		return "🔴"
	default:
		return "⚪"
	}
}

// ColorBox returns a colored box character
func ColorBox(hexColor string) string {
	// Just return a simple box - actual coloring would need terminal escape codes
	return "■"
}

// JoinNonEmpty joins non-empty strings with a separator
func JoinNonEmpty(sep string, parts ...string) string {
	nonEmpty := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			nonEmpty = append(nonEmpty, p)
		}
	}
	return strings.Join(nonEmpty, sep)
}
