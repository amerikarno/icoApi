package main

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

const (
	// Replace these values with your email server details
	smtpServer     = "smtp.gmail.com"
	senderEmail    = "karnake@gmail.com"
	senderPassword = "GG@23Apr1978"
	smtpPort       = 587
	// Replace the recipient email with the actual recipient's email address
	recipientEmail = "karnake.r@finansiada.com"
	// Email content in HTML format with a confirmation button
	subject = `แจ้งยืนยัน "อีเมล" การเปิดบัญชีกับ FINANSIA DIGITAL ASSET`

	token = "fda-authen-key"
)

func main() {

	// Generate a random confirmation code or token
	confirmationCode := generateConfirmationCode()
	now := time.Now().Local()
	date := convertDateInThai(now)
	name := "ทดสอบ ส่งอีเมล"
	htmlBody := fmt.Sprintf(`
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
		<p class="textBegin">วันที่ %s</p>
		<p class="textBegin">เรียน คุณ%s</p>
		<p class="textBegin">เรื่องแจ้งยืนยัน "อีเมล" การเปิดบัญชีกับ FINANSIA DIGITAL ASSET</p>
		<p class="textBody">กรุณายืนยันอีเมลของท่าน</p>
		<p class="textBody">ตามที่ท่านได้สมัครเปิดบัญชีกับ บล.ฟินันเซีย ดิจิทัล แอสเซท จำกัด</p>
		<p class="textBody">กรุณากดปุ่มด้านล่างเพื่อยืนยันอีเมล</p>
		<p><a href="http://localhost:1323/api/v1/updateCustomerConfirms/%s" header={Authorization: Bearer %s class="button"}>กดยืนยันอีเมล</a></p>

	</body>
	</html>
	`, date, name, confirmationCode, token)

	// Create a new message
	m := gomail.NewMessage()

	// Set sender and recipient
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)

	// Set subject
	m.SetHeader("Subject", subject)

	// Set HTML body with custom Authorization header
	m.SetBody("text/html", htmlBody)

	// Create a new mailer
	d := gomail.NewDialer(smtpServer, smtpPort, senderEmail, senderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	println("Email sent successfully!")
}

func generateConfirmationCode() string {
	// Implement your logic to generate a confirmation code (e.g., random string)
	return "ABC123"
}

func convertDateInThai(d time.Time) (thaiDate string) {
	var monthStr string
	date := d.Day()
	month := int(d.Month())
	year := d.Year() + 543
	switch month {
	case 1:
		monthStr = "มกราคม"
	case 2:
		monthStr = "กุมภาพันธ์"
	case 3:
		monthStr = "มีนาคม"
	case 4:
		monthStr = "เมษายน"
	case 5:
		monthStr = "พฤษภาคม"
	case 6:
		monthStr = "มิถุนายม"
	case 7:
		monthStr = "กรกฎาคม"
	case 8:
		monthStr = "สิงหาคม"
	case 9:
		monthStr = "กันยายน"
	case 10:
		monthStr = "ตุลาคม"
	case 11:
		monthStr = "พฤษจิกายน"
	case 12:
		monthStr = "ธันวาคม"
	}
	return fmt.Sprintf("%d %s %d", date, monthStr, year)
}
