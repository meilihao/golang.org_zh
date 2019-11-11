

# sql
`import "database/sql"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package sql provides a generic interface around SQL (or SQL-like)
databases.

The sql package must be used in conjunction with a database driver.
See <a href="https://golang.org/s/sqldrivers">https://golang.org/s/sqldrivers</a> for a list of drivers.

Drivers that do not support context cancellation will not return until
after the query is completed.

For usage examples, see the wiki page at
<a href="https://golang.org/s/sqlwiki">https://golang.org/s/sqlwiki</a>.



<a id="example__openDBCLI">Example (OpenDBCLI)</a>


```go
```

output:
```txt
```

<a id="example__openDBService">Example (OpenDBService)</a>


```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Drivers() []string](#Drivers)
* [func Register(name string, driver driver.Driver)](#Register)
* [type ColumnType](#ColumnType)
  * [func (ci *ColumnType) DatabaseTypeName() string](#ColumnType.DatabaseTypeName)
  * [func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)](#ColumnType.DecimalSize)
  * [func (ci *ColumnType) Length() (length int64, ok bool)](#ColumnType.Length)
  * [func (ci *ColumnType) Name() string](#ColumnType.Name)
  * [func (ci *ColumnType) Nullable() (nullable, ok bool)](#ColumnType.Nullable)
  * [func (ci *ColumnType) ScanType() reflect.Type](#ColumnType.ScanType)
* [type Conn](#Conn)
  * [func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)](#Conn.BeginTx)
  * [func (c *Conn) Close() error](#Conn.Close)
  * [func (c *Conn) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)](#Conn.ExecContext)
  * [func (c *Conn) PingContext(ctx context.Context) error](#Conn.PingContext)
  * [func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)](#Conn.PrepareContext)
  * [func (c *Conn) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)](#Conn.QueryContext)
  * [func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row](#Conn.QueryRowContext)
  * [func (c *Conn) Raw(f func(driverConn interface{}) error) (err error)](#Conn.Raw)
* [type DB](#DB)
  * [func Open(driverName, dataSourceName string) (*DB, error)](#Open)
  * [func OpenDB(c driver.Connector) *DB](#OpenDB)
  * [func (db *DB) Begin() (*Tx, error)](#DB.Begin)
  * [func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)](#DB.BeginTx)
  * [func (db *DB) Close() error](#DB.Close)
  * [func (db *DB) Conn(ctx context.Context) (*Conn, error)](#DB.Conn)
  * [func (db *DB) Driver() driver.Driver](#DB.Driver)
  * [func (db *DB) Exec(query string, args ...interface{}) (Result, error)](#DB.Exec)
  * [func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)](#DB.ExecContext)
  * [func (db *DB) Ping() error](#DB.Ping)
  * [func (db *DB) PingContext(ctx context.Context) error](#DB.PingContext)
  * [func (db *DB) Prepare(query string) (*Stmt, error)](#DB.Prepare)
  * [func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)](#DB.PrepareContext)
  * [func (db *DB) Query(query string, args ...interface{}) (*Rows, error)](#DB.Query)
  * [func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)](#DB.QueryContext)
  * [func (db *DB) QueryRow(query string, args ...interface{}) *Row](#DB.QueryRow)
  * [func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row](#DB.QueryRowContext)
  * [func (db *DB) SetConnMaxLifetime(d time.Duration)](#DB.SetConnMaxLifetime)
  * [func (db *DB) SetMaxIdleConns(n int)](#DB.SetMaxIdleConns)
  * [func (db *DB) SetMaxOpenConns(n int)](#DB.SetMaxOpenConns)
  * [func (db *DB) Stats() DBStats](#DB.Stats)
* [type DBStats](#DBStats)
* [type IsolationLevel](#IsolationLevel)
  * [func (i IsolationLevel) String() string](#IsolationLevel.String)
* [type NamedArg](#NamedArg)
  * [func Named(name string, value interface{}) NamedArg](#Named)
* [type NullBool](#NullBool)
  * [func (n *NullBool) Scan(value interface{}) error](#NullBool.Scan)
  * [func (n NullBool) Value() (driver.Value, error)](#NullBool.Value)
* [type NullFloat64](#NullFloat64)
  * [func (n *NullFloat64) Scan(value interface{}) error](#NullFloat64.Scan)
  * [func (n NullFloat64) Value() (driver.Value, error)](#NullFloat64.Value)
* [type NullInt32](#NullInt32)
  * [func (n *NullInt32) Scan(value interface{}) error](#NullInt32.Scan)
  * [func (n NullInt32) Value() (driver.Value, error)](#NullInt32.Value)
* [type NullInt64](#NullInt64)
  * [func (n *NullInt64) Scan(value interface{}) error](#NullInt64.Scan)
  * [func (n NullInt64) Value() (driver.Value, error)](#NullInt64.Value)
* [type NullString](#NullString)
  * [func (ns *NullString) Scan(value interface{}) error](#NullString.Scan)
  * [func (ns NullString) Value() (driver.Value, error)](#NullString.Value)
* [type NullTime](#NullTime)
  * [func (n *NullTime) Scan(value interface{}) error](#NullTime.Scan)
  * [func (n NullTime) Value() (driver.Value, error)](#NullTime.Value)
* [type Out](#Out)
* [type RawBytes](#RawBytes)
* [type Result](#Result)
* [type Row](#Row)
  * [func (r *Row) Scan(dest ...interface{}) error](#Row.Scan)
* [type Rows](#Rows)
  * [func (rs *Rows) Close() error](#Rows.Close)
  * [func (rs *Rows) ColumnTypes() ([]*ColumnType, error)](#Rows.ColumnTypes)
  * [func (rs *Rows) Columns() ([]string, error)](#Rows.Columns)
  * [func (rs *Rows) Err() error](#Rows.Err)
  * [func (rs *Rows) Next() bool](#Rows.Next)
  * [func (rs *Rows) NextResultSet() bool](#Rows.NextResultSet)
  * [func (rs *Rows) Scan(dest ...interface{}) error](#Rows.Scan)
* [type Scanner](#Scanner)
* [type Stmt](#Stmt)
  * [func (s *Stmt) Close() error](#Stmt.Close)
  * [func (s *Stmt) Exec(args ...interface{}) (Result, error)](#Stmt.Exec)
  * [func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (Result, error)](#Stmt.ExecContext)
  * [func (s *Stmt) Query(args ...interface{}) (*Rows, error)](#Stmt.Query)
  * [func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error)](#Stmt.QueryContext)
  * [func (s *Stmt) QueryRow(args ...interface{}) *Row](#Stmt.QueryRow)
  * [func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *Row](#Stmt.QueryRowContext)
* [type Tx](#Tx)
  * [func (tx *Tx) Commit() error](#Tx.Commit)
  * [func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)](#Tx.Exec)
  * [func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)](#Tx.ExecContext)
  * [func (tx *Tx) Prepare(query string) (*Stmt, error)](#Tx.Prepare)
  * [func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)](#Tx.PrepareContext)
  * [func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error)](#Tx.Query)
  * [func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)](#Tx.QueryContext)
  * [func (tx *Tx) QueryRow(query string, args ...interface{}) *Row](#Tx.QueryRow)
  * [func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row](#Tx.QueryRowContext)
  * [func (tx *Tx) Rollback() error](#Tx.Rollback)
  * [func (tx *Tx) Stmt(stmt *Stmt) *Stmt](#Tx.Stmt)
  * [func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt](#Tx.StmtContext)
* [type TxOptions](#TxOptions)


#### <a id="pkg-examples">Examples</a>
* [Conn.BeginTx](#example_Conn_BeginTx)
* [Conn.ExecContext](#example_Conn_ExecContext)
* [DB.ExecContext](#example_DB_ExecContext)
* [DB.PingContext](#example_DB_PingContext)
* [DB.Prepare](#example_DB_Prepare)
* [DB.QueryContext](#example_DB_QueryContext)
* [DB.QueryRowContext](#example_DB_QueryRowContext)
* [DB.Query (MultipleResultSets)](#example_DB_Query_multipleResultSets)
* [Rows](#example_Rows)
* [Stmt](#example_Stmt)
* [Stmt.QueryRowContext](#example_Stmt_QueryRowContext)
* [Tx.ExecContext](#example_Tx_ExecContext)
* [Tx.Prepare](#example_Tx_Prepare)
* [Tx.Rollback](#example_Tx_Rollback)
* [Package (OpenDBCLI)](#example__openDBCLI)
* [Package (OpenDBService)](#example__openDBService)


#### <a id="pkg-files">Package files</a>
[convert.go](https://golang.org/src/database/sql/convert.go) [ctxutil.go](https://golang.org/src/database/sql/ctxutil.go) [sql.go](https://golang.org/src/database/sql/sql.go) 




## <a id="pkg-variables">Variables</a>
ErrConnDone is returned by any operation that is performed on a connection
that has already been returned to the connection pool.


<pre>var <span id="ErrConnDone">ErrConnDone</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;sql: connection is already closed&#34;)</pre>ErrNoRows is returned by Scan when QueryRow doesn't return a
row. In such a case, QueryRow returns a placeholder *Row value that
defers this error until a Scan.


<pre>var <span id="ErrNoRows">ErrNoRows</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;sql: no rows in result set&#34;)</pre>ErrTxDone is returned by any operation that is performed on a transaction
that has already been committed or rolled back.


<pre>var <span id="ErrTxDone">ErrTxDone</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;sql: transaction has already been committed or rolled back&#34;)</pre>

## <a id="Drivers">func</a> [Drivers](https://golang.org/src/database/sql/sql.go?s=1526:1549#L54)
<pre>func Drivers() []<a href="/pkg/builtin/#string">string</a></pre>
Drivers returns a sorted list of the names of the registered drivers.



## <a id="Register">func</a> [Register](https://golang.org/src/database/sql/sql.go?s=1040:1088#L34)
<pre>func Register(name <a href="/pkg/builtin/#string">string</a>, driver <a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Driver">Driver</a>)</pre>
Register makes a database driver available by the provided name.
If Register is called twice with the same name or if driver is nil,
it panics.





## <a id="ColumnType">type</a> [ColumnType](https://golang.org/src/database/sql/sql.go?s=81182:81422#L2866)
ColumnType contains the name and type of a column.


<pre>type ColumnType struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ColumnType.DatabaseTypeName">func</a> (\*ColumnType) [DatabaseTypeName](https://golang.org/src/database/sql/sql.go?s=83005:83052#L2919)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) DatabaseTypeName() <a href="/pkg/builtin/#string">string</a></pre>
DatabaseTypeName returns the database system name of the column type. If an empty
string is returned the driver type name is not supported.
Consult your driver documentation for a list of driver data types. Length specifiers
are not included.
Common type include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", "BIGINT".




### <a id="ColumnType.DecimalSize">func</a> (\*ColumnType) [DecimalSize](https://golang.org/src/database/sql/sql.go?s=82075:82144#L2897)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) DecimalSize() (precision, scale <a href="/pkg/builtin/#int64">int64</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
DecimalSize returns the scale and precision of a decimal type.
If not applicable or if not supported ok is false.




### <a id="ColumnType.Length">func</a> (\*ColumnType) [Length](https://golang.org/src/database/sql/sql.go?s=81863:81917#L2891)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) Length() (length <a href="/pkg/builtin/#int64">int64</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Length returns the column type length for variable length column types such
as text and binary field types. If the type length is unbounded the value will
be math.MaxInt64 (any database limits will still apply).
If the column type is not variable length, such as an int, or if not supported
by the driver ok is false.




### <a id="ColumnType.Name">func</a> (\*ColumnType) [Name](https://golang.org/src/database/sql/sql.go?s=81473:81508#L2882)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name or alias of the column.




### <a id="ColumnType.Nullable">func</a> (\*ColumnType) [Nullable](https://golang.org/src/database/sql/sql.go?s=82566:82618#L2910)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) Nullable() (nullable, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Nullable reports whether the column may be null.
If a driver does not support this property ok will be false.




### <a id="ColumnType.ScanType">func</a> (\*ColumnType) [ScanType](https://golang.org/src/database/sql/sql.go?s=82379:82424#L2904)
<pre>func (ci *<a href="#ColumnType">ColumnType</a>) ScanType() <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a></pre>
ScanType returns a Go type suitable for scanning into using Rows.Scan.
If a driver does not support this property ScanType will return
the type of an empty interface.




## <a id="Conn">type</a> [Conn](https://golang.org/src/database/sql/sql.go?s=49337:49824#L1767)
Conn represents a single database connection rather than a pool of database
connections. Prefer running queries from DB unless there is a specific
need for a continuous single database connection.

A Conn must call Close to return the connection to the database pool
and may do so concurrently with a running query.

After a call to Close, all operations on the
connection fail with ErrConnDone.


<pre>type Conn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Conn.BeginTx">func</a> (\*Conn) [BeginTx](https://golang.org/src/database/sql/sql.go?s=53499:53572#L1894)
<pre>func (c *<a href="#Conn">Conn</a>) BeginTx(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, opts *<a href="#TxOptions">TxOptions</a>) (*<a href="#Tx">Tx</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
BeginTx starts a transaction.

The provided context is used until the transaction is committed or rolled back.
If the context is canceled, the sql package will roll back
the transaction. Tx.Commit will return an error if the context provided to
BeginTx is canceled.

The provided TxOptions is optional and may be nil if defaults should be used.
If a non-default isolation level is used that the driver doesn't support,
an error will be returned.



<a id="example_Conn_BeginTx">Example</a>


```go
```

output:
```txt
```


### <a id="Conn.Close">func</a> (\*Conn) [Close](https://golang.org/src/database/sql/sql.go?s=54634:54662#L1936)
<pre>func (c *<a href="#Conn">Conn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close returns the connection to the connection pool.
All operations after a Close will return with ErrConnDone.
Close is safe to call concurrently with other operations and will
block until all other operations finish. It may be useful to first
cancel any used context and then call close directly after.




### <a id="Conn.ExecContext">func</a> (\*Conn) [ExecContext](https://golang.org/src/database/sql/sql.go?s=50495:50593#L1806)
<pre>func (c *<a href="#Conn">Conn</a>) ExecContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExecContext executes a query without returning any rows.
The args are for any placeholder parameters in the query.



<a id="example_Conn_ExecContext">Example</a>


```go
```

output:
```txt
```


### <a id="Conn.PingContext">func</a> (\*Conn) [PingContext](https://golang.org/src/database/sql/sql.go?s=50207:50260#L1796)
<pre>func (c *<a href="#Conn">Conn</a>) PingContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>) <a href="/pkg/builtin/#error">error</a></pre>
PingContext verifies the connection to the database is still alive.




### <a id="Conn.PrepareContext">func</a> (\*Conn) [PrepareContext](https://golang.org/src/database/sql/sql.go?s=52004:52083#L1843)
<pre>func (c *<a href="#Conn">Conn</a>) PrepareContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>) (*<a href="#Stmt">Stmt</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
PrepareContext creates a prepared statement for later queries or executions.
Multiple queries or executions may be run concurrently from the
returned statement.
The caller must call the statement's Close method
when the statement is no longer needed.

The provided context is used for the preparation of the statement, not for the
execution of the statement.




### <a id="Conn.QueryContext">func</a> (\*Conn) [QueryContext](https://golang.org/src/database/sql/sql.go?s=50858:50956#L1816)
<pre>func (c *<a href="#Conn">Conn</a>) QueryContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
QueryContext executes a query that returns rows, typically a SELECT.
The args are for any placeholder parameters in the query.




### <a id="Conn.QueryRowContext">func</a> (\*Conn) [QueryRowContext](https://golang.org/src/database/sql/sql.go?s=51440:51531#L1830)
<pre>func (c *<a href="#Conn">Conn</a>) QueryRowContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRowContext executes a query that is expected to return at most one row.
QueryRowContext always returns a non-nil value. Errors are deferred until
Row's Scan method is called.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.




### <a id="Conn.Raw">func</a> (\*Conn) [Raw](https://golang.org/src/database/sql/sql.go?s=52451:52519#L1856)
<pre>func (c *<a href="#Conn">Conn</a>) Raw(f func(driverConn interface{}) <a href="/pkg/builtin/#error">error</a>) (err <a href="/pkg/builtin/#error">error</a>)</pre>
Raw executes f exposing the underlying driver connection for the
duration of f. The driverConn must not be used outside of f.

Once f returns and err is nil, the Conn will continue to be usable
until Conn.Close is called.




## <a id="DB">type</a> [DB](https://golang.org/src/database/sql/sql.go?s=10849:12619#L392)
DB is a database handle representing a pool of zero or more
underlying connections. It's safe for concurrent use by multiple
goroutines.

The sql package creates and frees connections automatically; it
also maintains a free pool of idle connections. If the database has
a concept of per-connection state, such state can be reliably observed
within a transaction (Tx) or connection (Conn). Once DB.Begin is called, the
returned Tx is bound to a single connection. Once Commit or
Rollback is called on the transaction, that transaction's
connection is returned to DB's idle connection pool. The pool size
can be controlled with SetMaxIdleConns.


<pre>type DB struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Open">func</a> [Open](https://golang.org/src/database/sql/sql.go?s=21573:21630#L735)
<pre>func Open(driverName, dataSourceName <a href="/pkg/builtin/#string">string</a>) (*<a href="#DB">DB</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens a database specified by its database driver name and a
driver-specific data source name, usually consisting of at least a
database name and connection information.

Most users will open a database via a driver-specific connection
helper function that returns a *DB. No database drivers are included
in the Go standard library. See <a href="https://golang.org/s/sqldrivers">https://golang.org/s/sqldrivers</a> for
a list of third-party drivers.

Open may just validate its arguments without creating a connection
to the database. To verify that the data source name is valid, call
Ping.

The returned DB is safe for concurrent use by multiple goroutines
and maintains its own pool of idle connections. Thus, the Open
function should be called just once. It is rarely necessary to
close a DB.




### <a id="OpenDB">func</a> [OpenDB](https://golang.org/src/database/sql/sql.go?s=20334:20369#L701)
<pre>func OpenDB(c <a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Connector">Connector</a>) *<a href="#DB">DB</a></pre>
OpenDB opens a database using a Connector, allowing drivers to
bypass a string based data source name.

Most users will open a database via a driver-specific connection
helper function that returns a *DB. No database drivers are included
in the Go standard library. See <a href="https://golang.org/s/sqldrivers">https://golang.org/s/sqldrivers</a> for
a list of third-party drivers.

OpenDB may just validate its arguments without creating a connection
to the database. To verify that the data source name is valid, call
Ping.

The returned DB is safe for concurrent use by multiple goroutines
and maintains its own pool of idle connections. Thus, the OpenDB
function should be called just once. It is rarely necessary to
close a DB.






### <a id="DB.Begin">func</a> (\*DB) [Begin](https://golang.org/src/database/sql/sql.go?s=46757:46791#L1679)
<pre>func (db *<a href="#DB">DB</a>) Begin() (*<a href="#Tx">Tx</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Begin starts a transaction. The default isolation level is dependent on
the driver.




### <a id="DB.BeginTx">func</a> (\*DB) [BeginTx](https://golang.org/src/database/sql/sql.go?s=46329:46401#L1662)
<pre>func (db *<a href="#DB">DB</a>) BeginTx(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, opts *<a href="#TxOptions">TxOptions</a>) (*<a href="#Tx">Tx</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
BeginTx starts a transaction.

The provided context is used until the transaction is committed or rolled back.
If the context is canceled, the sql package will roll back
the transaction. Tx.Commit will return an error if the context provided to
BeginTx is canceled.

The provided TxOptions is optional and may be nil if defaults should be used.
If a non-default isolation level is used that the driver doesn't support,
an error will be returned.




### <a id="DB.Close">func</a> (\*DB) [Close](https://golang.org/src/database/sql/sql.go?s=23258:23285#L799)
<pre>func (db *<a href="#DB">DB</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the database and prevents new queries from starting.
Close then waits for all queries that have started processing on the server
to finish.

It is rare to Close a DB, as the DB handle is meant to be
long-lived and shared between many goroutines.




### <a id="DB.Conn">func</a> (\*DB) [Conn](https://golang.org/src/database/sql/sql.go?s=48489:48543#L1733)
<pre>func (db *<a href="#DB">DB</a>) Conn(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>) (*<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Conn returns a single connection by either opening a new connection
or returning an existing connection from the connection pool. Conn will
block until either a connection is returned or ctx is canceled.
Queries run on the same Conn will be run in the same database session.

Every Conn must be returned to the database pool after use by
calling Conn.Close.




### <a id="DB.Driver">func</a> (\*DB) [Driver](https://golang.org/src/database/sql/sql.go?s=47836:47872#L1718)
<pre>func (db *<a href="#DB">DB</a>) Driver() <a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Driver">Driver</a></pre>
Driver returns the database's underlying driver.




### <a id="DB.Exec">func</a> (\*DB) [Exec](https://golang.org/src/database/sql/sql.go?s=40820:40889#L1480)
<pre>func (db *<a href="#DB">DB</a>) Exec(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Exec executes a query without returning any rows.
The args are for any placeholder parameters in the query.




### <a id="DB.ExecContext">func</a> (\*DB) [ExecContext](https://golang.org/src/database/sql/sql.go?s=40325:40422#L1463)
<pre>func (db *<a href="#DB">DB</a>) ExecContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExecContext executes a query without returning any rows.
The args are for any placeholder parameters in the query.



<a id="example_DB_ExecContext">Example</a>


```go
```

output:
```txt
```


### <a id="DB.Ping">func</a> (\*DB) [Ping](https://golang.org/src/database/sql/sql.go?s=22905:22931#L789)
<pre>func (db *<a href="#DB">DB</a>) Ping() <a href="/pkg/builtin/#error">error</a></pre>
Ping verifies a connection to the database is still alive,
establishing a connection if necessary.




### <a id="DB.PingContext">func</a> (\*DB) [PingContext](https://golang.org/src/database/sql/sql.go?s=22423:22475#L767)
<pre>func (db *<a href="#DB">DB</a>) PingContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>) <a href="/pkg/builtin/#error">error</a></pre>
PingContext verifies a connection to the database is still alive,
establishing a connection if necessary.



<a id="example_DB_PingContext">Example</a>


```go
```

output:
```txt
```


### <a id="DB.Prepare">func</a> (\*DB) [Prepare](https://golang.org/src/database/sql/sql.go?s=38604:38654#L1410)
<pre>func (db *<a href="#DB">DB</a>) Prepare(query <a href="/pkg/builtin/#string">string</a>) (*<a href="#Stmt">Stmt</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Prepare creates a prepared statement for later queries or executions.
Multiple queries or executions may be run concurrently from the
returned statement.
The caller must call the statement's Close method
when the statement is no longer needed.



<a id="example_DB_Prepare">Example</a>


```go
```

output:
```txt
```


### <a id="DB.PrepareContext">func</a> (\*DB) [PrepareContext](https://golang.org/src/database/sql/sql.go?s=37987:38065#L1390)
<pre>func (db *<a href="#DB">DB</a>) PrepareContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>) (*<a href="#Stmt">Stmt</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
PrepareContext creates a prepared statement for later queries or executions.
Multiple queries or executions may be run concurrently from the
returned statement.
The caller must call the statement's Close method
when the statement is no longer needed.

The provided context is used for the preparation of the statement, not for the
execution of the statement.




### <a id="DB.Query">func</a> (\*DB) [Query](https://golang.org/src/database/sql/sql.go?s=42773:42842#L1550)
<pre>func (db *<a href="#DB">DB</a>) Query(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Query executes a query that returns rows, typically a SELECT.
The args are for any placeholder parameters in the query.



<a id="example_DB_Query_multipleResultSets">Example (MultipleResultSets)</a>


```go
```

output:
```txt
```


### <a id="DB.QueryContext">func</a> (\*DB) [QueryContext](https://golang.org/src/database/sql/sql.go?s=42262:42359#L1533)
<pre>func (db *<a href="#DB">DB</a>) QueryContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
QueryContext executes a query that returns rows, typically a SELECT.
The args are for any placeholder parameters in the query.



<a id="example_DB_QueryContext">Example</a>


```go
```

output:
```txt
```


### <a id="DB.QueryRow">func</a> (\*DB) [QueryRow](https://golang.org/src/database/sql/sql.go?s=45722:45784#L1648)
<pre>func (db *<a href="#DB">DB</a>) QueryRow(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRow executes a query that is expected to return at most one row.
QueryRow always returns a non-nil value. Errors are deferred until
Row's Scan method is called.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.




### <a id="DB.QueryRowContext">func</a> (\*DB) [QueryRowContext](https://golang.org/src/database/sql/sql.go?s=45208:45298#L1637)
<pre>func (db *<a href="#DB">DB</a>) QueryRowContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRowContext executes a query that is expected to return at most one row.
QueryRowContext always returns a non-nil value. Errors are deferred until
Row's Scan method is called.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.



<a id="example_DB_QueryRowContext">Example</a>


```go
```

output:
```txt
```


### <a id="DB.SetConnMaxLifetime">func</a> (\*DB) [SetConnMaxLifetime](https://golang.org/src/database/sql/sql.go?s=25802:25851#L906)
<pre>func (db *<a href="#DB">DB</a>) SetConnMaxLifetime(d <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>)</pre>
SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

Expired connections may be closed lazily before reuse.

If d <= 0, connections are reused forever.




### <a id="DB.SetMaxIdleConns">func</a> (\*DB) [SetMaxIdleConns](https://golang.org/src/database/sql/sql.go?s=24444:24480#L854)
<pre>func (db *<a href="#DB">DB</a>) SetMaxIdleConns(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetMaxIdleConns sets the maximum number of connections in the idle
connection pool.

If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.

If n <= 0, no idle connections are retained.

The default max idle connections is currently 2. This may change in
a future release.




### <a id="DB.SetMaxOpenConns">func</a> (\*DB) [SetMaxOpenConns](https://golang.org/src/database/sql/sql.go?s=25374:25410#L888)
<pre>func (db *<a href="#DB">DB</a>) SetMaxOpenConns(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetMaxOpenConns sets the maximum number of open connections to the database.

If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
MaxIdleConns, then MaxIdleConns will be reduced to match the new
MaxOpenConns limit.

If n <= 0, then there is no limit on the number of open connections.
The default is 0 (unlimited).




### <a id="DB.Stats">func</a> (\*DB) [Stats](https://golang.org/src/database/sql/sql.go?s=28093:28122#L997)
<pre>func (db *<a href="#DB">DB</a>) Stats() <a href="#DBStats">DBStats</a></pre>
Stats returns database statistics.




## <a id="DBStats">type</a> [DBStats](https://golang.org/src/database/sql/sql.go?s=27342:28053#L981)
DBStats contains database statistics.


<pre>type DBStats struct {
<span id="DBStats.MaxOpenConnections"></span>    MaxOpenConnections <a href="/pkg/builtin/#int">int</a> <span class="comment">// Maximum number of open connections to the database.</span>

    <span class="comment">// Pool Status</span>
<span id="DBStats.OpenConnections"></span>    OpenConnections <a href="/pkg/builtin/#int">int</a> <span class="comment">// The number of established connections both in use and idle.</span>
<span id="DBStats.InUse"></span>    InUse           <a href="/pkg/builtin/#int">int</a> <span class="comment">// The number of connections currently in use.</span>
<span id="DBStats.Idle"></span>    Idle            <a href="/pkg/builtin/#int">int</a> <span class="comment">// The number of idle connections.</span>

    <span class="comment">// Counters</span>
<span id="DBStats.WaitCount"></span>    WaitCount         <a href="/pkg/builtin/#int64">int64</a>         <span class="comment">// The total number of connections waited for.</span>
<span id="DBStats.WaitDuration"></span>    WaitDuration      <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a> <span class="comment">// The total time blocked waiting for a new connection.</span>
<span id="DBStats.MaxIdleClosed"></span>    MaxIdleClosed     <a href="/pkg/builtin/#int64">int64</a>         <span class="comment">// The total number of connections closed due to SetMaxIdleConns.</span>
<span id="DBStats.MaxLifetimeClosed"></span>    MaxLifetimeClosed <a href="/pkg/builtin/#int64">int64</a>         <span class="comment">// The total number of connections closed due to SetConnMaxLifetime.</span>
}
</pre>











## <a id="IsolationLevel">type</a> [IsolationLevel](https://golang.org/src/database/sql/sql.go?s=3064:3087#L109)
IsolationLevel is the transaction isolation level used in TxOptions.


<pre>type IsolationLevel <a href="/pkg/builtin/#int">int</a></pre>


Various isolation levels that drivers may support in BeginTx.
If a driver does not support a given isolation level an error may be returned.

See <a href="https://en.wikipedia.org/wiki/Isolation_">https://en.wikipedia.org/wiki/Isolation_</a>(database_systems)#Isolation_levels.


<pre>const (
    <span id="LevelDefault">LevelDefault</span> <a href="#IsolationLevel">IsolationLevel</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="LevelReadUncommitted">LevelReadUncommitted</span>
    <span id="LevelReadCommitted">LevelReadCommitted</span>
    <span id="LevelWriteCommitted">LevelWriteCommitted</span>
    <span id="LevelRepeatableRead">LevelRepeatableRead</span>
    <span id="LevelSnapshot">LevelSnapshot</span>
    <span id="LevelSerializable">LevelSerializable</span>
    <span id="LevelLinearizable">LevelLinearizable</span>
)</pre>









### <a id="IsolationLevel.String">func</a> (IsolationLevel) [String](https://golang.org/src/database/sql/sql.go?s=3570:3609#L127)
<pre>func (i <a href="#IsolationLevel">IsolationLevel</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the name of the transaction isolation level.




## <a id="NamedArg">type</a> [NamedArg](https://golang.org/src/database/sql/sql.go?s=1963:2338#L71)
A NamedArg is a named argument. NamedArg values may be used as
arguments to Query or Exec and bind to the corresponding named
parameter in the SQL statement.

For a more concise way to create NamedArg values, see
the Named function.


<pre>type NamedArg struct {

<span id="NamedArg.Name"></span>    <span class="comment">// Name is the name of the parameter placeholder.</span>
    <span class="comment">//</span>
    <span class="comment">// If empty, the ordinal position in the argument list will be</span>
    <span class="comment">// used.</span>
    <span class="comment">//</span>
    <span class="comment">// Name must omit any symbol prefix.</span>
    Name <a href="/pkg/builtin/#string">string</a>

<span id="NamedArg.Value"></span>    <span class="comment">// Value is the value of the parameter.</span>
    <span class="comment">// It may be assigned the same value types as the query</span>
    <span class="comment">// arguments.</span>
    Value interface{}
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Named">func</a> [Named](https://golang.org/src/database/sql/sql.go?s=2672:2723#L100)
<pre>func Named(name <a href="/pkg/builtin/#string">string</a>, value interface{}) <a href="#NamedArg">NamedArg</a></pre>
Named provides a more concise way to create NamedArg values.

Example usage:


	db.ExecContext(ctx, `
	    delete from Invoice
	    where
	        TimeCreated < @end
	        and TimeCreated >= @start;`,
	    sql.Named("start", startTime),
	    sql.Named("end", endTime),
	)






## <a id="NullBool">type</a> [NullBool](https://golang.org/src/database/sql/sql.go?s=7492:7577#L282)
NullBool represents a bool that may be null.
NullBool implements the Scanner interface so
it can be used as a scan destination, similar to NullString.


<pre>type NullBool struct {
<span id="NullBool.Bool"></span>    Bool  <a href="/pkg/builtin/#bool">bool</a>
<span id="NullBool.Valid"></span>    Valid <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if Bool is not NULL</span>
}
</pre>











### <a id="NullBool.Scan">func</a> (\*NullBool) [Scan](https://golang.org/src/database/sql/sql.go?s=7621:7669#L288)
<pre>func (n *<a href="#NullBool">NullBool</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullBool.Value">func</a> (NullBool) [Value](https://golang.org/src/database/sql/sql.go?s=7846:7893#L298)
<pre>func (n <a href="#NullBool">NullBool</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="NullFloat64">type</a> [NullFloat64](https://golang.org/src/database/sql/sql.go?s=6845:6943#L256)
NullFloat64 represents a float64 that may be null.
NullFloat64 implements the Scanner interface so
it can be used as a scan destination, similar to NullString.


<pre>type NullFloat64 struct {
<span id="NullFloat64.Float64"></span>    Float64 <a href="/pkg/builtin/#float64">float64</a>
<span id="NullFloat64.Valid"></span>    Valid   <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if Float64 is not NULL</span>
}
</pre>











### <a id="NullFloat64.Scan">func</a> (\*NullFloat64) [Scan](https://golang.org/src/database/sql/sql.go?s=6987:7038#L262)
<pre>func (n *<a href="#NullFloat64">NullFloat64</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullFloat64.Value">func</a> (NullFloat64) [Value](https://golang.org/src/database/sql/sql.go?s=7217:7267#L272)
<pre>func (n <a href="#NullFloat64">NullFloat64</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="NullInt32">type</a> [NullInt32](https://golang.org/src/database/sql/sql.go?s=6202:6290#L230)
NullInt32 represents an int32 that may be null.
NullInt32 implements the Scanner interface so
it can be used as a scan destination, similar to NullString.


<pre>type NullInt32 struct {
<span id="NullInt32.Int32"></span>    Int32 <a href="/pkg/builtin/#int32">int32</a>
<span id="NullInt32.Valid"></span>    Valid <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if Int32 is not NULL</span>
}
</pre>











### <a id="NullInt32.Scan">func</a> (\*NullInt32) [Scan](https://golang.org/src/database/sql/sql.go?s=6334:6383#L236)
<pre>func (n *<a href="#NullInt32">NullInt32</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullInt32.Value">func</a> (NullInt32) [Value](https://golang.org/src/database/sql/sql.go?s=6558:6606#L246)
<pre>func (n <a href="#NullInt32">NullInt32</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="NullInt64">type</a> [NullInt64](https://golang.org/src/database/sql/sql.go?s=5571:5659#L204)
NullInt64 represents an int64 that may be null.
NullInt64 implements the Scanner interface so
it can be used as a scan destination, similar to NullString.


<pre>type NullInt64 struct {
<span id="NullInt64.Int64"></span>    Int64 <a href="/pkg/builtin/#int64">int64</a>
<span id="NullInt64.Valid"></span>    Valid <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if Int64 is not NULL</span>
}
</pre>











### <a id="NullInt64.Scan">func</a> (\*NullInt64) [Scan](https://golang.org/src/database/sql/sql.go?s=5703:5752#L210)
<pre>func (n *<a href="#NullInt64">NullInt64</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullInt64.Value">func</a> (NullInt64) [Value](https://golang.org/src/database/sql/sql.go?s=5927:5975#L220)
<pre>func (n <a href="#NullInt64">NullInt64</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="NullString">type</a> [NullString](https://golang.org/src/database/sql/sql.go?s=4921:5014#L178)
NullString represents a string that may be null.
NullString implements the Scanner interface so
it can be used as a scan destination:


	var s NullString
	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
	...
	if s.Valid {
	   // use s.String
	} else {
	   // NULL value
	}


<pre>type NullString struct {
<span id="NullString.String"></span>    String <a href="/pkg/builtin/#string">string</a>
<span id="NullString.Valid"></span>    Valid  <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if String is not NULL</span>
}
</pre>











### <a id="NullString.Scan">func</a> (\*NullString) [Scan](https://golang.org/src/database/sql/sql.go?s=5058:5109#L184)
<pre>func (ns *<a href="#NullString">NullString</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullString.Value">func</a> (NullString) [Value](https://golang.org/src/database/sql/sql.go?s=5291:5341#L194)
<pre>func (ns <a href="#NullString">NullString</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="NullTime">type</a> [NullTime](https://golang.org/src/database/sql/sql.go?s=8120:8210#L308)
NullTime represents a time.Time that may be null.
NullTime implements the Scanner interface so
it can be used as a scan destination, similar to NullString.


<pre>type NullTime struct {
<span id="NullTime.Time"></span>    Time  <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>
<span id="NullTime.Valid"></span>    Valid <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Valid is true if Time is not NULL</span>
}
</pre>











### <a id="NullTime.Scan">func</a> (\*NullTime) [Scan](https://golang.org/src/database/sql/sql.go?s=8254:8302#L314)
<pre>func (n *<a href="#NullTime">NullTime</a>) Scan(value interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan implements the Scanner interface.




### <a id="NullTime.Value">func</a> (NullTime) [Value](https://golang.org/src/database/sql/sql.go?s=8485:8532#L324)
<pre>func (n <a href="#NullTime">NullTime</a>) Value() (<a href="/pkg/database/sql/driver/">driver</a>.<a href="/pkg/database/sql/driver/#Value">Value</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Value implements the driver Valuer interface.




## <a id="Out">type</a> [Out](https://golang.org/src/database/sql/sql.go?s=9545:9941#L362)
Out may be used to retrieve OUTPUT value parameters from stored procedures.

Not all drivers and databases support OUTPUT value parameters.

Example usage:


	var outArg string
	_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))


<pre>type Out struct {

<span id="Out.Dest"></span>    <span class="comment">// Dest is a pointer to the value that will be set to the result of the</span>
    <span class="comment">// stored procedure&#39;s OUTPUT parameter.</span>
    Dest interface{}

<span id="Out.In"></span>    <span class="comment">// In is whether the parameter is an INOUT parameter. If so, the input value to the stored</span>
    <span class="comment">// procedure is the dereferenced value of Dest&#39;s pointer, which is then replaced with</span>
    <span class="comment">// the output value.</span>
    In <a href="/pkg/builtin/#bool">bool</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="RawBytes">type</a> [RawBytes](https://golang.org/src/database/sql/sql.go?s=4570:4590#L163)
RawBytes is a byte slice that holds a reference to memory owned by
the database itself. After a Scan into a RawBytes, the slice is only
valid until the next call to Next, Scan, or Close.


<pre>type RawBytes []<a href="/pkg/builtin/#byte">byte</a></pre>











## <a id="Result">type</a> [Result](https://golang.org/src/database/sql/sql.go?s=90076:90582#L3134)
A Result summarizes an executed SQL command.


<pre>type Result interface {
    <span class="comment">// LastInsertId returns the integer generated by the database</span>
    <span class="comment">// in response to a command. Typically this will be from an</span>
    <span class="comment">// &#34;auto increment&#34; column when inserting a new row. Not all</span>
    <span class="comment">// databases support this feature, and the syntax of such</span>
    <span class="comment">// statements varies.</span>
    LastInsertId() (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// RowsAffected returns the number of rows affected by an</span>
    <span class="comment">// update, insert, or delete. Not every database or database</span>
    <span class="comment">// driver may support this.</span>
    RowsAffected() (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="Row">type</a> [Row](https://golang.org/src/database/sql/sql.go?s=88387:88504#L3083)
Row is the result of calling QueryRow to select a single row.


<pre>type Row struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Row.Scan">func</a> (\*Row) [Scan](https://golang.org/src/database/sql/sql.go?s=88790:88835#L3094)
<pre>func (r *<a href="#Row">Row</a>) Scan(dest ...interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan copies the columns from the matched row into the values
pointed at by dest. See the documentation on Rows.Scan for details.
If more than one row matches the query,
Scan uses the first row and discards the rest. If no row matches
the query, Scan returns ErrNoRows.




## <a id="Rows">type</a> [Rows](https://golang.org/src/database/sql/sql.go?s=75693:76415#L2672)
Rows is the result of a query. Its cursor starts before the first row
of the result set. Use Next to advance from row to row.


<pre>type Rows struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>






<a id="example_Rows">Example</a>


```go
```

output:
```txt
```






### <a id="Rows.Close">func</a> (\*Rows) [Close](https://golang.org/src/database/sql/sql.go?s=87838:87867#L3048)
<pre>func (rs *<a href="#Rows">Rows</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Rows, preventing further enumeration. If Next is called
and returns false and there are no further result sets,
the Rows are closed automatically and it will suffice to check the
result of Err. Close is idempotent and does not affect the result of Err.




### <a id="Rows.ColumnTypes">func</a> (\*Rows) [ColumnTypes](https://golang.org/src/database/sql/sql.go?s=80790:80842#L2850)
<pre>func (rs *<a href="#Rows">Rows</a>) ColumnTypes() ([]*<a href="#ColumnType">ColumnType</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ColumnTypes returns column information such as column type, length,
and nullable. Some information may not be available from some drivers.




### <a id="Rows.Columns">func</a> (\*Rows) [Columns](https://golang.org/src/database/sql/sql.go?s=80337:80380#L2833)
<pre>func (rs *<a href="#Rows">Rows</a>) Columns() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Columns returns the column names.
Columns returns an error if the rows are closed.




### <a id="Rows.Err">func</a> (\*Rows) [Err](https://golang.org/src/database/sql/sql.go?s=80023:80050#L2822)
<pre>func (rs *<a href="#Rows">Rows</a>) Err() <a href="/pkg/builtin/#error">error</a></pre>
Err returns the error, if any, that was encountered during iteration.
Err may be called after an explicit or implicit Close.




### <a id="Rows.Next">func</a> (\*Rows) [Next](https://golang.org/src/database/sql/sql.go?s=77715:77742#L2732)
<pre>func (rs *<a href="#Rows">Rows</a>) Next() <a href="/pkg/builtin/#bool">bool</a></pre>
Next prepares the next result row for reading with the Scan method. It
returns true on success, or false if there is no next result row or an error
happened while preparing it. Err should be consulted to distinguish between
the two cases.

Every call to Scan, even the first one, must be preceded by a call to Next.




### <a id="Rows.NextResultSet">func</a> (\*Rows) [NextResultSet](https://golang.org/src/database/sql/sql.go?s=79274:79310#L2786)
<pre>func (rs *<a href="#Rows">Rows</a>) NextResultSet() <a href="/pkg/builtin/#bool">bool</a></pre>
NextResultSet prepares the next result set for reading. It reports whether
there is further result sets, or false if there is no further result set
or if there is an error advancing to it. The Err method should be consulted
to distinguish between the two cases.

After calling NextResultSet, the Next method should always be called before
scanning. If there are further result sets they may not have rows in the result
set.




### <a id="Rows.Scan">func</a> (\*Rows) [Scan](https://golang.org/src/database/sql/sql.go?s=86652:86699#L3011)
<pre>func (rs *<a href="#Rows">Rows</a>) Scan(dest ...interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Scan copies the columns in the current row into the values pointed
at by dest. The number of values in dest must be the same as the
number of columns in Rows.

Scan converts columns read from the database into the following
common Go types and special types provided by the sql package:


	*string
	*[]byte
	*int, *int8, *int16, *int32, *int64
	*uint, *uint8, *uint16, *uint32, *uint64
	*bool
	*float32, *float64
	*interface{}
	*RawBytes
	*Rows (cursor value)
	any type implementing Scanner (see Scanner docs)

In the most simple case, if the type of the value from the source
column is an integer, bool or string type T and dest is of type *T,
Scan simply assigns the value through the pointer.

Scan also converts between string and numeric types, as long as no
information would be lost. While Scan stringifies all numbers
scanned from numeric database columns into *string, scans into
numeric types are checked for overflow. For example, a float64 with
value 300 or a string with value "300" can scan into a uint16, but
not into a uint8, though float64(255) or "255" can scan into a
uint8. One exception is that scans of some float64 numbers to
strings may lose information when stringifying. In general, scan
floating point columns into *float64.

If a dest argument has type *[]byte, Scan saves in that argument a
copy of the corresponding data. The copy is owned by the caller and
can be modified and held indefinitely. The copy can be avoided by
using an argument of type *RawBytes instead; see the documentation
for RawBytes for restrictions on its use.

If an argument has type *interface{}, Scan copies the value
provided by the underlying driver without conversion. When scanning
from a source value of type []byte to *interface{}, a copy of the
slice is made and the caller owns the result.

Source values of type time.Time may be scanned into values of type
*time.Time, *interface{}, *string, or *[]byte. When converting to
the latter two, time.RFC3339Nano is used.

Source values of type bool may be scanned into types *bool,
*interface{}, *string, *[]byte, or *RawBytes.

For scanning into *bool, the source may be true, false, 1, 0, or
string inputs parseable by strconv.ParseBool.

Scan can also convert a cursor returned from a query, such as
"select cursor(select * from my_table) from dual", into a
*Rows value that can itself be scanned from. The parent
select query will close any cursor *Rows if the parent *Rows is closed.




## <a id="Scanner">type</a> [Scanner](https://golang.org/src/database/sql/sql.go?s=8635:9258#L332)
Scanner is an interface used by Scan.


<pre>type Scanner interface {
    <span class="comment">// Scan assigns a value from a database driver.</span>
    <span class="comment">//</span>
    <span class="comment">// The src value will be of one of the following types:</span>
    <span class="comment">//</span>
    <span class="comment">//    int64</span>
    <span class="comment">//    float64</span>
    <span class="comment">//    bool</span>
    <span class="comment">//    []byte</span>
    <span class="comment">//    string</span>
    <span class="comment">//    time.Time</span>
    <span class="comment">//    nil - for NULL values</span>
    <span class="comment">//</span>
    <span class="comment">// An error should be returned if the value cannot be stored</span>
    <span class="comment">// without loss of information.</span>
    <span class="comment">//</span>
    <span class="comment">// Reference types such as []byte are only valid until the next call to Scan</span>
    <span class="comment">// and should not be retained. Their underlying memory is owned by the driver.</span>
    <span class="comment">// If retention is necessary, copy their values before the next call to Scan.</span>
    Scan(src interface{}) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="Stmt">type</a> [Stmt](https://golang.org/src/database/sql/sql.go?s=66615:68033#L2343)
Stmt is a prepared statement.
A Stmt is safe for concurrent use by multiple goroutines.

If a Stmt is prepared on a Tx or Conn, it will be bound to a single
underlying connection forever. If the Tx or Conn closes, the Stmt will
become unusable and all operations will return an error.
If a Stmt is prepared on a DB, it will remain usable for the lifetime of the
DB. When the Stmt needs to execute on a new underlying connection, it will
prepare itself on the new connection automatically.


<pre>type Stmt struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>






<a id="example_Stmt">Example</a>


```go
```

output:
```txt
```






### <a id="Stmt.Close">func</a> (\*Stmt) [Close](https://golang.org/src/database/sql/sql.go?s=74827:74855#L2627)
<pre>func (s *<a href="#Stmt">Stmt</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the statement.




### <a id="Stmt.Exec">func</a> (\*Stmt) [Exec](https://golang.org/src/database/sql/sql.go?s=68901:68957#L2412)
<pre>func (s *<a href="#Stmt">Stmt</a>) Exec(args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Exec executes a prepared statement with the given arguments and
returns a Result summarizing the effect of the statement.




### <a id="Stmt.ExecContext">func</a> (\*Stmt) [ExecContext](https://golang.org/src/database/sql/sql.go?s=68170:68254#L2383)
<pre>func (s *<a href="#Stmt">Stmt</a>) ExecContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExecContext executes a prepared statement with the given arguments and
returns a Result summarizing the effect of the statement.




### <a id="Stmt.Query">func</a> (\*Stmt) [Query](https://golang.org/src/database/sql/sql.go?s=73205:73261#L2583)
<pre>func (s *<a href="#Stmt">Stmt</a>) Query(args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Query executes a prepared query statement with the given arguments
and returns the query results as a *Rows.




### <a id="Stmt.QueryContext">func</a> (\*Stmt) [QueryContext](https://golang.org/src/database/sql/sql.go?s=71802:71886#L2528)
<pre>func (s *<a href="#Stmt">Stmt</a>) QueryContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
QueryContext executes a prepared query statement with the given arguments
and returns the query results as a *Rows.




### <a id="Stmt.QueryRow">func</a> (\*Stmt) [QueryRow](https://golang.org/src/database/sql/sql.go?s=74684:74733#L2622)
<pre>func (s *<a href="#Stmt">Stmt</a>) QueryRow(args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRow executes a prepared query statement with the given arguments.
If an error occurs during the execution of the statement, that error will
be returned by a call to Scan on the returned *Row, which is always non-nil.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.

Example usage:


	var name string
	err := nameByUseridStmt.QueryRow(id).Scan(&name)




### <a id="Stmt.QueryRowContext">func</a> (\*Stmt) [QueryRowContext](https://golang.org/src/database/sql/sql.go?s=74004:74081#L2603)
<pre>func (s *<a href="#Stmt">Stmt</a>) QueryRowContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRowContext executes a prepared query statement with the given arguments.
If an error occurs during the execution of the statement, that error will
be returned by a call to Scan on the returned *Row, which is always non-nil.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.



<a id="example_Stmt_QueryRowContext">Example</a>


```go
```

output:
```txt
```


## <a id="Tx">type</a> [Tx](https://golang.org/src/database/sql/sql.go?s=55052:56025#L1950)
Tx is an in-progress database transaction.

A transaction must end with a call to Commit or Rollback.

After a call to Commit or Rollback, all operations on the
transaction fail with ErrTxDone.

The statements prepared for a transaction by calling
the transaction's Prepare or Stmt methods are closed
by the call to Commit or Rollback.


<pre>type Tx struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Tx.Commit">func</a> (\*Tx) [Commit](https://golang.org/src/database/sql/sql.go?s=58265:58293#L2068)
<pre>func (tx *<a href="#Tx">Tx</a>) Commit() <a href="/pkg/builtin/#error">error</a></pre>
Commit commits the transaction.




### <a id="Tx.Exec">func</a> (\*Tx) [Exec](https://golang.org/src/database/sql/sql.go?s=63749:63818#L2270)
<pre>func (tx *<a href="#Tx">Tx</a>) Exec(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Exec executes a query that doesn't return rows.
For example: an INSERT and UPDATE.




### <a id="Tx.ExecContext">func</a> (\*Tx) [ExecContext](https://golang.org/src/database/sql/sql.go?s=63429:63526#L2260)
<pre>func (tx *<a href="#Tx">Tx</a>) ExecContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (<a href="#Result">Result</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExecContext executes a query that doesn't return rows.
For example: an INSERT and UPDATE.



<a id="example_Tx_ExecContext">Example</a>


```go
```

output:
```txt
```


### <a id="Tx.Prepare">func</a> (\*Tx) [Prepare](https://golang.org/src/database/sql/sql.go?s=60491:60541#L2151)
<pre>func (tx *<a href="#Tx">Tx</a>) Prepare(query <a href="/pkg/builtin/#string">string</a>) (*<a href="#Stmt">Stmt</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Prepare creates a prepared statement for use within a transaction.

The returned statement operates within the transaction and can no longer
be used once the transaction has been committed or rolled back.

To use an existing prepared statement on this transaction, see Tx.Stmt.



<a id="example_Tx_Prepare">Example</a>


```go
```

output:
```txt
```


### <a id="Tx.PrepareContext">func</a> (\*Tx) [PrepareContext](https://golang.org/src/database/sql/sql.go?s=59846:59924#L2129)
<pre>func (tx *<a href="#Tx">Tx</a>) PrepareContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>) (*<a href="#Stmt">Stmt</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
PrepareContext creates a prepared statement for use within a transaction.

The returned statement operates within the transaction and will be closed
when the transaction has been committed or rolled back.

To use an existing prepared statement on this transaction, see Tx.Stmt.

The provided context will be used for the preparation of the context, not
for the execution of the returned statement. The returned statement
will run in the transaction context.




### <a id="Tx.Query">func</a> (\*Tx) [Query](https://golang.org/src/database/sql/sql.go?s=64263:64332#L2285)
<pre>func (tx *<a href="#Tx">Tx</a>) Query(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Query executes a query that returns rows, typically a SELECT.




### <a id="Tx.QueryContext">func</a> (\*Tx) [QueryContext](https://golang.org/src/database/sql/sql.go?s=63957:64054#L2275)
<pre>func (tx *<a href="#Tx">Tx</a>) QueryContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (*<a href="#Rows">Rows</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
QueryContext executes a query that returns rows, typically a SELECT.




### <a id="Tx.QueryRow">func</a> (\*Tx) [QueryRow](https://golang.org/src/database/sql/sql.go?s=65260:65322#L2306)
<pre>func (tx *<a href="#Tx">Tx</a>) QueryRow(query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRow executes a query that is expected to return at most one row.
QueryRow always returns a non-nil value. Errors are deferred until
Row's Scan method is called.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.




### <a id="Tx.QueryRowContext">func</a> (\*Tx) [QueryRowContext](https://golang.org/src/database/sql/sql.go?s=64746:64836#L2295)
<pre>func (tx *<a href="#Tx">Tx</a>) QueryRowContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, query <a href="/pkg/builtin/#string">string</a>, args ...interface{}) *<a href="#Row">Row</a></pre>
QueryRowContext executes a query that is expected to return at most one row.
QueryRowContext always returns a non-nil value. Errors are deferred until
Row's Scan method is called.
If the query selects no rows, the *Row's Scan will return ErrNoRows.
Otherwise, the *Row's Scan scans the first selected row and discards
the rest.




### <a id="Tx.Rollback">func</a> (\*Tx) [Rollback](https://golang.org/src/database/sql/sql.go?s=59298:59328#L2115)
<pre>func (tx *<a href="#Tx">Tx</a>) Rollback() <a href="/pkg/builtin/#error">error</a></pre>
Rollback aborts the transaction.



<a id="example_Tx_Rollback">Example</a>


```go
```

output:
```txt
```


### <a id="Tx.Stmt">func</a> (\*Tx) [Stmt](https://golang.org/src/database/sql/sql.go?s=63240:63276#L2254)
<pre>func (tx *<a href="#Tx">Tx</a>) Stmt(stmt *<a href="#Stmt">Stmt</a>) *<a href="#Stmt">Stmt</a></pre>
Stmt returns a transaction-specific prepared statement from
an existing statement.

Example:


	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	...
	tx, err := db.Begin()
	...
	res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)

The returned statement operates within the transaction and will be closed
when the transaction has been committed or rolled back.




### <a id="Tx.StmtContext">func</a> (\*Tx) [StmtContext](https://golang.org/src/database/sql/sql.go?s=61164:61228#L2170)
<pre>func (tx *<a href="#Tx">Tx</a>) StmtContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, stmt *<a href="#Stmt">Stmt</a>) *<a href="#Stmt">Stmt</a></pre>
StmtContext returns a transaction-specific prepared statement from
an existing statement.

Example:


	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	...
	tx, err := db.Begin()
	...
	res, err := tx.StmtContext(ctx, updateMoney).Exec(123.45, 98293203)

The provided context is used for the preparation of the statement, not for the
execution of the statement.

The returned statement operates within the transaction and will be closed
when the transaction has been committed or rolled back.




## <a id="TxOptions">type</a> [TxOptions](https://golang.org/src/database/sql/sql.go?s=4194:4372#L153)
TxOptions holds the transaction options to be used in DB.BeginTx.


<pre>type TxOptions struct {
<span id="TxOptions.Isolation"></span>    <span class="comment">// Isolation is the transaction isolation level.</span>
    <span class="comment">// If zero, the driver or database&#39;s default level is used.</span>
    Isolation <a href="#IsolationLevel">IsolationLevel</a>
<span id="TxOptions.ReadOnly"></span>    ReadOnly  <a href="/pkg/builtin/#bool">bool</a>
}
</pre>














