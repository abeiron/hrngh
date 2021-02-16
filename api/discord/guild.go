// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains an implementation of Guild.

package discord

// VerificationLevel type definition
type VerificationLevel int

// Constants for VerificationLevel levels from 0 to 4 inclusive
const (
  VerificationLevelNone VerificationLevel = iota
  VerificationLevelLow
  VerificationLevelMedium
  VerificationLevelHigh
  VerificationLevelVeryHigh
)

// ExplicitContentFilterLevel type definition
type ExplicitContentFilterLevel int

// Constants for ExplicitContentFilterLevel levels from 0 to 2 inclusive
const (
  ExplicitContentFilterDisabled ExplicitContentFilterLevel = iota
  ExplicitContentFilterMembersWithoutRoles
  ExplicitContentFilterAllMembers
)

// MfaLevel type definition
type MfaLevel int

// Constants for MfaLevel levels from 0 to 1 inclusive
const (
  MfaLevelNone MfaLevel = iota
  MfaLevelElevated
)

// PremiumTier type definition
type PremiumTier int

// Constants for PremiumTier levels from 0 to 3 inclusive
const (
  PremiumTierNone PremiumTier = iota
  PremiumTier1
  PremiumTier2
  PremiumTier3
)

// A Guild holds all data related to a specific Discord Guild.  Guilds are also
// sometimes referred to as Servers in the Discord client.
type Guild struct {
  // The ID of the guild.
  ID string `json:"id"`

  // The name of the guild. (2–100 characters)
  Name string `json:"name"`

  // The hash of the guild's icon. Use Session.GuildIcon
  // to retrieve the icon itself.
  Icon string `json:"icon"`

  // The voice region of the guild.
  Region string `json:"region"`

  // The ID of the AFK voice channel.
  AfkChannelID string `json:"afk_channel_id"`

  // The user ID of the owner of the guild.
  OwnerID string `json:"owner_id"`

  // If we are the owner of the guild
  Owner bool `json:"owner"`

  // The time at which the current user joined the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  JoinedAt Timestamp `json:"joined_at"`

  // The hash of the guild's discovery splash.
  DiscoverySplash string `json:"discovery_splash"`

  // The hash of the guild's splash.
  Splash string `json:"splash"`

  // The timeout, in seconds, before a user is considered AFK in voice.
  AfkTimeout int `json:"afk_timeout"`

  // The number of members in the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  MemberCount int `json:"member_count"`

  // The verification level required for the guild.
  VerificationLevel VerificationLevel `json:"verification_level"`

  // Whether the guild is considered large. This is
  // determined by a member threshold in the identify packet,
  // and is currently hard-coded at 250 members in the library.
  Large bool `json:"large"`

  // The default message notification setting for the guild.
  DefaultMessageNotifications MessageNotifications `json:"default_message_notifications"`

  // A list of roles in the guild.
  Roles []*Role `json:"roles"`

  // A list of the custom emojis present in the guild.
  Emojis []*Emoji `json:"emojis"`

  // A list of the members in the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  Members []*Member `json:"members"`

  // A list of partial presence objects for members in the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  Presences []*Presence `json:"presences"`

  // The maximum number of presences for the guild (the default value, currently 25000, is in effect when null is returned)
  MaxPresences int `json:"max_presences"`

  // The maximum number of members for the guild
  MaxMembers int `json:"max_members"`

  // A list of channels in the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  Channels []*Channel `json:"channels"`

  // A list of voice states for the guild.
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  VoiceStates []*VoiceState `json:"voice_states"`

  // Whether this guild is currently unavailable (most likely due to outage).
  // This field is only present in GUILD_CREATE events and websocket
  // update events, and thus is only present in state-cached guilds.
  Unavailable bool `json:"unavailable"`

  // The explicit content filter level
  ExplicitContentFilter ExplicitContentFilterLevel `json:"explicit_content_filter"`

  // The list of enabled guild features
  Features []string `json:"features"`

  // Required MFA level for the guild
  MfaLevel MfaLevel `json:"mfa_level"`

  // The application id of the guild if bot created.
  ApplicationID string `json:"application_id"`

  // Whether or not the Server Widget is enabled
  WidgetEnabled bool `json:"widget_enabled"`

  // The Channel ID for the Server Widget
  WidgetChannelID string `json:"widget_channel_id"`

  // The Channel ID to which system messages are sent (eg join and leave messages)
  SystemChannelID string `json:"system_channel_id"`

  // The System channel flags
  SystemChannelFlags SystemChannelFlag `json:"system_channel_flags"`

  // The ID of the rules channel ID, used for rules.
  RulesChannelID string `json:"rules_channel_id"`

  // the vanity url code for the guild
  VanityURLCode string `json:"vanity_url_code"`

  // the description for the guild
  Description string `json:"description"`

  // The hash of the guild's banner
  Banner string `json:"banner"`

  // The premium tier of the guild
  PremiumTier PremiumTier `json:"premium_tier"`

  // The total number of users currently boosting this server
  PremiumSubscriptionCount int `json:"premium_subscription_count"`

  // The preferred locale of a guild with the "PUBLIC" feature; used in server discovery and notices from Discord; defaults to "en-US"
  PreferredLocale string `json:"preferred_locale"`

  // The id of the channel where admins and moderators of guilds with the "PUBLIC" feature receive notices from Discord
  PublicUpdatesChannelID string `json:"public_updates_channel_id"`

  // The maximum amount of users in a video channel
  MaxVideoChannelUsers int `json:"max_video_channel_users"`

  // Approximate number of members in this guild, returned from the GET /guild/<id> endpoint when with_counts is true
  ApproximateMemberCount int `json:"approximate_member_count"`

  // Approximate number of non-offline members in this guild, returned from the GET /guild/<id> endpoint when with_counts is true
  ApproximatePresenceCount int `json:"approximate_presence_count"`

  // Permissions of our user
  Permissions int64 `json:"permissions,string"`
}

