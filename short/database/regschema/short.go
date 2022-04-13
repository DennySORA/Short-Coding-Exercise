package regschema

func regShortStmt() {
	regShortCreateSQL()
	regShortGetSQL()
}

func regShortCreateSQL() {
	RegContext.RegisterSchema("CreateShort",
		`
		INSERT INTO short(
				id,
				url,
				expires
			)
		VALUES ($1, $2, $3);
		`,
	)
}

func regShortGetSQL() {
	RegContext.RegisterSchema("GetShort",
		`
		SELECT id, url, expires FROM short WHERE id = $1;
		`,
	)
}
