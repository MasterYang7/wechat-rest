package wcf

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

type Client struct {
	WcfPath   string     // sdk.dll 路径
	WcfAddr   string     // wcf 监听地址
	WcfPort   int        // wcf 监听端口
	CmdClient *CmdClient // 命令客户端
	MsgClient *MsgClient // 消息客户端
}

// 启动 wcf 服务
// return error 错误信息
func (c *Client) Connect() error {
	// 设置默认值
	if c.WcfAddr == "" {
		c.WcfAddr = "127.0.0.1"
	}
	if c.WcfPort == 0 {
		c.WcfPort = 10080
	}
	// 启动 wcf 服务
	if err := c.wxInitSDK(); err != nil {
		return err
	}
	// 连接 wcf 服务
	c.CmdClient = &CmdClient{
		pbSocket: pbSocket{server: c.buildAddr(c.WcfAddr, c.WcfPort)},
	}
	c.MsgClient = &MsgClient{
		pbSocket: pbSocket{server: c.buildAddr(c.WcfAddr, c.WcfPort+1)},
	}
	return c.CmdClient.dial()
}

// 自动销毁 wcf 服务
func (c *Client) AutoDestory() {
	onquit.Register(func() {
		// 关闭 wcf 连接
		c.MsgClient.Close()
		c.CmdClient.Close()
		// 关闭 wcf 服务
		c.wxDestroySDK()
	})
}

// 启动消息接收器
// param pyq bool 是否接收朋友圈消息
// param fn ...MsgCallback 消息回调函数
// return error 错误信息
func (c *Client) EnrollReceiver(pyq bool, fn ...MsgCallback) error {
	if c.CmdClient.EnableMsgServer(true) != 0 {
		return errors.New("failed to enable msg server")
	}
	time.Sleep(1 * time.Second)
	c.MsgClient.Register(fn...)
	return nil
}

// 关闭消息接收器
// return error 错误信息
func (c *Client) DisableReceiver() error {
	if c.CmdClient.DisableMsgServer() != 0 {
		return errors.New("failed to disable msg server")
	}
	return c.MsgClient.Close()
}

// 构建地址
// param ip string IP地址
// param port int 端口
// return string IP地址和端口
func (c *Client) buildAddr(ip string, port int) string {
	if strings.Contains(ip, ":") {
		return fmt.Sprintf("tcp://[%s]:%d", ip, port)
	} else {
		return fmt.Sprintf("tcp://%s:%d", ip, port)
	}
}

// 调用 sdk.dll 中的函数
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	// 加载 sdk.dll 库
	sdk, err := syscall.LoadDLL(c.WcfPath)
	if err != nil {
		logman.Info("failed to load sdk.dll", "error", err)
		return err
	}
	defer sdk.Release()
	// 查找 fn 函数
	proc, err := sdk.FindProc(fn)
	if err != nil {
		logman.Info("failed to call "+fn, "error", err)
		return err
	}
	// 初始化 fn 服务
	r1, r2, err := proc.Call(a...)
	logman.Warn(fn, "r1", r1, "r2", r2, "error", err)
	return err
}

// 启动 wcf 服务并注入 wechat
// return error 错误信息
func (c *Client) wxInitSDK() error {
	if c.WcfPath == "" {
		return nil
	}
	cmd := exec.Command("tasklist")
	out := bytes.Buffer{}
	cmd.Stdout = &out
	if strings.Contains(out.String(), "WeChat") {
		return errors.New("please close wechat first")
	}
	err := c.sdkCall("WxInitSDK", uintptr(0), uintptr(c.WcfPort))
	time.Sleep(5 * time.Second)
	return err
}

// 关闭 wcf 服务
// return error 错误信息
func (c *Client) wxDestroySDK() error {
	if c.WcfPath == "" {
		return nil
	}
	logman.Info("killing wechat process")
	err := c.sdkCall("WxDestroySDK", uintptr(0))
	cmd := exec.Command("taskkill", "/IM", "WeChat.exe", "/F")
	if err := cmd.Run(); err != nil {
		logman.Warn("failed to kill wechat", "error", err)
	}
	return err
}
