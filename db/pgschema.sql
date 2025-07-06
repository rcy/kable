--
-- PostgreSQL database dump
--

-- Dumped from database version 15.13 (Debian 15.13-1.pgdg120+1)
-- Dumped by pg_dump version 15.13 (Debian 15.13-1.pgdg120+1)

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: attempts; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.attempts (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    quiz_id bigint NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.attempts OWNER TO appuser;

--
-- Name: attempts_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.attempts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.attempts_id_seq OWNER TO appuser;

--
-- Name: attempts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.attempts_id_seq OWNED BY public.attempts.id;


--
-- Name: bots; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.bots (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    owner_id bigint NOT NULL,
    assistant_id text NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    published boolean DEFAULT false NOT NULL
);


ALTER TABLE public.bots OWNER TO appuser;

--
-- Name: bots_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.bots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bots_id_seq OWNER TO appuser;

--
-- Name: bots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.bots_id_seq OWNED BY public.bots.id;


--
-- Name: codes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.codes (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    code text NOT NULL,
    nonce text NOT NULL,
    email text NOT NULL
);


ALTER TABLE public.codes OWNER TO appuser;

--
-- Name: codes_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.codes_id_seq OWNER TO appuser;

--
-- Name: codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.codes_id_seq OWNED BY public.codes.id;


--
-- Name: deliveries; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.deliveries (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    message_id bigint NOT NULL,
    room_id bigint NOT NULL,
    recipient_id bigint NOT NULL,
    sender_id bigint NOT NULL,
    sent_at timestamp with time zone
);


ALTER TABLE public.deliveries OWNER TO appuser;

--
-- Name: deliveries_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.deliveries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.deliveries_id_seq OWNER TO appuser;

--
-- Name: deliveries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.deliveries_id_seq OWNED BY public.deliveries.id;


--
-- Name: friends; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.friends (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    a_id bigint NOT NULL,
    b_id bigint NOT NULL,
    b_role text NOT NULL
);


ALTER TABLE public.friends OWNER TO appuser;

--
-- Name: friends_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.friends_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.friends_id_seq OWNER TO appuser;

--
-- Name: friends_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.friends_id_seq OWNED BY public.friends.id;


--
-- Name: gradients; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.gradients (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id bigint NOT NULL,
    gradient jsonb NOT NULL
);


ALTER TABLE public.gradients OWNER TO appuser;

--
-- Name: gradients_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.gradients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gradients_id_seq OWNER TO appuser;

--
-- Name: gradients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.gradients_id_seq OWNED BY public.gradients.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    url text NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.images OWNER TO appuser;

--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.images_id_seq OWNER TO appuser;

--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: kids_codes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.kids_codes (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    code text NOT NULL,
    nonce text NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.kids_codes OWNER TO appuser;

--
-- Name: kids_codes_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.kids_codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.kids_codes_id_seq OWNER TO appuser;

--
-- Name: kids_codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.kids_codes_id_seq OWNED BY public.kids_codes.id;


--
-- Name: kids_parents; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.kids_parents (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    kid_id bigint NOT NULL,
    parent_id bigint NOT NULL
);


ALTER TABLE public.kids_parents OWNER TO appuser;

--
-- Name: kids_parents_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.kids_parents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.kids_parents_id_seq OWNER TO appuser;

--
-- Name: kids_parents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.kids_parents_id_seq OWNED BY public.kids_parents.id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    sender_id bigint NOT NULL,
    room_id bigint NOT NULL,
    body text NOT NULL
);


ALTER TABLE public.messages OWNER TO appuser;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.messages_id_seq OWNER TO appuser;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: notes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.notes (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    owner_id bigint NOT NULL,
    body text NOT NULL,
    published boolean DEFAULT false NOT NULL
);


ALTER TABLE public.notes OWNER TO appuser;

--
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.notes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notes_id_seq OWNER TO appuser;

--
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- Name: postcards; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.postcards (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    sender bigint NOT NULL,
    recipient bigint NOT NULL,
    subject text NOT NULL,
    body text NOT NULL,
    state text DEFAULT 'draft'::text NOT NULL
);


ALTER TABLE public.postcards OWNER TO appuser;

--
-- Name: postcards_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.postcards_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.postcards_id_seq OWNER TO appuser;

--
-- Name: postcards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.postcards_id_seq OWNED BY public.postcards.id;


--
-- Name: questions; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.questions (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    quiz_id bigint NOT NULL,
    text text NOT NULL,
    answer text NOT NULL
);


ALTER TABLE public.questions OWNER TO appuser;

--
-- Name: questions_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.questions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questions_id_seq OWNER TO appuser;

--
-- Name: questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.questions_id_seq OWNED BY public.questions.id;


--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.quizzes (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    published boolean DEFAULT false NOT NULL
);


ALTER TABLE public.quizzes OWNER TO appuser;

--
-- Name: quizzes_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.quizzes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quizzes_id_seq OWNER TO appuser;

--
-- Name: quizzes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.quizzes_id_seq OWNED BY public.quizzes.id;


--
-- Name: responses; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.responses (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    quiz_id bigint NOT NULL,
    user_id bigint NOT NULL,
    attempt_id bigint NOT NULL,
    question_id bigint NOT NULL,
    text text NOT NULL
);


ALTER TABLE public.responses OWNER TO appuser;

--
-- Name: responses_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.responses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.responses_id_seq OWNER TO appuser;

--
-- Name: responses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.responses_id_seq OWNED BY public.responses.id;


--
-- Name: room_users; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.room_users (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    room_id bigint NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.room_users OWNER TO appuser;

--
-- Name: room_users_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.room_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.room_users_id_seq OWNER TO appuser;

--
-- Name: room_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.room_users_id_seq OWNED BY public.room_users.id;


--
-- Name: rooms; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.rooms (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    key text NOT NULL
);


ALTER TABLE public.rooms OWNER TO appuser;

--
-- Name: rooms_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.rooms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rooms_id_seq OWNER TO appuser;

--
-- Name: rooms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.rooms_id_seq OWNED BY public.rooms.id;


--
-- Name: schema_version; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.schema_version (
    version integer NOT NULL
);


ALTER TABLE public.schema_version OWNER TO appuser;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    key text NOT NULL
);


ALTER TABLE public.sessions OWNER TO appuser;

--
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sessions_id_seq OWNER TO appuser;

--
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


--
-- Name: threads; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.threads (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    thread_id text NOT NULL,
    assistant_id text NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.threads OWNER TO appuser;

--
-- Name: threads_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.threads_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.threads_id_seq OWNER TO appuser;

--
-- Name: threads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.threads_id_seq OWNED BY public.threads.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    username text NOT NULL,
    email text,
    avatar_url text DEFAULT 'https://www.gravatar.com/avatar/?d=mp'::text NOT NULL,
    is_parent boolean DEFAULT false NOT NULL,
    bio text DEFAULT ''::text NOT NULL,
    become_user_id bigint,
    admin boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users OWNER TO appuser;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: appuser
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO appuser;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: appuser
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: attempts id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.attempts ALTER COLUMN id SET DEFAULT nextval('public.attempts_id_seq'::regclass);


--
-- Name: bots id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.bots ALTER COLUMN id SET DEFAULT nextval('public.bots_id_seq'::regclass);


--
-- Name: codes id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.codes ALTER COLUMN id SET DEFAULT nextval('public.codes_id_seq'::regclass);


--
-- Name: deliveries id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries ALTER COLUMN id SET DEFAULT nextval('public.deliveries_id_seq'::regclass);


--
-- Name: friends id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends ALTER COLUMN id SET DEFAULT nextval('public.friends_id_seq'::regclass);


--
-- Name: gradients id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.gradients ALTER COLUMN id SET DEFAULT nextval('public.gradients_id_seq'::regclass);


--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Name: kids_codes id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_codes ALTER COLUMN id SET DEFAULT nextval('public.kids_codes_id_seq'::regclass);


--
-- Name: kids_parents id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_parents ALTER COLUMN id SET DEFAULT nextval('public.kids_parents_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- Name: postcards id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.postcards ALTER COLUMN id SET DEFAULT nextval('public.postcards_id_seq'::regclass);


--
-- Name: questions id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.questions ALTER COLUMN id SET DEFAULT nextval('public.questions_id_seq'::regclass);


--
-- Name: quizzes id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.quizzes ALTER COLUMN id SET DEFAULT nextval('public.quizzes_id_seq'::regclass);


--
-- Name: responses id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses ALTER COLUMN id SET DEFAULT nextval('public.responses_id_seq'::regclass);


--
-- Name: room_users id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.room_users ALTER COLUMN id SET DEFAULT nextval('public.room_users_id_seq'::regclass);


--
-- Name: rooms id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.rooms ALTER COLUMN id SET DEFAULT nextval('public.rooms_id_seq'::regclass);


--
-- Name: sessions id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


--
-- Name: threads id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.threads ALTER COLUMN id SET DEFAULT nextval('public.threads_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: attempts attempts_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.attempts
    ADD CONSTRAINT attempts_pkey PRIMARY KEY (id);


--
-- Name: bots bots_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.bots
    ADD CONSTRAINT bots_pkey PRIMARY KEY (id);


--
-- Name: codes codes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.codes
    ADD CONSTRAINT codes_pkey PRIMARY KEY (id);


--
-- Name: deliveries deliveries_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_pkey PRIMARY KEY (id);


--
-- Name: friends friends_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_pkey PRIMARY KEY (id);


--
-- Name: gradients gradients_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.gradients
    ADD CONSTRAINT gradients_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: kids_codes kids_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_codes
    ADD CONSTRAINT kids_codes_pkey PRIMARY KEY (id);


--
-- Name: kids_parents kids_parents_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.kids_parents
    ADD CONSTRAINT kids_parents_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- Name: postcards postcards_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.postcards
    ADD CONSTRAINT postcards_pkey PRIMARY KEY (id);


--
-- Name: questions questions_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_pkey PRIMARY KEY (id);


--
-- Name: quizzes quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_pkey PRIMARY KEY (id);


--
-- Name: responses responses_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.responses
    ADD CONSTRAINT responses_pkey PRIMARY KEY (id);


--
-- Name: room_users room_users_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.room_users
    ADD CONSTRAINT room_users_pkey PRIMARY KEY (id);


--
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: threads threads_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.threads
    ADD CONSTRAINT threads_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


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
    ADD CONSTRAINT friends_a_id_fkey FOREIGN KEY (a_id) REFERENCES public.users(id);


--
-- Name: friends friends_b_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_b_id_fkey FOREIGN KEY (b_id) REFERENCES public.users(id);


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

