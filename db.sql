CREATE TABLE
  public.posts (
    id serial NOT NULL,
    user_id integer,
    title text NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    short character varying(255) NOT NULL
  );

ALTER TABLE
  public.posts
ADD
  CONSTRAINT posts_pkey PRIMARY KEY (id);


CREATE TABLE public.users (
                              id serial NOT NULL,
                              email character varying(255) NOT NULL,
                              first_name character varying(255) NOT NULL,
                              last_name character varying(255) NOT NULL,
                              password character varying(255) NOT NULL,
                              is_active integer DEFAULT 0,
                              is_admin integer default 0,
                              created_at timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE
  public.users
ADD
  CONSTRAINT users_pkey PRIMARY KEY (id);
  
  
  ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE CASCADE;




INSERT INTO posts (user_id, title, description, short)
    VALUES 
	(2, 'How to do something', 'Lorem lorem Lorem lorem Lorem lorem Lorem lorem','jUST A SHORT description');