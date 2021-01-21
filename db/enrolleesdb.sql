-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2
-- PostgreSQL version: 12.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- object: enrolleesdbusr1 | type: ROLE --
-- DROP ROLE IF EXISTS enrolleesdbusr1;
CREATE ROLE enrolleesdbusr1 WITH 
	ENCRYPTED PASSWORD '12345';
-- ddl-end --


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: enrolleesdb | type: DATABASE --
-- -- DROP DATABASE IF EXISTS enrolleesdb;
-- CREATE DATABASE enrolleesdb;
-- -- ddl-end --
-- 

-- object: enrolleesdb | type: SCHEMA --
-- DROP SCHEMA IF EXISTS enrolleesdb CASCADE;
CREATE SCHEMA enrolleesdb;
-- ddl-end --
ALTER SCHEMA enrolleesdb OWNER TO enrolleesdbusr1;
-- ddl-end --

SET search_path TO pg_catalog,public,enrolleesdb;
-- ddl-end --

-- object: enrolleesdb.enrollees | type: TABLE --
-- DROP TABLE IF EXISTS enrolleesdb.enrollees CASCADE;
CREATE TABLE enrolleesdb.enrollees (
	enrolleeid uuid NOT NULL DEFAULT enrolleesdb.uuid_generate_v1(),
	name varchar NOT NULL,
	birthdate date NOT NULL,
	activated boolean NOT NULL,
	phonenumber varchar,
	CONSTRAINT enrolleeid_pk PRIMARY KEY (enrolleeid)

);
-- ddl-end --
COMMENT ON TABLE enrolleesdb.enrollees IS E'enrollees in a health care program\n- Enrollees must have an id, name, and activation status (true or false), and a birth date\n- Enrollees may have a phone number (although they do not have to supply this)';
-- ddl-end --
ALTER TABLE enrolleesdb.enrollees OWNER TO enrolleesdbusr1;
-- ddl-end --

-- object: enrolleesdb.dependents | type: TABLE --
-- DROP TABLE IF EXISTS enrolleesdb.dependents CASCADE;
CREATE TABLE enrolleesdb.dependents (
	dependentid uuid NOT NULL DEFAULT enrolleesdb.uuid_generate_v1(),
	enrolleeid uuid NOT NULL,
	name varchar NOT NULL,
	birthdate date NOT NULL,
	CONSTRAINT dependentid_pk PRIMARY KEY (dependentid)

);
-- ddl-end --
COMMENT ON TABLE enrolleesdb.dependents IS E'- Enrollees may have zero or more dependents\n- Each of an enrollee''s dependents must have an id, name, and birth date';
-- ddl-end --
ALTER TABLE enrolleesdb.dependents OWNER TO enrolleesdbusr1;
-- ddl-end --

-- object: "uuid-ossp" | type: EXTENSION --
-- DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;
CREATE EXTENSION "uuid-ossp"
WITH SCHEMA enrolleesdb;
-- ddl-end --

-- object: enrolleeid_fk | type: CONSTRAINT --
-- ALTER TABLE enrolleesdb.dependents DROP CONSTRAINT IF EXISTS enrolleeid_fk CASCADE;
ALTER TABLE enrolleesdb.dependents ADD CONSTRAINT enrolleeid_fk FOREIGN KEY (enrolleeid)
REFERENCES enrolleesdb.enrollees (enrolleeid) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


