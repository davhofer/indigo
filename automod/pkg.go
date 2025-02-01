package automod

import (
	"github.com/davhofer/indigo/automod/countstore"
	"github.com/davhofer/indigo/automod/engine"
)

type Engine = engine.Engine
type EngineConfig = engine.EngineConfig
type AccountMeta = engine.AccountMeta
type ProfileSummary = engine.ProfileSummary
type AccountPrivate = engine.AccountPrivate
type RuleSet = engine.RuleSet

type Notifier = engine.Notifier
type SlackNotifier = engine.SlackNotifier

type AccountContext = engine.AccountContext
type RecordContext = engine.RecordContext
type OzoneEventContext = engine.OzoneEventContext
type NotificationContext = engine.NotificationContext
type RecordOp = engine.RecordOp

type IdentityRuleFunc = engine.IdentityRuleFunc
type RecordRuleFunc = engine.RecordRuleFunc
type PostRuleFunc = engine.PostRuleFunc
type ProfileRuleFunc = engine.ProfileRuleFunc
type BlobRuleFunc = engine.BlobRuleFunc
type NotificationRuleFunc = engine.NotificationRuleFunc
type OzoneEventRuleFunc = engine.OzoneEventRuleFunc

var (
	ReportReasonSpam       = engine.ReportReasonSpam
	ReportReasonViolation  = engine.ReportReasonViolation
	ReportReasonMisleading = engine.ReportReasonMisleading
	ReportReasonSexual     = engine.ReportReasonSexual
	ReportReasonRude       = engine.ReportReasonRude
	ReportReasonOther      = engine.ReportReasonOther

	PeriodTotal = countstore.PeriodTotal
	PeriodDay   = countstore.PeriodDay
	PeriodHour  = countstore.PeriodHour

	CreateOp = engine.CreateOp
	UpdateOp = engine.UpdateOp
	DeleteOp = engine.DeleteOp
)
