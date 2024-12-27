package dreamtwodancetiktokstruct

type LiveCommentItem struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	Content   string `json:"content"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}

type LiveGiftItem struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	SecGiftId string `json:"sec_gift_id"`
	GiftNum   int    `json:"gift_num"`
	GiftValue int    `json:"gift_value"`
	Test      bool   `json:"test"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}

type LiveLikeItem struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	LikeNum   int64  `json:"like_num"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}

type LiveFansclubItem struct {
	MsgId              string `json:"msg_id"`
	SecOpenid          string `json:"sec_openid"`
	AvatarUrl          string `json:"avatar_url"`
	Nickname           string `json:"nickname"`
	Timestamp          int64  `json:"timestamp"`
	FansclubReasonType int    `json:"fansclub_reason_type"` // 粉丝团消息类型：1-升级，2-加团
	FansclubLevel      int    `json:"fansclub_level"`       // 用户粉丝团等级，加团消息下默认传1
}

type ListStruct struct {
	Timestamp string `json:"timestamp"`
	RoomId    string `json:"roomId"`
	MsgType   string `json:"msgType"`
	BodyStr   string `json:"bodyStr"`
}

type ListStructCommon struct {
	Timestamp int64  `json:"timestamp"`
	RoomId    string `json:"roomId"`
	MsgType   string `json:"msgType"`
	BodyStr   string `json:"bodyStr"`
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	Content   string `json:"content"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	SecGiftId string `json:"sec_gift_id"`
	GiftNum   int    `json:"gift_num"`
	GiftValue int    `json:"gift_value"`
	Test      bool   `json:"test"`
	LikeNum   int64  `json:"like_num"`
}

type EndModelCombinationItem struct {
	HostNickname     string `json:"host_nickname"`       //主播昵称
	HostNo           string `json:"host_no"`             //主播编号
	RoomId           string `json:"room_id"`             //直播间id
	RoomCreateAt     int64  `json:"room_create_at"`      //直播间id
	ModelId          string `json:"model_id"`            //局id
	ModelCreateAt    int64  `json:"model_create_at"`     //直播间id
	ZSetKey          string `json:"z_set_key"`           //主动at人的id
	HashDetailMapKey string `json:"hash_detail_map_key"` //组合详情
	BeginTime        string `json:"begin_time"`          //区间开始时间
	EndTime          string `json:"end_time"`            //区间结束时间
	UidActive        string `json:"uid_active"`          //主动at人的id
	UidPassive       string `json:"uid_passive"`         //主动at人的id
	ActiveScore      int64  `json:"active_score"`        //积分
	PassiveScore     int64  `json:"passive_score"`       //积分
	TotalScore       int64  `json:"total_score"`         //总积分
}

// FullServerSpeaker 全服喇叭
type FullServerSpeaker struct {
	FullType        int32  `json:"full_type"`         //0:单人10个神秘空投，全服特效;1:结婚模式组成cp神秘空投;2:离婚全服喇叭
	UserIdLeft      string `json:"user_id_left"`      //送礼人
	UserNameLeft    string `json:"user_name_left"`    //送礼人
	UserAvatarLeft  string `json:"user_avatar_left"`  //用户头像
	UserIdRight     string `json:"user_id_right"`     //蹭礼物的
	UserNameRight   string `json:"user_name_right"`   //蹭礼物的
	UserAvatarRight string `json:"user_avatar_right"` //蹭礼物的
	PropsNo         int64  `json:"props_no"`          //
	PropsExpireTime int64  `json:"props_expire_time"` //
	PropsNumber     int32  `json:"props_number"`      //
	DrawType        int32  `json:"draw_type"`         // 1.道具 2.星光值
	Score           int64  `json:"score"`             //星光值分数
}

type PackageDetailRedis struct {
	PropType      int   `json:"prop_type"`      //一级类型  1:服饰;2:道具;3:工具
	ElementType   int   `json:"element_type"`   //二级类型
	ElementNo     int64 `json:"element_no"`     //服饰或者道具编号
	ElementNumber int64 `json:"element_number"` //服饰默认为1，永远为1，道具叠加数量
	ExpiredTime   int64 `json:"expired_time"`   //服饰过期时间，不存道具过期时间
	ActivityId    int64 `json:"activity_id"`    //活动id，主要给碎片道具区分用的
}

type DrawPrizePackageDetail struct {
	PropType      int    `json:"prop_type"`      //一级类型  1:服饰;2:道具;3:工具
	ElementType   int    `json:"element_type"`   //二级类型
	ElementNo     int64  `json:"element_no"`     //服饰或者道具编号
	ElementNumber int64  `json:"element_number"` //服饰默认为1，永远为1，道具叠加数量
	ExpiredTime   int64  `json:"expired_time"`   //服饰过期时间，不存道具过期时间
	ActivityId    int64  `json:"activity_id"`    //活动id，主要给碎片道具区分用的
	DrawPrizeId   string `json:"drawPrizeId"`
}

