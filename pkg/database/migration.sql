create table admin(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255)
 
);

create table lendor(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255),
    gst varchar(255),
    status bool
);
create table users(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255),
    phone varchar(255),
    status bool
);
create table office(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    office_name varchar(255),
    place varchar(255),
    status bool


)
create table employ(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255),
    phone varchar(255),
    role_f_id INT REFERENCES role(id),
    office_f_id INT REFERENCES office(id),
    status bool
);

create table role(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar(255),
    status bool
);

create table privilages(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    mentioning_privilage varchar(255),
    status bool
);

create table privilages_role_factor(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    role_f_id INT REFERENCES role(id),
    privilages_f_id INT REFERENCES privilages(id),
    status bool

);

create table loan_type(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    loan_type_name varchar(255),
   
    status bool
);

create table rule(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_name varchar(255),
    status bool
);

create table eligible_contraints(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    contraint_key_name varchar(255),
    contraint_condition varchar(255),
    val float,
    status bool

);

create table rule_eligible_contraint(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_f_id INT REFERENCES rule(id),
    eligible_type_f_id INT REFERENCES eligible(id),
    status bool
);
create table loan_rule_contraint(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_f_id INT REFERENCES rule(id),
    loan_type_f_id INT REFERENCES loan_type(id),
    status bool
);

create table document_type(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    document_name varchar(255),
    status bool
);

create table rule_document_contraint(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_f_id INT REFERENCES rule(id),
    document_type_f_id INT REFERENCES document_type(id),
    status bool
);


create table loan(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    loan_type_f_id INT REFERENCES loan_type(id),
    user_type_f_id INT REFERENCES users(id),
    lendor_type_f_id INT REFERENCES lendor(id),
    employ_f_id INT REFERENCES employ(id),
    status varchar(255)

); 


create table documents_submission(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    loan_type_f_id INT REFERENCES loan(id),
    document_type_f_id INT REFERENCES document_type(id),
    status bool
);













