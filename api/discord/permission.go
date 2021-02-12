// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to Discord permissions.

package discord

// Constants for the different bit offsets of text channel permissions
const (
  // Deprecated: PermissionReadMessages has been replaced with PermissionViewChannel for text and voice channels
  PermissionReadMessages = 1 << (iota + 10)
  PermissionSendMessages
  PermissionSendTTSMessages
  PermissionManageMessages
  PermissionEmbedLinks
  PermissionAttachFiles
  PermissionReadMessageHistory
  PermissionMentionEveryone
  PermissionUseExternalEmojis
)

// Constants for the different bit offsets of voice permissions
const (
  PermissionVoiceConnect = 1 << (iota + 20)
  PermissionVoiceSpeak
  PermissionVoiceMuteMembers
  PermissionVoiceDeafenMembers
  PermissionVoiceMoveMembers
  PermissionVoiceUseVAD
  PermissionVoicePrioritySpeaker = 1 << (iota + 2)
)

// Constants for general management.
const (
  PermissionChangeNickname = 1 << (iota + 26)
  PermissionManageNicknames
  PermissionManageRoles
  PermissionManageWebhooks
  PermissionManageEmojis
)

// Constants for the different bit offsets of general permissions
const (
  PermissionCreateInstantInvite = 1 << iota
  PermissionKickMembers
  PermissionBanMembers
  PermissionAdministrator
  PermissionManageChannels
  PermissionManageServer
  PermissionAddReactions
  PermissionViewAuditLogs
  PermissionViewChannel = 1 << (iota + 2)

  PermissionAllText = PermissionViewChannel |
    PermissionSendMessages |
    PermissionSendTTSMessages |
    PermissionManageMessages |
    PermissionEmbedLinks |
    PermissionAttachFiles |
    PermissionReadMessageHistory |
    PermissionMentionEveryone
  PermissionAllVoice = PermissionViewChannel |
    PermissionVoiceConnect |
    PermissionVoiceSpeak |
    PermissionVoiceMuteMembers |
    PermissionVoiceDeafenMembers |
    PermissionVoiceMoveMembers |
    PermissionVoiceUseVAD |
    PermissionVoicePrioritySpeaker
  PermissionAllChannel = PermissionAllText |
    PermissionAllVoice |
    PermissionCreateInstantInvite |
    PermissionManageRoles |
    PermissionManageChannels |
    PermissionAddReactions |
    PermissionViewAuditLogs
  PermissionAll = PermissionAllChannel |
    PermissionKickMembers |
    PermissionBanMembers |
    PermissionManageServer |
    PermissionAdministrator |
    PermissionManageWebhooks |
    PermissionManageEmojis
)

// PermissionOverwriteType represents the type of resource on which
// a permission overwrite acts.
type PermissionOverwriteType int

// The possible permission overwrite types.
const (
  PermissionOverwriteTypeRole PermissionOverwriteType = iota
  PermissionOverwriteTypeMember
)

// A PermissionOverwrite holds permission overwrite data for a Channel
type PermissionOverwrite struct {
  ID    string                  `json:"id"`
  Type  PermissionOverwriteType `json:"type"`
  Deny  int64                   `json:"deny,string"`
  Allow int64                   `json:"allow,string"`
}
