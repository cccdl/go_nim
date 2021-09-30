package nim

import (
	"encoding/json"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

const (
	MsgTypeText      = 0   //文本消息
	MsgTypeImage     = 1   //图片消息
	MsgTypeVoice     = 2   //语音消息
	MsgTypeVideo     = 3   //视频消息
	MsgTypeLocation  = 4   //地理位置
	MsgTypeFile      = 6   //文件
	MsgTypeTips      = 10  //提示消息
	MsgTypeCustomize = 100 //自定义消息类型
)

const (
	sendMsgUrl            = "/msg/sendMsg.action"            //发送普通消息
	sendBatchMsgUrl       = "/msg/sendBatchMsg.action"       //批量发送点对点普通消息
	sendAttachMsgUrl      = "/msg/sendAttachMsg.action"      //发送自定义系统通知
	sendBatchAttachMsgUrl = "/msg/sendBatchAttachMsg.action" //批量发送点对点自定义系统通知
	uploadUrl             = "/msg/upload.action"             //文件上传
	fileUploadUrl         = "/msg/fileUpload.action"         //文件上传（multipart方式）
	recallUrl             = "/msg/recall.action"             //消息撤回
	delMsgOneWayUrl       = "/msg/delMsgOneWay.action"       //单向撤回消息
	broadcastMsgUrl       = "/msg/broadcastMsg.action"       //发送广播消息
	delRoamSessionUrl     = "/msg/delRoamSession.action"     //删除会话漫游
	jobNosDelUrl          = "/job/nos/del.action"            //上传NOS文件清理任务
)

// SendCustomMessage 发送自定义消息
func (c *ImClient) SendCustomMessage(fromID, toID string, msg interface{}, opt *ImSendMessageOption) error {
	bd, err := jsonTool.MarshalToString(msg)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return c.SendMessage(fromID, toID, bd, 0, MsgTypeCustomize, opt)
}

//SendMessage 发送普通消息
/**
 * @param fromID 发送者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param toID ope==0是表示accid即用户id，ope==1表示tid即群id
 * @param ope 0：点对点个人消息，1：群消息（高级群），其他返回414
 * @param msgType 0 表示文本消息,1 表示图片，2 表示语音，3 表示视频，4 表示地理位置信息，6 表示文件，100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
 * @param body 最大长度5000字符，为一个JSON串
 */
func (c *ImClient) SendMessage(fromID, toID, body string, ope, msgType int, opt *ImSendMessageOption) error {
	param := map[string]string{}
	param["from"] = fromID
	param["ope"] = strconv.Itoa(ope)
	param["to"] = toID
	param["type"] = strconv.Itoa(msgType)
	param["body"] = body

	if opt != nil {
		param["antispam"] = strconv.FormatBool(opt.Antispam)

		if opt.AntispamCustom != nil {
			param["antispamCustom"], _ = jsonTool.MarshalToString(opt.AntispamCustom)
		}

		if opt.Option != nil {
			param["option"], _ = jsonTool.MarshalToString(opt.Option)
		}

		if len(opt.Pushcontent) > 0 {
			param["pushcontent"] = opt.Pushcontent
		}

		if len(opt.Payload) > 0 {
			param["payload"] = opt.Payload
		}

		if len(opt.Ext) > 0 {
			param["ext"] = opt.Ext
		}

		if opt.ForcePushList != nil {
			param["forcepushlist"], _ = jsonTool.MarshalToString(opt.ForcePushList)
		}

		if len(opt.ForcePushContent) > 0 {
			param["forcepushcontent"] = opt.ForcePushContent
		}

		param["forcepushall"] = strconv.FormatBool(opt.ForcePushAll)

		if len(opt.Bid) > 0 {
			param["bid"] = opt.Bid
		}

		param["useYidun"] = strconv.Itoa(opt.UseYidun)

		if len(opt.YidunAntiCheating) > 0 {
			param["yidunAntiCheating"] = opt.YidunAntiCheating
		}

		param["markRead"] = strconv.Itoa(opt.MarkRead)

		param["checkFriend"] = strconv.FormatBool(opt.CheckFriend)

		if opt.SubType != 0 {
			param["subType"] = strconv.Itoa(opt.SubType)
		}

		param["msgSenderNoSense"] = strconv.Itoa(opt.MsgSenderNoSense)

		param["msgReceiverNoSense"] = strconv.Itoa(opt.MsgReceiverNoSense)

		if len(opt.Env) > 0 {
			param["env"] = opt.Env
		}
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(BaseUrl + sendMsgUrl)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}
