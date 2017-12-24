package SMSService

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type EnmOnLengthProblem string

const (
	EnmOnLengthProblemRejectAllPackage EnmOnLengthProblem = "RejectAllPackage"

	EnmOnLengthProblemSendOtherMessages EnmOnLengthProblem = "SendOtherMessages"

	EnmOnLengthProblemSendAllPackage EnmOnLengthProblem = "SendAllPackage"
)

type EnmGroupInfoType string

const (
	EnmGroupInfoTypeGroupName EnmGroupInfoType = "GroupName"

	EnmGroupInfoTypeGroupId EnmGroupInfoType = "GroupId"
)

type EnmACParameterTypes string

const (
	EnmACParameterTypesDT EnmACParameterTypes = "DT"

	EnmACParameterTypesGI EnmACParameterTypes = "GI"

	EnmACParameterTypesSI EnmACParameterTypes = "SI"

	EnmACParameterTypesSMS EnmACParameterTypes = "SMS"
)

type EnmGIFormat string

const (
	EnmGIFormatNUMERIC EnmGIFormat = "NUMERIC"

	EnmGIFormatALPHANUMERIC EnmGIFormat = "ALPHANUMERIC"

	EnmGIFormatPASSWORD_NUMERIC EnmGIFormat = "PASSWORD_NUMERIC"
)

type SendSMSSingle struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSSingle"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messages string `xml:"messages,omitempty"`

	Receipents string `xml:"receipents,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSSingleResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSSingleResponse"`

	SendSMSSingleResult *SendSMSSingleResult `xml:"SendSMSSingleResult,omitempty"`
}
type SendSMSSingleResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSSingleResult"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ID string `xml:"ID,omitempty"`
}
type SendSMSSeperatorAVS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSSeperatorAVS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messagewithsep string `xml:"messagewithsep,omitempty"`

	Receipentwithsep string `xml:"receipentwithsep,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`

	Sep string `xml:"sep,omitempty"`
}

type SendSMSSeperatorAVSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSSeperatorAVSResponse"`

	SendSMSSeperatorAVSResult *SendSMSResult `xml:"SendSMSSeperatorAVSResult,omitempty"`
}

type SendSMS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messages *ArrayOfString `xml:"messages,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSResponse"`

	SendSMSResult *SendSMSResult `xml:"SendSMSResult,omitempty"`
}

type SendSMS_V2 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMS_V2"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messages *ArrayOfString `xml:"messages,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	MessageIds *ArrayOfString `xml:"messageIds,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMS_V2Response struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMS_V2Response"`

	SendSMS_V2Result *ResponseFromHostV2 `xml:"SendSMS_V2Result,omitempty"`
}

type SendSMSBroadcast struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSBroadcast"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	BroadcastMessage string `xml:"broadcastMessage,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSBroadcastResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSBroadcastResponse"`

	SendSMSBroadcastResult *SendSMSResult `xml:"SendSMSBroadcastResult,omitempty"`
}

type SendSMSBroadcast_V2 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSBroadcast_V2"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	BroadcastMessage string `xml:"broadcastMessage,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	MessageIds *ArrayOfString `xml:"messageIds,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSBroadcast_V2Response struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSBroadcast_V2Response"`

	SendSMSBroadcast_V2Result *ResponseFromHostV2 `xml:"SendSMSBroadcast_V2Result,omitempty"`
}

type SendSMSToGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSToGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Message string `xml:"message,omitempty"`

	Groupinfotype *EnmGroupInfoType `xml:"groupinfotype,omitempty"`

	Group *ArrayOfString `xml:"group,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSToGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSToGroupResponse"`

	SendSMSToGroupResult *SendSMSResult `xml:"SendSMSToGroupResult,omitempty"`
}

type SendTCKNSMS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messages *ArrayOfString `xml:"messages,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendTCKNSMSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMSResponse"`

	SendTCKNSMSResult *SendSMSResult `xml:"SendTCKNSMSResult,omitempty"`
}

