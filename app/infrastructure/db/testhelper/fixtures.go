package testhelper

import (
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/kakkky/app/infrastructure/db"
)

func SetupFixtures(fixture_path ...string) {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db.GetDB()),
		testfixtures.Dialect("mysql"),
		testfixtures.SkipResetSequences(),
		testfixtures.Files(
			fixture_path...,
		),
	)
	if err != nil {
		log.Printf("testfixtures failed to create Loader instance:%v", err)
	}
	// テーブルのデータを削除&用意
	if err := fixtures.Load(); err != nil {
		log.Printf("testfixtures failed to load fixtures:%v", err)
	}
}
