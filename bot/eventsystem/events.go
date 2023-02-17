// GENERATED using events_gen.go

// Custom event handlers that adds a redis connection to the handler
// They will also recover from panics

package eventsystem

import (
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
)

type Event int

const (
	EventNewGuild                            Event = 0
	EventAll                                 Event = 1
	EventAllPre                              Event = 2
	EventAllPost                             Event = 3
	EventMemberFetched                       Event = 4
	EventYagShardReady                       Event = 5
	EventYagShardsAdded                      Event = 6
	EventYagShardRemoved                     Event = 7
	EventApplicationCommandCreate            Event = 8
	EventApplicationCommandDelete            Event = 9
	EventApplicationCommandPermissionsUpdate Event = 10
	EventApplicationCommandUpdate            Event = 11
	EventAutoModerationActionExecution       Event = 12
	EventAutoModerationRuleCreate            Event = 13
	EventAutoModerationRuleDelete            Event = 14
	EventAutoModerationRuleUpdate            Event = 15
	EventChannelCreate                       Event = 16
	EventChannelDelete                       Event = 17
	EventChannelPinsUpdate                   Event = 18
	EventChannelUpdate                       Event = 19
	EventConnect                             Event = 20
	EventDisconnect                          Event = 21
	EventGuildBanAdd                         Event = 22
	EventGuildBanRemove                      Event = 23
	EventGuildCreate                         Event = 24
	EventGuildDelete                         Event = 25
	EventGuildEmojisUpdate                   Event = 26
	EventGuildIntegrationsUpdate             Event = 27
	EventGuildMemberAdd                      Event = 28
	EventGuildMemberRemove                   Event = 29
	EventGuildMemberUpdate                   Event = 30
	EventGuildMembersChunk                   Event = 31
	EventGuildRoleCreate                     Event = 32
	EventGuildRoleDelete                     Event = 33
	EventGuildRoleUpdate                     Event = 34
	EventGuildStickersUpdate                 Event = 35
	EventGuildUpdate                         Event = 36
	EventInteractionCreate                   Event = 37
	EventInviteCreate                        Event = 38
	EventInviteDelete                        Event = 39
	EventMessageAck                          Event = 40
	EventMessageCreate                       Event = 41
	EventMessageDelete                       Event = 42
	EventMessageDeleteBulk                   Event = 43
	EventMessageReactionAdd                  Event = 44
	EventMessageReactionRemove               Event = 45
	EventMessageReactionRemoveAll            Event = 46
	EventMessageReactionRemoveEmoji          Event = 47
	EventMessageUpdate                       Event = 48
	EventPresenceUpdate                      Event = 49
	EventPresencesReplace                    Event = 50
	EventRateLimit                           Event = 51
	EventReady                               Event = 52
	EventRelationshipAdd                     Event = 53
	EventRelationshipRemove                  Event = 54
	EventResumed                             Event = 55
	EventStageInstanceCreate                 Event = 56
	EventStageInstanceDelete                 Event = 57
	EventStageInstanceUpdate                 Event = 58
	EventThreadCreate                        Event = 59
	EventThreadDelete                        Event = 60
	EventThreadListSync                      Event = 61
	EventThreadMemberUpdate                  Event = 62
	EventThreadMembersUpdate                 Event = 63
	EventThreadUpdate                        Event = 64
	EventTypingStart                         Event = 65
	EventUserGuildSettingsUpdate             Event = 66
	EventUserNoteUpdate                      Event = 67
	EventUserSettingsUpdate                  Event = 68
	EventUserUpdate                          Event = 69
	EventVoiceServerUpdate                   Event = 70
	EventVoiceStateUpdate                    Event = 71
	EventWebhooksUpdate                      Event = 72
)

