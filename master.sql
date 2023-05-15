-- public.transactions definition

-- Drop table

-- DROP TABLE public.transactions;

CREATE TABLE public.transactions (
	transaction_id varchar(50) NOT NULL,
	transaction_total numeric(10, 2) NULL,
	transaction_person_in_charge varchar(50) NULL,
	transaction_date date NULL,
	created_date date NULL,
	last_modified date NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id)
);

-- Dummy Data
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN001', 100.00, 'user1', '2023-05-01', '2023-05-01', '2023-05-01');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN002', 50.00, 'user2', '2023-05-02', '2023-05-02', '2023-05-02');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN003', 75.25, 'user3', '2023-05-03', '2023-05-03', '2023-05-03');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN004', 30.50, 'user4', '2023-05-04', '2023-05-04', '2023-05-04');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN005', 200.00, 'user5', '2023-05-05', '2023-05-05', '2023-05-05');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN006', 45.75, 'user1', '2023-05-06', '2023-05-06', '2023-05-06');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN007', 80.00, 'user2', '2023-05-07', '2023-05-07', '2023-05-07');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN008', 22.50, 'user3', '2023-05-08', '2023-05-08', '2023-05-08');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN009', 150.00, 'user4', '2023-05-09', '2023-05-09', '2023-05-09');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN010', 500.00, 'user5', '2023-05-10', '2023-05-10', '2023-05-10');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN011', 55.25, 'user1', '2023-05-11', '2023-05-11', '2023-05-11');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN012', 90.00, 'user2', '2023-05-12', '2023-05-12', '2023-05-12');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN013', 75.50, 'user3', '2023-05-13', '2023-05-13', '2023-05-13');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN014', 40.00, 'user4', '2023-05-14', '2023-05-14', '2023-05-14');
INSERT INTO public.transactions (transaction_id, transaction_total, transaction_person_in_charge, transaction_date, created_date, last_modified) VALUES('TXN015', 350.25, 'user5', '2023-05-15', '2023-05-15', '2023-05-15');