type SendTCKNSMSBroadcast struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMSBroadcast"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	BroadcastMessage string `xml:"broadcastMessage,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendTCKNSMSBroadcastResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMSBroadcastResponse"`

	SendTCKNSMSBroadcastResult *SendSMSResult `xml:"SendTCKNSMSBroadcastResult,omitempty"`
}

type SendTCKNSMSToGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMSToGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Message string `xml:"message,omitempty"`

	Groupinfotype *EnmGroupInfoType `xml:"groupinfotype,omitempty"`

	Group *ArrayOfString `xml:"group,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendTCKNSMSToGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendTCKNSMSToGroupResponse"`

	SendTCKNSMSToGroupResult *SendSMSResult `xml:"SendTCKNSMSToGroupResult,omitempty"`
}

type SendSMSAdvancedToGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSAdvancedToGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Message string `xml:"message,omitempty"`

	MessageType string `xml:"messageType,omitempty"`

	MessageHeader string `xml:"messageHeader,omitempty"`

	Groupinfotype *EnmGroupInfoType `xml:"groupinfotype,omitempty"`

	Group *ArrayOfString `xml:"group,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSAdvancedToGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSAdvancedToGroupResponse"`

	SendSMSAdvancedToGroupResult *SendSMSResult `xml:"SendSMSAdvancedToGroupResult,omitempty"`
}

type SendSMSAdvanced struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSAdvanced"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Messages *ArrayOfString `xml:"messages,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	MessageType string `xml:"messageType,omitempty"`

	MessageHeader *ArrayOfString `xml:"messageHeader,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`

	X string `xml:"x,omitempty"`

	Onlengthproblem *EnmOnLengthProblem `xml:"onlengthproblem,omitempty"`
}

type SendSMSAdvancedResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSAdvancedResponse"`

	SendSMSAdvancedResult *SendSMSResult `xml:"SendSMSAdvancedResult,omitempty"`
}

type ReportSMS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Groupid string `xml:"groupid,omitempty"`

	Status string `xml:"status,omitempty"`
}

type ReportSMSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSResponse"`

	ReportSMSResult *ArrayOfReportSMSResult `xml:"ReportSMSResult,omitempty"`
}

type ReportSMS_V2 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMS_V2"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MessageIds *ArrayOfString `xml:"messageIds,omitempty"`

	Status string `xml:"status,omitempty"`
}

type ReportSMS_V2Response struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMS_V2Response"`

	ReportSMS_V2Result *MessageResponseV2 `xml:"ReportSMS_V2Result,omitempty"`
}

type ReportLOKSMS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportLOKSMS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	ProcessId string `xml:"processId,omitempty"`
}

type ReportLOKSMSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportLOKSMSResponse"`

	ReportLOKSMSResult *ReportLOKSMSResult `xml:"ReportLOKSMSResult,omitempty"`
}

type ReportOriginator struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportOriginator"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`
}

type ReportOriginatorResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportOriginatorResponse"`

	ReportOriginatorResult *ArrayOfString `xml:"ReportOriginatorResult,omitempty"`
}

type ReportQuota struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportQuota"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`
}

type ReportQuotaResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportQuotaResponse"`

	ReportQuotaResult string `xml:"ReportQuotaResult,omitempty"`
}

type ReportGroups struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportGroups"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`
}

type ReportGroupsResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportGroupsResponse"`

	ReportGroupsResult *ArrayOfClsGroups `xml:"ReportGroupsResult,omitempty"`
}

type ReportGroupNumbers struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportGroupNumbers"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Groupinfotype *EnmGroupInfoType `xml:"groupinfotype,omitempty"`

	Group *ArrayOfString `xml:"group,omitempty"`
}

type ReportGroupNumbersResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportGroupNumbersResponse"`

	ReportGroupNumbersResult *ArrayOfClsGroupInfo `xml:"ReportGroupNumbersResult,omitempty"`
}

type ReportSMSReverse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSReverse"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`
}

type ReportSMSReverseResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSReverseResponse"`

	ReportSMSReverseResult *ClsReverse `xml:"ReportSMSReverseResult,omitempty"`
}

type AprroveReverse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AprroveReverse"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	GroupId string `xml:"groupId,omitempty"`
}

type AprroveReverseResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AprroveReverseResponse"`

	AprroveReverseResult bool `xml:"AprroveReverseResult,omitempty"`
}

type Optin struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Optin"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`
}

type OptinResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptinResponse"`

	OptinResult *SendSMSResult `xml:"OptinResult,omitempty"`
}

type OptinByTitles struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptinByTitles"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`

	OptinTitles *ArrayOfString `xml:"optinTitles,omitempty"`
}

type OptinByTitlesResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptinByTitlesResponse"`

	OptinByTitlesResult *SendSMSResult `xml:"OptinByTitlesResult,omitempty"`
}

type Optout struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Optout"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`
}

type OptoutResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptoutResponse"`

	OptoutResult *SendSMSResult `xml:"OptoutResult,omitempty"`
}

type OptoutByTitles struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptoutByTitles"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`

	OptinTitles *ArrayOfString `xml:"optinTitles,omitempty"`
}

type OptoutByTitlesResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ OptoutByTitlesResponse"`

	OptoutByTitlesResult *SendSMSResult `xml:"OptoutByTitlesResult,omitempty"`
}

type SendMMS struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendMMS"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Subject string `xml:"subject,omitempty"`

	Message string `xml:"message,omitempty"`

	MmsContentDataBase64 string `xml:"mmsContentDataBase64,omitempty"`

	MmsContentType string `xml:"mmsContentType,omitempty"`

	Exclusionstarttime string `xml:"exclusionstarttime,omitempty"`

	Exclusionexpiretime string `xml:"exclusionexpiretime,omitempty"`

	Channel string `xml:"channel,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	Optinfilter string `xml:"optinfilter,omitempty"`
}

type SendMMSResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendMMSResponse"`

	SendMMSResult *SendSMSResult `xml:"SendMMSResult,omitempty"`
}

type SendAnindaCevap struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendAnindaCevap"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	Receipents *ArrayOfString `xml:"receipents,omitempty"`

	Blacklistfilter string `xml:"blacklistfilter,omitempty"`

	AnindaCevapParameters *ArrayOfACParameter `xml:"anindaCevapParameters,omitempty"`

	AnindaCevapSeperator string `xml:"anindaCevapSeperator,omitempty"`
}

type SendAnindaCevapResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendAnindaCevapResponse"`

	SendAnindaCevapResult *SendSMSResult `xml:"SendAnindaCevapResult,omitempty"`
}

type CreateGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Title string `xml:"title,omitempty"`

	Desc string `xml:"desc,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`

	IsProfiled bool `xml:"isProfiled,omitempty"`

	IsMMS bool `xml:"isMMS,omitempty"`

	IsLocation bool `xml:"isLocation,omitempty"`
}

type CreateGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateGroupResponse"`

	CreateGroupResult *SendSMSResult `xml:"CreateGroupResult,omitempty"`
}

type CreateCampaign struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateCampaign"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Title string `xml:"title,omitempty"`

	Message string `xml:"message,omitempty"`

	GroupId int32 `xml:"groupId,omitempty"`

	Attributes string `xml:"attributes,omitempty"`
}

type CreateCampaignResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateCampaignResponse"`

	CreateCampaignResult *SendSMSResult `xml:"CreateCampaignResult,omitempty"`
}

type CreateLOKCampaign struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateLOKCampaign"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Title string `xml:"title,omitempty"`

	Message string `xml:"message,omitempty"`

	CampaignType int32 `xml:"campaignType,omitempty"`
}

type CreateLOKCampaignResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreateLOKCampaignResponse"`

	CreateLOKCampaignResult *SendSMSResult `xml:"CreateLOKCampaignResult,omitempty"`
}

type SendCampaign struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendCampaign"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	CampaignId int32 `xml:"campaignId,omitempty"`
}

type SendCampaignResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendCampaignResponse"`

	SendCampaignResult *SendSMSResult `xml:"SendCampaignResult,omitempty"`
}

