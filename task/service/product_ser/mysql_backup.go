package product_ser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Backup struct {
	DataBaseName string `json:"data_base_name"`
	TableName    string `json:"table_name"`
	WithDate     bool   `json:"with_date"`
	Where        string `json:"where"`
}

var jsonConf = `
[
 {"data_base_name": "xtaoke", "table_name": "brands", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "category", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "collect", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "coupons", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "goods", "with_date": true, "where": "id <= 413131"},
 {"data_base_name": "xtaoke", "table_name": "goods_image", "with_date": false, "where": null},
 {"data_base_name": "xtaoke", "table_name": "goods_tag", "with_date": false, "where": null},
 {"data_base_name": "xtaoke", "table_name": "grab_logs", "with_date": false, "where": null},
 {"data_base_name": "xtaoke", "table_name": "grab_records", "with_date": false, "where": null},
 {"data_base_name": "xtaoke", "table_name": "images", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "media", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "platform", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "promotion", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "shop", "with_date": true, "where": null},
 {"data_base_name": "xtaoke", "table_name": "tags", "with_date": true, "where": null},

 {"data_base_name":"gin_vue_admin","table_name":"casbin_rule","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"exa_customers","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"exa_file_chunks","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"exa_file_upload_and_downloads","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"exa_files","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"jwt_blacklists","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_apis","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_authorities","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_authority_btns","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_authority_menus","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_auto_code_histories","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_auto_codes","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_base_menu_btns","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_base_menu_parameters","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_base_menus","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_data_authority_id","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_dictionaries","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_dictionary_details","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_export_template_condition","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_export_template_join","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_export_templates","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_operation_records","with_date":false,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_user_authority","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_user_setting","with_date":true,"where":null},
 {"data_base_name":"gin_vue_admin","table_name":"sys_users","with_date":true,"where":null}
]
`

func (s *GoodsService) MysqlBackup() {
	var bs []Backup
	err := json.Unmarshal([]byte(jsonConf), &bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = backupGva(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("备份完成")
}

func backupGva(bs []Backup) error {
	// 远程地址
	remoteHost := "47.109.31.106"
	remotePort := "3307"
	username := "root"
	password := "a6381286"
	now := time.Now().Format("20060102_150405")
	for _, b := range bs {
		dir := fmt.Sprintf("backup/%s/%v", now, b.DataBaseName)
		// 创建目录
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录失败: %v\n", err)
			return err
		}
		// 创建备份文件
		backupFile := fmt.Sprintf("%s/%s.sql", dir, b.TableName)
		outfile, err := os.Create(backupFile)
		if err != nil {
			return fmt.Errorf("创建备份文件失败: %v", err)
		}
		defer outfile.Close()
		// 包含表数据

		// 设置 Docker exec mysqldump 命令
		args := []string{
			"exec", "mysql",
			"mysqldump",
			"-h", remoteHost,
			"-P", remotePort,
			"-u", username,
			"-p" + password,
		}

		// 带条件
		if b.Where != "" {
			args = append(args, "--where", b.Where) // 添加 WHERE 条件
		}
		// 不备份数据
		if !b.WithDate {
			args = append(args, "--no-data") // 添加 --no-data 参数
		}
		// 添加数据库和表名
		args = append(args, b.DataBaseName, b.TableName)

		// 构建命令
		cmd := exec.Command("docker", args...)
		// 写入文件
		cmd.Stdout = outfile
		// 运行备份命令
		fmt.Printf("开始备份远程数据库:%v 表: %v\n", b.DataBaseName, b.TableName)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("mysqldump 命令执行失败: %v, 错误详情: %s", err, stderr.String())
		}
		fmt.Println("远程数据库备份完成:", backupFile)
	}
	return nil
}

//func back()  {
//	package product_ser
//
//	import (
//		"fmt"
//	"os"
//	"os/exec"
//	"time"
//	)
//
//	func (s *GoodsService) MysqlBackupPartialData() error {
//		timestamp := time.Now().Format("20060102_150405")
//		backupFile := fmt.Sprintf("partial_backup_%s.sql", timestamp)
//
//		password := os.Getenv("MYSQL_PASSWORD")
//		if password == "" {
//		return fmt.Errorf("MySQL 密码未设置")
//	}
//
//		cmd := exec.Command("mysqldump", "-u", "root", "-p"+password, "xtaoke", "products", "--where", "price > 100")
//
//		outfile, err := os.Create(backupFile)
//		if err != nil {
//		return fmt.Errorf("创建备份文件失败: %v", err)
//	}
//		defer outfile.Close()
//
//		cmd.Stdout = outfile
//
//		fmt.Println("开始备份部分数据")
//		if err = cmd.Run(); err != nil {
//		return fmt.Errorf("mysqldump 命令执行失败: %v", err)
//	}
//
//		fmt.Println("部分数据备份完成:", backupFile)
//		return nil
//	}
//
//
//}
