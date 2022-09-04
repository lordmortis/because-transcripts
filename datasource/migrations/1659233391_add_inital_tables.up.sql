CREATE TABLE "podcasts" (
    id uuid NOT NULL PRIMARY KEY,
    name TEXT,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE "episodes" (
    id uuid NOT NULL PRIMARY KEY,
    podcast_id uuid NOT NULL,
    number INTEGER,
    name TEXT,
    aired_at DATE NOT NULL,
    patreon_only bool NOT NULL DEFAULT false,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT podcast_id_check FOREIGN KEY (podcast_id) REFERENCES podcasts (id)
);

CREATE INDEX episodes_by_podcast ON episodes(podcast_id);
CREATE INDEX episode_order ON episodes(number);

CREATE TABLE "speakers" (
    id uuid NOT NULL PRIMARY KEY,
    transcript_name TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE "turns" (
    id uuid NOT NULL PRIMARY KEY,
    episode_id uuid NOT NULL,
    sequence_no INTEGER NOT NULL,
    start_time INTEGER,
    end_time INTEGER,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT episode_id_check FOREIGN KEY (episode_id) REFERENCES episodes (id)
);

CREATE INDEX turns_by_episode ON turns(episode_id);

CREATE TABLE "utterances" (
    id uuid NOT NULL PRIMARY KEY,
    turn_id uuid NOT NULL,
    sequence_no INTEGER NOT NULL,
    is_paralinguistic bool NOT NULL,
    start_time INTEGER,
    end_time INTEGER,
    utterance TEXT,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT turn_id_check FOREIGN KEY (turn_id) REFERENCES turns (id)
);

CREATE INDEX utterance_order ON utterances(sequence_no);
CREATE INDEX utterance_time_order ON utterances(start_time);
CREATE INDEX utterance_sequence_order ON utterances(sequence_no);
CREATE INDEX utterances_by_turns ON utterances(turn_id);

CREATE TABLE "utterance_speakers" (
  utterance_id uuid NOT NULL,
  speaker_id uuid NOT NULL,
  PRIMARY KEY (utterance_id, speaker_id),
  CONSTRAINT utterance_id_check FOREIGN KEY (utterance_id) REFERENCES utterances (id),
  CONSTRAINT speaker_id_check FOREIGN KEY (speaker_id) REFERENCES speakers (id)
);
