package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type LeaderboardPlayer struct {
	Time1 int `json:"time1"`
	Time2 int `json:"time2"`
	Time3 int `json:"time3"`
}

type LeaderboardData struct {
	Leaderboard map[string]LeaderboardPlayer `json:"leaderboard"`
}

func getLeaderboard(c echo.Context, db *sql.DB) error {
	query, err := db.Query(`SELECT * FROM leaderboard`);

	if (err != nil) {
		panic(err.Error);
	}

	defer query.Close();

	data := LeaderboardData{Leaderboard:map[string]LeaderboardPlayer{}};

	for (query.Next()) {
		var name string;

		var time1, time2, time3 int;

		err := query.Scan(&name, &time1, &time2, &time3);

		if (err != nil) {
			panic(err.Error);
		}

		data.Leaderboard[name] = LeaderboardPlayer{Time1:time1, Time2:time2, Time3:time3};
	}

	return c.JSONPretty(http.StatusOK, data, "  ");
}

func main() {
	e := echo.New();

	e.Use(middleware.CORS());

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/NEAGameDB");

	if (err != nil) {
		panic(err.Error);
	}

	defer db.Close();

	e.GET("/leaderboard", func(c echo.Context) error {
		return getLeaderboard(c, db);
	});

	e.Logger.Fatal(e.Start(":1323"));
}