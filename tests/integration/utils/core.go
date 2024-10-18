package utils

//
//import (
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	"strings"
//	"time"
//
//	"github.com/Masterminds/squirrel"
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//type SortDirection int8
//
//type Postgres struct {
//	maxPoolSize  int
//	connAttempts int
//	connTimeout  time.Duration
//
//	Builder squirrel.StatementBuilderType
//	Pool    *sqlx.DB
//}
//
//type Option func(*Postgres)
//
//type SortOptions struct {
//	Direction SortDirection
//	Columns   []string
//}
//
//type FilterOptions struct {
//	Pattern string
//	Column  string
//}
//
//type Pagination struct {
//	PageNumber int64
//	PageSize   int64
//	Filter     FilterOptions
//	Sort       SortOptions
//}
//
//const (
//	ASC SortDirection = iota
//	DESC
//)
//
//const (
//	defaultMaxPoolSize  = 1
//	defaultConnAttempts = 10
//	defaultConnTimeout  = time.Second
//)
//
//func New(url string, opts ...Option) (*Postgres, error) {
//	pg := &Postgres{
//		maxPoolSize:  defaultMaxPoolSize,
//		connAttempts: defaultConnAttempts,
//		connTimeout:  defaultConnTimeout,
//	}
//
//	for _, opt := range opts {
//		opt(pg)
//	}
//
//	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
//
//	poolConfig, err := pgxpool.ParseConfig(url)
//	if err != nil {
//		return nil, fmt.Errorf("pgxpool parse config: %w", err)
//	}
//
//	poolConfig.MaxConns = int32(pg.maxPoolSize)
//
//	//for pg.connAttempts > 0 {
//	//	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
//	//	if err == nil {
//	//		break
//	//	}
//	//
//	//	time.Sleep(pg.connTimeout)
//	//
//	//	pg.connAttempts--
//	//}
//
//	//if err != nil {
//	//	return nil, fmt.Errorf("out of attempts to connect Postgres: %w", err)
//	//}
//
//	return pg, nil
//}
//
//func (p *Postgres) Close() {
//	if p.Pool != nil {
//		p.Pool.Close()
//	}
//}
//
//func MaxPoolSize(size int) Option {
//	return func(c *Postgres) {
//		c.maxPoolSize = size
//	}
//}
//
//func ConnAttempts(attempts int) Option {
//	return func(c *Postgres) {
//		c.connAttempts = attempts
//	}
//}
//
//func ConnTimeout(timeout time.Duration) Option {
//	return func(c *Postgres) {
//		c.connTimeout = timeout
//	}
//}
//
//func (s SortDirection) String() string {
//	switch s {
//	case DESC:
//		return "DESC"
//	default:
//		return "ASC"
//	}
//}
//
//func SortDirectionFromString(dir string) SortDirection {
//	switch dir {
//	case "ASC":
//		return ASC
//	default:
//		return DESC
//	}
//}
//
//func (s SortOptions) Format() string {
//	return fmt.Sprintf("%s %s", strings.Join(s.Columns, ","), s.Direction.String())
//}
//
//func (p *Pagination) ToSQL(s squirrel.SelectBuilder) squirrel.SelectBuilder {
//	if p.Sort.Columns[0] != "" {
//		s = s.OrderBy(p.Sort.Format())
//	}
//
//	if p.Filter.Column != "" {
//		s = s.Where(squirrel.ILike{p.Filter.Column + "::text": fmt.Sprintf("%%%s%%", p.Filter.Pattern)})
//	}
//	if p.PageSize > 0 {
//		s = s.Limit(uint64(p.PageSize))
//	}
//	if p.PageNumber >= 0 && p.PageSize > 0 {
//		s = s.Offset(uint64(p.PageNumber * p.PageSize))
//	}
//
//	return s
//}
