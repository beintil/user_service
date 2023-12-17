# gosql

## *Postgres SQL builder*

### Create table (support full [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-createtable.html)) examples

###### Table with named primary key constraint
```sql
CREATE TABLE films (
    code        char(5) CONSTRAINT firstkey PRIMARY KEY,
    title       varchar(40) NOT NULL,
    did         integer NOT NULL,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute
);
```
```go
films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)").Constraint().Name("firstkey").PrimaryKey()
films.AddColumn("title").Type("varchar(40)").Constraint().NotNull()
films.AddColumn("did").Type("integer").Constraint().NotNull()
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
```

######  Table with unique named constraint
```sql
CREATE TABLE films (
    code        char(5),
    title       varchar(40),
    did         integer,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute,
    CONSTRAINT production UNIQUE(date_prod)
);
```
```go
films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)")
films.AddColumn("title").Type("varchar(40)")
films.AddColumn("did").Type("integer")
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
films.AddConstraint().Name("production").Unique().Columns().Add("date_prod")
```

###### Table with primary key constraint
```sql
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    PRIMARY KEY(did)
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
distributors.AddConstraint().PrimaryKey().Columns().Add("did")
```

###### Table with primary key in column definition
```sql
CREATE TABLE distributors (
    did     integer PRIMARY KEY,
    name    varchar(40)
);
```
```go
distributors = gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().PrimaryKey()
distributors.AddColumn("name").Type("varchar(40)")
```

###### Table with named constraint not null
```sql
CREATE TABLE distributors (
    did     integer CONSTRAINT no_null NOT NULL,
    name    varchar(40) NOT NULL
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().Name("no_null").NotNull()
distributors.AddColumn("name").Type("varchar(40)").Constraint().NotNull()
```

###### Table with unique column
```sql
CREATE TABLE distributors (
    did     integer,
    name    varchar(40) UNIQUE
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)").Constraint().Unique()
```

###### Table with unique constraint with storage parameter
```sql
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    UNIQUE(name) WITH (fillfactor=70)
)
WITH (fillfactor=70);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
unique := distributors.AddConstraint().Unique()
unique.Columns().Add("name")
unique.IndexParameters().With().Add("fillfactor=70")
distributors.With().Expression().Add("fillfactor=70")
```

###### Table with name constraint primary key on multiple column
```sql
CREATE TABLE films (
    code        char(5),
    title       varchar(40),
    did         integer,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute,
    CONSTRAINT code_title PRIMARY KEY(code,title)
);
```
```go
films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)")
films.AddColumn("title").Type("varchar(40)")
films.AddColumn("did").Type("integer")
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
films.AddConstraint().Name("code_title").PrimaryKey().Columns().Add("code", "title")
```

###### Table with check constraint
```sql
CREATE TABLE distributors (
    did     integer CHECK (did > 100),
    name    varchar(40)
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().Check().Expression().Add("did > 100")
distributors.AddColumn("name").Type("varchar(40)")
```

###### Table with default values in column definition
```sql
CREATE TABLE distributors (
    name      varchar(40) DEFAULT 'Luso Films',
    did       integer DEFAULT nextval('distributors_serial'),
    modtime   timestamp DEFAULT current_timestamp
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("name").Type("varchar(40)").Constraint().Default("'Luso Films'")
distributors.AddColumn("did").Type("integer").Constraint().Default("nextval('distributors_serial')")
distributors.AddColumn("modtime").Type("timestamp").Constraint().Default("current_timestamp")
```

###### Table with tablespace
```sql
CREATE TABLE cinemas (
    id serial,
    name text,
    location text
) TABLESPACE diskvol1;
```
```go
cinemas := gosql.CreateTable("cinemas")
cinemas.AddColumn("id").Type("serial")
cinemas.AddColumn("name").Type("text")
cinemas.AddColumn("location").Type("text")
cinemas.TableSpace("diskvol1")
```

###### Table with options and default constraint
```sql
CREATE TABLE employees OF employee_type (
    PRIMARY KEY (name),
    salary WITH OPTIONS DEFAULT 1000
);
```
```go
employees := gosql.CreateTable("employees")
employees.OfType().Name("employee_type")
employees.OfType().Columns().AddColumn("name").Constraint().PrimaryKey()
salary := employees.OfType().Columns().AddColumn("salary")
salary.Constraint().Default("1000")
salary.WithOptions()
```

