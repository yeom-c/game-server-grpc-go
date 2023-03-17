package redis_session

type Session struct {
	SessionId       string `redis:"session_id"`
	WorldId         int32  `redis:"world_id"`
	AccountId       int32  `redis:"account_id"`
	AccountUserId   int32  `redis:"account_user_id"`
	BattleServerUrl string `redis:"battle_server_url"`
	BattleChannelId string `redis:"battle_channel_id"`
	GameDb          int32  `redis:"game_db"`
	Nickname        string `redis:"nickname"`
	SignedInAt      string `redis:"signed_in_at"`
}
