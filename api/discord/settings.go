// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to Discord's configuration.

package discord

// A Settings stores data for a specific users Discord client settings.
type Settings struct {
  RenderEmbeds           bool               `json:"render_embeds"`
  InlineEmbedMedia       bool               `json:"inline_embed_media"`
  InlineAttachmentMedia  bool               `json:"inline_attachment_media"`
  EnableTTSCommand       bool               `json:"enable_tts_command"`
  MessageDisplayCompact  bool               `json:"message_display_compact"`
  ShowCurrentGame        bool               `json:"show_current_game"`
  ConvertEmoticons       bool               `json:"convert_emoticons"`
  Locale                 string             `json:"locale"`
  Theme                  string             `json:"theme"`
  GuildPositions         []string           `json:"guild_positions"`
  RestrictedGuilds       []string           `json:"restricted_guilds"`
  FriendSourceFlags      *FriendSourceFlags `json:"friend_source_flags"`
  Status                 Status             `json:"status"`
  DetectPlatformAccounts bool               `json:"detect_platform_accounts"`
  DeveloperMode          bool               `json:"developer_mode"`
}


// A UserGuildSettingsChannelOverride stores data for a channel override for a users guild settings.
type UserGuildSettingsChannelOverride struct {
  Muted                bool   `json:"muted"`
  MessageNotifications int    `json:"message_notifications"`
  ChannelID            string `json:"channel_id"`
}

// A UserGuildSettings stores data for a users guild settings.
type UserGuildSettings struct {
  SupressEveryone      bool                                `json:"suppress_everyone"`
  Muted                bool                                `json:"muted"`
  MobilePush           bool                                `json:"mobile_push"`
  MessageNotifications int                                 `json:"message_notifications"`
  GuildID              string                              `json:"guild_id"`
  ChannelOverrides     []*UserGuildSettingsChannelOverride `json:"channel_overrides"`
}

// A UserGuildSettingsEdit stores data for editing UserGuildSettings
type UserGuildSettingsEdit struct {
  SupressEveryone      bool                                         `json:"suppress_everyone"`
  Muted                bool                                         `json:"muted"`
  MobilePush           bool                                         `json:"mobile_push"`
  MessageNotifications int                                          `json:"message_notifications"`
  ChannelOverrides     map[string]*UserGuildSettingsChannelOverride `json:"channel_overrides"`
}
