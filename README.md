# Course Golang SQL

### Database

```sql
CREATE TABLE public.mst_customer (
    id integer NOT NULL,
    address character varying(100),
    created_at date,
    date_of_birth date,
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    user_password character varying(255),
    status smallint,
    updated_at date,
    user_name character varying(255),
    domisili_id integer
);


ALTER TABLE public.mst_customer OWNER TO postgres;
```


### Running
> 1. Open file `main.go` and choose method
> 
> 2. Open terminal and running with `go run github.com/jutionck/course-golang-postgre`