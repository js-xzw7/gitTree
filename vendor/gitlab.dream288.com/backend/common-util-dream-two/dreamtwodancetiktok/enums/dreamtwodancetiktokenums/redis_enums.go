package dreamtwodancetiktokenums

const (
	UserInfoHash = "%s:%s:user:userinfo:%s" //陌陌用户信息表
	UserAllHas   = "%s:%s:user_hash"        //陌陌用户列表，存储在zSet中，score连胜场次作为积分

	//RoundZSet                = "%s:%s:zSet:currentRound:%s"         //当局连胜排行榜
	RoundString              = "%s:%s:String:currentRound:%s"       //当局连胜排行榜
	CurrentWinningNumberHash = "%s:%s:hash:currentWinningNumber:%s" //当局连胜次数
	TotalZSet                = "%s:%s:zSet:Total"                   //总排行榜
	WinningStreakZSet        = "%s:%s:zSet:WinningStreakZSet"       //历史连胜排行榜
	WinningTotalZSet         = "%s:%s:zSet:WinningTotalZSet"        //历史连胜排行榜

	ClientHashConnection     = "%s:%s:Hash:Connection"               //记录客户端连接状态
	NeedDeleteHostNoDataPool = "%s:%s:Hash:NeedDeleteHostNoDataPool" //记录客户端连接状态

	PlatformTikTok = "twoTiktok" //抖音
	PlatformSpeedy = "twoSpeedy" //快手

	RoomProblemHash     = "%s:%s:hash:roomProblem:%s"         //局对应的问题
	RoomProblemInfoHash = "%s:%s:hash:roomProblemInfo:%s:%+v" //局对应的问题

	PipelineBatchNumber = 1000 //快手

	ScoreTotalZSet = "%s:%s:zSet:ScoreTotalZSet" //历史星光值排行榜

	MonthlyRankingTotalZSet = "%s:%s:zSet:MonthlyRankingTotalZSet:%s" //月星光值排行榜  （%s,是当月地一天的日期  2024-07-01）//todo  最好不要用那个函数获取，不准确

	ScoreDayZSet        = "%s:%s:zSet:ScoreDayZSet:%s"     //最近几天星光值排行榜
	ExperienceTotalZSet = "%s:%s:zSet:ExperienceTotalZSet" //历史星光值排行榜

	HasChangeUserZSet    = "%s:%s:zSet:HasChangeUser"    //记录有变化的用户，方便更新
	HasChangeUserZSetNew = "%s:%s:zSet:HasChangeUserNew" //记录有变化的用户，方便更新

	RankSuitHash         = "%s:%s:hash:rankSuit"         //套装榜单
	RankWingsHash        = "%s:%s:hash:rankWings"        //翅膀榜单
	RankWingsHashOneGame = "%s:%s:hash:rankWingsOneGame" //翅膀榜单

	ComposeRankSuitHash         = "%s:%s:hash:composeRankSuit"         //组合套装榜单
	ComposeRankWingsHash        = "%s:%s:hash:composeRankWings"        //组合翅膀榜单
	ComposeRankWingsHashOneGame = "%s:%s:hash:composeRankWingsOneGame" //组合翅膀榜单

	//SpecialEffectsForUnsentGifts       = "%s:%s:hash:specialEffectsForUnsentGifts"       //还未发送的礼物特效后台，

	FullServiceZSet       = "%s:%s:zSet:fullServiceZSet"         //全部的队列，
	FullServiceSingleList = "%s:%s:list:fullServiceSingleList%s" //队列，专门用来管理对应的喇叭的，
	FullServiceCommonList = "%s:%s:list:fullServiceCommonList"   //队列，专门用来管理对应的喇叭的，

	CombinationTotalZSet     = "%s:%s:zSet:CombinationTotalZSet:%s"        //組合星光值排行榜
	CombinationHashDetailMap = "%s:%s:hash:CombinationHashDetailMap:%s:%s" //翅膀榜单
	CombinationAsyncList     = "%s:%s:list:CombinationAsyncList"           //同步组合数据的队列

	ComposeRankTitleHash = "%s:%s:hash:composeRankTitle" //组合套装榜单,发给用户的

	ScoreRaloTotalZSet            = "%s:%s:zSet:ScoreRaloZSet:%s"                    //活动排名，活动id
	ScoreStarPointsTotalZSet      = "%s:%s:zSet:ScoreStarPointsZSet:%s"              //活动总星光值
	ScoreRaloDistributeHash       = "%s:%s:hash:ScoreRaloDistributeHash:%s:%s"       //发送过散件的人，不需要再次发送
	ScoreStarPointsDistributeHash = "%s:%s:hash:ScoreStarPointsDistributeHash:%s:%s" //发送过散件的人，不需要再次发送
	ScoreRaloUserElementHash      = "%s:%s:hash:ScoreRaloUserElementHash"            //过期了，就删除
	ScoreRaloWingsUserElementHash = "%s:%s:hash:ScoreRaloWingsUserElementHash"       //过期了，就删除
	BarrageList                   = "%s:%s:list:barrage:%s"                          //过期了，就删除

	TestUserBarrage  = "%s:%s:hash:testUserBarrage" //这里面的用户弹幕才会发送到测试环境
	ListMiningReward = "%s:%s:list:MiningReward"    //发放奖励列表

	ScoreRaloHandsUserElementHash = "%s:%s:hash:ScoreRaloHandsUserElementHash" //过期了，就删除

	ScoreRaloCommonUserElementHash = "%s:%s:hash:ScoreRaloCommonUserElementHash:%s" //过期了，就删除

	UserPackageHash = "%s:%s:hash:UserPackageHash:%s" //用户背包

	UserHasRoomCode           = "%s:%s:hash:UserHasRoomCode:%s"           //用户所在直播间id映射
	UserPackageListHasChange  = "%s:%s:list:UserPackageHashHasChange:%s"  //改动的
	UserElementBag            = "%s:%s:hash:userElementBag:%s"            //用户抽奖次数
	UserDrawPrizeNumber       = "%s:%s:hash:userDrawPrizeNumber:%s"       //用户抽奖次数 :活动id
	DistributeDrawPrizeNumber = "%s:%s:hash:distributeDrawPrizeNumber:%s" //给用户发放抽奖次数

	UserDrawPrizeNumberInfo = "%s:%s:hash:UserDrawPrizeNumberInfo:%s:%s" //用户抽奖次数,filed:批次和累计次数 :活动id:用户id
	UserPieceGiveNumberPkg  = "%s:%s:hash:UserPieceGiveNumberPkg:%s:%s"  //用户碎片可赠送次数背包 :活动id:用户id

	LockStringPieceGive              = "%s:%s:string:PieceGive:%s"                  //碎片赠送加锁
	LockStringPieceExchange          = "%s:%s:string:PieceExchange:%s"              //碎片兑换服饰,道具，工具加锁
	LockStringPieceExchangeScore     = "%s:%s:string:PieceExchangeScore:%s"         //碎片兑换星光值加锁
	LockStringPieceExchangeTripScore = "%s:%s:string:PieceExchangeTripScore:%s"     //碎片兑换旅途点加锁
	LockStringPropNumGive            = "%s:%s:string:PropNumGive:%s"                //道具数量赠送加锁
	LockStringUseProp                = "%s:%s:string:UseProp:%s"                    //使用道具加锁
	LockStringUseLoudSpeakerProp     = "%s:%s:string:UseLoudSpeakerProp:%s"         //使用道具喇叭加锁
	LockStringPropResolve            = "%s:%s:string:LockStringPropResolve:%s"      //道具分解加锁
	LockStringBuyShopGood            = "%s:%s:string:BuyShopGood:%s"                //购买商品加锁
	LockStringOpenTreasureBox        = "%s:%s:string:LockStringOpenTreasureBox:%s"  //开宝箱加锁
	LockStringOpenPetEgg             = "%s:%s:string:LockStringOpenPetEgg:%s"       //开宠物蛋加锁
	LockStringEnterTeam              = "%s:%s:string:LockStringEnterTeam:%s"        //加入战队加锁
	LockStringExitTeam               = "%s:%s:string:LockStringExitTeam:%s"         //退出战队加锁
	LockStringEatAp                  = "%s:%s:string:LockStringEatAp:%s"            //给宠物喂饼干加锁
	LockStringEatHp                  = "%s:%s:string:LockStringEatHp:%s"            //给宠物喂生命丹加锁
	LockStringEatExp                 = "%s:%s:string:LockStringEatExp:%s"           //给宠物喂经验加锁
	LockStringSlgAttackCity          = "%s:%s:string:LockStringSlgAttackCity:%s"    //占领城池加锁
	LockStringSlgLeaveCity           = "%s:%s:string:LockStringSlgLeaveCity:%s"     //结束挖矿加锁
	LockStringPetEnterPosition       = "%s:%s:string:LockStringPetEnterPosition:%s" //宠物上阵,下阵加锁
	LockStringPetFollowRequest       = "%s:%s:string:LockStringPetFollowRequest:%s" //换跟随的宠物加锁

	RoomDrawPrizeActivityListChange = "%s:%s:list:RoomDrawPrizeActivityListChange:%s" //直播间活动变化

	RoomDynamic               = "%s:%s:zSet:RoomDynamic"                        //直播间活跃度
	UserDynamic               = "%s:%s:zSet:UserDynamic"                        //用户活跃度
	UserTripTotalScore        = "%s:%s:zSet:UserTripTotalScore:%s"              //用户总旅途点数 活动id
	UserTripDayTaskInfo       = "%s:%s:hash:UserTripDayTaskInfo:%s:%s:%s"       //用户旅途每日任务信息 活动id:日期（2024-07-07）:用户id;
	UserTripDayTaskInfoFinish = "%s:%s:hash:UserTripDayTaskInfoFinish:%s:%s:%s" //用户旅途每日任务是否完成 活动id:日期（2024-07-07）:用户id;
	UserTripLevelPrize        = "%s:%s:hash:UserTripLevelPrize:%s:%s"           //用户旅途等级奖品发放记录 活动id:用户id

	PetLevelScore = "%s:%s:hash:PetLevelScore:%s:%s" //宠物等级积分，第一个变量时 宠物类型，第二个变量时品质类型 （ss是1，S是2，A是3，B是4）

	UserTripLevelListHasChange = "%s:%s:list:UserTripLevelListHasChange:%s:%s" //用户旅途等级变化通知 活动id:房间id

	RoomModelInfo               = "%s:%s:hash:RoomModelInfo:%s"               //直播间比赛信息 :直播间id
	RoomTotalScoreZSet          = "%s:%s:zSet:RoomTotalScoreZSet"             //直播间总星光值 member:直播间id
	RoomUserTotalScoreZSet      = "%s:%s:zSet:RoomUserTotalScoreZSet:%s"      //直播间用户总星光值 :直播间id member:uid
	ModelTotalScoreZSet         = "%s:%s:zSet:ModelTotalScoreZSet"            //比赛总星光值  member:直播间id
	ModelUserTotalScoreZSet     = "%s:%s:zSet:ModelUserTotalScoreZSet:%s"     //比赛用户总星光值 :比赛id member:uid
	ModelRedBlueTotalScoreZSet  = "%s:%s:zSet:ModelRedBlueTotalScoreZSet:%s"  //比赛红方蓝方总星光值 :比赛id ,member只有两个 red,blue
	ModelRedUserTotalScoreZSet  = "%s:%s:zSet:ModelRedUserTotalScoreZSet:%s"  //比赛红方用户总星光值 :比赛id member:uid
	ModelBlueUserTotalScoreZSet = "%s:%s:zSet:ModelBlueUserTotalScoreZSet:%s" //比赛蓝方用户总星光值 :比赛id member:uid

	UserUseLoudSpeakerNumber = "%s:%s:hash:UserUseLoudSpeakerNumber:%s:%s" //用户每日使用喇叭个数 :日期（2024-07-07）:用户id;

	ShopMallGoodChangeListChange = "%s:%s:list:ShopMallGoodChangeListChange:%s" //商城商品变化通知 :直播间id

	UserPetPackageHash = "%s:%s:hash:UserPetPackageHash:%s" //用户宠物背包 :用户id;
	UserPetProduceZSet = "%s:%s:zSet:UserPetProduceZSet"    //用户宠物生成时间 member:宠物编号和用户id,score为生成宠物时间

	HostTeamHash                = "%s:%s:hash:HostTeamHash:%s:%s"                //主播战队 活动id:主播id
	HostTeamActivityZSet        = "%s:%s:zSet:HostTeamActivityZSet:%s"           //组建战队的主播 活动id member :主播id
	UserTeamHash                = "%s:%s:hash:UserTeamHash:%s"                   //用户加入战队 :活动id
	UserPetEnterPositionHash    = "%s:%s:hash:UserPetEnterPositionHash:%s:%s"    //用户宠物上阵详情 :活动id:用户id
	UserPetEnterPositionZSet    = "%s:%s:zSet:UserPetEnterPositionZSet:%s"       //宠物上阵的用户记录 :活动id member:用户id,score宠物上阵最新时间
	HostTeamActiveLeaveCityHash = "%s:%s:hash:HostTeamActiveLeaveCityHash:%s:%s" //主播主动离开城池时间 :活动id:城池编号; field:主播id value:离开城池时间
	PetApEmptyList              = "%s:%s:list:PetApEmptyList:%s"                 //用户宠物体力值为0 :活动id
	PetApEmptyZSet              = "%s:%s:zSet:PetApEmptyZSet:%s:%s"              //用户宠物体力值为0推送记录 :活动id:用户id member:宠物编号,score推送时间

	PetPowerRankingZSet = "%s:%s:zSet:PetPowerRankingZSet:%s" //宠物战力排行榜:活动id； member:用户id,score用户战力

	RoomHostDayGiftNum = "%s:%s:hash:RoomHostDayGiftNum:%s:%s" //主播每日累计礼物数量 :日期:主播id field:礼物编号 value:礼物数量

	UserSinglesDayCarnivalZSet = "%s:%s:zSet:UserSinglesDayCarnivalZSet" //双十一狂欢积分排行榜 member:用户id,score用户积分

	SingleScoreDayZSet     = "%s:%s:zSet:SingleScoreDayZSet:%s"     //最近几天星光值排行榜
	ModelUserScoreZSet     = "%s:%s:zSet:ModelUserScoreZSet:%s"     //比赛用户排行榜单
	FlowerCarUserScoreZSet = "%s:%s:zSet:FlowerCarUserScoreZSet:%s" //花车排行榜单
	WeekUserScoreZSet      = "%s:%s:zSet:WeekUserScoreZSet:%s"      //单人周排行榜单'

	SingleCpScoreDayZSet     = "%s:%s:zSet:SingleCpScoreDayZSet:%s"     //cp
	ModelCpUserScoreZSet     = "%s:%s:zSet:ModelCpUserScoreZSet:%s"     //cp比赛用户排行榜单
	FlowerCarCpUserScoreZSet = "%s:%s:zSet:FlowerCarCpUserScoreZSet:%s" //cp花车排行榜单
	WeekCpUserScoreZSet      = "%s:%s:zSet:WeekCpUserScoreZSet:%s"      //cp花车排行榜单

	UserLikeNumberZSet = "%s:%s:zSet:UserLikeNumberZSet:%s" //用户日点赞数

	SpecialEffectsList       = "%s:%s:list:SpecialEffectsList:%s:%s"    //主播编号，用户id
	ClientSpecialEffectsHash = "%s:%s:list:ClientSpecialEffectsHash:%s" //主播编号

	RedisKeyListRoom = "%s:%s:list:barrage:%s" //弹幕列表  环境id,项目id,直播间id

	UserRelationshipHash = "%s:%s:hash:UserRelationshipHash:%s" //用户关系hash :用户id field:userIdOther value:关系

	UserRelationSwitchingTimesString = "%s:%s:string:UserRelationSwitchingTimesString:%s" //user_id和user_id 用户关系切换次数

	RedisAccessTokenKey = "%s:%s:GetAccessToken"
)
