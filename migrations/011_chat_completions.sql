ALTER TABLE bots ADD COLUMN model text NOT NULL DEFAULT '';
ALTER TABLE bots DROP COLUMN assistant_id;

ALTER TABLE threads ADD COLUMN bot_id bigint REFERENCES bots(id);
ALTER TABLE threads DROP COLUMN assistant_id;

CREATE TABLE bot_messages (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    thread_id bigint REFERENCES threads(id),
    role text NOT NULL,
    content text NOT NULL
);

CREATE SEQUENCE bot_messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE bot_messages_id_seq OWNED BY bot_messages.id;
ALTER TABLE ONLY bot_messages ALTER COLUMN id SET DEFAULT nextval('bot_messages_id_seq'::regclass);
ALTER TABLE ONLY bot_messages ADD CONSTRAINT bot_messages_pkey PRIMARY KEY (id);

ALTER TABLE public.bot_messages OWNER TO appuser;