###### Table with excluding definition
```sql
CREATE TABLE circles (
    c circle,
    EXCLUDE USING gist (c WITH &&)
);
```
```go
circles := gosql.CreateTable("circles")
circles.AddColumn("c").Type("circle")
exclude := circles.AddConstraint().Exclude().Using("gist")
exclude.ExcludeElement().Expression().Add("c")
exclude.With().Add("&&")
```

###### Table with named check constraint with multiple condition
```sql
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    CONSTRAINT con1 CHECK (did > 100 AND name <> '')
);
```
```go
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
distributors.AddConstraint().Name("con1").
    Check().
    AddExpression("did > 100").
    AddExpression("name <> ''")
```

###### Table with partition by range and clause
```sql
CREATE TABLE measurement_year_month (
    logdate         date not null,
    peaktemp        int,
    unitsales       int
) PARTITION BY RANGE (EXTRACT(YEAR FROM logdate), EXTRACT(MONTH FROM logdate));
```
```go
measurement = gosql.CreateTable("measurement_year_month")
measurement.AddColumn("logdate").Type("date").Constraint().NotNull()
measurement.AddColumn("peaktemp").Type("int")
measurement.AddColumn("unitsales").Type("int")
measurement.Partition().By(gosql.PartitionByRange).Clause("EXTRACT(YEAR FROM logdate)", "EXTRACT(MONTH FROM logdate)")
```

###### Table with partition by hash
```sql
CREATE TABLE orders (
    order_id     bigint not null,
    cust_id      bigint not null,
    status       text
) PARTITION BY HASH (order_id);
```
```go
orders := gosql.CreateTable("orders")
orders.AddColumn("order_id").Type("bigint").Constraint().NotNull()
orders.AddColumn("cust_id").Type("bigint").Constraint().NotNull()
orders.AddColumn("status").Type("text")
orders.Partition().By(gosql.PartitionByHash).Clause("order_id")
```

###### Table with partition for values
```sql
CREATE TABLE measurement_y2016m07
    PARTITION OF measurement (
    unitsales DEFAULT 0
) FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');
```
```go
measurement := gosql.CreateTable("measurement_y2016m07")
measurement.OfPartition().Parent("measurement")
measurement.OfPartition().Columns().AddColumn("unitsales").Constraint().Default("0")
measurement.OfPartition().Values().From().Add("'2016-07-01'")
measurement.OfPartition().Values().To().Add("'2016-08-01'")
```

###### Table with partition for values with constant MINVALUE
```sql
CREATE TABLE measurement_ym_older
    PARTITION OF measurement_year_month
    FOR VALUES FROM (MINVALUE, MINVALUE) TO (2016, 11);
```
```go    
measurement = gosql.CreateTable("measurement_ym_older")
measurement.OfPartition().Parent("measurement_year_month")
measurement.OfPartition().Values().From().Add(PartitionBoundFromMin, PartitionBoundFromMin)
measurement.OfPartition().Values().To().Add("2016", "11")
```

###### Table with partition for values by range
```sql
CREATE TABLE cities_ab
    PARTITION OF cities (
    CONSTRAINT city_id_nonzero CHECK (city_id != 0)
) FOR VALUES IN ('a', 'b') PARTITION BY RANGE (population);
```
```go
cities = gosql.CreateTable("cities_ab")
cities.OfPartition().Parent("cities")
cities.OfPartition().Columns().AddConstraint().Name("city_id_nonzero").Check().AddExpression("city_id != 0")
cities.OfPartition().Values().In().Add("'a'", "'b'")
cities.Partition().By(PartitionByRange).Clause("population")
```

###### Table with partition for values from to
```sql
CREATE TABLE cities_ab_10000_to_100000
    PARTITION OF cities_ab FOR VALUES FROM (10000) TO (100000);
```
```go    
citiesAb := gosql.CreateTable("cities_ab_10000_to_100000")
citiesAb.OfPartition().Parent("cities_ab").Values().From().Add("10000")
citiesAb.OfPartition().Parent("cities_ab").Values().To().Add("100000")
```

###### Table with default partition
```sql
CREATE TABLE cities_partdef
    PARTITION OF cities DEFAULT;
```
```go    
citiesPartdef := gosql.CreateTable("cities_partdef")
citiesPartdef.OfPartition().Parent("cities")
```

