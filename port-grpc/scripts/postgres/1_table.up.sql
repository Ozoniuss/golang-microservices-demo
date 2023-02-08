CREATE TABLE Ports(
	id TEXT primary key,
	name TEXT,
	city TEXT,
	country TEXT,
    alias TEXT[],
    regions TEXT[],
	province TEXT,
	timezone TEXT,
	code TEXT,
	latitude double precision,
	longitude double precision
);
