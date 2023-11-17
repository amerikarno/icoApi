create table customer_bookbanks (
	customer_id VARCHAR(40) DEFAULT NULL,
	bank_name VARCHAR(40) DEFAULT NULL,
	bank_branch_name VARCHAR(60) DEFAULT NULL,
	bank_account_number VARCHAR(20) DEFAULT NULL,
	is_default BOOLEAN DEFAULT NULL,
	is_deposit BOOLEAN DEFAULT NULL,
	is_withdraw BOOLEAN DEFAULT NULL,
	create_at datetime DEFAULT NULL
);