###### Table with primary key generated by default
```sql
CREATE TABLE distributors (
    did    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name   varchar(40) NOT NULL CHECK (name <> '')
);
```
```go
distributors := gosql.CreateTable("distributors")
did := distributors.AddColumn("did").Type("integer")
did.Constraint().PrimaryKey()
did.Constraint().Generated().SetDetail(GeneratedByDefault)
name := distributors.AddColumn("name").Type("varchar(40)")
name.Constraint().NotNull()
name.Constraint().Check().Expression().Add("name <> ''")
```

### Create Index (support full [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-createindex.html)) examples

###### Create simple index
```sql
CREATE UNIQUE INDEX title_idx ON films (title);
```
```go
idx := gosql.CreateIndex("films", "title").Name("title_idx").Unique()
OR
idx = gosql.CreateIndex().Table("films").Name("title_idx").Unique()
idx.Expression().Add("title")
OR
idx = gosql.CreateIndex("films", "title").Unique().AutoName()
```

###### Create unique index
```sql
CREATE UNIQUE INDEX title_idx ON films (title) INCLUDE (director, rating);
```
```go
idx := gosql.CreateIndex("films", "title").Name("title_idx").Include("director", "rating").Unique()
OR
idx = gosql.CreateIndex("films", "title").AutoName().Include("director", "rating").Unique()
```

###### Create index with storage param
```sql
CREATE INDEX title_idx ON films (title) WITH (deduplicate_items = off);
```
```go
idx := gosql.CreateIndex("films", "title").Name("title_idx").With("deduplicate_items = off")
```

###### Create index with expression
```sql
CREATE INDEX ON films ((lower(title)));
```
```go
idx := gosql.CreateIndex("films", "(lower(title))")
```

###### Create index with collate
```sql
CREATE INDEX title_idx_german ON films (title COLLATE "de_DE");
```
```go
idx := gosql.CreateIndex("films", `title COLLATE "de_DE"`).Name("title_idx_german")
```

###### Create index nulls first
```sql
CREATE INDEX title_idx_nulls_low ON films (title NULLS FIRST);
```
```go
idx := gosql.CreateIndex("films", `title NULLS FIRST`).Name("title_idx_nulls_low")
```

###### Create index with using
```sql
CREATE INDEX pointloc ON points USING gist (box(location,location));
```
```go
idx := gosql.CreateIndex("points", "box(location,location)").Name("pointloc").Using("gist")
```

###### Create index concurrently
```sql
CREATE INDEX CONCURRENTLY sales_quantity_index ON sales_table (quantity);
```
```go
idx := gosql.CreateIndex("sales_table", "quantity").Name("sales_quantity_index").Concurrently()
```

### Comment examples

###### Comment column
```sql
COMMENT ON COLUMN table_name.column IS 'The column comment';
```
```go
c := gosql.Comment().Column("table_name.column", "The column comment")
```

###### Comment table
```sql
COMMENT ON TABLE table_name IS 'The table comment';
```
```go
c := gosql.Comment().Table("table_name", "The table comment")
```

### Delete query (support full [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-delete.html)) examples

###### Delete with condition
```sql
DELETE FROM films WHERE (kind <> ?);
```
```go
d := gosql.NewDelete().From("films")
d.Where().AddExpression("kind <> ?", "Musical")
```

###### Delete all from table
```sql
DELETE FROM films;
```
```go
d := gosql.NewDelete().From("films")
```

###### Delete with condition returning all
```sql
DELETE FROM tasks WHERE (status = ?) RETURNING *;
```
```go
d := gosql.NewDelete().From("tasks")
d.Returning().Add("*")
d.Where().AddExpression("status = ?", "DONE")
```

###### Delete where sub query
```sql
DELETE FROM tasks WHERE (producer_id IN (SELECT id FROM producers WHERE (name = ?)));
```
```go
sub := gosql.NewSelect()
sub.Columns().Add("id")
sub.From("producers")
sub.Where().AddExpression("name = ?", "foo")
sub.SubQuery = true

d := NewDelete().From("tasks")
d.Where().AddExpression("producer_id IN "+sub.String(), sub.GetArguments()...)
```

### Update query (support full [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-update.html)) examples

###### Update with condition
```sql
UPDATE films SET kind = ? WHERE (kind = ?);
```
```go
u := gosql.NewUpdate().Table("films")
u.Set().Append("kind = ?", "Dramatic")
u.Where().AddExpression("kind = ?", "Drama")
```

