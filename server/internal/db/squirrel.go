package db

import sq "github.com/Masterminds/squirrel"

var Qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
