package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("手机已连接......")
}

func (p *Phone) Stop() {
	fmt.Println("手机已断开连接.....")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("相机已连接.....")
}

func (c *Camera) Stop() {
	fmt.Println("相机已断开连接.....")
}

type Computer struct {
}

func (cm *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	phone := Phone{}
	camera := Camera{}
	computer := Computer{}

	computer.Working(&phone)
	computer.Working(&camera)
}
