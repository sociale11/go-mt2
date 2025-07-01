package out

import "go-mt2/internal/packets"

type ServerPacket interface {
	GetHeader() byte
	Parse(reader *packets.BufferReader) error
	GetType() ServerPacket
}

type ServerPacketType byte

const (

	/////////////////////////////////////////////////
	// From Server

	HEADER_GC_CHARACTER_ADD                 ServerPacketType = 1
	HEADER_GC_CHARACTER_DEL                 ServerPacketType = 2
	HEADER_GC_CHARACTER_MOVE                ServerPacketType = 3
	HEADER_GC_CHAT                          ServerPacketType = 4
	HEADER_GC_SYNC_POSITION                 ServerPacketType = 5
	HEADER_GC_LOGIN_SUCCESS3                ServerPacketType = 6
	HEADER_GC_LOGIN_FAILURE                 ServerPacketType = 7
	HEADER_GC_PLAYER_CREATE_SUCCESS         ServerPacketType = 8
	HEADER_GC_PLAYER_CREATE_FAILURE         ServerPacketType = 9
	HEADER_GC_PLAYER_DELETE_SUCCESS         ServerPacketType = 10
	HEADER_GC_PLAYER_DELETE_WRONG_SOCIAL_ID ServerPacketType = 11
	// 12
	HEADER_GC_STUN ServerPacketType = 13
	HEADER_GC_DEAD ServerPacketType = 14

	HEADER_GC_MAIN_CHARACTER      ServerPacketType = 15
	HEADER_GC_PLAYER_POINTS       ServerPacketType = 16
	HEADER_GC_PLAYER_POINT_CHANGE ServerPacketType = 17
	HEADER_GC_CHANGE_SPEED        ServerPacketType = 18
	HEADER_GC_CHARACTER_UPDATE    ServerPacketType = 19
	// #if defined(GAIDEN)
	HEADER_GC_ITEM_DEL   ServerPacketType = 20 // ������ â�� �߰�
	HEADER_GC_ITEM_SET_1 ServerPacketType = 21 // ������ â�� �߰�
	// #else
	HEADER_GC_ITEM_SET_GAIDEN  ServerPacketType = 20 // ������ â�� �߰�
	HEADER_GC_ITEM_SET2_GAIDEN ServerPacketType = 21 // ������ â�� �߰�
	// #endif
	HEADER_GC_ITEM_USE         ServerPacketType = 22 // ������ ��� (���� ����鿡�� �����ֱ� ����)
	HEADER_GC_ITEM_DROP        ServerPacketType = 23 // ������ ������
	HEADER_GC_ITEM_UPDATE      ServerPacketType = 25 // ������ ��ġ ������Ʈ
	HEADER_GC_ITEM_GROUND_ADD  ServerPacketType = 26 // �ٴڿ� ������ �߰�
	HEADER_GC_ITEM_GROUND_DEL  ServerPacketType = 27 // �ٴڿ��� ������ ����
	HEADER_GC_QUICKSLOT_ADD    ServerPacketType = 28
	HEADER_GC_QUICKSLOT_DEL    ServerPacketType = 29
	HEADER_GC_QUICKSLOT_SWAP   ServerPacketType = 30
	HEADER_GC_ITEM_OWNERSHIP   ServerPacketType = 31
	HEADER_GC_LOGIN_SUCCESS4   ServerPacketType = 32
	HEADER_GC_ITEM_UNBIND_TIME ServerPacketType = 33
	HEADER_GC_WHISPER          ServerPacketType = 34
	HEADER_GC_ALERT            ServerPacketType = 35

	HEADER_GC_MOTION ServerPacketType = 36

	HEADER_GC_SHOP      ServerPacketType = 38
	HEADER_GC_SHOP_SIGN ServerPacketType = 39

	// 39 ~ 41 Balnk
	HEADER_GC_DUEL_START         ServerPacketType = 40
	HEADER_GC_PVP                ServerPacketType = 41
	HEADER_GC_EXCHANGE           ServerPacketType = 42
	HEADER_GC_CHARACTER_POSITION ServerPacketType = 43

	HEADER_GC_PING ServerPacketType = 44

	HEADER_GC_SCRIPT        ServerPacketType = 45
	HEADER_GC_QUEST_CONFIRM ServerPacketType = 46

	HEADER_GC_MOUNT             ServerPacketType = 61
	HEADER_GC_OWNERSHIP         ServerPacketType = 62
	HEADER_GC_TARGET            ServerPacketType = 63
	HEADER_GC_WARP              ServerPacketType = 65
	HEADER_GC_ADD_FLY_TARGETING ServerPacketType = 69

	HEADER_GC_CREATE_FLY         ServerPacketType = 70
	HEADER_GC_FLY_TARGETING      ServerPacketType = 71
	HEADER_GC_SKILL_LEVEL        ServerPacketType = 72
	HEADER_GC_SKILL_COOLTIME_END ServerPacketType = 73
	HEADER_GC_MESSENGER          ServerPacketType = 74
	HEADER_GC_GUILD              ServerPacketType = 75
	HEADER_GC_SKILL_LEVEL_NEW    ServerPacketType = 76

	HEADER_GC_PARTY_INVITE ServerPacketType = 77
	HEADER_GC_PARTY_ADD    ServerPacketType = 78
	HEADER_GC_PARTY_UPDATE ServerPacketType = 79
	HEADER_GC_PARTY_REMOVE ServerPacketType = 80

	HEADER_GC_QUEST_INFO         ServerPacketType = 81
	HEADER_GC_REQUEST_MAKE_GUILD ServerPacketType = 82
	HEADER_GC_PARTY_PARAMETER    ServerPacketType = 83

	HEADER_GC_SAFEBOX_MONEY_CHANGE   ServerPacketType = 84
	HEADER_GC_SAFEBOX_SET            ServerPacketType = 85
	HEADER_GC_SAFEBOX_DEL            ServerPacketType = 86
	HEADER_GC_SAFEBOX_WRONG_PASSWORD ServerPacketType = 87
	HEADER_GC_SAFEBOX_SIZE           ServerPacketType = 88

	HEADER_GC_FISHING ServerPacketType = 89

	HEADER_GC_EMPIRE ServerPacketType = 90

	HEADER_GC_PARTY_LINK   ServerPacketType = 91
	HEADER_GC_PARTY_UNLINK ServerPacketType = 92

	HEADER_GC_REFINE_INFORMATION ServerPacketType = 95

	HEADER_GC_OBSERVER_ADD    ServerPacketType = 96
	HEADER_GC_OBSERVER_REMOVE ServerPacketType = 97
	HEADER_GC_OBSERVER_MOVE   ServerPacketType = 98
	HEADER_GC_VIEW_EQUIP      ServerPacketType = 99

	HEADER_GC_MARK_BLOCK     ServerPacketType = 100
	HEADER_GC_MARK_DIFF_DATA ServerPacketType = 101
	HEADER_GC_MARK_IDXLIST   ServerPacketType = 102

	//HEADER_GC_SLOW_TIMER						ServerPacketType = 105
	HEADER_GC_TIME        ServerPacketType = 106
	HEADER_GC_CHANGE_NAME ServerPacketType = 107

	HEADER_GC_DUNGEON            ServerPacketType = 110
	HEADER_GC_WALK_MODE          ServerPacketType = 111
	HEADER_GC_CHANGE_SKILL_GROUP ServerPacketType = 112

	// #if defined(GAIDEN)
	HEADER_GC_MAIN_CHARACTER_GAIDEN ServerPacketType = 113

	// #else
	// SUPPORT_BGM
	HEADER_GC_MAIN_CHARACTER2_EMPIRE ServerPacketType = 113
	// END_OF_SUPPORT_BGM
	// #endif

	HEADER_GC_SEPCIAL_EFFECT ServerPacketType = 114
	HEADER_GC_NPC_POSITION   ServerPacketType = 115

	HEADER_GC_CHARACTER_UPDATE2      ServerPacketType = 117
	HEADER_GC_LOGIN_KEY              ServerPacketType = 118
	HEADER_GC_REFINE_INFORMATION_NEW ServerPacketType = 119
	HEADER_GC_CHARACTER_ADD2         ServerPacketType = 120
	HEADER_GC_CHANNEL                ServerPacketType = 121

	HEADER_GC_MALL_OPEN         ServerPacketType = 122
	HEADER_GC_TARGET_UPDATE     ServerPacketType = 123
	HEADER_GC_TARGET_DELETE     ServerPacketType = 124
	HEADER_GC_TARGET_CREATE_NEW ServerPacketType = 125

	HEADER_GC_AFFECT_ADD    ServerPacketType = 126
	HEADER_GC_AFFECT_REMOVE ServerPacketType = 127

	HEADER_GC_MALL_SET          ServerPacketType = 128
	HEADER_GC_MALL_DEL          ServerPacketType = 129
	HEADER_GC_LAND_LIST         ServerPacketType = 130
	HEADER_GC_LOVER_INFO        ServerPacketType = 131
	HEADER_GC_LOVE_POINT_UPDATE ServerPacketType = 132
	HEADER_GC_GUILD_SYMBOL_DATA ServerPacketType = 133
	HEADER_GC_DIG_MOTION        ServerPacketType = 134

	HEADER_GC_DAMAGE_INFO          ServerPacketType = 135
	HEADER_GC_CHAR_ADDITIONAL_INFO ServerPacketType = 136

	// SUPPORT_BGM
	HEADER_GC_MAIN_CHARACTER3_BGM     ServerPacketType = 137
	HEADER_GC_MAIN_CHARACTER4_BGM_VOL ServerPacketType = 138
	// END_OF_SUPPORT_BGM

	HEADER_GC_AUTH_SUCCESS ServerPacketType = 150

	HEADER_GC_AUTH_SUCCESS_OPENID ServerPacketType = 154

	HEADER_GC_SPECIFIC_EFFECT       ServerPacketType = 208
	HEADER_GC_DRAGON_SOUL_REFINE    ServerPacketType = 209
	HEADER_GC_RESPOND_CHANNELSTATUS ServerPacketType = 210

	HEADER_GC_HANDSHAKE_OK ServerPacketType = 0xfc // 252
	HEADER_GC_PHASE        ServerPacketType = 0xfd // 253
	HEADER_GC_BINDUDP      ServerPacketType = 0xfe // 254
	HEADER_GC_HANDSHAKE    ServerPacketType = 0xff // 255
)