var EventNames = []string{
	"NewGuild",
	"All",
	"AllPre",
	"AllPost",
	"MemberFetched",
	"YagShardReady",
	"YagShardsAdded",
	"YagShardRemoved",
	"ApplicationCommandCreate",
	"ApplicationCommandDelete",
	"ApplicationCommandPermissionsUpdate",
	"ApplicationCommandUpdate",
	"AutoModerationActionExecution",
	"AutoModerationRuleCreate",
	"AutoModerationRuleDelete",
	"AutoModerationRuleUpdate",
	"ChannelCreate",
	"ChannelDelete",
	"ChannelPinsUpdate",
	"ChannelUpdate",
	"Connect",
	"Disconnect",
	"GuildBanAdd",
	"GuildBanRemove",
	"GuildCreate",
	"GuildDelete",
	"GuildEmojisUpdate",
	"GuildIntegrationsUpdate",
	"GuildMemberAdd",
	"GuildMemberRemove",
	"GuildMemberUpdate",
	"GuildMembersChunk",
	"GuildRoleCreate",
	"GuildRoleDelete",
	"GuildRoleUpdate",
	"GuildStickersUpdate",
	"GuildUpdate",
	"InteractionCreate",
	"InviteCreate",
	"InviteDelete",
	"MessageAck",
	"MessageCreate",
	"MessageDelete",
	"MessageDeleteBulk",
	"MessageReactionAdd",
	"MessageReactionRemove",
	"MessageReactionRemoveAll",
	"MessageReactionRemoveEmoji",
	"MessageUpdate",
	"PresenceUpdate",
	"PresencesReplace",
	"RateLimit",
	"Ready",
	"RelationshipAdd",
	"RelationshipRemove",
	"Resumed",
	"StageInstanceCreate",
	"StageInstanceDelete",
	"StageInstanceUpdate",
	"ThreadCreate",
	"ThreadDelete",
	"ThreadListSync",
	"ThreadMemberUpdate",
	"ThreadMembersUpdate",
	"ThreadUpdate",
	"TypingStart",
	"UserGuildSettingsUpdate",
	"UserNoteUpdate",
	"UserSettingsUpdate",
	"UserUpdate",
	"VoiceServerUpdate",
	"VoiceStateUpdate",
	"WebhooksUpdate",
}

func (e Event) String() string {
	return EventNames[e]
}

var AllDiscordEvents = []Event{
	EventApplicationCommandCreate,
	EventApplicationCommandDelete,
	EventApplicationCommandPermissionsUpdate,
	EventApplicationCommandUpdate,
	EventAutoModerationActionExecution,
	EventAutoModerationRuleCreate,
	EventAutoModerationRuleDelete,
	EventAutoModerationRuleUpdate,
	EventChannelCreate,
	EventChannelDelete,
	EventChannelPinsUpdate,
	EventChannelUpdate,
	EventConnect,
	EventDisconnect,
	EventGuildBanAdd,
	EventGuildBanRemove,
	EventGuildCreate,
	EventGuildDelete,
	EventGuildEmojisUpdate,
	EventGuildIntegrationsUpdate,
	EventGuildMemberAdd,
	EventGuildMemberRemove,
	EventGuildMemberUpdate,
	EventGuildMembersChunk,
	EventGuildRoleCreate,
	EventGuildRoleDelete,
	EventGuildRoleUpdate,
	EventGuildStickersUpdate,
	EventGuildUpdate,
	EventInteractionCreate,
	EventInviteCreate,
	EventInviteDelete,
	EventMessageAck,
	EventMessageCreate,
	EventMessageDelete,
	EventMessageDeleteBulk,
	EventMessageReactionAdd,
	EventMessageReactionRemove,
	EventMessageReactionRemoveAll,
	EventMessageReactionRemoveEmoji,
	EventMessageUpdate,
	EventPresenceUpdate,
	EventPresencesReplace,
	EventRateLimit,
	EventReady,
	EventRelationshipAdd,
	EventRelationshipRemove,
	EventResumed,
	EventStageInstanceCreate,
	EventStageInstanceDelete,
	EventStageInstanceUpdate,
	EventThreadCreate,
	EventThreadDelete,
	EventThreadListSync,
	EventThreadMemberUpdate,
	EventThreadMembersUpdate,
	EventThreadUpdate,
	EventTypingStart,
	EventUserGuildSettingsUpdate,
	EventUserNoteUpdate,
	EventUserSettingsUpdate,
	EventUserUpdate,
	EventVoiceServerUpdate,
	EventVoiceStateUpdate,
	EventWebhooksUpdate,
}