###### Update complex expression
```sql
UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT WHERE (city = ? AND date = ?);
```
```go
u := gosql.NewUpdate().Table("weather")
u.Set().Add("temp_lo = temp_lo+1", "temp_hi = temp_lo+15", "prcp = DEFAULT")
u.Where().
    AddExpression("city = ?", "San Francisco").
    AddExpression("date = ?", "2003-07-03")
```

###### Update with returning
```sql
UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT WHERE (city = ? AND date = ?) RETURNING temp_lo, temp_hi, prcp;
```
```go
u := gosql.NewUpdate().Table("weather")
u.Set().Add("temp_lo = temp_lo+1", "temp_hi = temp_lo+15", "prcp = DEFAULT")
u.Returning().Add("temp_lo", "temp_hi", "prcp")
u.Where().
    AddExpression("city = ?", "San Francisco").
    AddExpression("date = ?", "2003-07-03")
```

###### Update from
```sql
UPDATE employees SET sales_count = sales_count + 1 FROM accounts WHERE (accounts.name = ? AND employees.id = accounts.sales_person);
```
```go
u := gosql.NewUpdate().Table("employees").From("accounts")
u.Set().Add("sales_count = sales_count + 1")
u.Where().
    AddExpression("accounts.name = ?", "Acme Corporation").
    AddExpression("employees.id = accounts.sales_person")
```

###### Update sub select
```sql
UPDATE employees SET sales_count = sales_count + 1 WHERE (id = (SELECT sales_person FROM accounts WHERE (name = ?)));
```
```go
sub := gosql.NewSelect()
sub.From("accounts")
sub.Columns().Add("sales_person")
sub.Where().AddExpression("name = ?", "Acme Corporation")
sub.SubQuery = true

u := gosql.NewUpdate().Table("employees")
u.Set().Add("sales_count = sales_count + 1")
u.Where().AddExpression("id = "+sub.String(), sub.GetArguments()...)
```

### Insert query (support full [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-insert.html)) examples

###### Insert values
```sql
INSERT INTO user (name, entity_id, created_at) VALUES (?, ?, ?), (?, ?, ?) RETURNING id, created_at;
```
```go
i := gosql.NewInsert().Into("user")
i.Columns().Add("name", "entity_id", "created_at")
i.Returning().Add("id", "created_at")
i.Columns().Arg("foo", 10, "2021-01-01T10:10:00Z")
i.Columns().Arg("bar", 20, "2021-01-01T10:10:00Z")
```

###### Insert with
```sql
WITH dict AS (SELECT * FROM dictionary d JOIN relation r ON r.dictionary_id = d.id WHERE (some = ?)) INSERT INTO user (name, entity_id, created_at) RETURNING id, created_at;
```
```go
i := gosql.NewInsert().Into("user")
i.Columns().Add("name", "entity_id", "created_at")
i.Returning().Add("id", "created_at")

q := gosql.NewSelect()
q.From("dictionary d")
q.Columns().Add("*")
q.Where().AddExpression("some = ?", 1)
q.Relate("JOIN relation r ON r.dictionary_id = d.id")

i.With().Add("dict", q)
```

###### Insert conflict
```sql
INSERT INTO distributors (did, dname) VALUES (?, ?), (?, ?) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname;
```
```go
i := gosql.NewInsert().Into("distributors")
i.Columns().Add("did", "dname")
i.Columns().Arg(5, "Gizmo Transglobal")
i.Columns().Arg(6, "Associated Computing, Inc")
i.Conflict().Object("did").Action("UPDATE").Set().Add("dname = EXCLUDED.dname")
```

###### Insert conflict no action
```sql
INSERT INTO distributors (did, dname) VALUES (?, ?) ON CONFLICT (did) DO NOTHING;
```
```go
i := gosql.NewInsert().Into("distributors")
i.Columns().Add("did", "dname")
i.Columns().Arg(7, "Redline GmbH")
i.Conflict().Object("did").Action("NOTHING")
```

###### Insert conflict with condition
```sql
INSERT INTO distributors AS d (did, dname) VALUES (?, ?) ON CONFLICT (did) DO UPDATE SET dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')' WHERE (d.zipcode <> '21201');
```
```go
i := gosql.NewInsert().Into("distributors AS d")
i.Columns().Add("did", "dname")
i.Columns().Arg(8, "Anvil Distribution")
i.Conflict().Object("did").Action("UPDATE").Set().Add("dname = EXCLUDED.dname || ' (formerly ' || d.dname || ')'")
i.Conflict().Where().AddExpression("d.zipcode <> '21201'")
```

