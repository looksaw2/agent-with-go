package repository

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/looksaw2/ai-agent-with-go/cards/internal/db"
	"github.com/looksaw2/ai-agent-with-go/cards/model"
)

type PostgresRepository struct {
	q *db.Queries
	conn *pgx.Conn
}

//初始化PGREPOSITORY
func NewPostgresRepository(dbconn string)(*PostgresRepository,error) {
	conn , err := pgx.Connect(context.Background(),dbconn)
	if err != nil {
		slog.Info("Postgres DB Error",
		"Can't connect to the db instance",
			err,
		)
		return nil,err
	}
	q := db.New(conn)
	return &PostgresRepository{
		q: q,
		conn: conn,
	},nil
}
//实现接口
func(r *PostgresRepository)CreateTodo(ctx context.Context , todo *model.Todo ) error {
	tx ,err := r.conn.Begin(ctx)
	if err != nil {
		slog.Info("Start Transcation failed.....","Error",err)
		return err
	}
	defer func(){
		if rollBackErr := tx.Rollback(ctx);  rollBackErr != nil {
			if !errors.Is(rollBackErr,pgx.ErrTxClosed) && !strings.Contains(rollBackErr.Error(),"transaction already committed"){
				  slog.Error("事务回滚失败", "Error", rollBackErr)
			}
		}
	}()
	qtx := r.q.WithTx(tx)
	_ , err = qtx.CreateTodo(ctx,db.CreateTodoParams{
		Title: todo.Title,
		Description: pgtype.Text{Valid: true,String: todo.Description},
		Completed: pgtype.Bool{Valid: true,Bool: todo.Completed},
	})
	if err != nil {
		slog.Info("Create Todo failed","Err",err)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		slog.Info("Commit Todo failed","Err",err)
		return err 
	}
	return nil
}

//实现接口
func(r *PostgresRepository)GetTodoByID(ctx context.Context , id int)(*model.Todo ,error){
	tx ,err := r.conn.Begin(ctx)
	if err != nil {
		slog.Info("Start Transcation failed.....","Error",err)
		return nil ,err
	}
	defer func(){
		if rollBackErr := tx.Rollback(ctx); rollBackErr != nil {
			if ! errors.Is(rollBackErr,pgx.ErrTxClosed) && ! strings.Contains(rollBackErr.Error(),"transaction already committed"){
				slog.Info("事物回滚失败","Error",rollBackErr)
			}
		}
	}()

	qtx := r.q.WithTx(tx)
	r1 , err := qtx.GetTodoByID(ctx,int64(id))
	if err != nil {
		slog.Info("Get Todo By ID","failed",err)
		return nil ,err
	}
	if err = tx.Commit(ctx); err != nil {
		slog.Info("Commit Todo","failed",err)
		return nil ,err
	}
	return &model.Todo{
		ID: int(r1.ID),
		Title: r1.Title,
		Description: r1.Description.String,
		Completed: r1.Completed.Bool,
		CreatedAt: r1.CreatedAt.Time,
		UpdatedAt: r1.UpdatedAt.Time,
	},nil
}

//实现接口
func(r *PostgresRepository)GetAllTodos(ctx context.Context) ([]*model.Todo, error){
	tx ,err := r.conn.Begin(ctx)
	if err != nil {
		slog.Info("Start Transcation failed.....","Error",err)
		return nil ,err
	}
	defer func(){
		if rollBackErr := tx.Rollback(ctx); rollBackErr != nil {
			if ! errors.Is(rollBackErr,pgx.ErrTxClosed) && ! strings.Contains(rollBackErr.Error(),"transaction already committed"){
				slog.Info("事物回滚失败","Error",rollBackErr)
			}
		}
	}()

	qtx := r.q.WithTx(tx)
	rR , err := qtx.GetAllTodos(ctx)
	if err != nil {
		slog.Info("Get Todo All","failed",err)
		return nil ,err
	}
	todos := make([]*model.Todo,0)
	for _ ,r := range rR {
		todo := &model.Todo{
			ID: int(r.ID),
			Title: r.Title,
			Description: r.Description.String,
			Completed: r.Completed.Bool,
			CreatedAt: r.CreatedAt.Time,
			UpdatedAt: r.UpdatedAt.Time,
		}
		todos = append(todos, todo)
	}
	if err = tx.Commit(ctx); err != nil {
		slog.Info("Commit Todo","failed",err)
		return nil ,err
	}
	return todos ,nil
}

func(r *PostgresRepository)UpdateTodo(ctx context.Context ,id int, updates map[string]any)error {
	tx ,err := r.conn.Begin(ctx)
	if err != nil {
		slog.Info("Start Transcation failed.....","Error",err)
		return err
	}
	defer func(){
		if rollBackErr := tx.Rollback(ctx); rollBackErr != nil {
			if ! errors.Is(rollBackErr,pgx.ErrTxClosed) && ! strings.Contains(rollBackErr.Error(),"transaction already committed"){
				slog.Info("事物回滚失败","Error",rollBackErr)
			}
		}
	}()
	toUpdate := db.UpdateTodoParams{
		ID: int64(id),
		Description: pgtype.Text{Valid: false},
		Completed: pgtype.Bool{Valid: false},
		UpdatedAt: pgtype.Timestamp{Valid: true,Time: time.Now()},

	}
	title , ok:= updates["title"]; 
	if ok {
		toUpdate.Title = title.(string)
	}


	description , ok := updates["description"];
	if ok {
		toUpdate.Description = pgtype.Text{Valid: true,String: description.(string)}
	}
	
	completed , ok := updates["completed"]; 
	if ok {
		toUpdate.Completed = pgtype.Bool{Valid: true,Bool: completed.(bool)}
	}
	

	qtx := r.q.WithTx(tx)
	err = qtx.UpdateTodo(ctx,toUpdate)
	if err != nil {
		slog.Info("update Todo failed","Error",err)
		return err
	}
		if err = tx.Commit(ctx); err != nil {
		slog.Info("Commit Todo","failed",err)
		return err
	}
	return nil
}