// MessageNotifications is the notification level for a guild
// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
type MessageNotifications int

// Block containing known MessageNotifications values
const (
  MessageNotificationsAllMessages MessageNotifications = iota
  MessageNotificationsOnlyMentions
)

// SystemChannelFlag is the type of flags in the system channel (see SystemChannelFlag* consts)
// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
type SystemChannelFlag int

// Block containing known SystemChannelFlag values
const (
  SystemChannelFlagsSuppressJoin SystemChannelFlag = 1 << iota
  SystemChannelFlagsSuppressPremium
)

// IconURL returns a URL to the guild's icon.
func (g *Guild) IconURL() string {
  if g.Icon == "" {
    return ""
  }

  if strings.HasPrefix(g.Icon, "a_") {
    return EndpointGuildIconAnimated(g.ID, g.Icon)
  }

  return EndpointGuildIcon(g.ID, g.Icon)
}

// A UserGuild holds a brief version of a Guild
type UserGuild struct {
  ID          string `json:"id"`
  Name        string `json:"name"`
  Icon        string `json:"icon"`
  Owner       bool   `json:"owner"`
  Permissions int64  `json:"permissions,string"`
}

// A GuildParams stores all the data needed to update discord guild settings
type GuildParams struct {
  Name                        string             `json:"name,omitempty"`
  Region                      string             `json:"region,omitempty"`
  VerificationLevel           *VerificationLevel `json:"verification_level,omitempty"`
  DefaultMessageNotifications int                `json:"default_message_notifications,omitempty"` // TODO: Separate type?
  AfkChannelID                string             `json:"afk_channel_id,omitempty"`
  AfkTimeout                  int                `json:"afk_timeout,omitempty"`
  Icon                        string             `json:"icon,omitempty"`
  OwnerID                     string             `json:"owner_id,omitempty"`
  Splash                      string             `json:"splash,omitempty"`
  Banner                      string             `json:"banner,omitempty"`
}

// A GuildRole stores data for guild roles.
type GuildRole struct {
  Role    *Role  `json:"role"`
  GuildID string `json:"guild_id"`
}

// A GuildBan stores data for a guild ban.
type GuildBan struct {
  Reason string `json:"reason"`
  User   *User  `json:"user"`
}

// A GuildEmbed stores data for a guild embed.
type GuildEmbed struct {
  Enabled   bool   `json:"enabled"`
  ChannelID string `json:"channel_id"`
}

// A GuildAuditLog stores data for a guild audit log.
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type GuildAuditLog struct {
  Webhooks        []*Webhook       `json:"webhooks,omitempty"`
  Users           []*User          `json:"users,omitempty"`
  AuditLogEntries []*AuditLogEntry `json:"audit_log_entries"`
  Integrations    []*Integration   `json:"integrations"`
}

// AuditLogEntry for a GuildAuditLog
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-entry-structure
type AuditLogEntry struct {
  TargetID   string            `json:"target_id"`
  Changes    []*AuditLogChange `json:"changes"`
  UserID     string            `json:"user_id"`
  ID         string            `json:"id"`
  ActionType *AuditLogAction   `json:"action_type"`
  Options    *AuditLogOptions  `json:"options"`
  Reason     string            `json:"reason"`
}