###### Insert on conflict on constraint
```sql
INSERT INTO distributors (did, dname) VALUES (?, ?) ON CONFLICT ON CONSTRAINT distributors_pkey DO NOTHING;
```
```go
i := gosql.NewInsert().Into("distributors")
i.Columns().Add("did", "dname")
i.Columns().Arg(9, "Antwerp Design")
i.Conflict().Constraint("distributors_pkey").Action("NOTHING")
```

###### Insert and returning
```sql
INSERT INTO distributors (did, dname) VALUES (?, ?) RETURNING did;
```
```go
i := gosql.NewInsert().Into("distributors")
i.Columns().Add("did", "dname")
i.Columns().Arg(1, "XYZ Widgets")
i.Returning().Add("did")
```

### Select query (partial support [PG15 SQL specification](https://www.postgresql.org/docs/current/sql-select.html)) examples

###### Select from table
```sql
SELECT * FROM name
```
```go
s := NewSelect().From("name")
s.Columns().Add("*")
```

###### Select from join using
```sql
SELECT f.title, f.did, d.name, f.date_prod, f.kind
FROM distributors d 
    JOIN films f USING (did)
```
```go
s := NewSelect().From("distributors d").Relate("JOIN films f USING (did)")
s.Columns().Add("f.title", "f.did", "d.name", "f.date_prod", "f.kind")
```

###### Select sum group by
```sql
SELECT kind, sum(len) AS total FROM films GROUP BY kind
```
```go
s := NewSelect().From("films").GroupBy("kind")
s.Columns().Add("kind", "sum(len) AS total")
```

###### Select group by having
```sql
SELECT kind, sum(len) AS total
    FROM films
    GROUP BY kind
    HAVING sum(len) < interval '5 hours'
```
```go
s := NewSelect().From("films").GroupBy("kind")
s.Columns().Add("kind", "sum(len) AS total")
s.Having().AddExpression("sum(len) < interval '5 hours'")
```

###### Select order
```sql
SELECT * FROM distributors ORDER BY name
```
```go
s := NewSelect().From("distributors").AddOrder("name")
s.Columns().Add("*")
```

###### Select union
```sql
SELECT distributors.name
    FROM distributors
WHERE distributors.name LIKE 'W%'
UNION
SELECT actors.name
    FROM actors
WHERE actors.name LIKE 'W%'
```
```go
s := NewSelect().From("distributors")
s.Columns().Add("distributors.name")
s.Where().AddExpression("distributors.name LIKE 'W%'")
u := NewSelect().From("actors")
u.Columns().Add("actors.name")
u.Where().AddExpression("actors.name LIKE 'W%'")
```

###### Select from unnest
```sql
SELECT * FROM unnest(ARRAY['a','b','c','d','e','f']) WITH ORDINALITY
```
```go
s := NewSelect().From("unnest(ARRAY['a','b','c','d','e','f']) WITH ORDINALITY")
s.Columns().Add("*")
```

###### Select from tables
```sql
SELECT m.name AS mname, pname
FROM manufacturers m, LATERAL get_product_names(m.id) pname;
```
```go
s := NewSelect().From("manufacturers m", "LATERAL get_product_names(m.id) pname")
s.Columns().Add("m.name AS mname", "pname")
```

###### Select union intersect
```sql
WITH some AS (
    SELECT * FROM some_table 
    UNION (
        SELECT * FROM some_table_union_1 INTERSECT SELECT * FROM some_table_union_2
     )
) 
SELECT * FROM main_table
```
```go
m := NewSelect().From("main_table")
m.Columns().Add("*")
q := NewSelect().From("some_table")
q.Columns().Add("*")
u1 := NewSelect().From("some_table_union_1")
u1.Columns().Add("*")
u2 := NewSelect().From("some_table_union_2")
u2.Columns().Add("*")
u1.Intersect(u2)
u1.SubQuery = true
q.Union(u1)

m.With().Add("some", q)
```

