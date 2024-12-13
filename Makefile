up:
	goose -dir internal/database/migrations mysql "root:@tcp(localhost:3306)/go_rbac_db" up

down:
	goose -dir internal/database/migrations mysql "root:@tcp(localhost:3306)/go_rbac_db" down