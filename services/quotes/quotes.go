package quotes

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vkrayzee/testtask-wow/services"
)

const file = "./db/quotes.db"

// QuoteDB is a database of quotes
type QuoteDB struct {
	db   *sql.DB
	Iter int
}

// NewQuoteDB returns new QuoteDB
func New() services.QuoteDB {
	rand.Seed(time.Now().UnixNano())

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	// create table if not exists
	sqlStmt := `
	create table if not exists quotes (id integer not null primary key, quote text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	// insert test data
	// fill(db)

	return &QuoteDB{
		db: db,
	}
}

// GetRandomQuote returns random quote
func (q *QuoteDB) GetRandomQuote() string {
	// count quotes
	stmt := "select count(*) from quotes;"
	rows, err := q.db.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	n := rand.Intn(count) + 1

	// select quote
	stmt = "select quote from quotes where id = ?;"
	rows, err = q.db.Query(stmt, n)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var quote string
	for rows.Next() {
		err = rows.Scan(&quote)
		if err != nil {
			log.Fatal(err)
		}
	}

	return quote
}

func (q *QuoteDB) Close() {
	// close database
	q.db.Close()
}

// fill fills database with test quotes
func fill(db *sql.DB) {
	// truncate table quotes;
	_, err := db.Exec("delete from quotes;")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into quotes(quote) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	qq := []string{
		"When you can stand alone, do not try hard to fit in the crowd.",
		"A wise man will not recognize a fool, and a fool will not recognize a wise man.",
		"Never argue with fools; avoid them as they will harm you and take you down to their place.",
		"Keep attitude and gratitude with you. One will make you weak without another.",
		"Little things done for others will make bigger places for you in the people’s hearts.",
		"A Stone is broken by the last stroke of the hammer. This doesn’t mean that the 1st stroke is useless. Success is the result of continuous effort!",
		"Every single person on the planet has a story. Don’t judge people before you truly know them. The truth might surprise you.",
		"If you do something wrong, then accept it. Denying wrongdoings will never give you self-respect.",
		"There is a limit to intelligence, but no limit to foolishness. Foolish people drown themselves and everything with them.",
		"Friendship is like an umbrella; the more it rains, the more you will need it. Cherish your bonds with friends.",
		"It is not a shame that you are failing, but a shame that you are not standing back to achieve your goals.",
		"Choose people wisely in the path of your life. The right people will take you to your goals.",
		"The word ‘TRUST’ is the base of all relations, but a small mistake made can change its entire meaning. Like just a missing ‘T’ can ‘Rust’ the relation!",
		"Life does not provide Warranties and Guarantees. It only provides possibilities and opportunities for those who are there to make the best use of it!",
		"Lend your ears to people in need and try to hear what they cannot say to you. Read between the lines.",
		"Don’t waste words on people who deserve your silence. Sometimes the most powerful thing you can do is nothing at all.",
		"Successful people always have two things on their lips, Silence and Smile. “SILENCE” to avoid problems and “SMILE” to solve problems.",
		"Success is like your own shadow; if you try to catch it, you will never succeed, ignore it and walk in your own way.. it will follow you.",
		"Not every person is going to understand you and that’s okay. They have a right to their opinion and you have the right to ignore it.",
		"Just because you miss someone doesn’t mean you need them back in your life. Missing is just a part of moving on.",
		"Listen to your elder’s advice, not because they are always right but because they have more experiences of being wrong!",
		"Do not dump your woes upon people, keep the sad story of your life to yourself. Troubles grow by recounting them.",
		"Being alone is really better than laughing with the people who hate you but act like loving you!",
		"The trouble with a Rat race is that even if you win, you are still a Rat. So Always run with Lions. No matter even if you are defeated, you are still a Lion.",
		"Sometimes you need to add respect to love. Respect is the basis of any bond. It makes relationships strong",
		"If you can’t give something new to this world, then you are the burden of this universe. Be the wealth of the world.",
	}

	for i := 0; i < len(qq); i++ {
		_, err = stmt.Exec(qq[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
