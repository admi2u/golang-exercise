package main

import (
	"fmt"
	"log"
	"regexp"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	Val string
}

func main() {
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("postgresql database open error:", err)
	}
	// 表名默认就是结构体名称的复数
	// 这里也可以指定表名生成数据表
	// 更改表明的其他方法参见：https://cloud.tencent.com/developer/article/1667752
	db.Transaction(func(tx *gorm.DB) error {
		tx.Table("phone_numbers").Migrator().CreateTable(&PhoneNumber{})

		// 插入数据
		phs := []PhoneNumber{
			{Val: "1234567890"},
			{Val: "123 456 7891"},
			{Val: "(123) 456 7892"},
			{Val: "(123) 456-7893"},
			{Val: "123-456-7894"},
			{Val: "123-456-7890"},
			{Val: "1234567892"},
			{Val: "(123)456-7892"},
		}

		result := tx.Create(&phs)
		if result.Error != nil {
			return fmt.Errorf("insert data to database error: %v", result.Error)
		}
		log.Printf("insert %d records to database.\n", result.RowsAffected)

		// 读取数据
		var phones []PhoneNumber
		result = tx.Find(&phones)
		if result.Error != nil {
			return fmt.Errorf("retrive data from database error: %v", result.Error)
		}
		log.Printf("retrive %d rows from database.\n", result.RowsAffected)

		// 删除原先的数据
		// 全局删除必须加上条件，否则会报错
		tx.Where("1 = 1").Delete(&PhoneNumber{})

		// 格式化处理数据
		var mapPhoneNumbers = make(map[string]int)
		// 移除数字之外的字符
		re, err := regexp.Compile(`[^\d]`)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		// 将格式化后的字符串保存在map类型中去重
		for _, n := range phones {
			val := re.ReplaceAllString(n.Val, "")
			mapPhoneNumbers[val]++
		}

		var nums []PhoneNumber
		for p := range mapPhoneNumbers {
			num := PhoneNumber{p}
			nums = append(nums, num)
		}
		result = tx.Create(nums)
		if result.Error != nil {
			return fmt.Errorf("re insert data to database error: %v", err)
		}
		log.Printf("re insert %d records to database.", result.RowsAffected)

		return nil
	})

}
