CREATE TABLE "episodes" (
    id BLOB NOT NULL PRIMARY KEY,
    name TEXT,
    aired_at INTEGER NOT NULL,
    patreon_only INTEGER
);

CREATE TABLE "speaker" (
    id BLOB NOT NULL PRIMARY KEY,
    name TEXT
);

CREATE TABLE "utterances" (
    id BLOB NOT NULL PRIMARY KEY,
    speaker_id BLOB NOT NULL,
    episode_id BLOB NOT NULL,
    start_time INTEGER NOT NULL
);

CREATE INDEX utterances_by_episode ON utterances(episode_id);
CREATE INDEX utterances_by_speaker ON utterances(speaker_id);

CREATE TABLE "utterance_fragments" (
    id BLOB NOT NULL PRIMARY KEY,
    value TEXT
);

CREATE TABLE "utterance_fragment_links" (
    utterance_id BLOB NOT NULL,
    sequence_no INTEGER NOT NULL,
    utterance_fragment BLOB NOT NULL
);

CREATE UNIQUE INDEX utterance_fragment_links_key
    ON utterance_fragment_links(utterance_id, sequence_no);
CREATE INDEX utterance_fragment_links_by_utterance_fragment
    ON utterance_fragment_links(utterance_fragment);
