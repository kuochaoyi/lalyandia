package gsop

const (
	SendProtocolVersion uint16 = iota
	MoveBackwardToLocation
	Say
	RequestEnterWorld
	Action
	RequestLogin
	SendLogOut
	RequestAttack
	RequestCharacterCreate
	RequestCharacterDelete
	RequestGameStart
	RequestNewCharacter
	RequestItemList
	RequestEquipItem
	RequestUnEquipItem
	RequestDropItem
	RequestUseItem
	RequestTrade
	RequestAddTradeItem
	TradeDone
	RequestTeleport
	SocialAction
	ChangeMoveType
	ChangeWaitType
	RequestSellItem
	RequestBuyItem
	RequestLinkHtml
	RequestBypassToServer
	RequestBBSWrite
	RequestCreatePledge
	RequestJoinPledge
	RequestAnswerJoinPledge
	RequestWithDrawalPledge
	RequestOustPledgeMember
	RequestDismissPledge
	RequestJoinParty
	RequestAnswerJoinParty
	RequestWithDrawalParty
	RequestOustPartyMember
	RequestDismissParty
	UserAck
	RequestMagicSkillUse
	SendAppearing
	SendWareHouseDepositList
	SendWareHouseWithdrawList
	RequestShortCutReg
	RequestShortCutUse
	RequestShortCutDel
	CanNotMoveAnymore
	RequestTargetCancel
	Say2
	RequestPledgeMemberList
	RequestMagicList
	RequestSkillList
	MoveWithDelta
	GetOnVehicle
	GetOffVehicle
	AnswerTradeRequest
	RequestActionUse
	RequestRestart
	RequestSiegeInfo
	ValidateLocation
	RequestSEKCustom
	StartRotating
	FinishRotating
	RequestStartPledgeWar
	RequestReplyStartPledgeWar
	RequestStopPledgeWar
	RequestReplyStopPledgeWar
	RequestSurrenderPledgeWar
	RequestReplySurrenderPledgeWar
	RequestSetPledgeCrest
	RequestGiveNickName
	RequestShowboard
	RequestEnchantItem
	RequestDestroyItem
	SendBypassBuildCmd
	MoveToLocationInVehicle
	CanNotMoveAnymoreVehicle
	RequestFriendInvite
	RequestFriendAddReply
	RequestFriendInfoList
	RequestFriendDel
	RequestCharacterRestore
	RequestQuestList
	RequestDestroyQuest
	RequestPledgeInfo
	RequestPledgeExtendedInfo
	RequestPledgeCrest
	RequestSurrenderPersonally
	RequestRide
	RequestAcquireSkillInfo
	RequestAcquireSkill
	RequestRestartPoint
	RequestGMCommand
	RequestListPartyWaiting
	RequestManagePartyRoom
	RequestJoinPartyRoom
	RequestCrystallizeItem
	RequestPrivateStoreSellManageList
	SetPrivateStoreSellList
	RequestPrivateStoreSellManageCancel
	RequestPrivateStoreSellQuit
	SetPrivateStoreSellMsg
	SendPrivateStoreBuyList
	RequestReviveReply
	RequestTutorialLinkHtml
	RequestTutorialPassCmdToServer
	RequestTutorialQuestionMarkPressed
	RequestTutorialClientEvent
	RequestPetition
	RequestPetitionCancel
	RequestGMList
	RequestJoinAlly
	RequestAnswerJoinAlly
	RequestWithdrawAlly
	RequestOustAlly
	RequestDismissAlly
	RequestSetAllyCrest
	RequestAllyCrest
	RequestChangePetName
	RequestPetUseItem
	RequestGiveItemToPet
	RequestGetItemFromPet
	RequestAllyInfo
	RequestPetGetItem
	RequestPrivateStoreBuyManageList
	SetPrivateStoreBuyList
	RequestPrivateStoreBuyManageQuit
	SetPrivateStoreBuyMsg
	SendPrivateStoreSellList
	SendTimeCheck
	RequestStartAllianceWar
	ReplyStartAllianceWar
	RequestStopAllianceWar
	ReplyStopAllianceWar
	RequestSurrenderAllianceWar
	RequestSkillCoolTime
	RequestPackageSendableItemList
	RequestPackageSend
	RequestBlock
	RequestCastleSiegeInfo
	RequestCastleSiegeAttackerList
	RequestCastleSiegeDefenderList
	RequestJoinCastleSiege
	RequestConfirmCastleSiegeWaitingList
	RequestSetCastleSiegeTime
	RequestMultiSellChoose
	NetPing
	RequestRemainTime
	BypassUserCmd
	GMSnoopEnd
	RequestRecipeBookOpen
	RequestRecipeItemDelete
	RequestRecipeItemMakeInfo
	RequestRecipeItemMakeSelf
	RequestRecipeShopManageList
	RequestRecipeShopMessageSet
	RequestRecipeShopListSet
	RequestRecipeShopManageQuit
	RequestRecipeShopManageCancel
	RequestRecipeShopMakeInfo
	RequestRecipeShopMakeDo
	RequestRecipeShopSellList
	RequestObserverEnd
	VoteSociality
	RequestHennaItemList
	RequestHennaItemInfo
	RequestHennaEquip
	RequestHennaUnequipList
	RequestHennaUnequipInfo
	RequestHennaUnequip
	RequestPledgePower
	RequestMakeMacro
	RequestDeleteMacro
	RequestProcureCrop
	RequestBuySeed
	ConfirmDlg
	RequestPreviewItem
	RequestSSQStatus
	PetitionVote
	ReplyGameGuardQuery
	RequestSendL2FriendSay
	RequestOpenMinimap
	RequestSendMsnChatLog
	RequestReload
	RequestOustFromPartyRoom
	RequestDismissPartyRoom
	RequestWithdrawPartyRoom
	RequestHandOverPartyMaster
	RequestAutoSoulShot
	RequestExEnchantSkillInfo
	RequestExEnchantSkill
	RequestManorList
	RequestProcureCropList
	RequestSetSeed
	RequestSetCrop
	RequestWriteHeroWords
	RequestExAskJoinMPCC
	RequestExAcceptJoinMPCC
	RequestExOustFromMPCC
	RequestExPledgeCrestLarge
	RequestExSetPledgeCrestLarge
	RequestOlympiadObserverEnd
	RequestOlympiadMatchList
	RequestAskJoinPartyRoom
	AnswerJoinPartyRoom
	RequestListPartyMatchingWaitingRoom
	RequestExitPartyMatchingWaitingRoom
	RequestGetBossRecord
	RequestPledgeSetAcademyMaster
	RequestPledgePowerGradeList
	RequestPledgeMemberPowerInfo
	RequestPledgeSetMemberPowerGrade
	RequestPledgeMemberInfo
	RequestPledgeWarList
	RequestExFishRanking
	RequestPCCafeCouponUse
	RequestCursedWeaponList
	RequestCursedWeaponLocation
	RequestPledgeReorganizeMember
	RequestExMPCCShowPartyMembersInfo
	RequestDuelStart
	RequestDuelAnswerStart
	RequestConfirmTargetItem
	RequestConfirmRefinerItem
	RequestConfirmGemStone
	RequestRefine
	RequestConfirmCancelItem
	RequestRefineCancel
	RequestExMagicSkillUseGround
	RequestDuelSurrender
	RequestExChangeName
)
