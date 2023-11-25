package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/AnvilStriker/flagup/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "test-app"
	app.Version = "1"
	app.Usage = "test-app"
	app.Flags = []cli.Flag{
		utils.RedisServerFlag,
		utils.RedisPrefixFlag,
		utils.KafkaServersFlag,
		utils.ESIndexLatencyFlag,
		utils.ESIndexThresholdFlag,
	}
	app.Action = func(c *cli.Context) {
		fmt.Println(utils.RedisServerFlag.Name, c.String(utils.RedisServerFlag.Name))
		fmt.Println(utils.RedisPrefixFlag.Name, c.String(utils.RedisPrefixFlag.Name))
		fmt.Println(utils.KafkaServersFlag.Name, c.String(utils.KafkaServersFlag.Name))
		fmt.Println(utils.ESIndexLatencyFlag.Name, c.String(utils.ESIndexLatencyFlag.Name))
		fmt.Println(utils.ESIndexThresholdFlag.Name, c.String(utils.ESIndexThresholdFlag.Name))
	}
	app.Run(os.Args)
}

