CREATE TABLE "podcast" (
    id BLOB NOT NULL PRIMARY KEY,
    name TEXT
);

CREATE TABLE "episodes" (
    id BLOB NOT NULL PRIMARY KEY,
    podcast_id BLOB NOT NULL,
    number integer,
    name TEXT,
    aired_at INTEGER NOT NULL,
    patreon_only INTEGER,
    CONSTRAINT postcast_id_check FOREIGN KEY (podcast_id) REFERENCES podcast (id)
);

CREATE INDEX episodes_by_podcast ON episodes(podcast_id);
CREATE INDEX episode_order ON episodes(number);

CREATE TABLE "speaker" (
    id BLOB NOT NULL PRIMARY KEY,
    name TEXT,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
);

CREATE TABLE "utterances" (
    id BLOB NOT NULL PRIMARY KEY,
    speaker_id BLOB NOT NULL,
    episode_id BLOB NOT NULL,
    start_time INTEGER NOT NULL,
    end_time INTEGER NOT NULL,
    utterance TEXT,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    CONSTRAINT speaker_id_check FOREIGN KEY (speaker_id) REFERENCES speaker (id),
    CONSTRAINT episode_id_check FOREIGN KEY (episode_id) REFERENCES episodes (id)
);

CREATE INDEX utterance_order ON utterances(start_time);
CREATE INDEX utterances_by_episode ON utterances(episode_id);
CREATE INDEX utterances_by_speaker ON utterances(speaker_id);

CREATE TABLE "utterance_fragments" (
    id BLOB NOT NULL PRIMARY KEY,
    value TEXT,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
);

CREATE TABLE "utterance_fragment_links" (
    utterance_id BLOB NOT NULL,
    sequence_no INTEGER NOT NULL,
    utterance_fragment_id BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    PRIMARY KEY (utterance_id, sequence_no),
    CONSTRAINT utterance_id_check FOREIGN KEY (utterance_id) REFERENCES utterances (id),
    CONSTRAINT utterance_fragments_id_check FOREIGN KEY (utterance_fragment_id) REFERENCES utterance_fragments (id)
);

CREATE INDEX utterance_fragment_links_by_utterance_fragment
    ON utterance_fragment_links(utterance_fragment_id);