type UserPackageListHasChangeItemRedis struct {
	PropType    int    `json:"prop_type"`    //道具类型类型  1:散件道具;2:碎片
	UserId      string `json:"user_id"`      //散件编号
	PropNo      int64  `json:"prop_no"`      //散件编号
	PropNumber  int64  `json:"prop_number"`  //散件数量
	ExpiredTime int64  `json:"expired_time"` //散件过期时间
}

// 用户碎片可赠送次数详情
type UserPieceGiveNumberDetail struct {
	PieceNo         int64 `json:"piece_no"`          //碎片编号
	PieceGiveNumber int64 `json:"piece_give_number"` //碎片可赠送次数
}

// 直播间活动变化通知
type RoomDrawPrizeActivityChange struct {
	ActivityId int64 `json:"activity_id"` //活动id
}

// 旅途用户等级变化通知
type TripUserLevelChangeLogList struct {
	List []TripUserLevelChangeLog
}

type TripUserLevelChangeLog struct {
	Uid      string //等级
	Level    int64  //等级
	Exp      int64  //经验
	PushTime int64  //推送时间
}

// 用户等级奖品发放时间
type TripUserLevelPrizeTime struct {
	Level     int64 //等级
	PrizeTime int64 // 发放时间
}

// PetPackageDetail 用户宠物背包详情
type PetPackageDetail struct {
	PetNo               int64  `json:"pet_no"`                 //宠物编号
	PetStar             int64  `json:"pet_star"`               //宠物星数
	PetCurrentLevel     int    `json:"pet_current_level"`      //宠物当前经验等级
	PetCurrentExp       int64  `json:"pet_current_exp"`        //宠物当前等级  1100 +2000
	PetApNumber         int64  `json:"pet_ap_number"`          //宠物当前体力值
	PetCumulateApNumber int64  `json:"pet_cumulate_ap_number"` //宠物累计体力值
	PetExpendApNumber   int64  `json:"pet_expend_ap_number"`   //宠物消耗体力值
	PetHpNumber         int64  `json:"pet_hp_number"`          //宠物生命值
	PetCriticalHit      int64  `json:"pet_critical_hit"`       //宠物暴击率 (S级上才有)
	CreateTime          int64  `json:"create_time"`            //宠物创建时间
	MiningStartTime     int64  `json:"mining_start_time"`      //挖矿开始时间
	UsePhysicalTime     int64  `json:"use_physical_time"`      //消耗体力值时间
	ClothesElementNo    string `json:"clothes_element_no"`     //宠物装扮散件信息(多个,逗号拼接)
}

// PetLevelScoreDetail 等级积分表，key是等级，1-1000
type PetLevelScoreDetail struct {
	MinExp        int64 `json:"min_exp"`        //经验值开区间  1000
	MaxExp        int64 `json:"max_exp"`        //经验值闭区间  1200   1200-1500 1500-2000
	PhysicalScore int64 `json:"physical_score"` //体力值
	LifeScore     int64 `json:"life_score"`     //生命值
	AttackScore   int64 `json:"attack_score"`   //攻击力
}

// 主播战队人员列表
type TeamHostDetailList []TeamHostDetail

// TeamHostDetail 主播战队详情
type TeamHostDetail struct {
	UserId    string `json:"user_id"`    //用户id
	EnterTime int64  `json:"enter_time"` //加入时间
}

// TeamUserDetail 用户战队详情
type TeamUserDetail struct {
	ActivitiesId int64  `json:"activities_id"` //活动id
	HostId       string `json:"host_id"`       //主播id
	TeamId       int64  `json:"team_id"`       //战队id
	EnterTime    int64  `json:"enter_time"`    //加入时间
}

// PetEnterPositionDetail 宠物上阵详情
type PetEnterPositionDetail struct {
	PetNo           int64  `json:"pet_no"`            //宠物编号
	BattleStartTime string `json:"battle_start_time"` //战斗开始时间
	BattleEndTime   string `json:"battle_end_time"`   //战斗结束时间
}

// PetSSDetail SS级宠物详情
type PetSSDetail struct {
	PetNo      int64 `json:"pet_no"`      //宠物编号
	CreateTime int64 `json:"create_time"` //宠物生成时间
}

type ElementNoObject struct {
	UserId   string                `json:"user_id"`
	SendTime int64                 `json:"send_time"` //发放时间
	Children []ElementNoItemObject `json:"children"`  //概率
}
type ElementNoItemObject struct {
	TmpStoreUserId string `json:"tmp_store_user_id"` //临时存放用户id
	ElementNo      int64  `json:"element_no"`        //散件编号
	SendTime       int64  `json:"send_time"`         //发放时间
	TimeNumber     int    `json:"time_number"`       //发放多久 小时为单位
	SendNumber     int    `json:"send_number"`       //发放数量
	Weight         int64  `json:"weight"`            //权重
}

// PetApEmptyDetail 用户宠物体力值为0时,推送队列
type PetApEmptyDetail struct {
	ActivitiesId int64  `json:"activities_id"` //活动id
	CityNo       string `json:"city_no"`       //城市编号
	HostId       string `json:"host_id"`       //主播id
	TeamId       int64  `json:"team_id"`       //战队id
	UserId       string `json:"user_id"`       //用户id
	PetNo        int64  `json:"pet_no"`        //宠物编号
}

