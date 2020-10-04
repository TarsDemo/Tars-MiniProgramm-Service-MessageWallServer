package main

import (
	"context"

	"github.com/TarsDemo/Tars-MiniProgramm-Service-MessageWallServer/tars-protocol/LifeService"
)

// MessageWallImp servant implementation
type MessageWallImp struct {
	dataServiceProxy *LifeService.DataService
	dataServiceObj   string
}

// Init servant init
func (imp *MessageWallImp) Init() error {
	//initialize servant here:
	imp.dataServiceProxy = new(LifeService.DataService)
	imp.dataServiceObj = "LifeService.DataServer.DataServiceObj"

	comm.StringToProxy(imp.dataServiceObj, imp.dataServiceProxy)
	return nil
}

// Destroy servant destory
func (imp *MessageWallImp) Destroy() {
	//destroy servant here:
	//...
}

//PostMessage 发布表白
func (imp *MessageWallImp) PostMessage(ctx context.Context, Msg *LifeService.Message) (int32, error) {
	iRet, err := imp.dataServiceProxy.InsertMessage(Msg)

	if err != nil {
		SLOG.Error("Call Remote Server DataServer Error: ", err.Error())
		return iRet, err
	}

	return 0, nil
}

//GetMessageList 获取列表
func (imp *MessageWallImp) GetMessageList(ctx context.Context, Index int32, Date string, wxID string, NextIndex *int32, MsgList *[]LifeService.Message) (int32, error) {
	iRet, err := imp.dataServiceProxy.GetMsgList(Index, Date, wxID, NextIndex, MsgList)

	if err != nil {
		SLOG.Error("Call Remote Server DataServer::getMsgList error: ", err.Error())
		return iRet, err
	}

	return 0, nil
}

//AddLike 点赞
func (imp *MessageWallImp) AddLike(ctx context.Context, messageID string) (int32, error) {
	iRet, err := imp.dataServiceProxy.AddLike(messageID)
	if err != nil {
		SLOG.Error("Call Remote Server DataServer Error: ", err.Error())
		return iRet, err
	}
	return 0, nil
}

//GetLike 获取点赞数
func (imp *MessageWallImp) GetLike(ctx context.Context, messageID string, LikeCount *int32) (int32, error) {
	iRet, err := imp.dataServiceProxy.GetLike(messageID, LikeCount)
	if err != nil {
		SLOG.Error("Call Remote Server Error: ", err.Error())
		return iRet, err
	}
	return 0, nil
}