###### Select union intersect
```sql
WITH RECURSIVE employee_recursive(distance, employee_name, manager_name) AS (
    SELECT 1, employee_name, manager_name
    FROM employee
    WHERE manager_name = 'Mary'
  UNION
    SELECT er.distance + 1, e.employee_name, e.manager_name
    FROM employee_recursive er, employee e
    WHERE er.employee_name = e.manager_name
  )
SELECT distance, employee_name FROM employee_recursive;
```
```go
employee := NewSelect().From("employee")
employee.Columns().Add("1", "employee_name", "manager_name")
employee.Where().AddExpression("manager_name = ?", "Mary")

reqEmployee := NewSelect().From("employee_recursive er", "employee e")
reqEmployee.Columns().Add("er.distance + 1", "e.employee_name", "e.manager_name")
reqEmployee.Where().AddExpression("er.employee_name = e.manager_name")
employee.Union(reqEmployee)

s := NewSelect().From("employee_recursive")
s.Columns().Add("distance", "employee_name")
s.With().Recursive().Add("employee_recursive(distance, employee_name, manager_name)", employee)
```

###### Select left join group order having limit
```sql
SELECT t.id, t.name, c.code 
FROM table AS t
LEFT JOIN country AS c ON c.tid = t.id
GROUP BY t.id, t.name, c.code 
ORDER BY t.name
LIMIT 10 OFFSET 30
```
```go
s := NewSelect().
From("table AS t").
Relate("LEFT JOIN country AS c ON c.tid = t.id").
GroupBy("t.id", "t.name", "c.code").
AddOrder("t.name").
SetPagination(10, 30)
s.Columns().Add("t.id", "t.name", "c.code")
```

### Merge query (full support [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-merge.html)) examples

###### Merge update insert
```sql
MERGE INTO customer_account ca
USING recent_transactions t ON t.customer_id = ca.customer_id
 WHEN MATCHED THEN
    UPDATE SET balance = balance + transaction_value
 WHEN NOT MATCHED THEN
    INSERT (customer_id, balance) VALUES (t.customer_id, t.transaction_value);
```
```go
m := gosql.NewMerge().
    Into("customer_account ca").
    Using("recent_transactions t ON t.customer_id = ca.customer_id")
    
    m.When().Update().Add("balance = balance + transaction_value")
    m.When().Insert().Columns("customer_id", "balance").Values().Add("t.customer_id", "t.transaction_value")
```

###### Merge update insert using sub query
```sql
MERGE INTO customer_account ca
USING (SELECT customer_id, transaction_value FROM recent_transactions) AS t ON t.customer_id = ca.customer_id
 WHEN MATCHED THEN
    UPDATE SET balance = balance + transaction_value
 WHEN NOT MATCHED THEN
    INSERT (customer_id, balance) VALUES (t.customer_id, t.transaction_value);
```
```go
sub := NewSelect().From("recent_transactions")
sub.Columns().Add("customer_id", "transaction_value")
sub.SubQuery = true

m := gosql.NewMerge().
    Into("customer_account ca").
    Using(sub.String() + " AS t ON t.customer_id = ca.customer_id")

m.When().Update().Add("balance = balance + transaction_value")
m.When().Insert().Columns("customer_id", "balance").
    Values().Add("t.customer_id", "t.transaction_value")
```

###### Merge insert update delete
```sql
MERGE INTO wines w
USING wine_stock_changes s ON s.winename = w.winename
 WHEN NOT MATCHED AND s.stock_delta > 0 THEN
    INSERT VALUES(s.winename, s.stock_delta)
 WHEN MATCHED AND w.stock + s.stock_delta > 0 THEN
    UPDATE SET stock = w.stock + s.stock_delta
 WHEN MATCHED THEN
    DELETE;
```
```go
m := gosql.NewMerge().
    Into("wines w").
    Using("wine_stock_changes s ON s.winename = w.winename")

insertWhen := m.When()
insertWhen.Condition().AddExpression("s.stock_delta > 0")
insertWhen.Insert().Values().Add("s.winename", "s.stock_delta")

updateWhen := m.When()
updateWhen.Condition().AddExpression("w.stock + s.stock_delta > 0")
updateWhen.Update().Add("stock = w.stock + s.stock_delta")

m.When().Delete()
```

###### Merge update insert with default fields
```sql
MERGE INTO station_data_actual sda
USING station_data_new sdn ON sda.station_id = sdn.station_id
 WHEN MATCHED THEN
    UPDATE SET a = sdn.a, b = sdn.b, updated = DEFAULT
 WHEN NOT MATCHED THEN
    INSERT (station_id, a, b) VALUES (sdn.station_id, sdn.a, sdn.b);
```

