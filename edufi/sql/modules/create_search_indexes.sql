-- create a new colomn from the name column
-- name_searchable contains the searchable tokens from name column
ALTER TABLE modules ADD COLUMN name_searchable tsvector 
    GENERATED ALWAYS AS (to_tsvector('english',name)) 
    STORED
;

-- create an index of the name_searchable column
-- the index helps to speed up the search operation
CREATE INDEX module_search_index ON modules USING GIN (name_searchable);


-- create a query to search e.g
-- SELECT name from modules WHERE name_searchable @@ websearch_to_tsquery('web');