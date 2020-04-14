package templatemethod

import "fmt"

// 现有2种游戏，基本流程相同，分为初始化，进行游戏，结束游戏

// Playter 作为统一对外接口
// GameFlow 作为步骤行为扩展接口

// Base implement Player
// (GameFlowA + Base) implement GameFlow

type GameFlow interface {
	Initialize()
	StartPlay()
	EndPlay()
}

type Player interface {
	Play()
}

type Base struct {
	GameFlow
	name string
}

func (p *Base) Play() {
	fmt.Println("init game")
	p.GameFlow.Initialize() // 目前 GameFlow 装载了 GameFlowA 这个类型
	fmt.Println("start game")
	p.GameFlow.StartPlay()
	fmt.Println("end game")
	p.GameFlow.EndPlay() // GameFlowA 没有实现，Base 提供了默认实现
}

func (*Base) EndPlay() {
	fmt.Println("i am Base EndPlay()")
}

// 因为Golang实现的接口没有一定放在 这个类型旁边，所以看起来比较麻烦，尽量放一起吧
type GameFlowA struct {
	*Base
}

func NewGameFlowA() Player {
	gameFlow := &GameFlowA{}
	base := &Base{gameFlow, "gameflowA"}
	gameFlow.Base = base
	return gameFlow
}

func (*GameFlowA) Initialize() {
	fmt.Println("GameFlowA Initialize")
}

func (*GameFlowA) StartPlay() {
	fmt.Println("GameFlowA StartPlay")
}
// 子类没有实现，使用 Base 的 EndPlay
//func (*GameFlowA) EndPlay() {
//	fmt.Println("GameFlowA EndPlay")
//}
