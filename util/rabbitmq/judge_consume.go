package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"xoj_judgehost/util/setting"
)

func NewRabbitMQConnect(rabbitmqSetting *setting.RabbitMQSettingS) (*amqp.Channel, *amqp.Connection, error) {

	url := fmt.Sprintf("amqp://%s:%s@%s:%d/judge", rabbitmqSetting.Username, rabbitmqSetting.Password, rabbitmqSetting.Host, rabbitmqSetting.Port)
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		connection.Close()
		return nil, nil, err
	}
	return channel, connection, nil
}
