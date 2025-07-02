--
-- PostgreSQL database dump
--

-- Dumped from database version 15.13 (Debian 15.13-1.pgdg120+1)
-- Dumped by pg_dump version 15.6

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: attempts; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.attempts (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    quiz_id bigint,
    user_id bigint
);


ALTER TABLE public.attempts OWNER TO appuser;

--
-- Name: bots; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.bots (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    owner_id bigint,
    assistant_id text,
    name text,
    description text,
    published boolean DEFAULT false
);


ALTER TABLE public.bots OWNER TO appuser;

--
-- Name: codes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.codes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    code text,
    nonce text,
    email text
);


ALTER TABLE public.codes OWNER TO appuser;

--
-- Name: deliveries; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.deliveries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    message_id bigint,
    room_id bigint,
    recipient_id bigint,
    sender_id bigint,
    sent_at timestamp with time zone
);


ALTER TABLE public.deliveries OWNER TO appuser;

--
-- Name: friends; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.friends (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    a_id bigint,
    b_id bigint,
    b_role text
);


ALTER TABLE public.friends OWNER TO appuser;

--
-- Name: gradients; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.gradients (
    id bigint NOT NULL,
    created_at text,
    user_id bigint,
    gradient bytea
);


ALTER TABLE public.gradients OWNER TO appuser;

--
-- Name: images; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    url text,
    user_id bigint
);


ALTER TABLE public.images OWNER TO appuser;

--
-- Name: kids_codes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.kids_codes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    code text,
    nonce text,
    user_id bigint
);


ALTER TABLE public.kids_codes OWNER TO appuser;

--
-- Name: kids_parents; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.kids_parents (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    kid_id bigint,
    parent_id bigint
);


ALTER TABLE public.kids_parents OWNER TO appuser;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    sender_id bigint,
    room_id bigint,
    body text
);


ALTER TABLE public.messages OWNER TO appuser;

--
-- Name: migration_version; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.migration_version (
    version bigint
);


ALTER TABLE public.migration_version OWNER TO appuser;

--
-- Name: notes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.notes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    owner_id bigint,
    body text,
    published boolean DEFAULT false
);


ALTER TABLE public.notes OWNER TO appuser;

--
-- Name: postcards; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.postcards (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    sender bigint,
    recipient bigint,
    subject text,
    body text,
    state text DEFAULT 'draft'::text
);


ALTER TABLE public.postcards OWNER TO appuser;

--
-- Name: questions; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.questions (
    id bigint NOT NULL,
    created_at text,
    quiz_id bigint,
    text text,
    answer text
);


ALTER TABLE public.questions OWNER TO appuser;

--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.quizzes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    name text,
    description text,
    published boolean DEFAULT false
);


ALTER TABLE public.quizzes OWNER TO appuser;

--
-- Name: responses; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.responses (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    quiz_id bigint,
    user_id bigint,
    attempt_id bigint,
    question_id bigint,
    text text
);


ALTER TABLE public.responses OWNER TO appuser;

--
-- Name: room_users; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.room_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    room_id bigint,
    user_id bigint
);


ALTER TABLE public.room_users OWNER TO appuser;

--
-- Name: rooms; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.rooms (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    key text
);


ALTER TABLE public.rooms OWNER TO appuser;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    user_id bigint,
    key text
);


ALTER TABLE public.sessions OWNER TO appuser;

--
-- Name: threads; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.threads (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    thread_id text,
    assistant_id text,
    user_id bigint
);


ALTER TABLE public.threads OWNER TO appuser;

--
-- Name: users; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    username text,
    email text,
    avatar_url text DEFAULT 'https://www.gravatar.com/avatar/?d=mp'::text,
    is_parent boolean DEFAULT false,
    bio text DEFAULT ''::text,
    become_user_id bigint,
    admin boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO appuser;

--
-- Name: codes idx_16393_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.codes
    ADD CONSTRAINT idx_16393_codes_pkey PRIMARY KEY (id);


--
-- Name: sessions idx_16398_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT idx_16398_sessions_pkey PRIMARY KEY (id);


--
-- Name: kids_codes idx_16403_kids_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_codes
    ADD CONSTRAINT idx_16403_kids_codes_pkey PRIMARY KEY (id);


--
-- Name: kids_parents idx_16408_kids_parents_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_parents
    ADD CONSTRAINT idx_16408_kids_parents_pkey PRIMARY KEY (id);


--
-- Name: gradients idx_16411_gradients_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.gradients
    ADD CONSTRAINT idx_16411_gradients_pkey PRIMARY KEY (id);


--
-- Name: rooms idx_16416_rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT idx_16416_rooms_pkey PRIMARY KEY (id);


--
-- Name: messages idx_16421_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT idx_16421_messages_pkey PRIMARY KEY (id);


--
-- Name: room_users idx_16426_room_users_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.room_users
    ADD CONSTRAINT idx_16426_room_users_pkey PRIMARY KEY (id);


--
-- Name: deliveries idx_16429_deliveries_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT idx_16429_deliveries_pkey PRIMARY KEY (id);


--
-- Name: friends idx_16432_friends_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT idx_16432_friends_pkey PRIMARY KEY (id);


--
-- Name: images idx_16437_images_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT idx_16437_images_pkey PRIMARY KEY (id);


