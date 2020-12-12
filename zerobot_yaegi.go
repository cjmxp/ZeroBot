// generated by the yaegi extract
package zero

import (
	"reflect"
)

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/wdvxdr1123/ZeroBot"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AdminPermission":         reflect.ValueOf(AdminPermission),
		"CallAction":              reflect.ValueOf(CallAction),
		"CheckUser":               reflect.ValueOf(CheckUser),
		"CommandRule":             reflect.ValueOf(CommandRule),
		"DeleteMessage":           reflect.ValueOf(DeleteMessage),
		"FinishResponse":          reflect.ValueOf(FinishResponse),
		"FullMatchRule":           reflect.ValueOf(FullMatchRule),
		"GetForwardMessage":       reflect.ValueOf(GetForwardMessage),
		"GetFriendList":           reflect.ValueOf(GetFriendList),
		"GetGroupHonorInfo":       reflect.ValueOf(GetGroupHonorInfo),
		"GetGroupInfo":            reflect.ValueOf(GetGroupInfo),
		"GetGroupList":            reflect.ValueOf(GetGroupList),
		"GetGroupMemberInfo":      reflect.ValueOf(GetGroupMemberInfo),
		"GetGroupMemberList":      reflect.ValueOf(GetGroupMemberList),
		"GetGroupSystemMessage":   reflect.ValueOf(GetGroupSystemMessage),
		"GetImage":                reflect.ValueOf(GetImage),
		"GetLoginInfo":            reflect.ValueOf(GetLoginInfo),
		"GetMessage":              reflect.ValueOf(GetMessage),
		"GetRecord":               reflect.ValueOf(GetRecord),
		"GetStrangerInfo":         reflect.ValueOf(GetStrangerInfo),
		"GetVersionInfo":          reflect.ValueOf(GetVersionInfo),
		"GetWordSlices":           reflect.ValueOf(GetWordSlices),
		"KeywordRule":             reflect.ValueOf(KeywordRule),
		"OCRImage":                reflect.ValueOf(OCRImage),
		"On":                      reflect.ValueOf(On),
		"OnCommand":               reflect.ValueOf(OnCommand),
		"OnCommandGroup":          reflect.ValueOf(OnCommandGroup),
		"OnFullMatch":             reflect.ValueOf(OnFullMatch),
		"OnFullMatchGroup":        reflect.ValueOf(OnFullMatchGroup),
		"OnKeyword":               reflect.ValueOf(OnKeyword),
		"OnKeywordGroup":          reflect.ValueOf(OnKeywordGroup),
		"OnMessage":               reflect.ValueOf(OnMessage),
		"OnMetaEvent":             reflect.ValueOf(OnMetaEvent),
		"OnNotice":                reflect.ValueOf(OnNotice),
		"OnPrefix":                reflect.ValueOf(OnPrefix),
		"OnPrefixGroup":           reflect.ValueOf(OnPrefixGroup),
		"OnRegex":                 reflect.ValueOf(OnRegex),
		"OnRequest":               reflect.ValueOf(OnRequest),
		"OnSuffix":                reflect.ValueOf(OnSuffix),
		"OnSuffixGroup":           reflect.ValueOf(OnSuffixGroup),
		"OnlyGroup":               reflect.ValueOf(OnlyGroup),
		"OnlyPrivate":             reflect.ValueOf(OnlyPrivate),
		"OnlyToMe":                reflect.ValueOf(OnlyToMe),
		"OwnerPermission":         reflect.ValueOf(OwnerPermission),
		"PrefixRule":              reflect.ValueOf(PrefixRule),
		"RegexRule":               reflect.ValueOf(RegexRule),
		"RegisterPlugin":          reflect.ValueOf(RegisterPlugin),
		"RejectResponse":          reflect.ValueOf(RejectResponse),
		"Run":                     reflect.ValueOf(Run),
		"Send":                    reflect.ValueOf(Send),
		"SendGroupForwardMessage": reflect.ValueOf(SendGroupForwardMessage),
		"SendGroupMessage":        reflect.ValueOf(SendGroupMessage),
		"SendPrivateMessage":      reflect.ValueOf(SendPrivateMessage),
		"SetFriendAddRequest":     reflect.ValueOf(SetFriendAddRequest),
		"SetGroupAddRequest":      reflect.ValueOf(SetGroupAddRequest),
		"SetGroupAdmin":           reflect.ValueOf(SetGroupAdmin),
		"SetGroupAnonymous":       reflect.ValueOf(SetGroupAnonymous),
		"SetGroupBan":             reflect.ValueOf(SetGroupBan),
		"SetGroupCard":            reflect.ValueOf(SetGroupCard),
		"SetGroupKick":            reflect.ValueOf(SetGroupKick),
		"SetGroupLeave":           reflect.ValueOf(SetGroupLeave),
		"SetGroupName":            reflect.ValueOf(SetGroupName),
		"SetGroupPortrait":        reflect.ValueOf(SetGroupPortrait),
		"SetGroupSpecialTitle":    reflect.ValueOf(SetGroupSpecialTitle),
		"SetGroupWholeBan":        reflect.ValueOf(SetGroupWholeBan),
		"SuccessResponse":         reflect.ValueOf(SuccessResponse),
		"SuffixRule":              reflect.ValueOf(SuffixRule),
		"SuperUserPermission":     reflect.ValueOf(SuperUserPermission),

		// type definitions
		"APIResponse":      reflect.ValueOf((*APIResponse)(nil)),
		"Event":            reflect.ValueOf((*Event)(nil)),
		"File":             reflect.ValueOf((*File)(nil)),
		"Group":            reflect.ValueOf((*Group)(nil)),
		"Handler":          reflect.ValueOf((*Handler)(nil)),
		"IPlugin":          reflect.ValueOf((*IPlugin)(nil)),
		"Matcher":          reflect.ValueOf((*Matcher)(nil)),
		"Message":          reflect.ValueOf((*Message)(nil)),
		"Option":           reflect.ValueOf((*Option)(nil)),
		"Params":           reflect.ValueOf((*Params)(nil)),
		"PluginInfo":       reflect.ValueOf((*PluginInfo)(nil)),
		"Response":         reflect.ValueOf((*Response)(nil)),
		"Rule":             reflect.ValueOf((*Rule)(nil)),
		"State":            reflect.ValueOf((*State)(nil)),
		"User":             reflect.ValueOf((*User)(nil)),
		"WebSocketRequest": reflect.ValueOf((*WebSocketRequest)(nil)),

		// interface wrapper definitions
		"_IPlugin": reflect.ValueOf((*_github_com_wdvxdr1123_ZeroBot_IPlugin)(nil)),
	}
}

// _github_com_wdvxdr1123_ZeroBot_IPlugin is an interface wrapper for IPlugin type
type _github_com_wdvxdr1123_ZeroBot_IPlugin struct {
	WGetPluginInfo func() PluginInfo
	WStart         func()
}

func (W _github_com_wdvxdr1123_ZeroBot_IPlugin) GetPluginInfo() PluginInfo { return W.WGetPluginInfo() }
func (W _github_com_wdvxdr1123_ZeroBot_IPlugin) Start()                    { W.WStart() }