var AllEvents = []Event{
	EventNewGuild,
	EventAll,
	EventAllPre,
	EventAllPost,
	EventMemberFetched,
	EventYagShardReady,
	EventYagShardsAdded,
	EventYagShardRemoved,
	EventApplicationCommandCreate,
	EventApplicationCommandDelete,
	EventApplicationCommandPermissionsUpdate,
	EventApplicationCommandUpdate,
	EventAutoModerationActionExecution,
	EventAutoModerationRuleCreate,
	EventAutoModerationRuleDelete,
	EventAutoModerationRuleUpdate,
	EventChannelCreate,
	EventChannelDelete,
	EventChannelPinsUpdate,
	EventChannelUpdate,
	EventConnect,
	EventDisconnect,
	EventGuildBanAdd,
	EventGuildBanRemove,
	EventGuildCreate,
	EventGuildDelete,
	EventGuildEmojisUpdate,
	EventGuildIntegrationsUpdate,
	EventGuildMemberAdd,
	EventGuildMemberRemove,
	EventGuildMemberUpdate,
	EventGuildMembersChunk,
	EventGuildRoleCreate,
	EventGuildRoleDelete,
	EventGuildRoleUpdate,
	EventGuildStickersUpdate,
	EventGuildUpdate,
	EventInteractionCreate,
	EventInviteCreate,
	EventInviteDelete,
	EventMessageAck,
	EventMessageCreate,
	EventMessageDelete,
	EventMessageDeleteBulk,
	EventMessageReactionAdd,
	EventMessageReactionRemove,
	EventMessageReactionRemoveAll,
	EventMessageReactionRemoveEmoji,
	EventMessageUpdate,
	EventPresenceUpdate,
	EventPresencesReplace,
	EventRateLimit,
	EventReady,
	EventRelationshipAdd,
	EventRelationshipRemove,
	EventResumed,
	EventStageInstanceCreate,
	EventStageInstanceDelete,
	EventStageInstanceUpdate,
	EventThreadCreate,
	EventThreadDelete,
	EventThreadListSync,
	EventThreadMemberUpdate,
	EventThreadMembersUpdate,
	EventThreadUpdate,
	EventTypingStart,
	EventUserGuildSettingsUpdate,
	EventUserNoteUpdate,
	EventUserSettingsUpdate,
	EventUserUpdate,
	EventVoiceServerUpdate,
	EventVoiceStateUpdate,
	EventWebhooksUpdate,
}

var handlers = make([][][]*Handler, 73)

