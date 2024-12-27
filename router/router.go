package router

import (
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"github.com/aceld/zinx/ziface"
)

type Router struct {
	ziface.IRouter
}

func GetRouters() map[api.MsgIds]ziface.IRouter {

	return map[api.MsgIds]ziface.IRouter{
		api.MsgIds_CREATE_ROOM_RESPONSE: &RoomRouter{},
		//api.MsgIds_PING_RESPONSE:         &PingRouter{Router{Logger: l}},
		api.MsgIds_CREATE_MODEL_RESPONSE: &ModelRouter{},
		//api.MsgIds_END_MODEL_RESPONSE:    &EndModelRouter{Router{Logger: l}},
		//api.MsgIds_CREATE_ROUND_RESPONSE: &RoundRouter{Router{Logger: l}},
		//api.MsgIds_ROUND_END_RESPONSE:    &EndRoundRouter{Router{Logger: l}},
		//api.MsgIds_BARRAGE_RESPONSE:      &BarrageRouter{Router{Logger: l}},
		//api.MsgIds_DANCE_LEVEL_RESPONSE:  &DanceLevelRouter{Router{Logger: l}},
		//
		//api.MsgIds_DANCE_USER_TOP_RANK_RESPONSE:    &UserTopRankRouter{Router{Logger: l}},
		//api.MsgIds_GET_BUFFOON_STATUS_RESPONSE:     &BuffoonStatusRouter{Router{Logger: l}},
		//api.MsgIds_DANCE_USER_TOP_CP_RANK_RESPONSE: &UserTopRankRouter{Router{Logger: l}},
		//api.MsgIds_USER_LUCKY_DRAW_RESPONSE:        &UserDrawListRouter{Router{Logger: l}},
		//api.MsgIds_USER_BOARD_RESPONSE:             &UserInfoBoredRouter{Router{Logger: l}},
		//
		//api.MsgIds_BAG_RESPONSE:               &BAGRouter{Router{Logger: l}},
		//api.MsgIds_LUCKY_DRAW_RECORD_RESPONSE: &LuckyRouter{Router{Logger: l}},
		//api.MsgIds_LUCK_DRAW_RULE_RESPONSE:    &LuckyRuleRouter{Router{Logger: l}},
		//api.MsgIds_TEST_BARRAGE_SEND_RESPONSE: &BarageSendRouter{Router{Logger: l}},
		//api.MsgIds_GET_USER_PACKAGE_RESPONSE:  &PackageRouter{Router{Logger: l}},
		//api.MsgIds_PACKAGE_CHANGE_RESPONSE:    &PackageChangeRouter{Router{Logger: l}},
		//api.MsgIds_ELEMENT_GIVE_RESPONSE:      &ElementGiveRouter{Router{Logger: l}},
		//
		//api.MsgIds_DAILY_STORE_RESPONSE:       &StoreRouter{Router{Logger: l}},
		//api.MsgIds_DAILY_STORE_OTHER_RESPONSE: &StoreOtherRouter{Router{Logger: l}},
		//api.MsgIds_DAILY_STORE_PUSH_RESPONSE:  &StorePush{Router{Logger: l}},
		//api.MsgIds_FRIEND_PUSH_RESPONSE:       &FriendPush{Router{Logger: l}},
		//api.MsgIds_SAVE_BADGE_RESPONSE:        &UserBadge{Router{Logger: l}},
		//api.MsgIds_PET_PANEL_RESPONSE:         &PetPanelRouter{Router{Logger: l}},
		//api.MsgIds_PET_CHARM_RANK_RESPONSE:    &PetCharmRankRouter{Router{Logger: l}},
	}
}
