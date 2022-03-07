package repository

const (
	createPort = `INSERT INTO all_ports (
		Name,        -- 01
		RefName,     -- 02
		City,        -- 03
		Country,     -- 04
		Alias,       -- 05
		Regions,     -- 06
		Coordinates, -- 07
		Province,    -- 08
		Timezone,    -- 09
		Unlocs,      -- 10
		Code         -- 11
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
	VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	ON CONFLICT (name, ref_name)
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
		, code        = $11
	WHERE all_ports.name = $1 or all_ports.ref_name = $2
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
		COALESCE  (i         ,NULL)
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
	FROM companies WHERE id = $1;`

	delPort = `DELETE FROM all_ports WHERE id=$1;`

	allPorts = `SELECT * FROM all_ports;`
)