func (data *EventData) ApplicationCommandCreate() *discordgo.ApplicationCommandCreate {
	return data.EvtInterface.(*discordgo.ApplicationCommandCreate)
}
func (data *EventData) ApplicationCommandDelete() *discordgo.ApplicationCommandDelete {
	return data.EvtInterface.(*discordgo.ApplicationCommandDelete)
}
func (data *EventData) ApplicationCommandPermissionsUpdate() *discordgo.ApplicationCommandPermissionsUpdate {
	return data.EvtInterface.(*discordgo.ApplicationCommandPermissionsUpdate)
}
func (data *EventData) ApplicationCommandUpdate() *discordgo.ApplicationCommandUpdate {
	return data.EvtInterface.(*discordgo.ApplicationCommandUpdate)
}
func (data *EventData) AutoModerationActionExecution() *discordgo.AutoModerationActionExecution {
	return data.EvtInterface.(*discordgo.AutoModerationActionExecution)
}
func (data *EventData) AutoModerationRuleCreate() *discordgo.AutoModerationRuleCreate {
	return data.EvtInterface.(*discordgo.AutoModerationRuleCreate)
}
func (data *EventData) AutoModerationRuleDelete() *discordgo.AutoModerationRuleDelete {
	return data.EvtInterface.(*discordgo.AutoModerationRuleDelete)
}
func (data *EventData) AutoModerationRuleUpdate() *discordgo.AutoModerationRuleUpdate {
	return data.EvtInterface.(*discordgo.AutoModerationRuleUpdate)
}
func (data *EventData) ChannelCreate() *discordgo.ChannelCreate {
	return data.EvtInterface.(*discordgo.ChannelCreate)
}
func (data *EventData) ChannelDelete() *discordgo.ChannelDelete {
	return data.EvtInterface.(*discordgo.ChannelDelete)
}
func (data *EventData) ChannelPinsUpdate() *discordgo.ChannelPinsUpdate {
	return data.EvtInterface.(*discordgo.ChannelPinsUpdate)
}
func (data *EventData) ChannelUpdate() *discordgo.ChannelUpdate {
	return data.EvtInterface.(*discordgo.ChannelUpdate)
}
func (data *EventData) Connect() *discordgo.Connect {
	return data.EvtInterface.(*discordgo.Connect)
}
func (data *EventData) Disconnect() *discordgo.Disconnect {
	return data.EvtInterface.(*discordgo.Disconnect)
}
func (data *EventData) GuildBanAdd() *discordgo.GuildBanAdd {
	return data.EvtInterface.(*discordgo.GuildBanAdd)
}
func (data *EventData) GuildBanRemove() *discordgo.GuildBanRemove {
	return data.EvtInterface.(*discordgo.GuildBanRemove)
}
func (data *EventData) GuildCreate() *discordgo.GuildCreate {
	return data.EvtInterface.(*discordgo.GuildCreate)
}
func (data *EventData) GuildDelete() *discordgo.GuildDelete {
	return data.EvtInterface.(*discordgo.GuildDelete)
}
func (data *EventData) GuildEmojisUpdate() *discordgo.GuildEmojisUpdate {
	return data.EvtInterface.(*discordgo.GuildEmojisUpdate)
}
func (data *EventData) GuildIntegrationsUpdate() *discordgo.GuildIntegrationsUpdate {
	return data.EvtInterface.(*discordgo.GuildIntegrationsUpdate)
}
func (data *EventData) GuildMemberAdd() *discordgo.GuildMemberAdd {
	return data.EvtInterface.(*discordgo.GuildMemberAdd)
}
func (data *EventData) GuildMemberRemove() *discordgo.GuildMemberRemove {
	return data.EvtInterface.(*discordgo.GuildMemberRemove)
}
func (data *EventData) GuildMemberUpdate() *discordgo.GuildMemberUpdate {
	return data.EvtInterface.(*discordgo.GuildMemberUpdate)
}
func (data *EventData) GuildMembersChunk() *discordgo.GuildMembersChunk {
	return data.EvtInterface.(*discordgo.GuildMembersChunk)
}
func (data *EventData) GuildRoleCreate() *discordgo.GuildRoleCreate {
	return data.EvtInterface.(*discordgo.GuildRoleCreate)
}
func (data *EventData) GuildRoleDelete() *discordgo.GuildRoleDelete {
	return data.EvtInterface.(*discordgo.GuildRoleDelete)
}
func (data *EventData) GuildRoleUpdate() *discordgo.GuildRoleUpdate {
	return data.EvtInterface.(*discordgo.GuildRoleUpdate)
}
func (data *EventData) GuildStickersUpdate() *discordgo.GuildStickersUpdate {
	return data.EvtInterface.(*discordgo.GuildStickersUpdate)
}
func (data *EventData) GuildUpdate() *discordgo.GuildUpdate {
	return data.EvtInterface.(*discordgo.GuildUpdate)
}
func (data *EventData) InteractionCreate() *discordgo.InteractionCreate {
	return data.EvtInterface.(*discordgo.InteractionCreate)
}
func (data *EventData) InviteCreate() *discordgo.InviteCreate {
	return data.EvtInterface.(*discordgo.InviteCreate)
}
func (data *EventData) InviteDelete() *discordgo.InviteDelete {
	return data.EvtInterface.(*discordgo.InviteDelete)
}
func (data *EventData) MessageAck() *discordgo.MessageAck {
	return data.EvtInterface.(*discordgo.MessageAck)
}
func (data *EventData) MessageCreate() *discordgo.MessageCreate {
	return data.EvtInterface.(*discordgo.MessageCreate)
}
func (data *EventData) MessageDelete() *discordgo.MessageDelete {
	return data.EvtInterface.(*discordgo.MessageDelete)
}
func (data *EventData) MessageDeleteBulk() *discordgo.MessageDeleteBulk {
	return data.EvtInterface.(*discordgo.MessageDeleteBulk)
}
func (data *EventData) MessageReactionAdd() *discordgo.MessageReactionAdd {
	return data.EvtInterface.(*discordgo.MessageReactionAdd)
}
func (data *EventData) MessageReactionRemove() *discordgo.MessageReactionRemove {
	return data.EvtInterface.(*discordgo.MessageReactionRemove)
}
func (data *EventData) MessageReactionRemoveAll() *discordgo.MessageReactionRemoveAll {
	return data.EvtInterface.(*discordgo.MessageReactionRemoveAll)
}
func (data *EventData) MessageReactionRemoveEmoji() *discordgo.MessageReactionRemoveEmoji {
	return data.EvtInterface.(*discordgo.MessageReactionRemoveEmoji)
}
func (data *EventData) MessageUpdate() *discordgo.MessageUpdate {
	return data.EvtInterface.(*discordgo.MessageUpdate)
}
func (data *EventData) PresenceUpdate() *discordgo.PresenceUpdate {
	return data.EvtInterface.(*discordgo.PresenceUpdate)
}
func (data *EventData) PresencesReplace() *discordgo.PresencesReplace {
	return data.EvtInterface.(*discordgo.PresencesReplace)
}
func (data *EventData) RateLimit() *discordgo.RateLimit {
	return data.EvtInterface.(*discordgo.RateLimit)
}
func (data *EventData) Ready() *discordgo.Ready {
	return data.EvtInterface.(*discordgo.Ready)
}
func (data *EventData) RelationshipAdd() *discordgo.RelationshipAdd {
	return data.EvtInterface.(*discordgo.RelationshipAdd)
}
func (data *EventData) RelationshipRemove() *discordgo.RelationshipRemove {
	return data.EvtInterface.(*discordgo.RelationshipRemove)
}
func (data *EventData) Resumed() *discordgo.Resumed {
	return data.EvtInterface.(*discordgo.Resumed)
}
func (data *EventData) StageInstanceCreate() *discordgo.StageInstanceCreate {
	return data.EvtInterface.(*discordgo.StageInstanceCreate)
}
func (data *EventData) StageInstanceDelete() *discordgo.StageInstanceDelete {
	return data.EvtInterface.(*discordgo.StageInstanceDelete)
}
func (data *EventData) StageInstanceUpdate() *discordgo.StageInstanceUpdate {
	return data.EvtInterface.(*discordgo.StageInstanceUpdate)
}
func (data *EventData) ThreadCreate() *discordgo.ThreadCreate {
	return data.EvtInterface.(*discordgo.ThreadCreate)
}
func (data *EventData) ThreadDelete() *discordgo.ThreadDelete {
	return data.EvtInterface.(*discordgo.ThreadDelete)
}
func (data *EventData) ThreadListSync() *discordgo.ThreadListSync {
	return data.EvtInterface.(*discordgo.ThreadListSync)
}
func (data *EventData) ThreadMemberUpdate() *discordgo.ThreadMemberUpdate {
	return data.EvtInterface.(*discordgo.ThreadMemberUpdate)
}
func (data *EventData) ThreadMembersUpdate() *discordgo.ThreadMembersUpdate {
	return data.EvtInterface.(*discordgo.ThreadMembersUpdate)
}
func (data *EventData) ThreadUpdate() *discordgo.ThreadUpdate {
	return data.EvtInterface.(*discordgo.ThreadUpdate)
}
func (data *EventData) TypingStart() *discordgo.TypingStart {
	return data.EvtInterface.(*discordgo.TypingStart)
}
func (data *EventData) UserGuildSettingsUpdate() *discordgo.UserGuildSettingsUpdate {
	return data.EvtInterface.(*discordgo.UserGuildSettingsUpdate)
}
func (data *EventData) UserNoteUpdate() *discordgo.UserNoteUpdate {
	return data.EvtInterface.(*discordgo.UserNoteUpdate)
}
func (data *EventData) UserSettingsUpdate() *discordgo.UserSettingsUpdate {
	return data.EvtInterface.(*discordgo.UserSettingsUpdate)
}
func (data *EventData) UserUpdate() *discordgo.UserUpdate {
	return data.EvtInterface.(*discordgo.UserUpdate)
}
func (data *EventData) VoiceServerUpdate() *discordgo.VoiceServerUpdate {
	return data.EvtInterface.(*discordgo.VoiceServerUpdate)
}
func (data *EventData) VoiceStateUpdate() *discordgo.VoiceStateUpdate {
	return data.EvtInterface.(*discordgo.VoiceStateUpdate)
}
func (data *EventData) WebhooksUpdate() *discordgo.WebhooksUpdate {
	return data.EvtInterface.(*discordgo.WebhooksUpdate)
}

