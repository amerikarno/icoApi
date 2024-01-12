package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

// type EmailData struct {
// 	Date             string
// 	Name             string
// 	ConfirmationCode string
// 	Token            string
// }

func GetHTMLTemplate() (htmlBody string) {
	dateLayout := "02 Jan 2006"
	data := EmailData{
		Date: time.Now().Local().Format(dateLayout),
	}

	tmpl, err := template.New(htmlBody).Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>แจ้งยืนยัน "อีเมล" การเปิดบัญชีกับ FINANSIA DIGITAL ASSET</title>
		<style>
		.button {
			position: absolute;
			left: 35%%;
			right: 35%%;
			background-color: #1c87c9;
			border: none;
			border-radius: 10px;
			color: white;
			padding: 10px 10px;
			text-align: center;
			text-decoration: none;
			display: inline-block;
			font-size: 20px;
			margin: 4px 2px;
			cursor: pointer;
		  }
		  .textBegin {
			color: black;
			font-size: 12px;
		  }
		  .textBody {
			color: blue;
			font-weight: 600;
			text-align: center;
			font-size: 20px;
		  }
		</style>
	</head>
	<body>
		<p class="textBegin">วันที่ {{.Date}}</p>
		<p class="textBegin">เรียน คุณ{{.Name}}</p>
		<p class="textBegin">เรื่องแจ้งยืนยัน "อีเมล" การเปิดบัญชีกับ FINANSIA DIGITAL ASSET</p>
		<p class="textBody">กรุณายืนยันอีเมลของท่าน</p>
		<p class="textBody">ตามที่ท่านได้สมัครเปิดบัญชีกับ บล.ฟินันเซีย ดิจิทัล แอสเซท จำกัด</p>
		<p class="textBody">กรุณากดปุ่มด้านล่างเพื่อยืนยันอีเมล</p>
		<p><a href="http://localhost:1323/api/v1/updateCustomerConfirms/{{.ConfirmationCode}}" header={Authorization: Bearer {{.Token}}} class="button">กดยืนยันอีเมล</a></p>

	</body>
	</html>
	`)
	if err != nil {
		log.Fatalf("Failed to create email template: %v", err)
	}

	if err = tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	return
}