// HostBattleLogObject 战队列表参数
type HostBattleLogObject struct {
	CityChangeLogId          string                      `json:"city_change_log_id"`            //城池流转记录id
	ActivitiesId             int64                       `json:"activities_id"`                 //活动id
	CityNo                   int64                       `json:"city_no"`                       //城市编号
	HostId                   string                      `json:"host_id"`                       //主播id
	TeamId                   int64                       `json:"team_id"`                       //战队id
	OtherHostId              string                      `json:"other_host_id"`                 //对方主播id
	OtherTeamId              int64                       `json:"other_team_id"`                 //对方战队id
	TeamTotalFighting        int64                       `json:"team_total_fighting"`           //团队总战力
	TeamTotalHp              int64                       `json:"team_total_hp"`                 //团队当前总血量
	TeamLevelMaxTotalHp      int64                       `json:"team_level_max_total_hp"`       //团队总等级上限血量
	UserPetHpLvMap           string                      `json:"user_pet_hp_lv_map"`            //团队每个宠物血量占比(json)
	OtherTeamTotalFighting   int64                       `json:"other_team_total_fighting"`     //对方团队总战力
	OtherTeamTotalHp         int64                       `json:"other_team_total_hp"`           //对方团队当前总血量
	OtherTeamLevelMaxTotalHp int64                       `json:"other_team_level_max_total_hp"` //对方团队总等级上限血量
	OtherUserPetHpLvMap      string                      `json:"other_user_pet_hp_lv_map"`      //对方团队每个宠物血量占比(json)
	BattleType               int64                       `json:"battle_type"`                   //战报类型 1:进攻空城(一定没奖励) 2:守城 3:离开城池 4:日清 5:月清 6:解散战队
	BattleResult             int64                       `json:"battle_result"`                 //战报结果 1:成功;2:失败
	UserRewardList           []ElementNoItemObject       `json:"user_reward_list"`              //用户维度奖励列表
	PetRewardList            []HostBattleDetailLogObject `json:"pet_reward_list"`               //宠物维度奖励列表
}

// HostBattleDetailLogObject 战队详情参数
type HostBattleDetailLogObject struct {
	CityChangeLogId            string                `json:"city_change_log_id"`             //城池流转记录id
	TeamHostBattleLogId        string                `json:"team_host_battle_log_id"`        //主播战队战报记录id
	UserId                     string                `json:"user_id"`                        //用户id
	PetNo                      int64                 `json:"pet_no"`                         //宠物编号
	PetStar                    int64                 `json:"pet_star"`                       //宠物星数
	PetCriticalHit             int64                 `json:"pet_critical_hit"`               //宠物暴击率 (S级上才有)
	PetApNumber                int64                 `json:"pet_ap_number"`                  //宠物当前体力值
	PetCumulateApNumber        int64                 `json:"pet_cumulate_ap_number"`         //宠物累计体力值
	PetExpendApNumber          int64                 `json:"pet_expend_ap_number"`           //宠物消耗体力值
	PetHpNumber                int64                 `json:"pet_hp_number"`                  //宠物生命值
	MiningStartTime            int64                 `json:"mining_start_time"`              //挖矿开始时间
	UsePhysicalTime            int64                 `json:"use_physical_time"`              //消耗体力值时间
	CityStamina                int64                 `json:"city_stamina"`                   //城池每分钟消耗的体力值
	WorkMinutes                int64                 `json:"work_minutes"`                   //挖矿时长(分钟)
	BattleReduceHp             int64                 `json:"battle_reduce_hp"`               //战斗扣减血量
	UpdatedPetApNumber         int64                 `json:"updated_pet_ap_number"`          //更新后宠物当前体力值
	UpdatedPetCumulateApNumber int64                 `json:"updated_pet_cumulate_ap_number"` //更新后宠物累计体力值
	UpdatedPetExpendApNumber   int64                 `json:"updated_pet_expend_ap_number"`   //更新后宠物消耗体力值
	UpdatedPetHpNumber         int64                 `json:"updated_pet_hp_number"`          //更新后宠物生命值
	UpdatedMiningStartTime     int64                 `json:"updated_mining_start_time"`      //更新后挖矿开始时间
	UpdatedUsePhysicalTime     int64                 `json:"updated_use_physical_time"`      //更新后消耗体力值时间
	RewardList                 []ElementNoItemObject `json:"reward_list"`                    //每个宠物的奖励列表
	HasReward                  bool                  `json:"has_reward"`                     //是否有奖励 ture有奖励;false无奖励
}

type SpecialEffectsListObject struct {
	GiftId    int `json:"gift_id"`    //礼物客户端id
	GiftCount int `json:"gift_count"` //礼物数量
	Number    int `json:"number"`     //几个特效
	Source    int `json:"source"`     //1.客户端推的；2:管理后台补的
}