--
-- Name: responses idx_16442_responses_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT idx_16442_responses_pkey PRIMARY KEY (id);


--
-- Name: questions idx_16447_questions_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT idx_16447_questions_pkey PRIMARY KEY (id);


--
-- Name: attempts idx_16452_attempts_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.attempts
    ADD CONSTRAINT idx_16452_attempts_pkey PRIMARY KEY (id);


--
-- Name: users idx_16455_users_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT idx_16455_users_pkey PRIMARY KEY (id);


--
-- Name: quizzes idx_16464_quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT idx_16464_quizzes_pkey PRIMARY KEY (id);


--
-- Name: postcards idx_16470_postcards_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.postcards
    ADD CONSTRAINT idx_16470_postcards_pkey PRIMARY KEY (id);


--
-- Name: threads idx_16476_threads_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.threads
    ADD CONSTRAINT idx_16476_threads_pkey PRIMARY KEY (id);


--
-- Name: bots idx_16481_bots_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.bots
    ADD CONSTRAINT idx_16481_bots_pkey PRIMARY KEY (id);


--
-- Name: notes idx_16487_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT idx_16487_notes_pkey PRIMARY KEY (id);


--
-- Name: idx_16398_sqlite_autoindex_sessions_1; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16398_sqlite_autoindex_sessions_1 ON public.sessions USING btree (key);


--
-- Name: idx_16426_sqlite_autoindex_room_users_1; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16426_sqlite_autoindex_room_users_1 ON public.room_users USING btree (room_id, user_id);


--
-- Name: idx_16429_sqlite_autoindex_deliveries_1; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16429_sqlite_autoindex_deliveries_1 ON public.deliveries USING btree (message_id, recipient_id);


--
-- Name: idx_16432_uidx_friends_a_b; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16432_uidx_friends_a_b ON public.friends USING btree (a_id, b_id);


--
-- Name: idx_16442_sqlite_autoindex_responses_1; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16442_sqlite_autoindex_responses_1 ON public.responses USING btree (attempt_id, question_id);


--
-- Name: idx_16455_sqlite_autoindex_users_1; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16455_sqlite_autoindex_users_1 ON public.users USING btree (username);


--
-- Name: idx_16455_sqlite_autoindex_users_2; Type: INDEX; Schema: public; Owner: appuser
--

CREATE UNIQUE INDEX idx_16455_sqlite_autoindex_users_2 ON public.users USING btree (email);


--
-- Name: attempts attempts_quiz_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.attempts
    ADD CONSTRAINT attempts_quiz_id_fkey FOREIGN KEY (quiz_id) REFERENCES public.quizzes(id);


--
-- Name: attempts attempts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.attempts
    ADD CONSTRAINT attempts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: bots bots_owner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.bots
    ADD CONSTRAINT bots_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES public.users(id);


--
-- Name: deliveries deliveries_message_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_message_id_fkey FOREIGN KEY (message_id) REFERENCES public.messages(id);


--
-- Name: deliveries deliveries_recipient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_recipient_id_fkey FOREIGN KEY (recipient_id) REFERENCES public.users(id);


--
-- Name: deliveries deliveries_room_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(id);


--
-- Name: deliveries deliveries_sender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES public.users(id);


--
-- Name: friends friends_a_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_a_id_fkey FOREIGN KEY (a_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: friends friends_b_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_b_id_fkey FOREIGN KEY (b_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: gradients gradients_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.gradients
    ADD CONSTRAINT gradients_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: images images_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: kids_codes kids_codes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_codes
    ADD CONSTRAINT kids_codes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: kids_parents kids_parents_kid_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_parents
    ADD CONSTRAINT kids_parents_kid_id_fkey FOREIGN KEY (kid_id) REFERENCES public.users(id);


--
-- Name: kids_parents kids_parents_parent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_parents
    ADD CONSTRAINT kids_parents_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.users(id);


--
-- Name: messages messages_room_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(id);


--
-- Name: messages messages_sender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES public.users(id);


--
-- Name: notes notes_owner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES public.users(id);


--
-- Name: postcards postcards_recipient_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.postcards
    ADD CONSTRAINT postcards_recipient_fkey FOREIGN KEY (recipient) REFERENCES public.users(id);


--
-- Name: postcards postcards_sender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.postcards
    ADD CONSTRAINT postcards_sender_fkey FOREIGN KEY (sender) REFERENCES public.users(id);


--
-- Name: questions questions_quiz_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_quiz_id_fkey FOREIGN KEY (quiz_id) REFERENCES public.quizzes(id);


--
-- Name: responses responses_attempt_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT responses_attempt_id_fkey FOREIGN KEY (attempt_id) REFERENCES public.attempts(id);


--
-- Name: responses responses_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT responses_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.questions(id);


--
-- Name: responses responses_quiz_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT responses_quiz_id_fkey FOREIGN KEY (quiz_id) REFERENCES public.quizzes(id);


--
-- Name: responses responses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT responses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: room_users room_users_room_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.room_users
    ADD CONSTRAINT room_users_room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(id);


--
-- Name: room_users room_users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.room_users
    ADD CONSTRAINT room_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: threads threads_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.threads
    ADD CONSTRAINT threads_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users users_become_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_become_user_id_fkey FOREIGN KEY (become_user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