// AuditLogChange for an AuditLogEntry
type AuditLogChange struct {
  NewValue interface{}        `json:"new_value"`
  OldValue interface{}        `json:"old_value"`
  Key      *AuditLogChangeKey `json:"key"`
}

// AuditLogChangeKey value for AuditLogChange.
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-key
type AuditLogChangeKey string

// Block of valid AuditLogChangeKey
const (
  AuditLogChangeKeyName                       AuditLogChangeKey = "name"
  AuditLogChangeKeyIconHash                   AuditLogChangeKey = "icon_hash"
  AuditLogChangeKeySplashHash                 AuditLogChangeKey = "splash_hash"
  AuditLogChangeKeyOwnerID                    AuditLogChangeKey = "owner_id"
  AuditLogChangeKeyRegion                     AuditLogChangeKey = "region"
  AuditLogChangeKeyAfkChannelID               AuditLogChangeKey = "afk_channel_id"
  AuditLogChangeKeyAfkTimeout                 AuditLogChangeKey = "afk_timeout"
  AuditLogChangeKeyMfaLevel                   AuditLogChangeKey = "mfa_level"
  AuditLogChangeKeyVerificationLevel          AuditLogChangeKey = "verification_level"
  AuditLogChangeKeyExplicitContentFilter      AuditLogChangeKey = "explicit_content_filter"
  AuditLogChangeKeyDefaultMessageNotification AuditLogChangeKey = "default_message_notifications"
  AuditLogChangeKeyVanityURLCode              AuditLogChangeKey = "vanity_url_code"
  AuditLogChangeKeyRoleAdd                    AuditLogChangeKey = "$add"
  AuditLogChangeKeyRoleRemove                 AuditLogChangeKey = "$remove"
  AuditLogChangeKeyPruneDeleteDays            AuditLogChangeKey = "prune_delete_days"
  AuditLogChangeKeyWidgetEnabled              AuditLogChangeKey = "widget_enabled"
  AuditLogChangeKeyWidgetChannelID            AuditLogChangeKey = "widget_channel_id"
  AuditLogChangeKeySystemChannelID            AuditLogChangeKey = "system_channel_id"
  AuditLogChangeKeyPosition                   AuditLogChangeKey = "position"
  AuditLogChangeKeyTopic                      AuditLogChangeKey = "topic"
  AuditLogChangeKeyBitrate                    AuditLogChangeKey = "bitrate"
  AuditLogChangeKeyPermissionOverwrite        AuditLogChangeKey = "permission_overwrites"
  AuditLogChangeKeyNSFW                       AuditLogChangeKey = "nsfw"
  AuditLogChangeKeyApplicationID              AuditLogChangeKey = "application_id"
  AuditLogChangeKeyRateLimitPerUser           AuditLogChangeKey = "rate_limit_per_user"
  AuditLogChangeKeyPermissions                AuditLogChangeKey = "permissions"
  AuditLogChangeKeyColor                      AuditLogChangeKey = "color"
  AuditLogChangeKeyHoist                      AuditLogChangeKey = "hoist"
  AuditLogChangeKeyMentionable                AuditLogChangeKey = "mentionable"
  AuditLogChangeKeyAllow                      AuditLogChangeKey = "allow"
  AuditLogChangeKeyDeny                       AuditLogChangeKey = "deny"
  AuditLogChangeKeyCode                       AuditLogChangeKey = "code"
  AuditLogChangeKeyChannelID                  AuditLogChangeKey = "channel_id"
  AuditLogChangeKeyInviterID                  AuditLogChangeKey = "inviter_id"
  AuditLogChangeKeyMaxUses                    AuditLogChangeKey = "max_uses"
  AuditLogChangeKeyUses                       AuditLogChangeKey = "uses"
  AuditLogChangeKeyMaxAge                     AuditLogChangeKey = "max_age"
  AuditLogChangeKeyTempoary                   AuditLogChangeKey = "temporary"
  AuditLogChangeKeyDeaf                       AuditLogChangeKey = "deaf"
  AuditLogChangeKeyMute                       AuditLogChangeKey = "mute"
  AuditLogChangeKeyNick                       AuditLogChangeKey = "nick"
  AuditLogChangeKeyAvatarHash                 AuditLogChangeKey = "avatar_hash"
  AuditLogChangeKeyID                         AuditLogChangeKey = "id"
  AuditLogChangeKeyType                       AuditLogChangeKey = "type"
  AuditLogChangeKeyEnableEmoticons            AuditLogChangeKey = "enable_emoticons"
  AuditLogChangeKeyExpireBehavior             AuditLogChangeKey = "expire_behavior"
  AuditLogChangeKeyExpireGracePeriod          AuditLogChangeKey = "expire_grace_period"
)

