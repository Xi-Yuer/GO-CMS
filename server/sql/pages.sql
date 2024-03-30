ALTER TABLE pages
    ADD CONSTRAINT pages_parent_page_fk FOREIGN KEY (parent_page) REFERENCES pages (page_id);