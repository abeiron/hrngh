// Discord bindings for the Hrngh bot.
// Available at https://github.com/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to error-handling.

package discord

// An APIErrorMessage is an api error message returned from discord
type APIErrorMessage struct {
  Code    int    `json:"code"`
  Message string `json:"message"`
}


// Block contains Discord JSON Error Response codes
const (
  ErrCodeUnknownAccount     = 10001
  ErrCodeUnknownApplication = 10002
  ErrCodeUnknownChannel     = 10003
  ErrCodeUnknownGuild       = 10004
  ErrCodeUnknownIntegration = 10005
  ErrCodeUnknownInvite      = 10006
  ErrCodeUnknownMember      = 10007
  ErrCodeUnknownMessage     = 10008
  ErrCodeUnknownOverwrite   = 10009
  ErrCodeUnknownProvider    = 10010
  ErrCodeUnknownRole        = 10011
  ErrCodeUnknownToken       = 10012
  ErrCodeUnknownUser        = 10013
  ErrCodeUnknownEmoji       = 10014
  ErrCodeUnknownWebhook     = 10015
  ErrCodeUnknownBan         = 10026

  ErrCodeBotsCannotUseEndpoint  = 20001
  ErrCodeOnlyBotsCanUseEndpoint = 20002

  ErrCodeMaximumGuildsReached     = 30001
  ErrCodeMaximumFriendsReached    = 30002
  ErrCodeMaximumPinsReached       = 30003
  ErrCodeMaximumGuildRolesReached = 30005
  ErrCodeTooManyReactions         = 30010

  ErrCodeUnauthorized = 40001

  ErrCodeMissingAccess                             = 50001
  ErrCodeInvalidAccountType                        = 50002
  ErrCodeCannotExecuteActionOnDMChannel            = 50003
  ErrCodeEmbedDisabled                             = 50004
  ErrCodeCannotEditFromAnotherUser                 = 50005
  ErrCodeCannotSendEmptyMessage                    = 50006
  ErrCodeCannotSendMessagesToThisUser              = 50007
  ErrCodeCannotSendMessagesInVoiceChannel          = 50008
  ErrCodeChannelVerificationLevelTooHigh           = 50009
  ErrCodeOAuth2ApplicationDoesNotHaveBot           = 50010
  ErrCodeOAuth2ApplicationLimitReached             = 50011
  ErrCodeInvalidOAuthState                         = 50012
  ErrCodeMissingPermissions                        = 50013
  ErrCodeInvalidAuthenticationToken                = 50014
  ErrCodeNoteTooLong                               = 50015
  ErrCodeTooFewOrTooManyMessagesToDelete           = 50016
  ErrCodeCanOnlyPinMessageToOriginatingChannel     = 50019
  ErrCodeCannotExecuteActionOnSystemMessage        = 50021
  ErrCodeMessageProvidedTooOldForBulkDelete        = 50034
  ErrCodeInvalidFormBody                           = 50035
  ErrCodeInviteAcceptedToGuildApplicationsBotNotIn = 50036

  ErrCodeReactionBlocked = 90001
)