// AuditLogOptions optional data for the AuditLog.
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogOptions struct {
  DeleteMemberDays string               `json:"delete_member_days"`
  MembersRemoved   string               `json:"members_removed"`
  ChannelID        string               `json:"channel_id"`
  MessageID        string               `json:"message_id"`
  Count            string               `json:"count"`
  ID               string               `json:"id"`
  Type             *AuditLogOptionsType `json:"type"`
  RoleName         string               `json:"role_name"`
}

// AuditLogOptionsType of the AuditLogOption.
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogOptionsType string

// Valid Types for AuditLogOptionsType.
const (
  AuditLogOptionsTypeMember AuditLogOptionsType = "member"
  AuditLogOptionsTypeRole   AuditLogOptionsType = "role"
)

// AuditLogAction is the Action of the AuditLog (see AuditLogAction* consts).
//
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
type AuditLogAction int

// Block contains Discord Audit Log Action Types
const (
  AuditLogActionGuildUpdate AuditLogAction = 1

  AuditLogActionChannelCreate          AuditLogAction = 10
  AuditLogActionChannelUpdate          AuditLogAction = 11
  AuditLogActionChannelDelete          AuditLogAction = 12
  AuditLogActionChannelOverwriteCreate AuditLogAction = 13
  AuditLogActionChannelOverwriteUpdate AuditLogAction = 14
  AuditLogActionChannelOverwriteDelete AuditLogAction = 15

  AuditLogActionMemberKick       AuditLogAction = 20
  AuditLogActionMemberPrune      AuditLogAction = 21
  AuditLogActionMemberBanAdd     AuditLogAction = 22
  AuditLogActionMemberBanRemove  AuditLogAction = 23
  AuditLogActionMemberUpdate     AuditLogAction = 24
  AuditLogActionMemberRoleUpdate AuditLogAction = 25

  AuditLogActionRoleCreate AuditLogAction = 30
  AuditLogActionRoleUpdate AuditLogAction = 31
  AuditLogActionRoleDelete AuditLogAction = 32

  AuditLogActionInviteCreate AuditLogAction = 40
  AuditLogActionInviteUpdate AuditLogAction = 41
  AuditLogActionInviteDelete AuditLogAction = 42

  AuditLogActionWebhookCreate AuditLogAction = 50
  AuditLogActionWebhookUpdate AuditLogAction = 51
  AuditLogActionWebhookDelete AuditLogAction = 52

  AuditLogActionEmojiCreate AuditLogAction = 60
  AuditLogActionEmojiUpdate AuditLogAction = 61
  AuditLogActionEmojiDelete AuditLogAction = 62

  AuditLogActionMessageDelete     AuditLogAction = 72
  AuditLogActionMessageBulkDelete AuditLogAction = 73
  AuditLogActionMessagePin        AuditLogAction = 74
  AuditLogActionMessageUnpin      AuditLogAction = 75

  AuditLogActionIntegrationCreate AuditLogAction = 80
  AuditLogActionIntegrationUpdate AuditLogAction = 81
  AuditLogActionIntegrationDelete AuditLogAction = 82
)

// A VoiceRegion stores data for a specific voice region server.
type VoiceRegion struct {
  ID       string `json:"id"`
  Name     string `json:"name"`
  Hostname string `json:"sample_hostname"`
  Port     int    `json:"sample_port"`
}

// A ICEServer stores data for a specific voice ICE server.
type ICEServer struct {
  URL        string `json:"url"`
  Username   string `json:"username"`
  Credential string `json:"credential"`
}

// A Invite stores all data related to a specific Discord Guild or Channel invite.
type Invite struct {
  Guild          *Guild         `json:"guild"`
  Channel        *Channel       `json:"channel"`
  Inviter        *User          `json:"inviter"`
  Code           string         `json:"code"`
  CreatedAt      Timestamp      `json:"created_at"`
  MaxAge         int            `json:"max_age"`
  Uses           int            `json:"uses"`
  MaxUses        int            `json:"max_uses"`
  Revoked        bool           `json:"revoked"`
  Temporary      bool           `json:"temporary"`
  Unique         bool           `json:"unique"`
  TargetUser     *User          `json:"target_user"`
  TargetUserType TargetUserType `json:"target_user_type"`

  // will only be filled when using InviteWithCounts
  ApproximatePresenceCount int `json:"approximate_presence_count"`
  ApproximateMemberCount   int `json:"approximate_member_count"`
}

// TargetUserType is the type of the target user.
//
// https://discord.com/developers/docs/resources/invite#invite-object-target-user-types
type TargetUserType int

// Block contains known TargetUserType values
const (
  TargetUserTypeStream TargetUserType = iota
)
