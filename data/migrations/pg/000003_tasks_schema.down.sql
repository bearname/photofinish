BEGIN TRANSACTION;
ALTER TABLE pictures
    DROP task_id;
DROP TABLE IF EXISTS tasks;
COMMIT;