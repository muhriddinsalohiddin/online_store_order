CREATE TABLE orders(
	id UUID PRIMARY KEY,
	book_id UUID,
	description VARCHAR(128),
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp	
);