type SendLOKCampaign struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendLOKCampaign"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Originator string `xml:"originator,omitempty"`

	Reference string `xml:"reference,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Expiredate string `xml:"expiredate,omitempty"`

	CampaignId int32 `xml:"campaignId,omitempty"`

	GroupId int32 `xml:"groupId,omitempty"`

	LokParameter *LOKParameter `xml:"lokParameter,omitempty"`
}

type SendLOKCampaignResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendLOKCampaignResponse"`

	SendLOKCampaignResult *SendSMSResult `xml:"SendLOKCampaignResult,omitempty"`
}

type GetAttributeValues struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetAttributeValues"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	AttributeKey string `xml:"attributeKey,omitempty"`
}

type GetAttributeValuesResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetAttributeValuesResponse"`

	GetAttributeValuesResult *AttributeValuesResult `xml:"GetAttributeValuesResult,omitempty"`
}

type AddMsisdnsToGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AddMsisdnsToGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`

	GroupId string `xml:"groupId,omitempty"`
}

type AddMsisdnsToGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AddMsisdnsToGroupResponse"`

	AddMsisdnsToGroupResult *GroupResult `xml:"AddMsisdnsToGroupResult,omitempty"`
}

type RemoveMsisdnsFromGroup struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RemoveMsisdnsFromGroup"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`

	GroupId string `xml:"groupId,omitempty"`
}

type RemoveMsisdnsFromGroupResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ RemoveMsisdnsFromGroupResponse"`

	RemoveMsisdnsFromGroupResult *GroupResult `xml:"RemoveMsisdnsFromGroupResult,omitempty"`
}

type GetBioTLCredit struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetBioTLCredit"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`
}

type GetBioTLCreditResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetBioTLCreditResponse"`

	GetBioTLCreditResult *CreditResult `xml:"GetBioTLCreditResult,omitempty"`
}

type GetSMSCredit struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetSMSCredit"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Channel string `xml:"channel,omitempty"`
}

type GetSMSCreditResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetSMSCreditResponse"`

	GetSMSCreditResult *CreditResult `xml:"GetSMSCreditResult,omitempty"`
}

type ReportSMSDetail struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSDetail"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Startdate string `xml:"startdate,omitempty"`

	Enddate string `xml:"enddate,omitempty"`

	Gsmnumber string `xml:"gsmnumber,omitempty"`

	Originator string `xml:"originator,omitempty"`
}

type ReportSMSDetailResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSDetailResponse"`

	ReportSMSDetailResult *ArrayOfReportSMSDetailedResult `xml:"ReportSMSDetailResult,omitempty"`
}

type CheckBlacklist struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CheckBlacklist"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	MsisdnList *ArrayOfString `xml:"msisdnList,omitempty"`
}

type CheckBlacklistResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CheckBlacklistResponse"`

	CheckBlacklistResult *CheckBLResult `xml:"CheckBlacklistResult,omitempty"`
}

type SendSMSResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SendSMSResult"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ID string `xml:"ID,omitempty"`
}

type ArrayOfString struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type ResponseFromHostV2 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ResponseFromHostV2"`

	ResponseCode string `xml:"ResponseCode,omitempty"`

	ResponseMessage string `xml:"ResponseMessage,omitempty"`

	ResponseGroupIdArray *ArrayOfResponseGroupId `xml:"ResponseGroupIdArray,omitempty"`
}

type ArrayOfResponseGroupId struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfResponseGroupId"`

	ResponseGroupId []*ResponseGroupId `xml:"ResponseGroupId,omitempty"`
}

type ResponseGroupId struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ResponseGroupId"`

	GsmNumber string `xml:"GsmNumber,omitempty"`

	MessageId string `xml:"MessageId,omitempty"`

	Status string `xml:"Status,omitempty"`
}

type ArrayOfReportSMSResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfReportSMSResult"`

	ReportSMSResult []*ReportSMSResult `xml:"ReportSMSResult,omitempty"`
}

type ReportSMSResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSResult"`

	ID string `xml:"ID,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Status string `xml:"Status,omitempty"`

	SendDate string `xml:"SendDate,omitempty"`

	DeliveredDate string `xml:"DeliveredDate,omitempty"`
}

type MessageResponseV2 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ MessageResponseV2"`

	ResponseCode string `xml:"ResponseCode,omitempty"`

	ResponseMessage string `xml:"ResponseMessage,omitempty"`

	MessageReportArray *ArrayOfMessageReport `xml:"MessageReportArray,omitempty"`
}

type ArrayOfMessageReport struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfMessageReport"`

	MessageReport []*MessageReport `xml:"MessageReport,omitempty"`
}

type MessageReport struct {
	XMLName xml.Name `xml:"http://tempuri.org/ MessageReport"`

	GsmNumber string `xml:"GsmNumber,omitempty"`

	Status string `xml:"Status,omitempty"`

	SentDate string `xml:"SentDate,omitempty"`

	DeliveredDate string `xml:"DeliveredDate,omitempty"`

	MessageId string `xml:"MessageId,omitempty"`

	ResponseCode string `xml:"ResponseCode,omitempty"`
}

type ReportLOKSMSResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportLOKSMSResult"`

	CampaignId string `xml:"CampaignId,omitempty"`

	MessageCount string `xml:"MessageCount,omitempty"`

	SuccessfulMessage string `xml:"SuccessfulMessage,omitempty"`

	FailedMessage string `xml:"FailedMessage,omitempty"`

	CatchedMsisdn string `xml:"CatchedMsisdn,omitempty"`
}

type ArrayOfClsGroups struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfClsGroups"`

	ClsGroups []*ClsGroups `xml:"clsGroups,omitempty"`
}

type ClsGroups struct {
	XMLName xml.Name `xml:"http://tempuri.org/ clsGroups"`

	Groupname string `xml:"Groupname,omitempty"`

	Groupid int32 `xml:"Groupid,omitempty"`
}

type ArrayOfClsGroupInfo struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfClsGroupInfo"`

	ClsGroupInfo []*ClsGroupInfo `xml:"clsGroupInfo,omitempty"`
}

type ClsGroupInfo struct {
	XMLName xml.Name `xml:"http://tempuri.org/ clsGroupInfo"`

	GroupId int32 `xml:"GroupId,omitempty"`

	Groupname string `xml:"Groupname,omitempty"`

	Ad string `xml:"Ad,omitempty"`

	Soyad string `xml:"Soyad,omitempty"`

	Note string `xml:"Note,omitempty"`

	Gsmnumber string `xml:"Gsmnumber,omitempty"`
}

type ClsReverse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ clsReverse"`

	Messages *ArrayOfClsMessage `xml:"Messages,omitempty"`

	Unqkey string `xml:"Unqkey,omitempty"`
}

type ArrayOfClsMessage struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfClsMessage"`

	ClsMessage []*ClsMessage `xml:"clsMessage,omitempty"`
}

type ClsMessage struct {
	XMLName xml.Name `xml:"http://tempuri.org/ clsMessage"`

	MessageId int32 `xml:"MessageId,omitempty"`

	Number string `xml:"Number,omitempty"`

	Status int32 `xml:"Status,omitempty"`

	GroupId string `xml:"GroupId,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	Message string `xml:"Message,omitempty"`

	SendDate string `xml:"SendDate,omitempty"`
}

type ArrayOfACParameter struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfACParameter"`

	ACParameter []*ACParameter `xml:"ACParameter,omitempty"`
}

type ACParameter struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ACParameter"`

	Type *EnmACParameterTypes `xml:"Type,omitempty"`

	DT_DisplayText string `xml:"DT_DisplayText,omitempty"`

	DT_SMSText string `xml:"DT_SMSText,omitempty"`

	GI_Text string `xml:"GI_Text,omitempty"`

	GI_Format *EnmGIFormat `xml:"GI_Format,omitempty"`

	SI_Text string `xml:"SI_Text,omitempty"`

	GI_MinLength int32 `xml:"GI_MinLength,omitempty"`

	GI_MaxLength int32 `xml:"GI_MaxLength,omitempty"`

	SI_OptionTextList *ArrayOfString `xml:"SI_OptionTextList,omitempty"`

	SI_OptionValueList *ArrayOfString `xml:"SI_OptionValueList,omitempty"`
}

