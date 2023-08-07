// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ptonlix/drinkGPT/common/globalkey"
	"github.com/ptonlix/drinkGPT/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	goodsFieldNames          = builder.RawFieldNames(&Goods{})
	goodsRows                = strings.Join(goodsFieldNames, ",")
	goodsRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	goodsRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDrinksgptGoodsIdPrefix = "cache:drinksgpt:goods:id:"
)

type (
	goodsModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Goods) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Goods, error)
		Update(ctx context.Context, session sqlx.Session, data *Goods) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Goods) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultGoodsModel struct {
		sqlc.CachedConn
		table string
	}

	Goods struct {
		Id          int64     `db:"id"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
		DeleteTime  time.Time `db:"delete_time"`
		DelState    int64     `db:"del_state"`
		Version     int64     `db:"version"`      // 版本号
		CategoryId  int64     `db:"category_id"`  // 饮品类型ID
		Name        string    `db:"name"`         // 饮品名称
		Ingredients string    `db:"ingredients"`  // 饮品成分
		Tea         string    `db:"tea"`          // 饮品茶底
		CupCapacity string    `db:"cup_capacity"` // 饮品容量
	}
)

func newGoodsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultGoodsModel {
	return &defaultGoodsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`goods`",
	}
}

func (m *defaultGoodsModel) Insert(ctx context.Context, session sqlx.Session, data *Goods) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	drinksgptGoodsIdKey := fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, goodsRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity)
	}, drinksgptGoodsIdKey)
}

func (m *defaultGoodsModel) FindOne(ctx context.Context, id int64) (*Goods, error) {
	drinksgptGoodsIdKey := fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, id)
	var resp Goods
	err := m.QueryRowCtx(ctx, &resp, drinksgptGoodsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", goodsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) Update(ctx context.Context, session sqlx.Session, data *Goods) (sql.Result, error) {
	drinksgptGoodsIdKey := fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity, data.Id)
	}, drinksgptGoodsIdKey)
}

func (m *defaultGoodsModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Goods) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	drinksgptGoodsIdKey := fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, goodsRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.CategoryId, data.Name, data.Ingredients, data.Tea, data.CupCapacity, data.Id, oldVersion)
	}, drinksgptGoodsIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultGoodsModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	drinksgptGoodsIdKey := fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, drinksgptGoodsIdKey)
	return err
}

func (m *defaultGoodsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDrinksgptGoodsIdPrefix, primary)
}
func (m *defaultGoodsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", goodsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultGoodsModel) tableName() string {
	return m.table
}
