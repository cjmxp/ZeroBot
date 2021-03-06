package zero

import (
	"regexp"
	"strconv"
	"strings"
)

// PrefixRule check if the message has the prefix and trim the prefix
func PrefixRule(prefixes ...string) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		if event.Message == nil && event.Message[0].Type != "text" { // 确保无空指针
			return false
		}
		first := event.Message[0]
		firstMessage := first.Data["text"]
		for _, prefix := range prefixes {
			if strings.HasPrefix(firstMessage, prefix) {
				state["prefix"] = prefix
				state["args"] = strings.TrimLeft(strings.TrimPrefix(firstMessage, prefix), " ")
				return true
			}
		}
		return false
	}
}

// SuffixRule check if the message has the suffix and trim the suffix
func SuffixRule(suffixes ...string) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		if event.Message == nil { // 确保无空指针
			return false
		}
		last := event.Message[len(event.Message)-1]
		if last.Type != "text" {
			return false
		}
		lastMessage := last.Data["text"]
		for _, suffix := range suffixes {
			if strings.HasSuffix(lastMessage, suffix) {
				state["suffix"] = suffix
				state["args"] = strings.TrimLeft(strings.TrimPrefix(lastMessage, suffix), " ")
				return true
			}
		}
		return false
	}
}

// CommandRule check if the message is a command and trim the command name
func CommandRule(commands ...string) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		if event.Message == nil && event.Message[0].Type != "text" {
			return false
		}
		first := event.Message[0]
		firstMessage := first.Data["text"]
		cmdMessage := strings.TrimPrefix(firstMessage, zeroBot.commandPrefix)
		if cmdMessage == firstMessage {
			return false
		}
		for _, command := range commands {
			if strings.HasPrefix(cmdMessage, command) {
				state["command"] = command
				state["args"] = strings.TrimLeft(cmdMessage[len(command):], " ")
				return true
			}
		}
		return false
	}
}

// RegexRule check if the message can be matched by the regex pattern
func RegexRule(regexPattern string) func(event *Event, state State) bool {
	regex := regexp.MustCompile(regexPattern)
	return func(event *Event, state State) bool {
		msg := event.Message.CQString()
		if regex.MatchString(msg) {
			state["regex_matched"] = regex.FindStringSubmatch(msg)
			return true
		}
		return false
	}
}

// KeywordRule check if the message has a keyword or keywords
func KeywordRule(src ...string) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		msg := event.Message.CQString()
		for _, str := range src {
			if strings.Contains(msg, str) {
				state["keyword"] = str
				return true
			}
		}
		return false
	}
}

// FullMatchRule check if src has the same copy of the message
func FullMatchRule(src ...string) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		msg := event.Message.CQString()
		for _, str := range src {
			if str == msg {
				return true
			}
		}
		return false
	}
}

// OnlyToMe only triggered in conditions of @bot or begin with the nicknames
func OnlyToMe(event *Event, _ State) bool {
	return event.IsToMe == true
}

// CheckUser only triggered by specific person
func CheckUser(userId ...int64) func(event *Event, state State) bool {
	return func(event *Event, state State) bool {
		for _, uid := range userId {
			if event.UserID == uid {
				return true
			}
		}
		return false
	}
}

// OnlyPrivate requres that the event is private message
func OnlyPrivate(event *Event, _ State) bool {
	return event.PostType == "message" && event.DetailType == "private"
}

// OnlyPrivate requres that the event is public/group message
func OnlyGroup(event *Event, _ State) bool {
	return event.PostType == "message" && event.DetailType == "group"
}

// SuperUserPermission only triggered by the bot's owner
func SuperUserPermission(event *Event, _ State) bool {
	uid := strconv.FormatInt(event.UserID, 10)
	for _, su := range zeroBot.SuperUsers {
		if su == uid {
			return true
		}
	}
	return false
}

// AdminPermission only triggered by the group admins or higher permission
func AdminPermission(event *Event, state State) bool {
	return SuperUserPermission(event, state) || event.Sender.Role != "member"
}

// OwnerPermission only triggered by the group owner or higher permission
func OwnerPermission(event *Event, state State) bool {
	return SuperUserPermission(event, state) ||
		(event.Sender.Role != "member" && event.Sender.Role != "admin")
}
