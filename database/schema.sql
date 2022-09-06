--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: app_hidden; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_hidden;


--
-- Name: app_private; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_private;


--
-- Name: app_public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_public;


--
-- Name: create_user_authentication(text, text, text, jsonb); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.create_user_authentication(name text, service text, identifier text, details jsonb) RETURNS uuid
    LANGUAGE plpgsql STRICT
    AS $$
declare
  user_id uuid;
begin
  insert into app_public.users(name) values(name) returning id into user_id;
  insert into app_public.authentications(service, identifier, user_id, details) values(service, identifier, user_id, details);
  return user_id;
end;
$$;


--
-- Name: user_id(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.user_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
  select nullif(current_setting('user.id', true), '')::uuid;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: passport_sessions; Type: TABLE; Schema: app_private; Owner: -
--

CREATE TABLE app_private.passport_sessions (
    sid character varying NOT NULL,
    sess json NOT NULL,
    expire timestamp(6) without time zone NOT NULL
);


--
-- Name: authentications; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.authentications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL,
    service text NOT NULL,
    identifier text NOT NULL,
    details jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    avatar_url text,
    CONSTRAINT users_avatar_url_check CHECK ((avatar_url ~ '^https?://[^/]+'::text))
);


--
-- Name: passport_sessions session_pkey; Type: CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.passport_sessions
    ADD CONSTRAINT session_pkey PRIMARY KEY (sid);


--
-- Name: authentications authentications_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_pkey PRIMARY KEY (id);


--
-- Name: authentications authentications_service_identifier_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_service_identifier_key UNIQUE (service, identifier);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_session_expire; Type: INDEX; Schema: app_private; Owner: -
--

CREATE INDEX idx_session_expire ON app_private.passport_sessions USING btree (expire);


--
-- Name: authentications authentications_user_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_public.users(id);


--
-- Name: authentications; Type: ROW SECURITY; Schema: app_public; Owner: -
--

ALTER TABLE app_public.authentications ENABLE ROW LEVEL SECURITY;

--
-- Name: users select_all; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY select_all ON app_public.users FOR SELECT USING (true);


--
-- Name: authentications select_own; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY select_own ON app_public.authentications FOR SELECT USING ((user_id = app_public.user_id()));


--
-- Name: users update_own; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY update_own ON app_public.users FOR UPDATE USING ((id = app_public.user_id()));


--
-- Name: users; Type: ROW SECURITY; Schema: app_public; Owner: -
--

ALTER TABLE app_public.users ENABLE ROW LEVEL SECURITY;

--
-- Name: SCHEMA app_public; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON SCHEMA app_public TO visitor;


--
-- Name: TABLE authentications; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.authentications TO visitor;


--
-- Name: TABLE users; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT,UPDATE ON TABLE app_public.users TO visitor;


--
-- PostgreSQL database dump complete
--

