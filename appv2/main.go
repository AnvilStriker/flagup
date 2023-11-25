package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/AnvilStriker/flagup/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "test-app"
	app.Version = "2"
	app.Usage = "test-app"
	app.Flags = []cli.Flag{
		utils.AsV2StringFlag(utils.RedisServerFlag),
		utils.AsV2StringFlag(utils.RedisPrefixFlag),
		utils.AsV2StringSliceFlag(utils.KafkaServersFlag),
		utils.AsV2IntFlag(utils.ESIndexLatencyFlag),
		utils.AsV2IntFlag(utils.ESIndexThresholdFlag),		
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println(utils.RedisServerFlag.Name, c.String(utils.RedisServerFlag.Name))
		fmt.Println(utils.RedisPrefixFlag.Name, c.String(utils.RedisPrefixFlag.Name))
		fmt.Println(utils.KafkaServersFlag.Name, c.String(utils.KafkaServersFlag.Name))
		fmt.Println(utils.ESIndexLatencyFlag.Name, c.String(utils.ESIndexLatencyFlag.Name))
		fmt.Println(utils.ESIndexThresholdFlag.Name, c.String(utils.ESIndexThresholdFlag.Name))
		return nil
	}
	app.Run(os.Args)
}
