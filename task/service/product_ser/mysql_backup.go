package product_ser

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func (s *GoodsService) MysqlBackup() error {
	// 定义备份文件名，使用当前时间戳来命名
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("backup_%s.sql", timestamp)

	// 设置 MySQL 备份命令
	cmd := exec.Command("mysqldump", "-u", "your_username", "-p", "your_password", "your_database")

	// 创建备份文件
	outfile, err := os.Create(backupFile)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %v", err)
	}
	defer outfile.Close()

	// 将命令的输出写入文件
	cmd.Stdout = outfile

	// 运行备份命令
	fmt.Println("mysql backup start")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mysqldump command failed: %v", err)
	}

	fmt.Println("mysql backup completed:", backupFile)
	return nil
}
