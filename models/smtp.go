package models

type EmailData struct {
	Date             string
	Name             string
	ConfirmationCode string
	Token            string
}

type SMTPServer struct {
	Host    string `yaml:"host"`
	Port      int    `yaml:"port"`
	Sender    string `yaml:"sender"`
	Password  string `yaml:"password"`
	Recipient string `yaml:"recipient"`
}

type SMTPConfig struct {
	Server SMTPServer `yaml:"server"`
}
