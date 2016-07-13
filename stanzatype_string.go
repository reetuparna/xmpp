// Code generated by "stringer -type=iqType,errorType,messageType,presenceType -output stanzatype_string.go"; DO NOT EDIT

package xmpp

import "fmt"

const _iqType_name = "GetIQSetIQResultIQErrorIQ"

var _iqType_index = [...]uint8{0, 5, 10, 18, 25}

func (i iqType) String() string {
	if i < 0 || i >= iqType(len(_iqType_index)-1) {
		return fmt.Sprintf("iqType(%d)", i)
	}
	return _iqType_name[_iqType_index[i]:_iqType_index[i+1]]
}

const _errorType_name = "CancelAuthContinueModifyWait"

var _errorType_index = [...]uint8{0, 6, 10, 18, 24, 28}

func (i errorType) String() string {
	if i < 0 || i >= errorType(len(_errorType_index)-1) {
		return fmt.Sprintf("errorType(%d)", i)
	}
	return _errorType_name[_errorType_index[i]:_errorType_index[i+1]]
}

const _messageType_name = "NormalMessageChatMessageErrorMessageGroupChatMessageHeadlineMessage"

var _messageType_index = [...]uint8{0, 13, 24, 36, 52, 67}

func (i messageType) String() string {
	if i < 0 || i >= messageType(len(_messageType_index)-1) {
		return fmt.Sprintf("messageType(%d)", i)
	}
	return _messageType_name[_messageType_index[i]:_messageType_index[i+1]]
}

const _presenceType_name = "NoTypePresenceErrorPresenceProbePresenceSubscribePresenceSubscribedPresenceUnavailablePresenceUnsubscribePresenceUnsubscribedPresence"

var _presenceType_index = [...]uint8{0, 14, 27, 40, 57, 75, 94, 113, 133}

func (i presenceType) String() string {
	if i < 0 || i >= presenceType(len(_presenceType_index)-1) {
		return fmt.Sprintf("presenceType(%d)", i)
	}
	return _presenceType_name[_presenceType_index[i]:_presenceType_index[i+1]]
}