```go
m := gosql.NewMerge().
    Into("station_data_actual sda").
    Using("station_data_new sdn ON sda.station_id = sdn.station_id")

m.When().Update().Add("a = sdn.a", "b = sdn.b", "updated = DEFAULT")

insertWhen := m.When().Insert()
insertWhen.Columns("station_id", "a", "b")
insertWhen.Values().Add("sdn.station_id", "sdn.a", "sdn.b")
```

### Alter table query (full support [PG16 SQL specification](https://www.postgresql.org/docs/current/sql-altertable.html)) examples

###### Add column
```sql
ALTER TABLE distributors ADD COLUMN address varchar(30);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().Column("address", "varchar(30)")
```

###### Add column with default constraint
```sql
ALTER TABLE measurements
    ADD COLUMN mtime timestamp with time zone DEFAULT now();
```
```go
alter := gosql.AlterTable("measurements")
alter.Action().Add().Column("mtime", "timestamp with time zone").Constraint().Default("now()")
```

###### Add and alter column with default constraint
```sql
ALTER TABLE transactions
    ADD COLUMN status varchar(30) DEFAULT 'old',
    ALTER COLUMN status SET default 'current';
```
```go
alter := gosql.AlterTable("transactions")
alter.Action().Add().Column("status", "varchar(30)").Constraint().Default("'old'")
alter.Action().AlterColumn("status").Set().Default("'current'")
```

###### Drop column restrict
```sql
ALTER TABLE distributors DROP COLUMN address RESTRICT;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Drop().Column("address").Restrict()
```

###### Change type of columns
```sql
ALTER TABLE distributors
    ALTER COLUMN address SET DATA TYPE varchar(80),
    ALTER COLUMN name SET DATA TYPE varchar(100);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().AlterColumn("address").Set().DataType("varchar(80)")
alter.Action().AlterColumn("name").Set().DataType("varchar(100)")
```

###### Change type of column using 
```sql
ALTER TABLE foo
    ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone
    USING
    timestamp with time zone 'epoch' + foo_timestamp * interval '1 second';
```
```go
alter := gosql.AlterTable("foo")
alter.Action().AlterColumn("foo_timestamp").Set().DataType("timestamp with time zone").
    Using("timestamp with time zone 'epoch' + foo_timestamp * interval '1 second'")
```

###### Change type of column using with default expression
```sql
ALTER TABLE foo
    ALTER COLUMN foo_timestamp DROP DEFAULT,
    ALTER COLUMN foo_timestamp TYPE timestamp with time zone
        USING
        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second',
    ALTER COLUMN foo_timestamp SET DEFAULT now();
```
```go
alter := gosql.AlterTable("foo")
alter.Action().AlterColumn("foo_timestamp").Drop().Default()
alter.Action().AlterColumn("foo_timestamp").Set().DataType("timestamp with time zone").
    Using("timestamp with time zone 'epoch' + foo_timestamp * interval '1 second'")
alter.Action().AlterColumn("foo_timestamp").Set().Default("now()")
```

###### Rename existing column
```sql
ALTER TABLE distributors RENAME COLUMN address TO city;
```
```go
alter := gosql.AlterTable("distributors").RenameColumn("address", "city")
```

###### Rename table
```sql
ALTER TABLE distributors RENAME TO suppliers;
```
```go
alter := gosql.AlterTable("distributors").Rename("suppliers")
```

###### Rename constraint
```sql
ALTER TABLE distributors RENAME CONSTRAINT zipchk TO zip_check;
```
```go
alter := gosql.AlterTable("distributors").RenameConstraint("zipchk", "zip_check")
```

###### Set column not null 
```sql
ALTER TABLE distributors ALTER COLUMN street SET NOT NULL;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().AlterColumn("street").Set().NotNull()
```

###### Drop column not null 
```sql
ALTER TABLE distributors ALTER COLUMN street DROP NOT NULL;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().AlterColumn("street").Drop().NotNull()
```

###### Add constraint
```sql
ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().TableConstraint().Name("zipchk").Check().AddExpression("char_length(zipcode) = 5")
```

###### Add constraint no inherit
```sql
ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5) NO INHERIT;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().TableConstraint().Name("zipchk").NoInherit().Check().AddExpression("char_length(zipcode) = 5")
```

###### Remove constraint
```sql
ALTER TABLE distributors DROP CONSTRAINT zipchk;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Drop().Constraint("zipchk")
```

