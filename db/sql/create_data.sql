CREATE TABLE locations (
	location_id VARCHAR(50) PRIMARY KEY,
	full_address VARCHAR(50), 
	province VARCHAR(50),
	district VARCHAR(50),
	ward VARCHAR(50)
);

CREATE TABLE users (
	user_id VARCHAR(50) PRIMARY KEY,
	user_name VARCHAR(50), 
	age INTEGER,
	gender INTEGER,
	user_type INTEGER,
	status VARCHAR(50),
	location_id VARCHAR(50), 
	FOREIGN KEY (location_id) REFERENCES locations(location_id)
);