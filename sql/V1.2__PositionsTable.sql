CREATE TABLE IF NOT EXISTS metrobusdb.positions (
	id          VARCHAR(100) PRIMARY KEY,
    vehicleId   VARCHAR(100),      
	startDate   VARCHAR(100),      
	lastUpdate  VARCHAR(100),      
	longitude         FLOAT,
	parentTripId     INTEGER, 
	positionspeed     INTEGER,
	latitude          FLOAT,
	routeId         VARCHAR(100),  
	label             VARCHAR(100),
	postitionOdometer INTEGER,
	tripId            VARCHAR(100),
	vehicleStatus     INTEGER,
	alcaldia VARCHAR(100)
) ENGINE=INNODB;