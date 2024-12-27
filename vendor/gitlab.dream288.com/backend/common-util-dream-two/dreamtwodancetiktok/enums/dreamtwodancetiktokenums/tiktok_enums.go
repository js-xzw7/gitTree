package dreamtwodancetiktokenums

const (
	BarrageLiveComment = "live_comment"
	BarrageLiveGift    = "live_gift"
	BarrageLiveLike    = "live_like"

	BarrageLiveFansclub = "live_fansclub"

	BaseScore      = 100
	BaseScoreError = 10

	OnRedTeam  = 1
	OnBlueTeam = 2

	ARight = 1
	BRight = 2

	WhoWinRedTeam  = 1
	WhoWinBlueTeam = 2
	WhoWinSomeTeam = -1

	UserSexMan   = 1
	UserSexWoman = 2

	// 推送中奖记录通知,道具类型
	LuckyDrawPropTypeClothesProps = 1 // 1:服饰,道具,工具
	LuckyDrawPropTypeScore        = 2 // 2:星光值
	LuckyDrawPropTypeTripScore    = 3 // 3:旅途点

	SyncBarrageTmpData                = 1
	SyncUserRedis2Mongo               = 2
	SyncRankData                      = 3
	SyncRankDataWings                 = 4
	SyncActivitiesUser                = 5
	SyncCombinationCleaning           = 6
	SyncRankCoupleTitleToPkg          = 7
	SyncRankComposeData               = 8
	SynSingleAssembleData             = 9
	SyncBalloonUserRevenue            = 10
	SyncDistributeHalo                = 11
	SyncDistributeStarPoints          = 12
	SyncElementReissue                = 13
	SyncDanceModelEndUserDetailsToDay = 14
	SyncMongoDrawUserLogToCk          = 15
	SyncTripUserLevelChange           = 16
	SyncActivityDownloadTask          = 17
	SyncActivitySplinterDownloadTask  = 18
	SyncAggregateData                 = 19
	SyncResetPetScore                 = 20
	SyncResetPetGrade                 = 21
	SyncResetPetBattle                = 22
	SyncBetaMonthClean                = 23
	SyncBetaSaveUserUsedDan           = 24
	SyncBetaSendUserUsedDan           = 25
	SyncBetaSaveUserPetGrade          = 26
	SyncBetaSendUserPetGradePrize     = 27
)
