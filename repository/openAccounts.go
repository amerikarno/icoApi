package repository

import (
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amerikarno/icoApi/models"
	"gorm.io/gorm"
)

type OpenAccountsRepository struct {
	db *gorm.DB
}

func NewOpenAccountsRepository(db *gorm.DB) *OpenAccountsRepository {
	return &OpenAccountsRepository{db: db}
}

func (e *OpenAccountsRepository) CreateCustomerInformation(datas models.CustomerInformations) error {
	columns := []string{"id", "th_title", "th_name", "th_surname", "en_title", "en_name", "en_surname", "email", "mobile_no", "personal_agreement", "birth_date", "marriage_status", "id_card", "laser_code", "create_at"}
	return e.db.Debug().Select(columns).Create(datas).Error
	// return e.db.Debug().Select("id", "th_title", "th_name", "th_surname", "en_title", "en_name", "en_surname", "email", "mobile_no", "personal_agreement", "birth_date", "marriage_status", "id_card", "laser_code", "create_at").Create(datas).Error
}

func (e *OpenAccountsRepository) CreateCustomerAddresses(datas models.CustomerAddresses) error {
	return e.db.Debug().Create(datas).Error
}

func (e *OpenAccountsRepository) CreateCustomerBookbanks(datas models.CustomerBookbanks) error {
	return e.db.Debug().Create(datas).Error
}

func (e *OpenAccountsRepository) CheckReisteredEmail(email string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("email=?", email).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking registered email: %v", err)
	}
	return cust
}

func (e *OpenAccountsRepository) CheckReisteredMobileNo(mobileno string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("mobile_no=?", mobileno).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking registered mobile number: %v", err)
	}
	return cust
}
func (e *OpenAccountsRepository) CheckReisteredCitizenID(cid string) models.CustomerInformations {
	cust := models.CustomerInformations{}
	if err := e.db.Debug().Where("id_card=?", cid).First(&cust).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking citizen id: %v", err)
	}
	return cust
}

func (e *OpenAccountsRepository) UpdatePersonalInformation(personalInfos models.PersonalInformations, cid string) error {
	tx := e.db.Begin()

	if err := tx.Updates(personalInfos.CustomerInformation).Where(cid).Error; err != nil {
		log.Printf("error1: %v", err)
		return err
	}

	if err := tx.Create(personalInfos.CustomerAddresseLists).Error; err != nil {
		log.Printf("error2: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Create(personalInfos.CustomerBookbankLists).Error; err != nil {
		log.Printf("error3: %v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) CreateCustomerExams(customerExams models.CustomerExamsRequest) error {
	tx := e.db.Begin()

	if err := tx.Create(customerExams).Error; err != nil {
		log.Printf("customer exams update error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) CreateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	columns := []string{
		"id",
		"token",
		"is_confirm",
		"confirm_types",
		"create_at",
		"expire_at",
	}
	tx := e.db.Begin()

	if err := tx.Select(columns).Create(customerConfirms).Error; err != nil {
		log.Printf("create customer confirm error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) UpdateCustomerConfirms(customerConfirms models.CustomerConfirmsRequest) error {
	columns := []string{
		"is_confirm",
		"confirm_at",
	}
	tx := e.db.Begin()

	if err := tx.Select(columns).Updates(customerConfirms).Where(customerConfirms.TokenID).Error; err != nil {
		log.Printf("create customer confirm error: %v", err)
		return err
	}

	tx.Commit()

	return nil
}

func (e *OpenAccountsRepository) QueryCustomerConfirmsExpireDT(tokenID string) models.CustomerConfirmsRequest {
	custConfirmDetail := models.CustomerConfirmsRequest{}

	if err := e.db.Debug().Where("token_id = ?", tokenID).First(&custConfirmDetail).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while checking customer token %s, error: %v", tokenID, err)
	}

	return custConfirmDetail
}

func (e *OpenAccountsRepository) GetAllRiskCountry() (riskCountries []models.RiskCountryModel) {

	if err := e.db.Debug().Find(&riskCountries).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while all risk countries, error: %v", err)
	}

	return
}

func (e *OpenAccountsRepository) GetAllRiskOccupation() (riskOccupations []models.RiskOccupationModel) {

	if err := e.db.Debug().Find(&riskOccupations).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while all risk occupation, error: %v", err)
	}

	return
}
func (e *OpenAccountsRepository) GetRiskCountryBy(country string) (riskCountry models.RiskCountryModel) {

	if err := e.db.Debug().Where("country_name = ?", country).First(&riskCountry).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while risk countries by %s, error: %v", country, err)
	}

	return
}

func (e *OpenAccountsRepository) GetRiskOccupation(occupation string, business string) (riskOccupation models.RiskOccupationModel) {
	var condition string
	if len(business) > 0 {
		condition = "business_type = ?"
	} else {
		condition = "business_type is null"
	}

	if err := e.db.Debug().Where("occupation_name = ?", occupation).Where(condition, business).First(&riskOccupation).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("error while risk occupation by %s|%s, error: %v", occupation, business, err)
	}

	return
}

func (e *OpenAccountsRepository) GetHTMLTemplate(thaiName, uid, token string) (htmlBody string) {
	thaiDate := convertDateInThai(time.Now().Local())
	data := models.EmailData{
		Date:             thaiDate,
		Name:             thaiName,
		ConfirmationCode: uid,
		Token:            token,
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

func convertDateInThai(d time.Time) (thaiDate string) {
	var (
		monthStr string
		builder  strings.Builder
	)

	dateStr := strconv.Itoa(d.Day())
	month := int(d.Month())
	yearStr := strconv.Itoa(d.Year() + 543)

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
	builder.WriteString(dateStr)
	builder.WriteString(" ")
	builder.WriteString(monthStr)
	builder.WriteString(" ")
	builder.WriteString(yearStr)

	return builder.String()
}
