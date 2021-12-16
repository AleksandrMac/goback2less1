CREATE TABLE public.account (
    id uuid NOT NULL DEFAULT UUID_GENERATE_V4(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    visited_at TIMESTAMP NOT NULL DEFAULT NOW(),
    reseted_at TIMESTAMP NOT NULL DEFAULT NOW(),
	
    name VARCHAR(150),
	username VARCHAR(150),
	email VARCHAR(150) NOT NULL,
    password VARCHAR(150),
    block BOOLEAN NOT NULL DEFAULT false,
    
    send_email BOOLEAN NOT NULL DEFAULT false,
    otp_key VARCHAR(1500),
    otep VARCHAR(1500),
    
    CONSTRAINT account_pkey PRIMARY KEY (id),
);

COMMENT ON COLUMN public.account.otp_key IS 'Two factor authentication encrypted keys'
COMMENT ON COLUMN public.account.otep IS 'One time emergency passwords'