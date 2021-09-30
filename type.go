package nim

type Message struct {
}

type Empty struct {
}

//ImSendMessageOption .
type ImSendMessageOption struct {
	msgDesc            string          //消息描述文本，针对非Text、Tip类型的消息有效，最大长度500字符。该描述信息可用于云端历史消息关键词检索。
	Antispam           bool            //对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。true或false, 默认false。只对消息类型为：100 自定义消息类型 的消息生效。
	AntispamCustom     *AntiSpamCustom //在antispam参数为true时生效。
	Option             *MessageOption  //发消息时特殊指定的行为选项
	Pushcontent        string          //ios推送内容，不超过150字符，option选项中允许推送（push=true），此字段可以指定推送内容
	Payload            string          //ios 推送对应的payload,必须是JSON,不能超过2k字符
	Ext                string          //开发者扩展字段，长度限制1024字符
	ForcePushList      []string        //发送群消息时的强推（@操作）用户列表，格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
	ForcePushContent   string          //发送群消息时，针对强推（@操作）列表forcepushlist中的用户，强制推送的内容
	ForcePushAll       bool            //发送群消息时，强推（@操作）列表是否为群里除发送者外的所有有效成员，true或false，默认为false
	Bid                string          //可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
	UseYidun           int             //可选，单条消息是否使用易盾反垃圾，可选值为0。 0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。 若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
	YidunAntiCheating  string          //可选，易盾反垃圾增强反作弊专属字段，限制json，长度限制1024字符（详见易盾反垃圾接口文档反垃圾防刷版专属字段）
	MarkRead           int             //可选，群消息是否需要已读业务（仅对群消息有效），0:不需要，1:需要
	CheckFriend        bool            //是否为好友关系才发送消息，默认否 注：使用该参数需要先开通功能服务
	SubType            int             //自定义消息子类型，大于0
	MsgSenderNoSense   int             //发送方是否无感知。0-有感知，1-无感知。若无感知，则消息发送者无该消息的多端、漫游、历史记录等。
	MsgReceiverNoSense int             //接受方是否无感知。0-有感知，1-无感知。若无感知，则消息接收者者无该消息的多端、漫游、历史记录等
	Env                string          //所属环境，根据env可以配置不同的抄送地址
}

//MessageOption 发消息时特殊指定的行为选项
type MessageOption struct {
	Roam          *bool `json:"roam,omitempty"`          //该消息是否需要漫游，默认true（需要app开通漫游消息功能）
	History       *bool `json:"history,omitempty"`       //该消息是否存云端历史，默认true
	Sendersync    *bool `json:"sendersync,omitempty"`    //该消息是否需要发送方多端同步，默认true
	Push          *bool `json:"push,omitempty"`          //该消息是否需要APNS推送或安卓系统通知栏推送，默认true
	Route         *bool `json:"route,omitempty"`         //该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
	Badge         *bool `json:"badge,omitempty"`         //该消息是否需要计入到未读计数中，默认true
	NeedPushNick  *bool `json:"needPushNick,omitempty"`  //推送文案是否需要带上昵称，不设置该参数时默认true
	Persistent    *bool `json:"persistent,omitempty"`    //是否需要存离线消息，不设置该参数时默认true
	SessionUpdate *bool `json:"sessionUpdate,omitempty"` //是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
}

//AntiSpamCustom 自定义的反垃圾检测内容, JSON格式，不能超过5000字符
type AntiSpamCustom struct {
	Type int    `json:"type"` //1：文本，2：图片。
	Data string `json:"data"` // 文本内容or图片地址
}
