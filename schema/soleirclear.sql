CREATE TABLE public.areacodes (
                                  id bigserial NOT NULL,
                                  code varchar NULL,
                                  state bigint NULL,
                                  latitude float8 NULL,
                                  longitude float8 NULL,
                                  created_at date NULL,
                                  updated_at date NULL,
                                  deleted_at date NULL,
                                  CONSTRAINT areacodes_pk PRIMARY KEY (id)
);

CREATE TABLE public.states (
                               id bigserial NOT NULL,
                               abbrev varchar NOT NULL,
                               "name" varchar NULL,
                               capitol varchar NULL,
                               region varchar NULL,
                               created_at date NULL,
                               updated_at date NULL,
                               deleted_at date NULL,
                               CONSTRAINT states_pk PRIMARY KEY (id),
                               CONSTRAINT states_unique UNIQUE (abbrev)
);
CREATE INDEX states_name_idx ON public.states ("name");

ALTER TABLE public.areacodes ADD CONSTRAINT areacodes_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON DELETE CASCADE ON UPDATE CASCADE;


CREATE TABLE public.cities (
                               id bigserial NOT NULL,
                               "name" varchar NOT NULL,
                               state bigint NOT NULL,
                               created_at date NULL,
                               updated_at date NULL,
                               deleted_at date NULL,
                               CONSTRAINT cities_pk PRIMARY KEY (id),
                               CONSTRAINT cities_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX cities_name_idx ON public.cities ("name");

CREATE TABLE public.zipcodes (
                                id bigserial NOT NULL,
                                zipcode varchar NULL,
                                city bigint NULL,
                                state bigint NULL,
                                population bigint NULL,
                                latitude float8 NULL,
                                longitude float8 NULL,
                                created_at date NULL,
                                updated_at date NULL,
                                deleted_at date NULL,
                                CONSTRAINT zipcode_pk PRIMARY KEY (id),
                                CONSTRAINT zipcode_cities_fk FOREIGN KEY (city) REFERENCES public.cities(id) ON DELETE CASCADE ON UPDATE CASCADE,
                                CONSTRAINT zipcode_states_fk FOREIGN KEY (state) REFERENCES public.states(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX zipcode_state_idx ON public.zipcodes (state,population);

CREATE TABLE public.cityareacodes (
                                      id bigserial NOT NULL,
                                      areacode bigint NOT NULL,
                                      city bigint NOT NULL,
                                      created_at date NULL,
                                      updated_at date NULL,
                                      deleted_at date NULL,
                                      CONSTRAINT cityareacodes_cities_fk FOREIGN KEY (city) REFERENCES public.cities(id) ON DELETE CASCADE ON UPDATE CASCADE,
                                      CONSTRAINT cityareacodes_cities_fk_1 FOREIGN KEY (city) REFERENCES public.cities(id)
);
CREATE INDEX cityareacodes_city_idx ON public.cityareacodes (city,created_at,updated_at,deleted_at);

CREATE TABLE public.keywords (
                                 id bigserial NOT NULL,
                                 keyword varchar NULL,
                                 created_at date NULL,
                                 updated_at date NULL,
                                 deleted_at date NULL,
                                 CONSTRAINT keywords_pk PRIMARY KEY (id)
);



commit;