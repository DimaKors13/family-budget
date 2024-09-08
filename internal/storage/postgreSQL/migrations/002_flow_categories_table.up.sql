CREATE TABLE IF NOT EXISTS flow_categories (
 id SERIAL PRIMARY KEY,
 name VARCHAR(64) NOT NULL,
 multiplier INT NOT NULL DEFAULT 1,
 parent_id INT
);
CREATE INDEX miltiplier_flow_cat ON flow_categories (multiplier); 