func fillEvent(evtData *EventData) {

	switch evtData.EvtInterface.(type) {
	case *discordgo.ApplicationCommandCreate:
		evtData.Type = Event(8)
	case *discordgo.ApplicationCommandDelete:
		evtData.Type = Event(9)
	case *discordgo.ApplicationCommandPermissionsUpdate:
		evtData.Type = Event(10)
	case *discordgo.ApplicationCommandUpdate:
		evtData.Type = Event(11)
	case *discordgo.AutoModerationActionExecution:
		evtData.Type = Event(12)
	case *discordgo.AutoModerationRuleCreate:
		evtData.Type = Event(13)
	case *discordgo.AutoModerationRuleDelete:
		evtData.Type = Event(14)
	case *discordgo.AutoModerationRuleUpdate:
		evtData.Type = Event(15)
	case *discordgo.ChannelCreate:
		evtData.Type = Event(16)
	case *discordgo.ChannelDelete:
		evtData.Type = Event(17)
	case *discordgo.ChannelPinsUpdate:
		evtData.Type = Event(18)
	case *discordgo.ChannelUpdate:
		evtData.Type = Event(19)
	case *discordgo.Connect:
		evtData.Type = Event(20)
	case *discordgo.Disconnect:
		evtData.Type = Event(21)
	case *discordgo.GuildBanAdd:
		evtData.Type = Event(22)
	case *discordgo.GuildBanRemove:
		evtData.Type = Event(23)
	case *discordgo.GuildCreate:
		evtData.Type = Event(24)
	case *discordgo.GuildDelete:
		evtData.Type = Event(25)
	case *discordgo.GuildEmojisUpdate:
		evtData.Type = Event(26)
	case *discordgo.GuildIntegrationsUpdate:
		evtData.Type = Event(27)
	case *discordgo.GuildMemberAdd:
		evtData.Type = Event(28)
	case *discordgo.GuildMemberRemove:
		evtData.Type = Event(29)
	case *discordgo.GuildMemberUpdate:
		evtData.Type = Event(30)
	case *discordgo.GuildMembersChunk:
		evtData.Type = Event(31)
	case *discordgo.GuildRoleCreate:
		evtData.Type = Event(32)
	case *discordgo.GuildRoleDelete:
		evtData.Type = Event(33)
	case *discordgo.GuildRoleUpdate:
		evtData.Type = Event(34)
	case *discordgo.GuildStickersUpdate:
		evtData.Type = Event(35)
	case *discordgo.GuildUpdate:
		evtData.Type = Event(36)
	case *discordgo.InteractionCreate:
		evtData.Type = Event(37)
	case *discordgo.InviteCreate:
		evtData.Type = Event(38)
	case *discordgo.InviteDelete:
		evtData.Type = Event(39)
	case *discordgo.MessageAck:
		evtData.Type = Event(40)
	case *discordgo.MessageCreate:
		evtData.Type = Event(41)
	case *discordgo.MessageDelete:
		evtData.Type = Event(42)
	case *discordgo.MessageDeleteBulk:
		evtData.Type = Event(43)
	case *discordgo.MessageReactionAdd:
		evtData.Type = Event(44)
	case *discordgo.MessageReactionRemove:
		evtData.Type = Event(45)
	case *discordgo.MessageReactionRemoveAll:
		evtData.Type = Event(46)
	case *discordgo.MessageReactionRemoveEmoji:
		evtData.Type = Event(47)
	case *discordgo.MessageUpdate:
		evtData.Type = Event(48)
	case *discordgo.PresenceUpdate:
		evtData.Type = Event(49)
	case *discordgo.PresencesReplace:
		evtData.Type = Event(50)
	case *discordgo.RateLimit:
		evtData.Type = Event(51)
	case *discordgo.Ready:
		evtData.Type = Event(52)
	case *discordgo.RelationshipAdd:
		evtData.Type = Event(53)
	case *discordgo.RelationshipRemove:
		evtData.Type = Event(54)
	case *discordgo.Resumed:
		evtData.Type = Event(55)
	case *discordgo.StageInstanceCreate:
		evtData.Type = Event(56)
	case *discordgo.StageInstanceDelete:
		evtData.Type = Event(57)
	case *discordgo.StageInstanceUpdate:
		evtData.Type = Event(58)
	case *discordgo.ThreadCreate:
		evtData.Type = Event(59)
	case *discordgo.ThreadDelete:
		evtData.Type = Event(60)
	case *discordgo.ThreadListSync:
		evtData.Type = Event(61)
	case *discordgo.ThreadMemberUpdate:
		evtData.Type = Event(62)
	case *discordgo.ThreadMembersUpdate:
		evtData.Type = Event(63)
	case *discordgo.ThreadUpdate:
		evtData.Type = Event(64)
	case *discordgo.TypingStart:
		evtData.Type = Event(65)
	case *discordgo.UserGuildSettingsUpdate:
		evtData.Type = Event(66)
	case *discordgo.UserNoteUpdate:
		evtData.Type = Event(67)
	case *discordgo.UserSettingsUpdate:
		evtData.Type = Event(68)
	case *discordgo.UserUpdate:
		evtData.Type = Event(69)
	case *discordgo.VoiceServerUpdate:
		evtData.Type = Event(70)
	case *discordgo.VoiceStateUpdate:
		evtData.Type = Event(71)
	case *discordgo.WebhooksUpdate:
		evtData.Type = Event(72)
	default:
		return
	}

	return
}
