CREATE EXTENSION postgis;
CREATE EXTENSION postgis_topology;
CREATE TABLE spatial_data (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    geom GEOMETRY(POINT, 4326),
    properties JSONB
);
