--blog

INSERT INTO blog.blog_posts (
    "id", "author_id", "tour_id", "title", "description", "creation_date", "images", "comments", "ratings", "status"
) VALUES
(gen_random_uuid(), 1, 101, 'Prva avantura u planinama', 'Opis mog prvog putovanja u Alpe.', '2025-04-01 10:00:00+00', 
 '["https://example.com/img1.jpg"]'::jsonb,
 '[{"AuthorID": 2, "AuthorUsername": "", "Text": "Prelepo iskustvo!", "CreationTime": "2025-04-01T10:30:00Z", "LastUpdatedTime": "2025-04-01T10:30:00Z"}]'::jsonb,
 '[{"AuthorID": 2, "CreationTime": "2025-04-01T10:35:00Z", "IsPositive": true}]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 2, 102, 'Obilazak jezera', 'Priča o danu provedenom na jezeru.', '2025-04-02 14:00:00+00', 
 '["https://example.com/img2.jpg", "https://example.com/img3.jpg"]'::jsonb,
 '[{"AuthorID": 3, "AuthorUsername": "", "Text": "Voda je bila kristalno čista!", "CreationTime": "2025-04-02T14:20:00Z", "LastUpdatedTime": "2025-04-02T14:20:00Z"}]'::jsonb,
 '[{"AuthorID": 3, "CreationTime": "2025-04-02T14:25:00Z", "IsPositive": true}, {"AuthorID": 4, "CreationTime": "2025-04-02T15:00:00Z", "IsPositive": false}]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 3, 103, 'Gradska tura', 'Istraživanje starog grada.', '2025-04-03 09:00:00+00', 
 '[]'::jsonb,
 '[]'::jsonb,
 '[{"AuthorID": 5, "CreationTime": "2025-04-03T09:15:00Z", "IsPositive": true}]'::jsonb, 'DRAFT'),

(gen_random_uuid(), 4, 104, 'Pustinjska ekspedicija', 'Putovanje kroz pustinju.', '2025-04-04 12:00:00+00', 
 '["https://example.com/img4.jpg"]'::jsonb,
 '[{"AuthorID": 6, "AuthorUsername": "", "Text": "Vrućina je bila nepodnošljiva!", "CreationTime": "2025-04-04T12:30:00Z", "LastUpdatedTime": "2025-04-04T12:30:00Z"}]'::jsonb,
 '[]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 5, 105, 'Planinarenje na vrh', 'Uspon na najviši vrh.', '2025-04-05 07:00:00+00', 
 '["https://example.com/img5.jpg"]'::jsonb,
 '[{"AuthorID": 7, "AuthorUsername": "", "Text": "Pogled je bio neverovatan!", "CreationTime": "2025-04-05T07:45:00Z", "LastUpdatedTime": "2025-04-05T07:45:00Z"}]'::jsonb,
 '[{"AuthorID": 7, "CreationTime": "2025-04-05T08:00:00Z", "IsPositive": true}]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 1, 106, 'Druga tura u šumu', 'Šetnja kroz gustu šumu.', '2025-04-06 15:00:00+00', 
 '[]'::jsonb,
 '[]'::jsonb,
 '[{"AuthorID": 8, "CreationTime": "2025-04-06T15:20:00Z", "IsPositive": false}]'::jsonb, 'DRAFT'),

(gen_random_uuid(), 2, 107, 'Obala mora', 'Dan na plaži.', '2025-04-07 11:00:00+00', 
 '["https://example.com/img6.jpg"]'::jsonb,
 '[{"AuthorID": 9, "AuthorUsername": "", "Text": "More je bilo prelepo!", "CreationTime": "2025-04-07T11:30:00Z", "LastUpdatedTime": "2025-04-07T11:30:00Z"}]'::jsonb,
 '[{"AuthorID": 9, "CreationTime": "2025-04-07T11:40:00Z", "IsPositive": true}]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 3, 108, 'Selo u brdima', 'Poseta tradicionalnom selu.', '2025-04-08 13:00:00+00', 
 '["https://example.com/img7.jpg"]'::jsonb,
 '[{"AuthorID": 10, "AuthorUsername": "", "Text": "Hrana je bila odlična!", "CreationTime": "2025-04-08T13:15:00Z", "LastUpdatedTime": "2025-04-08T13:15:00Z"}]'::jsonb,
 '[]'::jsonb, 'PUBLISHED'),

(gen_random_uuid(), 4, 109, 'Noć pod zvezdama', 'Kampovanje u prirodi.', '2025-04-09 20:00:00+00', 
 '[]'::jsonb,
 '[]'::jsonb,
 '[{"AuthorID": 11, "CreationTime": "2025-04-09T20:30:00Z", "IsPositive": true}]'::jsonb, 'DRAFT'),

(gen_random_uuid(), 5, 110, 'Reka i kanu', 'Vožnja kanuom niz reku.', '2025-04-10 08:00:00+00', 
 '["https://example.com/img8.jpg"]'::jsonb,
 '[{"AuthorID": 12, "AuthorUsername": "", "Text": "Avantura za pamćenje!", "CreationTime": "2025-04-10T08:45:00Z", "LastUpdatedTime": "2025-04-10T08:45:00Z"}]'::jsonb,
 '[{"AuthorID": 12, "CreationTime": "2025-04-10T09:00:00Z", "IsPositive": true}]'::jsonb, 'PUBLISHED');