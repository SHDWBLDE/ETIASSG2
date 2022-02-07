CREATE TABLE learning_objectives (
    id SERIAL PRIMARY KEY,
    module INTEGER,
    CONSTRAINT learning_objectives_module_fkey FOREIGN KEY (module) REFERENCES modules(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(64) NOT NULL UNIQUE
);