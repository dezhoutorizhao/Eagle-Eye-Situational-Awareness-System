package detect_result

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// 订阅主题
func Subscribe(sub *gin.Context) {
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("订阅成功")
	sub.String(200, "订阅成功")
}

func Release() {
	// 发布消息
	token := c.Publish("testtopic/1", 0, false, byteData)
	token.Wait()

	fmt.Println("发送成功")
	time.Sleep(6 * time.Second)
}

func Un_Subscribe(unsub *gin.Context) {
	// 取消订阅
	if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Println("取消订阅成功")
	// 断开连接
	c.Disconnect(250)
	fmt.Println("断开链接成功")
	unsub.String(200, "取消订阅成功")

	time.Sleep(1 * time.Second)
}