###### Remove constraint only from distributors table
```sql
ALTER TABLE ONLY distributors DROP CONSTRAINT zipchk;
```
```go
alter := gosql.AlterTable("distributors").Only()
alter.Action().Drop().Constraint("zipchk")
```

###### Add constraint foreign key
```sql
ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().TableConstraint().Name("distfk").
    ForeignKey().Column("address").References().RefTable("addresses").Column("address")
```

###### Add constraint and validate
```sql
ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address) NOT VALID;
ALTER TABLE distributors VALIDATE CONSTRAINT distfk;
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().NotValid().TableConstraint().Name("distfk").
    ForeignKey().Column("address").References().RefTable("addresses").Column("address")

alter = gosql.AlterTable("distributors")
alter.Action().ValidateConstraint("distfk")
```

###### Add multicolumn unique constraint
```sql
ALTER TABLE distributors ADD CONSTRAINT dist_id_zipcode_key UNIQUE (dist_id, zipcode);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().TableConstraint().Name("dist_id_zipcode_key").Unique().Column("dist_id", "zipcode")
```

###### Add primary key
```sql
ALTER TABLE distributors ADD PRIMARY KEY (dist_id);
```
```go
alter := gosql.AlterTable("distributors")
alter.Action().Add().TableConstraint().PrimaryKey().Column("dist_id")
```

###### Set tablespace
```sql
ALTER TABLE distributors SET TABLESPACE fasttablespace;
```
```go
alter := gosql.AlterTable("distributors").SetTableSpace("fasttablespace")
```

###### Set schema
```sql
ALTER TABLE myschema.distributors SET SCHEMA yourschema;
```
```go
alter := gosql.AlterTable("myschema.distributors").SetSchema("yourschema")
```

###### Recreate primary key without blocking updates while the index is rebuilt
```sql
CREATE UNIQUE INDEX CONCURRENTLY dist_id_temp_idx ON distributors (dist_id);
ALTER TABLE distributors DROP CONSTRAINT distributors_pkey,
                         ADD CONSTRAINT distributors_pkey PRIMARY KEY USING INDEX dist_id_temp_idx;
```
```go
unique := gosql.CreateIndex("distributors", "dist_id").Name("dist_id_temp_idx").Concurrently().Unique()

alter := gosql.AlterTable("distributors")
alter.Action().Drop().Constraint("distributors_pkey")
alter.Action().Add().TableConstraintUsingIndex().Name("distributors_pkey").PrimaryKey().Using("dist_id_temp_idx")
```

###### Attach partition
```sql
ALTER TABLE measurement
    ATTACH PARTITION measurement_y2016m07 FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');
```
```go
alter := gosql.AlterTable("distributors")
bound := alter.AttachPartition("measurement_y2016m07")
bound.From().Add("'2016-07-01'")
bound.To().Add("'2016-08-01'")
```

###### Attach partition to a list-partitioned table
```sql
ALTER TABLE cities
    ATTACH PARTITION cities_ab FOR VALUES IN ('a', 'b');
```
```go
alter := gosql.AlterTable("cities")
alter.AttachPartition("cities_ab").In().Add("'a'", "'b'")
```

###### Attach partition for values with
```sql
ALTER TABLE orders
    ATTACH PARTITION orders_p4 FOR VALUES WITH (MODULUS 4, REMAINDER 3);
```
```go
alter := gosql.AlterTable("orders")
alter.AttachPartition("orders_p4").With().Add("MODULUS 4", "REMAINDER 3")
```

###### Attach partition default
```sql
ALTER TABLE cities
    ATTACH PARTITION cities_partdef DEFAULT;
```
```go
alter := gosql.AlterTable("cities")
alter.AttachDefaultPartition("cities_partdef")
```

###### Attach partition default
```sql
ALTER TABLE measurement
    DETACH PARTITION measurement_y2015m12;;
```
```go
alter := gosql.AlterTable("cities")
alter.DetachPartition("measurement_y2015m12")
```

#### If you find this project useful or want to support the author, you can send tokens to any of these wallets
- Bitcoin: bc1qgx5c3n7q26qv0tngculjz0g78u6mzavy2vg3tf
- Ethereum: 0x62812cb089E0df31347ca32A1610019537bbFe0D
- Dogecoin: DET7fbNzZftp4sGRrBehfVRoi97RiPKajV