type LOKParameter struct {
	XMLName xml.Name `xml:"http://tempuri.org/ LOKParameter"`

	MaxMsisdnCount int32 `xml:"MaxMsisdnCount,omitempty"`

	DayHourInterval string `xml:"DayHourInterval,omitempty"`

	LokArea *ArrayOfLOKArea `xml:"LokArea,omitempty"`
}

type ArrayOfLOKArea struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfLOKArea"`

	LOKArea []*LOKArea `xml:"LOKArea,omitempty"`
}

type LOKArea struct {
	XMLName xml.Name `xml:"http://tempuri.org/ LOKArea"`

	AreaId int32 `xml:"AreaId,omitempty"`

	AreaType int32 `xml:"AreaType,omitempty"`
}

type AttributeValuesResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AttributeValuesResult"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ErrorDesc string `xml:"ErrorDesc,omitempty"`

	Values *ArrayOfAttributeValue `xml:"Values,omitempty"`
}

type ArrayOfAttributeValue struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfAttributeValue"`

	AttributeValue []*AttributeValue `xml:"AttributeValue,omitempty"`
}

type AttributeValue struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AttributeValue"`

	ID string `xml:"ID,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type GroupResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GroupResult"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ErrorDesc string `xml:"ErrorDesc,omitempty"`

	Count string `xml:"Count,omitempty"`
}

type CreditResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CreditResult"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ErrorDesc string `xml:"ErrorDesc,omitempty"`

	Balance float64 `xml:"Balance,omitempty"`

	BalanceDateTime string `xml:"BalanceDateTime,omitempty"`
}

type ArrayOfReportSMSDetailedResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfReportSMSDetailedResult"`

	ReportSMSDetailedResult []*ReportSMSDetailedResult `xml:"ReportSMSDetailedResult,omitempty"`
}

type ReportSMSDetailedResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ReportSMSDetailedResult"`

	ID string `xml:"ID,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Status string `xml:"Status,omitempty"`

	SendDate string `xml:"SendDate,omitempty"`

	DeliveredDate string `xml:"DeliveredDate,omitempty"`

	Originator string `xml:"Originator,omitempty"`

	CreateDate string `xml:"CreateDate,omitempty"`

	MessageText string `xml:"MessageText,omitempty"`
}

type CheckBLResult struct {
	XMLName xml.Name `xml:"http://tempuri.org/ CheckBLResult"`

	*SendSMSResult

	MsisdnBlacklistMatch *ArrayOfMsisdnBL `xml:"MsisdnBlacklistMatch,omitempty"`
}

type ArrayOfMsisdnBL struct {
	XMLName xml.Name `xml:"http://tempuri.org/ ArrayOfMsisdnBL"`

	MsisdnBL []*MsisdnBL `xml:"MsisdnBL,omitempty"`
}

type MsisdnBL struct {
	XMLName xml.Name `xml:"http://tempuri.org/ MsisdnBL"`

	Msisdn string `xml:"Msisdn,omitempty"`

	Blacklist bool `xml:"Blacklist,omitempty"`
}

type SMSServiceSoap struct {
	client *SOAPClient
}

func NewSMSServiceSoap(url string, tls bool, auth *BasicAuth) *SMSServiceSoap {
	if url == "" {
		url = "https://www.jetsms.net/ws/soapsms.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SMSServiceSoap{
		client: client,
	}
}

func NewSMSServiceSoapWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SMSServiceSoap {
	if url == "" {
		url = "https://www.jetsms.net/ws/soapsms.asmx"
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &SMSServiceSoap{
		client: client,
	}
}

func (service *SMSServiceSoap) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *SMSServiceSoap) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *SMSServiceSoap) SendSMSSingle(request *SendSMSSingle) (*SendSMSSingleResponse, error) {
	response := new(SendSMSSingleResponse)
	err := service.client.Call("http://tempuri.org/SendSMSSingle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSSeperatorAVS(request *SendSMSSeperatorAVS) (*SendSMSSeperatorAVSResponse, error) {
	response := new(SendSMSSeperatorAVSResponse)
	err := service.client.Call("http://tempuri.org/SendSMSSeperatorAVS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMS(request *SendSMS) (*SendSMSResponse, error) {
	response := new(SendSMSResponse)
	err := service.client.Call("http://tempuri.org/SendSMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMS_V2(request *SendSMS_V2) (*SendSMS_V2Response, error) {
	response := new(SendSMS_V2Response)
	err := service.client.Call("http://tempuri.org/SendSMS_V2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSBroadcast(request *SendSMSBroadcast) (*SendSMSBroadcastResponse, error) {
	response := new(SendSMSBroadcastResponse)
	err := service.client.Call("http://tempuri.org/SendSMSBroadcast", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSBroadcast_V2(request *SendSMSBroadcast_V2) (*SendSMSBroadcast_V2Response, error) {
	response := new(SendSMSBroadcast_V2Response)
	err := service.client.Call("http://tempuri.org/SendSMSBroadcast_V2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSToGroup(request *SendSMSToGroup) (*SendSMSToGroupResponse, error) {
	response := new(SendSMSToGroupResponse)
	err := service.client.Call("http://tempuri.org/SendSMSToGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendTCKNSMS(request *SendTCKNSMS) (*SendTCKNSMSResponse, error) {
	response := new(SendTCKNSMSResponse)
	err := service.client.Call("http://tempuri.org/SendTCKNSMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendTCKNSMSBroadcast(request *SendTCKNSMSBroadcast) (*SendTCKNSMSBroadcastResponse, error) {
	response := new(SendTCKNSMSBroadcastResponse)
	err := service.client.Call("http://tempuri.org/SendTCKNSMSBroadcast", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendTCKNSMSToGroup(request *SendTCKNSMSToGroup) (*SendTCKNSMSToGroupResponse, error) {
	response := new(SendTCKNSMSToGroupResponse)
	err := service.client.Call("http://tempuri.org/SendTCKNSMSToGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSAdvancedToGroup(request *SendSMSAdvancedToGroup) (*SendSMSAdvancedToGroupResponse, error) {
	response := new(SendSMSAdvancedToGroupResponse)
	err := service.client.Call("http://tempuri.org/SendSMSAdvancedToGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendSMSAdvanced(request *SendSMSAdvanced) (*SendSMSAdvancedResponse, error) {
	response := new(SendSMSAdvancedResponse)
	err := service.client.Call("http://tempuri.org/SendSMSAdvanced", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportSMS(request *ReportSMS) (*ReportSMSResponse, error) {
	response := new(ReportSMSResponse)
	err := service.client.Call("http://tempuri.org/ReportSMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportSMS_V2(request *ReportSMS_V2) (*ReportSMS_V2Response, error) {
	response := new(ReportSMS_V2Response)
	err := service.client.Call("http://tempuri.org/ReportSMS_V2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportLOKSMS(request *ReportLOKSMS) (*ReportLOKSMSResponse, error) {
	response := new(ReportLOKSMSResponse)
	err := service.client.Call("http://tempuri.org/ReportLOKSMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportOriginator(request *ReportOriginator) (*ReportOriginatorResponse, error) {
	response := new(ReportOriginatorResponse)
	err := service.client.Call("http://tempuri.org/ReportOriginator", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportQuota(request *ReportQuota) (*ReportQuotaResponse, error) {
	response := new(ReportQuotaResponse)
	err := service.client.Call("http://tempuri.org/ReportQuota", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportGroups(request *ReportGroups) (*ReportGroupsResponse, error) {
	response := new(ReportGroupsResponse)
	err := service.client.Call("http://tempuri.org/ReportGroups", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportGroupNumbers(request *ReportGroupNumbers) (*ReportGroupNumbersResponse, error) {
	response := new(ReportGroupNumbersResponse)
	err := service.client.Call("http://tempuri.org/ReportGroupNumbers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportSMSReverse(request *ReportSMSReverse) (*ReportSMSReverseResponse, error) {
	response := new(ReportSMSReverseResponse)
	err := service.client.Call("http://tempuri.org/ReportSMSReverse", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) AprroveReverse(request *AprroveReverse) (*AprroveReverseResponse, error) {
	response := new(AprroveReverseResponse)
	err := service.client.Call("http://tempuri.org/AprroveReverse", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) Optin(request *Optin) (*OptinResponse, error) {
	response := new(OptinResponse)
	err := service.client.Call("http://tempuri.org/Optin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) OptinByTitles(request *OptinByTitles) (*OptinByTitlesResponse, error) {
	response := new(OptinByTitlesResponse)
	err := service.client.Call("http://tempuri.org/OptinByTitles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) Optout(request *Optout) (*OptoutResponse, error) {
	response := new(OptoutResponse)
	err := service.client.Call("http://tempuri.org/Optout", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) OptoutByTitles(request *OptoutByTitles) (*OptoutByTitlesResponse, error) {
	response := new(OptoutByTitlesResponse)
	err := service.client.Call("http://tempuri.org/OptoutByTitles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendMMS(request *SendMMS) (*SendMMSResponse, error) {
	response := new(SendMMSResponse)
	err := service.client.Call("http://tempuri.org/SendMMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendAnindaCevap(request *SendAnindaCevap) (*SendAnindaCevapResponse, error) {
	response := new(SendAnindaCevapResponse)
	err := service.client.Call("http://tempuri.org/SendAnindaCevap", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) CreateGroup(request *CreateGroup) (*CreateGroupResponse, error) {
	response := new(CreateGroupResponse)
	err := service.client.Call("http://tempuri.org/CreateGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) CreateCampaign(request *CreateCampaign) (*CreateCampaignResponse, error) {
	response := new(CreateCampaignResponse)
	err := service.client.Call("http://tempuri.org/CreateCampaign", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) CreateLOKCampaign(request *CreateLOKCampaign) (*CreateLOKCampaignResponse, error) {
	response := new(CreateLOKCampaignResponse)
	err := service.client.Call("http://tempuri.org/CreateLOKCampaign", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendCampaign(request *SendCampaign) (*SendCampaignResponse, error) {
	response := new(SendCampaignResponse)
	err := service.client.Call("http://tempuri.org/SendCampaign", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) SendLOKCampaign(request *SendLOKCampaign) (*SendLOKCampaignResponse, error) {
	response := new(SendLOKCampaignResponse)
	err := service.client.Call("http://tempuri.org/SendLOKCampaign", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) GetAttributeValues(request *GetAttributeValues) (*GetAttributeValuesResponse, error) {
	response := new(GetAttributeValuesResponse)
	err := service.client.Call("http://tempuri.org/GetAttributeValues", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) AddMsisdnsToGroup(request *AddMsisdnsToGroup) (*AddMsisdnsToGroupResponse, error) {
	response := new(AddMsisdnsToGroupResponse)
	err := service.client.Call("http://tempuri.org/AddMsisdnsToGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) RemoveMsisdnsFromGroup(request *RemoveMsisdnsFromGroup) (*RemoveMsisdnsFromGroupResponse, error) {
	response := new(RemoveMsisdnsFromGroupResponse)
	err := service.client.Call("http://tempuri.org/RemoveMsisdnsFromGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) GetBioTLCredit(request *GetBioTLCredit) (*GetBioTLCreditResponse, error) {
	response := new(GetBioTLCreditResponse)
	err := service.client.Call("http://tempuri.org/GetBioTLCredit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) GetSMSCredit(request *GetSMSCredit) (*GetSMSCreditResponse, error) {
	response := new(GetSMSCreditResponse)
	err := service.client.Call("http://tempuri.org/GetSMSCredit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) ReportSMSDetail(request *ReportSMSDetail) (*ReportSMSDetailResponse, error) {
	response := new(ReportSMSDetailResponse)
	err := service.client.Call("http://tempuri.org/ReportSMSDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SMSServiceSoap) CheckBlacklist(request *CheckBlacklist) (*CheckBlacklistResponse, error) {
	response := new(CheckBlacklistResponse)
	err := service.client.Call("http://tempuri.org/CheckBlacklist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
