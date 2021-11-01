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
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

type Taskman struct {
	Db *bolt.DB // 数据库连接
}

type Task struct {
	Id  int
	Val string
}

var (
	cfgFile string
	taskman Taskman
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "taskman is a CLI for managing your TODOs.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 命令执行之前，打开bolt数据库的文件句柄，获得数据库的连接
		db, err := bolt.Open("taskman.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		// 如果TaskmanBucket不存在，就创建TaskmanBucket
		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("TaskmanBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})
		taskman.Db = db
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 命令执行完毕以后，关闭数据库的文件句柄连接
		if taskman.Db != nil {
			taskman.Db.Close()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskman.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".taskman" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".taskman")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// utils function
// int 转 []byte
// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// []byte 转 uint64
func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
