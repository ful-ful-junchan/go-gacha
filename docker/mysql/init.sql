CREATE TABLE gachas (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name          VARCHAR(100) NOT NULL,
  pity_threshold INT UNSIGNED NOT NULL DEFAULT 100,
  starts_at     DATETIME NOT NULL,
  ends_at       DATETIME NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_pity_threshold_positive CHECK (pity_threshold > 0),
  CONSTRAINT chk_period_valid CHECK (starts_at < ends_at)
) ENGINE=InnoDB;

CREATE TABLE items (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name       VARCHAR(100) NOT NULL,
  rarity     ENUM('SSR','SR','R') NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE gacha_items (
  gacha_id BIGINT UNSIGNED NOT NULL,
  item_id  BIGINT UNSIGNED NOT NULL,
  weight   INT UNSIGNED NOT NULL,
  PRIMARY KEY (gacha_id, item_id),
  KEY idx_gacha (gacha_id)
) ENGINE=InnoDB;

CREATE TABLE user_pity_counters (
  user_id    BIGINT UNSIGNED NOT NULL,
  gacha_id   BIGINT UNSIGNED NOT NULL,
  count      INT UNSIGNED NOT NULL DEFAULT 0,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, gacha_id)
) ENGINE=InnoDB;

CREATE TABLE gacha_histories (
  id       BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id  BIGINT UNSIGNED NOT NULL,
  gacha_id BIGINT UNSIGNED NOT NULL,
  item_id  BIGINT UNSIGNED NOT NULL,
  is_pity  TINYINT(1) NOT NULL DEFAULT 0,
  drawn_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY idx_user_gacha (user_id, gacha_id, drawn_at)
) ENGINE=InnoDB;
