CREATE TABLE "podcasts" (
    id BLOB NOT NULL PRIMARY KEY,
    name TEXT
);

CREATE TABLE "episodes" (
    id BLOB NOT NULL PRIMARY KEY,
    podcast_id BLOB NOT NULL,
    number INTEGER,
    name TEXT,
    aired_at DATE NOT NULL,
    patreon_only INTEGER,
    CONSTRAINT postcast_id_check FOREIGN KEY (podcast_id) REFERENCES podcasts (id)
);

CREATE INDEX episodes_by_podcast ON episodes(podcast_id);
CREATE INDEX episode_order ON episodes(number);

CREATE TABLE "speakers" (
    id BLOB NOT NULL PRIMARY KEY,
    transcript_name TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE "utterances" (
    id BLOB NOT NULL PRIMARY KEY,
    episode_id BLOB NOT NULL,
    sequence_no INTEGER NOT NULL,
    is_paralinguistic INTEGER NOT NULL,
    start_time INTEGER,
    end_time INTEGER,
    utterance TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    CONSTRAINT episode_id_check FOREIGN KEY (episode_id) REFERENCES episodes (id)
);

CREATE TABLE "utterance_speakers" (
  utterance_id BLOB NOT NULL,
  speaker_id BLOB NOT NULL,
  PRIMARY KEY (utterance_id, speaker_id),
  CONSTRAINT utterance_id_check FOREIGN KEY (utterance_id) REFERENCES utterances (id),
  CONSTRAINT speaker_id_check FOREIGN KEY (speaker_id) REFERENCES speakers (id)
);

CREATE INDEX utterance_order ON utterances(sequence_no);
CREATE INDEX utterance_time_order ON utterances(start_time);
CREATE INDEX utterances_by_episode ON utterances(episode_id);

CREATE TABLE "utterance_fragments" (
    id BLOB NOT NULL PRIMARY KEY,
    value TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE "utterance_fragment_links" (
    utterance_id BLOB NOT NULL,
    sequence_no INTEGER NOT NULL,
    utterance_fragment_id BLOB NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (utterance_id, sequence_no),
    CONSTRAINT utterance_id_check FOREIGN KEY (utterance_id) REFERENCES utterances (id),
    CONSTRAINT utterance_fragments_id_check FOREIGN KEY (utterance_fragment_id) REFERENCES utterance_fragments (id)
);

CREATE INDEX utterance_fragment_links_by_utterance_fragment
    ON utterance_fragment_links(utterance_fragment_id);
