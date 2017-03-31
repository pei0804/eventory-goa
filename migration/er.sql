SET SESSION FOREIGN_KEY_CHECKS=0;

/* Drop Indexes */

DROP INDEX search_index ON event;



/* Drop Tables */

DROP TABLE IF EXISTS event_genre;
DROP TABLE IF EXISTS user_keep_status;
DROP TABLE IF EXISTS event;
DROP TABLE IF EXISTS user_follow_genre;
DROP TABLE IF EXISTS genre;
DROP TABLE IF EXISTS user_follow_pref;
DROP TABLE IF EXISTS pref;
DROP TABLE IF EXISTS user_terminals;
DROP TABLE IF EXISTS user;




/* Create Tables */

-- イベント
CREATE TABLE event
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'イベントID',
	api_type enum('atdn','connpass','doorkeeper') NOT NULL COMMENT 'APIの種類',
	identifier varchar(10) NOT NULL COMMENT '識別子(api-event_id)',
	title varchar(200) NOT NULL COMMENT 'イベント名',
	description text NOT NULL COMMENT '説明',
	url text NOT NULL COMMENT 'イベントページURL',
	limits int NOT NULL COMMENT '参加人数上限',
	wait int NOT NULL COMMENT 'キャンセル待ち人数',
	accepte int NOT NULL COMMENT '参加済み人数',
	pref_id int(2) unsigned COMMENT '都道府県ID',
	address text NOT NULL COMMENT '住所',
	start_at datetime NOT NULL COMMENT '開催日時',
	end_at datetime NOT NULL COMMENT '終了日時',
	data_hash char(64) NOT NULL COMMENT 'データ識別Hash',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id),
	UNIQUE (identifier)
) COMMENT = 'イベント' DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;


-- イベントジャンル
CREATE TABLE event_genre
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	genre_id bigint(20) unsigned NOT NULL COMMENT 'ジャンルID',
	event_id bigint(20) unsigned NOT NULL COMMENT 'イベントID',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'イベントジャンル';


-- ジャンル
CREATE TABLE genre
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ジャンルID',
	name varchar(30) NOT NULL COMMENT 'ジャンル名(表示用)',
	keyword varchar(50) NOT NULL COMMENT 'キーワード',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ジャンル';


-- 都道府県
CREATE TABLE pref
(
	id int(2) unsigned NOT NULL AUTO_INCREMENT COMMENT '都道府県ID',
	name char(4) NOT NULL COMMENT '都道府県名',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = '都道府県';


-- ユーザー
CREATE TABLE user
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
	name varchar(30) COMMENT 'ユーザー名',
	email text COMMENT 'メールアドレス',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザー';


-- ジャンル
CREATE TABLE user_follow_genre
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ジャンルID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	genre_id bigint(20) unsigned NOT NULL COMMENT 'ジャンルID',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ジャンル';


-- 都道府県
CREATE TABLE user_follow_pref
(
	id int(2) unsigned NOT NULL AUTO_INCREMENT COMMENT '都道府県ID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	pref_id int(2) unsigned NOT NULL COMMENT '都道府県ID',
	name char(4) NOT NULL COMMENT '都道府県名',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = '都道府県';


-- ユーザーのキープ状態
CREATE TABLE user_keep_status
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	event_id bigint(20) unsigned NOT NULL COMMENT 'イベントID',
	status enum('keep','nokeep') NOT NULL COMMENT '状態',
	batch_processed boolean DEFAULT 'false' NOT NULL COMMENT 'バッチ処理済み',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザーのキープ状態';


-- ユーザー
CREATE TABLE user_terminals
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	platform varchar(20) NOT NULL COMMENT 'OSとバージョン',
	client_version varchar(10) NOT NULL COMMENT 'アプリのバージョン',
	token char(64) NOT NULL COMMENT 'トークン',
	identifier char(36) NOT NULL COMMENT '識別子(android:Android_ID, ios:IDFV)',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザー';



/* Create Foreign Keys */

ALTER TABLE event_genre
	ADD FOREIGN KEY (event_id)
	REFERENCES event (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_keep_status
	ADD FOREIGN KEY (event_id)
	REFERENCES event (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE event_genre
	ADD FOREIGN KEY (genre_id)
	REFERENCES genre (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_genre
	ADD FOREIGN KEY (genre_id)
	REFERENCES genre (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE event
	ADD FOREIGN KEY (pref_id)
	REFERENCES pref (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_pref
	ADD FOREIGN KEY (pref_id)
	REFERENCES pref (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_genre
	ADD FOREIGN KEY (user_id)
	REFERENCES user (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_pref
	ADD FOREIGN KEY (user_id)
	REFERENCES user (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_keep_status
	ADD FOREIGN KEY (user_id)
	REFERENCES user (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_terminals
	ADD FOREIGN KEY (user_id)
	REFERENCES user (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;



/* Create Indexes */

CREATE INDEX search_index USING BTREE ON event (end_at ASC, updated_at ASC, address ASC);



