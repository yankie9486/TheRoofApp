package query

const (
	EstimateInsert = `
		INSERT INTO estimates (
			id,
			address,
			address2,
			city,
			state,
			zip,
			create_time,
			update_time,
			status
		) VALUES (
			:id,
			:address,
			:address2,
			:city,
			:state,
			:zip,
			:create_time,
			:update_time,
			:status
		) RETURNING id;
	`
)
