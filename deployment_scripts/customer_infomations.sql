create DATABASE open_accounts;

CREATE TABLE customer_informations (
  id varchar(40) DEFAULT NULL,
  th_title varchar(6) DEFAULT NULL,
  th_name varchar(40) DEFAULT NULL,
  th_surname varchar(60) DEFAULT NULL,
  en_title varchar(4) DEFAULT NULL,
  en_name varchar(40) DEFAULT NULL,
  en_surname varchar(60) DEFAULT NULL,
  email varchar(60) DEFAULT NULL,
  mobile_no varchar(11) DEFAULT NULL,
  birth_date varchar(16) DEFAULT NULL,
  marriage_status varchar(10) DEFAULT NULL,
  id_card varchar(16) DEFAULT NULL,
  laser_code varchar(16) DEFAULT NULL,
  personal_agreement tinyint(1) DEFAULT NULL,
  personal_pages tinyint(1) DEFAULT NULL,
  create_at datetime DEFAULT NULL
);
