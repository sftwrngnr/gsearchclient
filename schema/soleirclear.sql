--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.3 (Ubuntu 17.3-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: areacodes; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.areacodes (
    id bigint NOT NULL,
    code character varying,
    state bigint,
    created_at date,
    updated_at date,
    deleted_at date,
    latitude double precision,
    longitude double precision
);


ALTER TABLE public.areacodes OWNER TO crawler;

--
-- Name: areacodes_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.areacodes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.areacodes_id_seq OWNER TO crawler;

--
-- Name: areacodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.areacodes_id_seq OWNED BY public.areacodes.id;


--
-- Name: cities; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    name character varying NOT NULL,
    state bigint NOT NULL,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.cities OWNER TO crawler;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cities_id_seq OWNER TO crawler;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: cityareacodes; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.cityareacodes (
    id bigint NOT NULL,
    areacode bigint NOT NULL,
    city bigint NOT NULL,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.cityareacodes OWNER TO crawler;

--
-- Name: cityareacodes_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.cityareacodes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cityareacodes_id_seq OWNER TO crawler;

--
-- Name: cityareacodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.cityareacodes_id_seq OWNED BY public.cityareacodes.id;


--
-- Name: query_results; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.query_results (
    id bigint NOT NULL,
    query_id bigint,
    resultseq bigint,
    result_type bigint,
    result jsonb
);


ALTER TABLE public.query_results OWNER TO crawler;

--
-- Name: crawler_results_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.crawler_results_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.crawler_results_id_seq OWNER TO crawler;

--
-- Name: crawler_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.crawler_results_id_seq OWNED BY public.query_results.id;


--
-- Name: keywords; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.keywords (
    id bigint NOT NULL,
    keyword character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.keywords OWNER TO crawler;

--
-- Name: keywords_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.keywords_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.keywords_id_seq OWNER TO crawler;

--
-- Name: keywords_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.keywords_id_seq OWNED BY public.keywords.id;


--
-- Name: qry_ac; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.qry_ac (
    id bigint NOT NULL,
    qry_id bigint,
    qry_ac bigint
);


ALTER TABLE public.qry_ac OWNER TO crawler;

--
-- Name: qry_ac_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.qry_ac_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.qry_ac_id_seq OWNER TO crawler;

--
-- Name: qry_ac_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.qry_ac_id_seq OWNED BY public.qry_ac.id;


--
-- Name: qry_kwds; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.qry_kwds (
    id bigint NOT NULL,
    query_id bigint,
    keyword_id bigint
);


ALTER TABLE public.qry_kwds OWNER TO crawler;

--
-- Name: qry_kwds_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.qry_kwds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.qry_kwds_id_seq OWNER TO crawler;

--
-- Name: qry_kwds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.qry_kwds_id_seq OWNED BY public.qry_kwds.id;


--
-- Name: qry_zips; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.qry_zips (
    id bigint NOT NULL,
    query_id bigint,
    zip_id bigint
);


ALTER TABLE public.qry_zips OWNER TO crawler;

--
-- Name: qry_zips_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.qry_zips_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.qry_zips_id_seq OWNER TO crawler;

--
-- Name: qry_zips_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.qry_zips_id_seq OWNED BY public.qry_zips.id;


--
-- Name: query; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.query (
    id bigint NOT NULL,
    state bigint,
    kwds bigint,
    zips bigint,
    acs bigint
);


ALTER TABLE public.query OWNER TO crawler;

--
-- Name: query_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.query_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.query_id_seq OWNER TO crawler;

--
-- Name: query_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.query_id_seq OWNED BY public.query.id;


--
-- Name: states; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.states (
    id bigint NOT NULL,
    abbrev character varying NOT NULL,
    name character varying,
    capitol character varying,
    region character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.states OWNER TO crawler;

--
-- Name: states_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.states_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.states_id_seq OWNER TO crawler;

--
-- Name: states_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.states_id_seq OWNED BY public.states.id;


--
-- Name: zipcodes; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.zipcodes (
    id bigint NOT NULL,
    zipcode character varying,
    city bigint,
    state bigint,
    population bigint,
    created_at date,
    updated_at date,
    deleted_at date,
    latitude double precision,
    longitude double precision
);


ALTER TABLE public.zipcodes OWNER TO crawler;

--
-- Name: zipcode_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.zipcode_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.zipcode_id_seq OWNER TO crawler;

--
-- Name: zipcode_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.zipcode_id_seq OWNED BY public.zipcodes.id;


--
-- Name: areacodes id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.areacodes ALTER COLUMN id SET DEFAULT nextval('public.areacodes_id_seq'::regclass);


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: cityareacodes id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cityareacodes ALTER COLUMN id SET DEFAULT nextval('public.cityareacodes_id_seq'::regclass);


--
-- Name: keywords id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.keywords ALTER COLUMN id SET DEFAULT nextval('public.keywords_id_seq'::regclass);


--
-- Name: qry_ac id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_ac ALTER COLUMN id SET DEFAULT nextval('public.qry_ac_id_seq'::regclass);


--
-- Name: qry_kwds id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds ALTER COLUMN id SET DEFAULT nextval('public.qry_kwds_id_seq'::regclass);


--
-- Name: qry_zips id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips ALTER COLUMN id SET DEFAULT nextval('public.qry_zips_id_seq'::regclass);


--
-- Name: query id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query ALTER COLUMN id SET DEFAULT nextval('public.query_id_seq'::regclass);


--
-- Name: query_results id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results ALTER COLUMN id SET DEFAULT nextval('public.crawler_results_id_seq'::regclass);


--
-- Name: states id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.states ALTER COLUMN id SET DEFAULT nextval('public.states_id_seq'::regclass);


--
-- Name: zipcodes id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes ALTER COLUMN id SET DEFAULT nextval('public.zipcode_id_seq'::regclass);


--
-- Data for Name: areacodes; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.areacodes (id, code, state, created_at, updated_at, deleted_at, latitude, longitude) FROM stdin;
\.


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.cities (id, name, state, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: cityareacodes; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.cityareacodes (id, areacode, city, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: keywords; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.keywords (id, keyword, created_at, updated_at, deleted_at) FROM stdin;
1	Clear Aligner	2025-02-10	2025-02-10	\N
2	Clear Aligners	2025-02-10	2025-02-10	\N
3	Teeth Straightening	2025-02-10	2025-02-10	\N
4	Orthodontist	2025-02-10	2025-02-10	\N
5	Dentist	2025-02-10	2025-02-10	\N
6	Dental clinic	2025-02-10	2025-02-10	\N
7	Dentist near me	2025-02-10	2025-02-10	\N
8	Braces	2025-02-10	2025-02-10	\N
9	Invisalign	2025-02-10	2025-02-10	\N
10	Aligners	2025-02-10	2025-02-10	\N
11	Teeth	2025-02-10	2025-02-10	\N
12	straight Teeth	2025-02-10	2025-02-10	\N
13	fix my teeth	2025-02-10	2025-02-10	\N
14	smile	2025-02-10	2025-02-10	\N
15	beautiful smile	2025-02-10	2025-02-10	\N
16	pretty teeth	2025-02-10	2025-02-10	\N
17	beautiful teeth	2025-02-10	2025-02-10	\N
18	Clear Aligner	2025-02-11	2025-02-11	\N
19	Clear Aligners	2025-02-11	2025-02-11	\N
20	Teeth Straightening	2025-02-11	2025-02-11	\N
21	Orthodontist	2025-02-11	2025-02-11	\N
22	Dentist	2025-02-11	2025-02-11	\N
23	Dental clinic	2025-02-11	2025-02-11	\N
24	Dentist near me	2025-02-11	2025-02-11	\N
25	Braces	2025-02-11	2025-02-11	\N
26	Invisalign	2025-02-11	2025-02-11	\N
27	Aligners	2025-02-11	2025-02-11	\N
28	Teeth	2025-02-11	2025-02-11	\N
29	straight Teeth	2025-02-11	2025-02-11	\N
30	fix my teeth	2025-02-11	2025-02-11	\N
31	smile	2025-02-11	2025-02-11	\N
32	beautiful smile	2025-02-11	2025-02-11	\N
33	pretty teeth	2025-02-11	2025-02-11	\N
34	beautiful teeth	2025-02-11	2025-02-11	\N
35	Dentist in {zipcode}	2025-02-11	2025-02-11	\N
36	Dentist in {state}	2025-02-11	2025-02-11	\N
37	Dentist in {area code}	2025-02-11	2025-02-11	\N
\.


--
-- Data for Name: qry_ac; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.qry_ac (id, qry_id, qry_ac) FROM stdin;
\.


--
-- Data for Name: qry_kwds; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.qry_kwds (id, query_id, keyword_id) FROM stdin;
\.


--
-- Data for Name: qry_zips; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.qry_zips (id, query_id, zip_id) FROM stdin;
\.


--
-- Data for Name: query; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.query (id, state, kwds, zips, acs) FROM stdin;
\.


--
-- Data for Name: query_results; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.query_results (id, query_id, resultseq, result_type, result) FROM stdin;
\.


--
-- Data for Name: states; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.states (id, abbrev, name, capitol, region, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: zipcodes; Type: TABLE DATA; Schema: public; Owner: crawler
--

COPY public.zipcodes (id, zipcode, city, state, population, created_at, updated_at, deleted_at, latitude, longitude) FROM stdin;
\.


--
-- Name: areacodes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.areacodes_id_seq', 1, false);


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.cities_id_seq', 1, false);


--
-- Name: cityareacodes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.cityareacodes_id_seq', 1, false);


--
-- Name: crawler_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.crawler_results_id_seq', 1, false);


--
-- Name: keywords_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.keywords_id_seq', 37, true);


--
-- Name: qry_ac_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.qry_ac_id_seq', 1, false);


--
-- Name: qry_kwds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.qry_kwds_id_seq', 1, false);


--
-- Name: qry_zips_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.qry_zips_id_seq', 1, false);


--
-- Name: query_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.query_id_seq', 1, false);


--
-- Name: states_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.states_id_seq', 1, false);


--
-- Name: zipcode_id_seq; Type: SEQUENCE SET; Schema: public; Owner: crawler
--

SELECT pg_catalog.setval('public.zipcode_id_seq', 1, false);


--
-- Name: areacodes areacodes_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.areacodes
    ADD CONSTRAINT areacodes_pk PRIMARY KEY (id);


--
-- Name: cities cities_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pk PRIMARY KEY (id);


--
-- Name: cityareacodes cityareacodes_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cityareacodes
    ADD CONSTRAINT cityareacodes_pk PRIMARY KEY (id);


--
-- Name: query_results crawler_results_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT crawler_results_pk PRIMARY KEY (id);


--
-- Name: keywords keywords_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.keywords
    ADD CONSTRAINT keywords_pk PRIMARY KEY (id);


--
-- Name: qry_ac qry_ac_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_ac
    ADD CONSTRAINT qry_ac_pk PRIMARY KEY (id);


--
-- Name: qry_kwds qry_kwds_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds
    ADD CONSTRAINT qry_kwds_pk PRIMARY KEY (id);


--
-- Name: qry_zips qry_zips_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips
    ADD CONSTRAINT qry_zips_pk PRIMARY KEY (id);


--
-- Name: query query_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_pk PRIMARY KEY (id);


--
-- Name: states states_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.states
    ADD CONSTRAINT states_pk PRIMARY KEY (id);


--
-- Name: states states_unique; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.states
    ADD CONSTRAINT states_unique UNIQUE (abbrev);


--
-- Name: zipcodes zipcode_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes
    ADD CONSTRAINT zipcode_pk PRIMARY KEY (id);


--
-- Name: cities_name_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX cities_name_idx ON public.cities USING btree (name);


--
-- Name: cityareacodes_city_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX cityareacodes_city_idx ON public.cityareacodes USING btree (city, created_at, updated_at, deleted_at);


--
-- Name: states_name_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX states_name_idx ON public.states USING btree (name);


--
-- Name: zipcode_state_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX zipcode_state_idx ON public.zipcodes USING btree (state, population);


--
-- Name: areacodes areacodes_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.areacodes
    ADD CONSTRAINT areacodes_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: cities cities_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: cityareacodes cityareacodes_cities_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.cityareacodes
    ADD CONSTRAINT cityareacodes_cities_fk FOREIGN KEY (city) REFERENCES public.cities(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_ac qry_ac_areacodes_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_ac
    ADD CONSTRAINT qry_ac_areacodes_fk FOREIGN KEY (qry_ac) REFERENCES public.areacodes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_ac qry_ac_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_ac
    ADD CONSTRAINT qry_ac_query_fk FOREIGN KEY (qry_id) REFERENCES public.query(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_kwds qry_kwds_keywords_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds
    ADD CONSTRAINT qry_kwds_keywords_fk FOREIGN KEY (keyword_id) REFERENCES public.keywords(id);


--
-- Name: qry_kwds qry_kwds_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds
    ADD CONSTRAINT qry_kwds_query_fk FOREIGN KEY (query_id) REFERENCES public.query(id);


--
-- Name: qry_zips qry_zips_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips
    ADD CONSTRAINT qry_zips_query_fk FOREIGN KEY (query_id) REFERENCES public.query(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_zips qry_zips_zipcodes_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips
    ADD CONSTRAINT qry_zips_zipcodes_fk FOREIGN KEY (zip_id) REFERENCES public.zipcodes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query query_qry_ac_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_qry_ac_fk FOREIGN KEY (acs) REFERENCES public.qry_ac(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query query_qry_kwds_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_qry_kwds_fk FOREIGN KEY (kwds) REFERENCES public.qry_kwds(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query query_qry_zips_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_qry_zips_fk FOREIGN KEY (zips) REFERENCES public.qry_zips(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query_results query_results_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT query_results_query_fk FOREIGN KEY (query_id) REFERENCES public.query(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query query_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query
    ADD CONSTRAINT query_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: zipcodes zipcode_cities_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes
    ADD CONSTRAINT zipcode_cities_fk FOREIGN KEY (city) REFERENCES public.cities(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: zipcodes zipcode_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes
    ADD CONSTRAINT zipcode_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

