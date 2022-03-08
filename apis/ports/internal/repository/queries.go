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
		, code         -- 11
		)
	VALUES( $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	ON CONFLICT (name)
	DO UPDATE SET 
		name          = $2
		, ref_name    = $3
		, city        = $4
		, country     = $5
		, alias       = $6
		, regions     = $7
		, coordinates = $8
		, province    = $9
		, timezone    = $10
		, unlocs      = $11
		, code        = $12
	WHERE all_ports.id = $1
	RETURNING
		COALESCE (id,NULL)
		,COALESCE (name,'')
		,COALESCE (ref_name,'')
		,COALESCE (city,'')
		,COALESCE (country,'')
		,COALESCE (alias,'')
		,COALESCE (regions,'')
		,COALESCE (coordinates,'')
		,COALESCE (province,'')
		,COALESCE (timezone,'')
		,COALESCE (unlocs,'')
		,COALESCE (code,'')
	;`  // on conflict updates the existing row; create a new one otherwise

	getPortByID = `SELECT 
		COALESCE  (id        ,NULL)
		,COALESCE (name        ,'')
		,COALESCE (ref_name    ,'')
		,COALESCE (city        ,'')
		,COALESCE (country     ,'')
		,COALESCE (alias       ,'')
		,COALESCE (regions     ,'')
		,COALESCE (coordinates ,'')
		,COALESCE (province    ,'')
		,COALESCE (timezone    ,'')
		,COALESCE (unlocs      ,'')
		,COALESCE (code        ,'')
	FROM all_ports WHERE id = $1;`

	delPort = `DELETE FROM all_ports WHERE id=$1;`

	allPorts   = `SELECT * FROM all_ports;`
	jsonParsed = `SELECT parsed FROM json_control limit 1;`
	setParsed  = `INSERT INTO json_control (parsed) VALUES ($1)`
)
