CREATE TABLE customer_addresses (
	customer_id VARCHAR(40) DEFAULT NULL,
	home_number VARCHAR(20) DEFAULT NULL,
	village_number SMALLINT DEFAULT NULL,
	village_name	VARCHAR(50) DEFAULT NULL,
	sub_street_name VARCHAR(50) DEFAULT NULL,
	street_name VARCHAR(50) DEFAULT NULL,
	sub_district_name VARCHAR(50) DEFAULT NULL,
	district_name VARCHAR(50) DEFAULT NULL,
	province_name VARCHAR(50) DEFAULT NULL,
	zipcode INT DEFAULT NULL,
	country_name VARCHAR(20) DEFAULT NULL,
	is_registered_address BOOLEAN DEFAULT NULL,
	is_current_address BOOLEAN DEFAULT NULL,
	is_office_address BOOLEAN DEFAULT NULL,
	create_at datetime DEFAULT NULL
);
