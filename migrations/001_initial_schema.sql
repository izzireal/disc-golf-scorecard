-- 1. Courses
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Course Holes
CREATE TABLE course_holes (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    hole_number INT NOT NULL,
    tee_color VARCHAR(50) NOT NULL DEFAULT 'Red',
    par INT NOT NULL,
    distance_feet INT,
    difficulty_rating DECIMAL(3,1),
    description TEXT,
    UNIQUE(course_id, hole_number, tee_color)
);

-- 3. Players
CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nickname VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Rounds
CREATE TABLE rounds (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id),
    date_played DATE NOT NULL,
    group_code VARCHAR(20) UNIQUE,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Round Participants
CREATE TABLE round_participants (
    id SERIAL PRIMARY KEY,
    round_id INT REFERENCES rounds(id) ON DELETE CASCADE,
    player_id INT REFERENCES players(id),
    final_score INT,
    UNIQUE(round_id, player_id)
);

-- 6. Hole Scores (Summary - ALWAYS exists)
CREATE TABLE hole_scores (
    id SERIAL PRIMARY KEY,
    round_id INT REFERENCES rounds(id) ON DELETE CASCADE,
    player_id INT REFERENCES players(id),
    hole_number INT NOT NULL,
    tee_color VARCHAR(50) NOT NULL DEFAULT 'Red',
    score INT NOT NULL,                    -- Always required
    putts INT,
    notes TEXT,
    UNIQUE(round_id, player_id, hole_number)
);

-- 7. Throws (OPTIONAL - only when user wants detail)
CREATE TABLE throws (
    id SERIAL PRIMARY KEY,
    hole_score_id INT REFERENCES hole_scores(id) ON DELETE CASCADE,
    throw_number INT NOT NULL,
    disc_used VARCHAR(100),
    throw_type VARCHAR(30),           -- Backhand, Forehand, Roller, etc.
    outcome VARCHAR(50),              -- Fairway, OB, Tree, Approach, In Basket, etc.
    distance_feet INT,
    notes TEXT,
    UNIQUE(hole_score_id, throw_number)
);

-- Indexes
CREATE INDEX idx_hole_scores_round ON hole_scores(round_id);
CREATE INDEX idx_throws_hole_score ON throws(hole_score_id);