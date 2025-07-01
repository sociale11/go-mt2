package in

import "go-mt2/internal/packets"

type ClientPacket interface {
	GetHeader() byte
	Parse(reader *packets.BufferReader) error
	GetType() byte
}

type ClientPacketType byte

const (
	HEADER_CG_LOGIN          ClientPacketType = 1
	HEADER_CG_ATTACK         ClientPacketType = 2
	HEADER_CG_CHAT           ClientPacketType = 3
	HEADER_CG_PLAYER_CREATE  ClientPacketType = 4 // ���ο� �÷��̾ ����
	HEADER_CG_PLAYER_DESTROY ClientPacketType = 5 // �÷��̾ ����.
	HEADER_CG_PLAYER_SELECT  ClientPacketType = 6
	HEADER_CG_CHARACTER_MOVE ClientPacketType = 7
	HEADER_CG_SYNC_POSITION  ClientPacketType = 8
	HEADER_CG_DIRECT_ENTER   ClientPacketType = 9
	HEADER_CG_ENTERGAME      ClientPacketType = 10
	HEADER_CG_ITEM_USE       ClientPacketType = 11
	HEADER_CG_ITEM_DROP      ClientPacketType = 12
	HEADER_CG_ITEM_MOVE      ClientPacketType = 13
	HEADER_CG_ITEM_PICKUP    ClientPacketType = 15
	HEADER_CG_QUICKSLOT_ADD  ClientPacketType = 16
	HEADER_CG_QUICKSLOT_DEL  ClientPacketType = 17
	HEADER_CG_QUICKSLOT_SWAP ClientPacketType = 18
	HEADER_CG_WHISPER        ClientPacketType = 19
	HEADER_CG_ITEM_DROP2     ClientPacketType = 20
	//HEADER_BLANK21								ClientPacketType = 21
	//HEADER_BLANK22								ClientPacketType = 22
	//HEADER_BLANK22								ClientPacketType = 23
	//HEADER_BLANK24								ClientPacketType = 24
	//HEADER_BLANK25								ClientPacketType = 25
	HEADER_CG_ON_CLICK           ClientPacketType = 26
	HEADER_CG_EXCHANGE           ClientPacketType = 27
	HEADER_CG_CHARACTER_POSITION ClientPacketType = 28
	HEADER_CG_SCRIPT_ANSWER      ClientPacketType = 29
	HEADER_CG_QUEST_INPUT_STRING ClientPacketType = 30
	HEADER_CG_QUEST_CONFIRM      ClientPacketType = 31
	//HEADER_BLANK32								ClientPacketType = 32
	//HEADER_BLANK33								ClientPacketType = 33
	//HEADER_BLANK34								ClientPacketType = 34
	//HEADER_BLANK35								ClientPacketType = 35
	//HEADER_BLANK36								ClientPacketType = 36
	//HEADER_BLANK37								ClientPacketType = 37
	//HEADER_BLANK38								ClientPacketType = 38
	//HEADER_BLANK39								ClientPacketType = 39
	//HEADER_BLANK40								ClientPacketType = 40
	HEADER_CG_PVP ClientPacketType = 41
	//HEADER_BLANK42								ClientPacketType = 42
	//HEADER_BLANK43								ClientPacketType = 43
	//HEADER_BLANK44								ClientPacketType = 44
	//HEADER_BLANK45								ClientPacketType = 45
	//HEADER_BLANK46								ClientPacketType = 46
	//HEADER_BLANK47								ClientPacketType = 47
	//HEADER_BLANK48								ClientPacketType = 48
	//HEADER_BLANK49								ClientPacketType = 49
	HEADER_CG_SHOP              ClientPacketType = 50
	HEADER_CG_FLY_TARGETING     ClientPacketType = 51
	HEADER_CG_USE_SKILL         ClientPacketType = 52
	HEADER_CG_ADD_FLY_TARGETING ClientPacketType = 53
	HEADER_CG_SHOOT             ClientPacketType = 54
	HEADER_CG_MYSHOP            ClientPacketType = 55
	//HEADER_BLANK56								ClientPacketType = 56
	//HEADER_BLANK57								ClientPacketType = 57
	//HEADER_BLANK58								ClientPacketType = 58
	//HEADER_BLANK59								ClientPacketType = 59
	HEADER_CG_ITEM_USE_TO_ITEM ClientPacketType = 60
	HEADER_CG_TARGET           ClientPacketType = 61
	//HEADER_BLANK62								ClientPacketType = 62
	//HEADER_BLANK63								ClientPacketType = 63
	//HEADER_BLANK64								ClientPacketType = 64
	HEADER_CG_WARP          ClientPacketType = 65
	HEADER_CG_SCRIPT_BUTTON ClientPacketType = 66
	HEADER_CG_MESSENGER     ClientPacketType = 67
	//HEADER_BLANK68								ClientPacketType = 68
	HEADER_CG_MALL_CHECKOUT       ClientPacketType = 69
	HEADER_CG_SAFEBOX_CHECKIN     ClientPacketType = 70 // �������� â���� �ִ´�.
	HEADER_CG_SAFEBOX_CHECKOUT    ClientPacketType = 71 // �������� â���� ���� ���´�.
	HEADER_CG_PARTY_INVITE        ClientPacketType = 72
	HEADER_CG_PARTY_INVITE_ANSWER ClientPacketType = 73
	HEADER_CG_PARTY_REMOVE        ClientPacketType = 74
	HEADER_CG_PARTY_SET_STATE     ClientPacketType = 75
	HEADER_CG_PARTY_USE_SKILL     ClientPacketType = 76
	HEADER_CG_SAFEBOX_ITEM_MOVE   ClientPacketType = 77
	HEADER_CG_PARTY_PARAMETER     ClientPacketType = 78
	//HEADER_BLANK68								ClientPacketType = 79
	HEADER_CG_GUILD             ClientPacketType = 80
	HEADER_CG_ANSWER_MAKE_GUILD ClientPacketType = 81
	HEADER_CG_FISHING           ClientPacketType = 82
	HEADER_CG_GIVE_ITEM         ClientPacketType = 83
	//HEADER_BLANK84								ClientPacketType = 84
	//HEADER_BLANK85								ClientPacketType = 85
	//HEADER_BLANK86								ClientPacketType = 86
	//HEADER_BLANK87								ClientPacketType = 87
	//HEADER_BLANK88								ClientPacketType = 88
	//HEADER_BLANK89								ClientPacketType = 89
	HEADER_CG_EMPIRE ClientPacketType = 90
	//HEADER_BLANK91								ClientPacketType = 91
	//HEADER_BLANK92								ClientPacketType = 92
	//HEADER_BLANK93								ClientPacketType = 93
	//HEADER_BLANK94								ClientPacketType = 94
	//HEADER_BLANK95								ClientPacketType = 95
	HEADER_CG_REFINE ClientPacketType = 96
	//HEADER_BLANK97								ClientPacketType = 97
	//HEADER_BLANK98								ClientPacketType = 98
	//HEADER_BLANK99								ClientPacketType = 99

	HEADER_CG_MARK_LOGIN   ClientPacketType = 100
	HEADER_CG_MARK_CRCLIST ClientPacketType = 101
	HEADER_CG_MARK_UPLOAD  ClientPacketType = 102
	HEADER_CG_MARK_IDXLIST ClientPacketType = 104

	HEADER_CG_CRC_REPORT ClientPacketType = 103

	HEADER_CG_HACK                ClientPacketType = 105
	HEADER_CG_CHANGE_NAME         ClientPacketType = 106
	HEADER_CG_LOGIN2              ClientPacketType = 109
	HEADER_CG_DUNGEON             ClientPacketType = 110
	HEADER_CG_LOGIN3              ClientPacketType = 111
	HEADER_CG_GUILD_SYMBOL_UPLOAD ClientPacketType = 112
	HEADER_CG_GUILD_SYMBOL_CRC    ClientPacketType = 113
	HEADER_CG_SCRIPT_SELECT_ITEM  ClientPacketType = 114
	HEADER_CG_LOGIN4              ClientPacketType = 115
	HEADER_CG_LOGIN5_OPENID       ClientPacketType = 116 //OpenID : ����� ���� ����Ű�� ������ ����.

	HEADER_CG_DRAGON_SOUL_REFINE ClientPacketType = 205
	HEADER_CG_STATE_CHECKER      ClientPacketType = 206

	HEADER_CG_TIME_SYNC       ClientPacketType = 0xfc
	HEADER_CG_CLIENT_VERSION  ClientPacketType = 0xfd
	HEADER_CG_CLIENT_VERSION2 ClientPacketType = 0xf1
	HEADER_CG_PONG            ClientPacketType = 0xfe
	HEADER_CG_HANDSHAKE       ClientPacketType = 0xff
)
