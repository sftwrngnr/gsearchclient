--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.3 (Ubuntu 17.3-3.pgdg24.04+1)

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

ALTER TABLE ONLY public.zipcodes DROP CONSTRAINT zipcode_states_fk;
ALTER TABLE ONLY public.zipcodes DROP CONSTRAINT zipcode_cities_fk;
ALTER TABLE ONLY public.urls DROP CONSTRAINT urls_queries_fk;
ALTER TABLE ONLY public.searchcampaign DROP CONSTRAINT searchcampaign_company_fk;
ALTER TABLE ONLY public.queries DROP CONSTRAINT query_states_fk;
ALTER TABLE ONLY public.query_results DROP CONSTRAINT query_results_query_fk;
ALTER TABLE ONLY public.queries DROP CONSTRAINT queries_searchcampaign_fk;
ALTER TABLE ONLY public.search_metadata DROP CONSTRAINT qrysummary_queries_fk;
ALTER TABLE ONLY public.qry_zips DROP CONSTRAINT qry_zips_zipcodes_fk;
ALTER TABLE ONLY public.qry_zips DROP CONSTRAINT qry_zips_query_fk;
ALTER TABLE ONLY public.qry_kwds DROP CONSTRAINT qry_kwds_query_fk;
ALTER TABLE ONLY public.qry_kwds DROP CONSTRAINT qry_kwds_keywords_fk;
ALTER TABLE ONLY public.qry_acs DROP CONSTRAINT qry_ac_query_fk;
ALTER TABLE ONLY public.phonenumber DROP CONSTRAINT phonenumber_areacodes_fk;
ALTER TABLE ONLY public.pcontact DROP CONSTRAINT pcontact_phonenumber_fk;
ALTER TABLE ONLY public.paddress DROP CONSTRAINT paddress_zipcodes_fk;
ALTER TABLE ONLY public.paddress DROP CONSTRAINT paddress_states_fk;
ALTER TABLE ONLY public.paddress DROP CONSTRAINT paddress_queries_fk;
ALTER TABLE ONLY public.paddress DROP CONSTRAINT paddress_phonenumber_fk;
ALTER TABLE ONLY public.crawlerresults DROP CONSTRAINT crawlerresults_searchcampaign_fk;
ALTER TABLE ONLY public.crawlerresults DROP CONSTRAINT crawlerresults_crawlerprofile_fk;
ALTER TABLE ONLY public.crawlerprofile DROP CONSTRAINT crawlercampaign_searchcampaign_fk;
ALTER TABLE ONLY public.crawlerprofile DROP CONSTRAINT crawlercampaign_company_fk;
ALTER TABLE ONLY public.cityareacodes DROP CONSTRAINT cityareacodes_cities_fk;
ALTER TABLE ONLY public.cities DROP CONSTRAINT cities_states_fk;
ALTER TABLE ONLY public.areacodes DROP CONSTRAINT areacodes_states_fk;
DROP INDEX public.zipcode_state_idx;
DROP INDEX public.urls_id_idx;
DROP INDEX public.states_name_idx;
DROP INDEX public.crawlercampaign_company_idx;
DROP INDEX public.cityareacodes_city_idx;
DROP INDEX public.cities_name_idx;
DROP INDEX public.areacodes_code_idx;
ALTER TABLE ONLY public.zipcodes DROP CONSTRAINT zipcode_pk;
ALTER TABLE ONLY public.urls DROP CONSTRAINT urls_pk;
ALTER TABLE ONLY public.states DROP CONSTRAINT states_unique;
ALTER TABLE ONLY public.states DROP CONSTRAINT states_pk;
ALTER TABLE ONLY public.searchcampaign DROP CONSTRAINT searchcampaign_pk;
ALTER TABLE ONLY public.queries DROP CONSTRAINT query_pk;
ALTER TABLE ONLY public.search_metadata DROP CONSTRAINT qrysummary_pk;
ALTER TABLE ONLY public.qry_zips DROP CONSTRAINT qry_zips_pk;
ALTER TABLE ONLY public.qry_kwds DROP CONSTRAINT qry_kwds_pk;
ALTER TABLE ONLY public.qry_acs DROP CONSTRAINT qry_ac_pk;
ALTER TABLE ONLY public.phonenumber DROP CONSTRAINT phonenumber_unique;
ALTER TABLE ONLY public.keywords DROP CONSTRAINT keywords_pk;
ALTER TABLE ONLY public.crawlerresults DROP CONSTRAINT crawlerresults_pk;
ALTER TABLE ONLY public.crawlerprofile DROP CONSTRAINT crawlerconfig_pk;
ALTER TABLE ONLY public.query_results DROP CONSTRAINT crawler_results_pk;
ALTER TABLE ONLY public.companies DROP CONSTRAINT company_pk;
ALTER TABLE ONLY public.cityareacodes DROP CONSTRAINT cityareacodes_pk;
ALTER TABLE ONLY public.cities DROP CONSTRAINT cities_pk;
ALTER TABLE ONLY public.areacodes DROP CONSTRAINT areacodes_pk;
ALTER TABLE public.zipcodes ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.urls ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.states ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.searchcampaign ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.search_metadata ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.query_results ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.queries ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.qry_zips ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.qry_kwds ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.qry_acs ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.phonenumber ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.pcontact ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.paddress ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.keywords ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.crawlerresults ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.crawlerprofile ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.companies ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.cityareacodes ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.cities ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.areacodes ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE public.zipcode_id_seq;
DROP TABLE public.zipcodes;
DROP SEQUENCE public.urls_id_seq;
DROP TABLE public.urls;
DROP SEQUENCE public.states_id_seq;
DROP TABLE public.states;
DROP SEQUENCE public.searchcampaign_id_seq;
DROP TABLE public.searchcampaign;
DROP SEQUENCE public.query_id_seq;
DROP TABLE public.queries;
DROP SEQUENCE public.qrysummary_id_seq;
DROP TABLE public.search_metadata;
DROP SEQUENCE public.qry_zips_id_seq;
DROP TABLE public.qry_zips;
DROP SEQUENCE public.qry_kwds_id_seq;
DROP TABLE public.qry_kwds;
DROP SEQUENCE public.qry_ac_id_seq;
DROP TABLE public.qry_acs;
DROP SEQUENCE public.phonenumber_id_seq;
DROP TABLE public.phonenumber;
DROP SEQUENCE public.pcontact_id_seq;
DROP TABLE public.pcontact;
DROP SEQUENCE public.paddress_id_seq;
DROP TABLE public.paddress;
DROP SEQUENCE public.keywords_id_seq;
DROP TABLE public.keywords;
DROP SEQUENCE public.crawlerresults_id_seq;
DROP TABLE public.crawlerresults;
DROP SEQUENCE public.crawlerconfig_id_seq;
DROP TABLE public.crawlerprofile;
DROP SEQUENCE public.crawler_results_id_seq;
DROP TABLE public.query_results;
DROP SEQUENCE public.company_id_seq;
DROP TABLE public.companies;
DROP SEQUENCE public.cityareacodes_id_seq;
DROP TABLE public.cityareacodes;
DROP SEQUENCE public.cities_id_seq;
DROP TABLE public.cities;
DROP SEQUENCE public.areacodes_id_seq;
DROP TABLE public.areacodes;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: areacodes; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.areacodes (
    id bigint NOT NULL,
    code character varying,
    state bigint,
    latitude double precision,
    longitude double precision,
    created_at date,
    updated_at date,
    deleted_at date
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
-- Name: companies; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.companies (
    id bigint NOT NULL,
    name character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.companies OWNER TO crawler;

--
-- Name: company_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.company_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.company_id_seq OWNER TO crawler;

--
-- Name: company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.company_id_seq OWNED BY public.companies.id;


--
-- Name: query_results; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.query_results (
    id bigint NOT NULL,
    query_id bigint,
    resultseq bigint,
    result_type bigint,
    result jsonb,
    created_at date,
    updated_at date,
    deleted_at date
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
-- Name: crawlerprofile; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.crawlerprofile (
    id bigint NOT NULL,
    maxdepth bigint DEFAULT 5 NOT NULL,
    maxtime bigint DEFAULT 60 NOT NULL,
    maxpages bigint DEFAULT 10 NOT NULL,
    maxthreads bigint DEFAULT 5 NOT NULL,
    storepages boolean DEFAULT true NOT NULL,
    buildseolist boolean DEFAULT false NOT NULL,
    extractaddress boolean DEFAULT true NOT NULL,
    extractphone boolean DEFAULT true NOT NULL,
    specialextract boolean DEFAULT true NOT NULL,
    extractfunc character varying,
    name character varying DEFAULT 'test'::character varying NOT NULL,
    dailymaxcrawl bigint DEFAULT 100 NOT NULL,
    extractexternallinks boolean DEFAULT false NOT NULL,
    extractwordcloud boolean DEFAULT false,
    company bigint NOT NULL,
    searchcampaign bigint NOT NULL,
    multicrawl boolean DEFAULT false NOT NULL,
    agenttype bigint,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.crawlerprofile OWNER TO crawler;

--
-- Name: crawlerconfig_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.crawlerconfig_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.crawlerconfig_id_seq OWNER TO crawler;

--
-- Name: crawlerconfig_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.crawlerconfig_id_seq OWNED BY public.crawlerprofile.id;


--
-- Name: crawlerresults; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.crawlerresults (
    id bigint NOT NULL,
    queryid bigint,
    url character varying,
    pagescrawled bigint,
    crawldepth bigint,
    totalduration real,
    alloweddomains jsonb,
    success boolean,
    crawldate date,
    pagecrawlsucc real,
    profile bigint,
    urlimportdate date,
    status bigint,
    created_at date,
    updated_at date,
    deleted_at date,
    crawled boolean DEFAULT false,
    crawler bigint NOT NULL,
    campaign bigint NOT NULL
);


ALTER TABLE public.crawlerresults OWNER TO crawler;

--
-- Name: crawlerresults_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.crawlerresults_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.crawlerresults_id_seq OWNER TO crawler;

--
-- Name: crawlerresults_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.crawlerresults_id_seq OWNED BY public.crawlerresults.id;


--
-- Name: keywords; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.keywords (
    id bigint NOT NULL,
    keyword character varying,
    req boolean DEFAULT false NOT NULL,
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
-- Name: paddress; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.paddress (
    id bigint NOT NULL,
    query_id bigint NOT NULL,
    name character varying,
    address1 character varying,
    address2 character varying,
    city character varying,
    state bigint NOT NULL,
    zip bigint NOT NULL,
    zipstr character varying,
    phoneid bigint,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.paddress OWNER TO crawler;

--
-- Name: paddress_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.paddress_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.paddress_id_seq OWNER TO crawler;

--
-- Name: paddress_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.paddress_id_seq OWNED BY public.paddress.id;


--
-- Name: pcontact; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.pcontact (
    id bigint NOT NULL,
    paddy bigint,
    prefix character varying,
    firstname character varying,
    middle character varying,
    lastname character varying,
    suffix character varying,
    phoneid bigint NOT NULL,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.pcontact OWNER TO crawler;

--
-- Name: pcontact_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.pcontact_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pcontact_id_seq OWNER TO crawler;

--
-- Name: pcontact_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.pcontact_id_seq OWNED BY public.pcontact.id;


--
-- Name: phonenumber; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.phonenumber (
    id bigint NOT NULL,
    area bigint NOT NULL,
    number character varying,
    extension character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.phonenumber OWNER TO crawler;

--
-- Name: phonenumber_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.phonenumber_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.phonenumber_id_seq OWNER TO crawler;

--
-- Name: phonenumber_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.phonenumber_id_seq OWNED BY public.phonenumber.id;


--
-- Name: qry_acs; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.qry_acs (
    id bigint NOT NULL,
    query_id bigint,
    qry_ac character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.qry_acs OWNER TO crawler;

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

ALTER SEQUENCE public.qry_ac_id_seq OWNED BY public.qry_acs.id;


--
-- Name: qry_kwds; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.qry_kwds (
    id bigint NOT NULL,
    query_id bigint,
    keyword_id bigint,
    created_at date,
    updated_at date,
    deleted_at date
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
    zip_id bigint,
    created_at date,
    updated_at date,
    deleted_at date
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
-- Name: search_metadata; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.search_metadata (
    id bigint NOT NULL,
    query_id bigint,
    status character varying,
    searchid character varying,
    total_time_taken real,
    screated_at date,
    google_url character varying,
    json_endpoint character varying,
    processed_at date,
    raw_html_file character varying,
    created_at date,
    updated_at date,
    deleted_at date
);


ALTER TABLE public.search_metadata OWNER TO crawler;

--
-- Name: qrysummary_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.qrysummary_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.qrysummary_id_seq OWNER TO crawler;

--
-- Name: qrysummary_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.qrysummary_id_seq OWNED BY public.search_metadata.id;


--
-- Name: queries; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.queries (
    id bigint NOT NULL,
    state bigint,
    query_string character varying,
    created_at date,
    updated_at date,
    deleted_at character varying,
    campaign bigint
);


ALTER TABLE public.queries OWNER TO crawler;

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

ALTER SEQUENCE public.query_id_seq OWNED BY public.queries.id;


--
-- Name: searchcampaign; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.searchcampaign (
    id bigint NOT NULL,
    company bigint NOT NULL,
    name character varying DEFAULT 'Default'::character varying NOT NULL
);


ALTER TABLE public.searchcampaign OWNER TO crawler;

--
-- Name: searchcampaign_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.searchcampaign_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.searchcampaign_id_seq OWNER TO crawler;

--
-- Name: searchcampaign_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.searchcampaign_id_seq OWNED BY public.searchcampaign.id;


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
-- Name: urls; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.urls (
    id bigint NOT NULL,
    query_id bigint,
    query_src bigint,
    seq_id bigint,
    url character varying,
    crawldate date,
    crawlsuccess boolean DEFAULT false,
    importdate date DEFAULT now(),
    "position" bigint,
    source character varying,
    created_at date,
    updated_at date,
    deleted_at date,
    transferred boolean DEFAULT false NOT NULL
);


ALTER TABLE public.urls OWNER TO crawler;

--
-- Name: urls_id_seq; Type: SEQUENCE; Schema: public; Owner: crawler
--

CREATE SEQUENCE public.urls_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.urls_id_seq OWNER TO crawler;

--
-- Name: urls_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: crawler
--

ALTER SEQUENCE public.urls_id_seq OWNED BY public.urls.id;


--
-- Name: zipcodes; Type: TABLE; Schema: public; Owner: crawler
--

CREATE TABLE public.zipcodes (
    id bigint NOT NULL,
    zipcode character varying,
    city bigint,
    state bigint,
    population bigint,
    latitude double precision,
    longitude double precision,
    created_at date,
    updated_at date,
    deleted_at date
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
-- Name: companies id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);


--
-- Name: crawlerprofile id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerprofile ALTER COLUMN id SET DEFAULT nextval('public.crawlerconfig_id_seq'::regclass);


--
-- Name: crawlerresults id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerresults ALTER COLUMN id SET DEFAULT nextval('public.crawlerresults_id_seq'::regclass);


--
-- Name: keywords id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.keywords ALTER COLUMN id SET DEFAULT nextval('public.keywords_id_seq'::regclass);


--
-- Name: paddress id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.paddress ALTER COLUMN id SET DEFAULT nextval('public.paddress_id_seq'::regclass);


--
-- Name: pcontact id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.pcontact ALTER COLUMN id SET DEFAULT nextval('public.pcontact_id_seq'::regclass);


--
-- Name: phonenumber id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.phonenumber ALTER COLUMN id SET DEFAULT nextval('public.phonenumber_id_seq'::regclass);


--
-- Name: qry_acs id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_acs ALTER COLUMN id SET DEFAULT nextval('public.qry_ac_id_seq'::regclass);


--
-- Name: qry_kwds id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds ALTER COLUMN id SET DEFAULT nextval('public.qry_kwds_id_seq'::regclass);


--
-- Name: qry_zips id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips ALTER COLUMN id SET DEFAULT nextval('public.qry_zips_id_seq'::regclass);


--
-- Name: queries id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.queries ALTER COLUMN id SET DEFAULT nextval('public.query_id_seq'::regclass);


--
-- Name: query_results id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results ALTER COLUMN id SET DEFAULT nextval('public.crawler_results_id_seq'::regclass);


--
-- Name: search_metadata id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.search_metadata ALTER COLUMN id SET DEFAULT nextval('public.qrysummary_id_seq'::regclass);


--
-- Name: searchcampaign id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.searchcampaign ALTER COLUMN id SET DEFAULT nextval('public.searchcampaign_id_seq'::regclass);


--
-- Name: states id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.states ALTER COLUMN id SET DEFAULT nextval('public.states_id_seq'::regclass);


--
-- Name: urls id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.urls ALTER COLUMN id SET DEFAULT nextval('public.urls_id_seq'::regclass);


--
-- Name: zipcodes id; Type: DEFAULT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes ALTER COLUMN id SET DEFAULT nextval('public.zipcode_id_seq'::regclass);


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
-- Name: companies company_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT company_pk PRIMARY KEY (id);


--
-- Name: query_results crawler_results_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT crawler_results_pk PRIMARY KEY (id);


--
-- Name: crawlerprofile crawlerconfig_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerprofile
    ADD CONSTRAINT crawlerconfig_pk PRIMARY KEY (id);


--
-- Name: crawlerresults crawlerresults_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerresults
    ADD CONSTRAINT crawlerresults_pk PRIMARY KEY (id);


--
-- Name: keywords keywords_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.keywords
    ADD CONSTRAINT keywords_pk PRIMARY KEY (id);


--
-- Name: phonenumber phonenumber_unique; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.phonenumber
    ADD CONSTRAINT phonenumber_unique UNIQUE (id);


--
-- Name: qry_acs qry_ac_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_acs
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
-- Name: search_metadata qrysummary_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.search_metadata
    ADD CONSTRAINT qrysummary_pk PRIMARY KEY (id);


--
-- Name: queries query_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT query_pk PRIMARY KEY (id);


--
-- Name: searchcampaign searchcampaign_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.searchcampaign
    ADD CONSTRAINT searchcampaign_pk PRIMARY KEY (id);


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
-- Name: urls urls_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_pk PRIMARY KEY (id);


--
-- Name: zipcodes zipcode_pk; Type: CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.zipcodes
    ADD CONSTRAINT zipcode_pk PRIMARY KEY (id);


--
-- Name: areacodes_code_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX areacodes_code_idx ON public.areacodes USING btree (code);


--
-- Name: cities_name_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX cities_name_idx ON public.cities USING btree (name);


--
-- Name: cityareacodes_city_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX cityareacodes_city_idx ON public.cityareacodes USING btree (city, created_at, updated_at, deleted_at);


--
-- Name: crawlercampaign_company_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX crawlercampaign_company_idx ON public.crawlerprofile USING btree (company);


--
-- Name: states_name_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX states_name_idx ON public.states USING btree (name);


--
-- Name: urls_id_idx; Type: INDEX; Schema: public; Owner: crawler
--

CREATE INDEX urls_id_idx ON public.urls USING btree (id, query_id, query_src, seq_id);


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
-- Name: crawlerprofile crawlercampaign_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerprofile
    ADD CONSTRAINT crawlercampaign_company_fk FOREIGN KEY (company) REFERENCES public.companies(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: crawlerprofile crawlercampaign_searchcampaign_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerprofile
    ADD CONSTRAINT crawlercampaign_searchcampaign_fk FOREIGN KEY (searchcampaign) REFERENCES public.searchcampaign(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: crawlerresults crawlerresults_crawlerprofile_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerresults
    ADD CONSTRAINT crawlerresults_crawlerprofile_fk FOREIGN KEY (crawler) REFERENCES public.crawlerprofile(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: crawlerresults crawlerresults_searchcampaign_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.crawlerresults
    ADD CONSTRAINT crawlerresults_searchcampaign_fk FOREIGN KEY (campaign) REFERENCES public.searchcampaign(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: paddress paddress_phonenumber_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.paddress
    ADD CONSTRAINT paddress_phonenumber_fk FOREIGN KEY (phoneid) REFERENCES public.phonenumber(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: paddress paddress_queries_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.paddress
    ADD CONSTRAINT paddress_queries_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: paddress paddress_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.paddress
    ADD CONSTRAINT paddress_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: paddress paddress_zipcodes_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.paddress
    ADD CONSTRAINT paddress_zipcodes_fk FOREIGN KEY (zip) REFERENCES public.zipcodes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: pcontact pcontact_phonenumber_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.pcontact
    ADD CONSTRAINT pcontact_phonenumber_fk FOREIGN KEY (phoneid) REFERENCES public.phonenumber(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: phonenumber phonenumber_areacodes_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.phonenumber
    ADD CONSTRAINT phonenumber_areacodes_fk FOREIGN KEY (area) REFERENCES public.areacodes(id);


--
-- Name: qry_acs qry_ac_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_acs
    ADD CONSTRAINT qry_ac_query_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_kwds qry_kwds_keywords_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds
    ADD CONSTRAINT qry_kwds_keywords_fk FOREIGN KEY (keyword_id) REFERENCES public.keywords(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_kwds qry_kwds_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_kwds
    ADD CONSTRAINT qry_kwds_query_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_zips qry_zips_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips
    ADD CONSTRAINT qry_zips_query_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: qry_zips qry_zips_zipcodes_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.qry_zips
    ADD CONSTRAINT qry_zips_zipcodes_fk FOREIGN KEY (zip_id) REFERENCES public.zipcodes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: search_metadata qrysummary_queries_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.search_metadata
    ADD CONSTRAINT qrysummary_queries_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: queries queries_searchcampaign_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_searchcampaign_fk FOREIGN KEY (campaign) REFERENCES public.searchcampaign(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: query_results query_results_query_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT query_results_query_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: queries query_states_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT query_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: searchcampaign searchcampaign_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.searchcampaign
    ADD CONSTRAINT searchcampaign_company_fk FOREIGN KEY (company) REFERENCES public.companies(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: urls urls_queries_fk; Type: FK CONSTRAINT; Schema: public; Owner: crawler
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_queries_fk FOREIGN KEY (query_id) REFERENCES public.queries(id) ON UPDATE CASCADE ON DELETE CASCADE;


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

