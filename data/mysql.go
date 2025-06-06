/*
 * Copyright 2023 CoolFunAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"coolfun-admin/configs"

	"ariga.io/atlas/sql/migrate"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"

	"coolfun-admin/data/ent"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mySqlDsn = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

// NewClient returns a new ent client.
// Only support mysql and postgres
func NewClient(config configs.Config) (client *ent.Client, err error) {
	switch config.Database.Type {
	case "mysql":
		return NewMySQLClient(config)
	case "postgres":
		return NewPostgresClient(config)
	}
	return nil, fmt.Errorf("unsupported database: %s. only support mysql and postgres", config.Database.DBName)
}

// NewMySQLClient returns a new ent client.
func NewMySQLClient(config configs.Config) (client *ent.Client, err error) {
	drv, err := sql.Open("mysql", fmt.Sprintf(mySqlDsn,
		config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName))
	if err != nil {
		hlog.Fatalf("failed opening connection to mysql: %v", err)
		return
	}

	var drive dialect.Driver
	if config.IsProd {
		hlog.Info("prod mode, use default mysql driver")
		drive = drv
	} else {
		// Debug driver.
		hlog.Info("dev mode, use debug mysql driver")
		drive = &DebugTimeDriver{
			Driver: drv,
			log: func(ctx context.Context, info ...any) {
				hlog.Info(info...)
			},
		}
	}

	client = ent.NewClient(ent.Driver(drive))

	// Create a local migration directory.
	migrationsPath := "./data/ent/migrations"
	_, err = os.Stat(migrationsPath)
	if err != nil {
		if !os.IsNotExist(err) {
			hlog.Fatalf("failed to stat migrations path: %v", err)
			return nil, err
		}
		// Create the directory if it doesn't exist.
		err = os.MkdirAll(migrationsPath, os.ModePerm)
		if err != nil {
			hlog.Fatalf("failed creating migrations path: %v", err)
			return nil, err
		}
	}
	dir, err := migrate.NewLocalDir(migrationsPath)
	if err != nil {
		hlog.Fatalf("failed creating atlas migration directory: %v", err)
		return nil, err
	}
	// Write migration diff.
	err = client.Schema.Diff(context.Background(), schema.WithDir(dir), schema.WithForeignKeys(false))
	if err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	return
}

// DebugTimeDriver is a driver that logs all driver operations.
type DebugTimeDriver struct {
	dialect.Driver                               // underlying driver.
	log            func(context.Context, ...any) // log function. defaults to log.Println.
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugTimeDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	d.log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, time.Since(start)))
	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *DebugTimeDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	start := time.Now()
	result, err := drv.ExecContext(ctx, query, args...)
	d.log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return result, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugTimeDriver) Query(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	d.log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v time=%v", query, args, time.Since(start)))
	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *DebugTimeDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	start := time.Now()
	rows, err := drv.QueryContext(ctx, query, args...)
	d.log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return rows, err
}
