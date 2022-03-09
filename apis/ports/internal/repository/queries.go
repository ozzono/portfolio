package repository

const (
	createPort = `INSERT INTO all_ports (
		name,        -- 01
		ref_name,    -- 02
		city,        -- 03
		country,     -- 04
		alias,       -- 05
		regions,     -- 06
		coordinates, -- 07
		province,    -- 08
		timezone,    -- 09
		unlocs,      -- 10
		code         -- 11
	) values (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
	);`

	upsertPort = `INSERT INTO all_ports (
		name          -- 01
		, ref_name    -- 02
		, city        -- 03
		, country     -- 04
		, alias       -- 05
		, regions     -- 06
		, coordinates -- 07
		, province    -- 08
		, timezone    -- 09
		, unlocs      -- 10
		, code        -- 11
		)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	ON CONFLICT (code)
	DO UPDATE SET 
		name          = $1
		, ref_name    = $2
		, city        = $3
		, country     = $4
		, alias       = $5
		, regions     = $6
		, coordinates = $7
		, province    = $8
		, timezone    = $9
		, unlocs      = $10
	WHERE all_ports.code = $11
	;`  // on code conflict updates the existing row; create a new one otherwise

	getPortByID = `SELECT 
		COALESCE  (id        ,NULL) -- 01
		,COALESCE (name        ,'') -- 02
		,COALESCE (ref_name    ,'') -- 03
		,COALESCE (city        ,'') -- 04
		,COALESCE (country     ,'') -- 05
		,COALESCE (alias       ,'') -- 06
		,COALESCE (regions     ,'') -- 07
		,COALESCE (coordinates ,'') -- 08
		,COALESCE (province    ,'') -- 09
		,COALESCE (timezone    ,'') -- 10
		,COALESCE (unlocs      ,'') -- 11
		,COALESCE (code        ,'') -- 12
	FROM all_ports WHERE id = $1;`

	getPortByCode = `SELECT 
		COALESCE  (id        ,NULL) -- 01
		,COALESCE (name        ,'') -- 02
		,COALESCE (ref_name    ,'') -- 03
		,COALESCE (city        ,'') -- 04
		,COALESCE (country     ,'') -- 05
		,COALESCE (alias       ,'') -- 06
		,COALESCE (regions     ,'') -- 07
		,COALESCE (coordinates ,'') -- 08
		,COALESCE (province    ,'') -- 09
		,COALESCE (timezone    ,'') -- 10
		,COALESCE (unlocs      ,'') -- 11
		,COALESCE (code        ,'') -- 12
	FROM all_ports WHERE code = $1;`

	delPort = `DELETE FROM all_ports WHERE id=$1;`

	allPorts   = `SELECT * FROM all_ports;`
	jsonParsed = `SELECT parsed FROM json_control limit 1;`
	setParsed  = `INSERT INTO json_control (parsed) VALUES ($1)`
)
