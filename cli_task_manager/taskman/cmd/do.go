/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
// 根据用户提供的任务ID，将已完成的任务从数据库中删除
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskman.Db.Update(func(tx *bolt.Tx) error {
			// 用户输入的id参数为字符串，这里转为byte类型
			id := args[0]
			int64Id, _ := strconv.ParseInt(id, 10, 8)
			intId := int(int64Id)
			byteId := itob(intId)

			b := tx.Bucket([]byte("TaskmanBucket"))
			v := b.Get(byteId)

			if len(v) > 0 {
				// v的类型为[]byte,
				// 如果获取到的v值不为空，说明这个id对应的任务存在，则删除它
				err := b.Delete(byteId)
				log.Printf("You have completed the \"%s\" task.", v)
				return err
			} else {
				log.Fatalln("输入的任务ID有误.")
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
